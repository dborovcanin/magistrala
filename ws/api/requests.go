package api

import (
	"time"

	"github.com/mainflux/mainflux/internal/apiutil"
)

type publishReq struct {
	protocol string
	channel  string
	subtopic string
	payload  []byte
	created  time.Time
	token    string
}

func (req publishReq) validate() error {
	if req.token == "" {
		return apiutil.ErrBearerToken
	}

	return nil
}
