package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-zoo/bone"
	"github.com/mainflux/mainflux/internal/apiutil"
	"github.com/mainflux/mainflux/pkg/errors"
	"github.com/mainflux/mainflux/ws"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func work(svc ws.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decodeRequest(nil, r)
		if err != nil {
			encodeError(context.TODO(), err, w)
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Warn(fmt.Sprintf("Failed to upgrade connection to websocket: %s", err.Error()))
			encodeError(context.TODO(), err, w)
		}
		conn.SetCloseHandler(func(code int, text string) error {
			return svc.Unsubscribe(context.TODO(), req.token, req.channel, req.subtopic)
		})
		c := ws.NewClient(req.token, req.channel, req.subtopic, conn)
		if err := svc.Subscribe(context.TODO(), req.channel, req.subtopic, c); err != nil {
			logger.Warn(err.Error())
			encodeError(context.TODO(), err, w)
		}

		msgs := make(chan []byte)
		go func() {
			for msg := range msgs {
				svc.Publish(context.Background(), req.token, req.channel, req.subtopic, msg)
			}
		}()
		go c.Publish(req.channel, req.subtopic, msgs)
	}
}

func decodeRequest(ctx context.Context, r *http.Request) (publishReq, error) {
	channelParts := channelPartRegExp.FindStringSubmatch(r.RequestURI)
	if len(channelParts) < 2 {
		return publishReq{}, errors.ErrMalformedEntity
	}

	subtopic, err := parseSubtopic(channelParts[2])
	if err != nil {
		return publishReq{}, err
	}

	var token string
	_, pass, ok := r.BasicAuth()
	switch {
	case ok:
		token = pass
	case !ok:
		token = apiutil.ExtractThingKey(r)
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return publishReq{}, errors.ErrMalformedEntity
	}
	defer r.Body.Close()

	req := publishReq{
		protocol: protocol,
		channel:  bone.GetValue(r, "id"),
		subtopic: subtopic,
		payload:  payload,
		created:  time.Now(),
		token:    token,
	}

	return req, nil
}

func parseSubtopic(subtopic string) (string, error) {
	if subtopic == "" {
		return subtopic, nil
	}

	subtopic, err := url.QueryUnescape(subtopic)
	if err != nil {
		return "", errMalformedSubtopic
	}
	subtopic = strings.Replace(subtopic, "/", ".", -1)

	elems := strings.Split(subtopic, ".")
	filteredElems := []string{}
	for _, elem := range elems {
		if elem == "" {
			continue
		}

		if len(elem) > 1 && (strings.Contains(elem, "*") || strings.Contains(elem, ">")) {
			return "", errMalformedSubtopic
		}

		filteredElems = append(filteredElems, elem)
	}

	subtopic = strings.Join(filteredElems, ".")
	return subtopic, nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch {
	case errors.Contains(err, errors.ErrAuthentication),
		err == apiutil.ErrBearerToken:
		w.WriteHeader(http.StatusUnauthorized)
	case errors.Contains(err, errors.ErrAuthorization):
		w.WriteHeader(http.StatusForbidden)
	case errors.Contains(err, errors.ErrUnsupportedContentType):
		w.WriteHeader(http.StatusUnsupportedMediaType)
	case errors.Contains(err, errMalformedSubtopic),
		errors.Contains(err, errors.ErrMalformedEntity):
		w.WriteHeader(http.StatusBadRequest)

	default:
		switch e, ok := status.FromError(err); {
		case ok:
			switch e.Code() {
			case codes.Unauthenticated:
				w.WriteHeader(http.StatusUnauthorized)
			case codes.PermissionDenied:
				w.WriteHeader(http.StatusForbidden)
			case codes.Internal:
				w.WriteHeader(http.StatusInternalServerError)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	if errorVal, ok := err.(errors.Error); ok {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(apiutil.ErrorRes{Err: errorVal.Msg()}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
