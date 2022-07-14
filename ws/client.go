// Copyright (c) Mainflux
// SPDX-Licence-Identifier: Apache-2.0

package ws

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mainflux/mainflux/pkg/messaging"
)

// Client wraps WS client.
type Client struct {
	pubID    string
	token    string
	chanID   string
	subtopic string
	conn     *websocket.Conn
}

func NewClient(id, token, chanID, subtopic string, conn *websocket.Conn) Client {
	return Client{
		pubID:    id,
		token:    token,
		chanID:   chanID,
		subtopic: subtopic,
		conn:     conn,
	}
}

func (c Client) Publish(channelID, subtopic string, msgs chan<- messaging.Message) {
	for {
		_, payload, err := c.conn.ReadMessage()
		if websocket.IsUnexpectedCloseError(err) {
			c.conn.Close()
			return
		}
		if err != nil {
			return
		}

		msg := messaging.Message{
			Protocol: "ws",
			Channel:  channelID,
			Subtopic: subtopic,
			Payload:  payload,
			Created:  time.Now().UnixNano(),
		}
		fmt.Println("sending ws")
		msgs <- msg
	}
}

func (c Client) Handle(msg messaging.Message) error {
	format := websocket.TextMessage
	if msg.Publisher == c.pubID {
		return nil
	}

	if err := c.conn.WriteMessage(format, msg.Payload); err != nil {
		// logger.Warn(fmt.Sprintf("Failed to broadcast message to thing: %s", err))
		return err
	}
	return nil
}

func (c Client) Cancel() error {
	return c.conn.Close()
}
