// Copyright (c) Mainflux
// SPDX-Licence-Identifier: Apache-2.0

// Package ws contains the domain concept definitions needed to support
// Mainflux ws adapter service functionality

package ws

import (
	"context"
	"fmt"
	"time"

	"github.com/mainflux/mainflux"
	"github.com/mainflux/mainflux/pkg/errors"
	"github.com/mainflux/mainflux/pkg/messaging"
)

var (
	// ErrFailedMessagePublish indicates that message publishing failed.
	ErrFailedMessagePublish = errors.New("failed to publish message")

	// ErrFailedSubscription indicates that client couldn't subscriber to specified channel
	ErrFailedSubscription = errors.New("failed to subscribe to a channel")

	// ErrFailedConnection indicates that service couldn't connect to message broker.
	ErrFailedConnection = errors.New("failed to connect to message broker")
)

// Service specifies web socket service API.
type Service interface {
	// Publish Message
	Publish(ctx context.Context, token, channel, subtopic string, payload []byte) error

	// Subscribes to a channel with specified id.
	Subscribe(ctx context.Context, chanID, subtopic string, c Client) error

	// Unsubscribe method is used to stop observing resource.
	Unsubscribe(ctx context.Context, thingKey, chanID, subtopic string) error
}

var _ Service = (*adapterService)(nil)

type adapterService struct {
	auth   mainflux.ThingsServiceClient
	pubsub messaging.PubSub
}

// New instantiates the WS adapter implementation
func New(auth mainflux.ThingsServiceClient, pubsub messaging.PubSub) Service {
	return &adapterService{
		auth:   auth,
		pubsub: pubsub,
	}
}

func (svc *adapterService) Publish(ctx context.Context, token, channel, subtopic string, payload []byte) error {
	ar := &mainflux.AccessByKeyReq{
		Token:  token,
		ChanID: channel,
	}

	id, err := svc.auth.CanAccessByKey(ctx, ar)
	if err != nil {
		return errors.Wrap(errors.ErrAuthorization, err)
	}

	msg := messaging.Message{
		Protocol:  "ws",
		Channel:   channel,
		Subtopic:  subtopic,
		Payload:   payload,
		Created:   time.Now().UnixNano(),
		Publisher: id.GetValue(),
	}

	return svc.pubsub.Publish(channel, msg)
}

func (svc *adapterService) Subscribe(ctx context.Context, chanID, subtopic string, c Client) error {
	ar := &mainflux.AccessByKeyReq{
		Token:  c.token,
		ChanID: c.chanID,
	}

	id, err := svc.auth.CanAccessByKey(ctx, ar)
	if err != nil {
		return errors.Wrap(errors.ErrAuthorization, err)
	}

	subject := fmt.Sprintf("%s.%s", "channels", chanID)
	if subtopic != "" {
		subject = fmt.Sprintf("%s.%s", subject, subtopic)
	}
	handler := msgHandler{thingID: id.GetValue(), c: c}

	return svc.pubsub.Subscribe(id.GetValue(), subject, handler)
}

type msgHandler struct {
	thingID string
	c       Client
}

func (h msgHandler) Handle(msg messaging.Message) error {
	if h.thingID == msg.Publisher {
		return nil
	}
	return h.c.Handle(msg.Payload)
}

func (h msgHandler) Cancel() error {
	return h.c.Cancel()
}

func (svc *adapterService) Unsubscribe(ctx context.Context, thingKey, chanID, subtopic string) error {
	ar := &mainflux.AccessByKeyReq{
		Token:  thingKey,
		ChanID: chanID,
	}

	if _, err := svc.auth.CanAccessByKey(ctx, ar); err != nil {
		return errors.Wrap(errors.ErrAuthorization, err)
	}

	subject := fmt.Sprintf("%s.%s", "channels", chanID)
	if subtopic != "" {
		subject = fmt.Sprintf("%s.%s", subject, subtopic)
	}

	return svc.pubsub.Unsubscribe(thingKey, subject)
}
