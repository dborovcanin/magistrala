// Copyright (c) Mainflux
// SPDX-Licence-Identifier: Apache-2.0

package ws

import (
	"github.com/gorilla/websocket"
)

// Client wraps WS client.
type Client struct {
	token    string
	chanID   string
	subtopic string
	conn     *websocket.Conn
}

func NewClient(token, chanID, subtopic string, conn *websocket.Conn) Client {
	return Client{
		token:    token,
		chanID:   chanID,
		subtopic: subtopic,
		conn:     conn,
	}
}

func (c Client) Publish(channelID, subtopic string, msgs chan<- []byte) {
	for {
		_, payload, err := c.conn.ReadMessage()
		if websocket.IsUnexpectedCloseError(err) {
			c.conn.Close()
			return
		}
		if err != nil {
			return
		}

		msgs <- payload
	}
}

func (c Client) Handle(payload []byte) error {
	if err := c.conn.WriteMessage(websocket.TextMessage, payload); err != nil {
		// logger.Warn(fmt.Sprintf("Failed to broadcast message to thing: %s", err))
		return err
	}
	return nil
}

func (c Client) Cancel() error {
	return c.conn.Close()
}
