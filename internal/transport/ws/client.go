package ws

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/service"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"time"
)

type Client struct {
	hub            *Hub
	conn           *websocket.Conn
	send           chan entity.Message
	messageService service.Messages
}

func NewClient(conn *websocket.Conn, hub *Hub, messageService service.Messages) *Client {
	client := &Client{
		hub:            hub,
		conn:           conn,
		send:           make(chan entity.Message),
		messageService: messageService,
	}
	client.hub.register <- client

	return client
}

func (c *Client) HandleMessage() {
	defer func() {
		c.hub.unregister <- c
		if err := c.conn.Close(); err != nil {
			logrus.Errorf("error occured on web socket client close: %s", err.Error())
			return
		}
	}()

	for {
		_, p, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("unexpected error occured on web socket client close: %s", err.Error())
				return
			}
		}

		message := entity.Message{
			Sender: "a",
			SentAt: time.Now(),
			Text:   string(p),
		}

		c.hub.broadcast <- message

		if err := c.messageService.Save(context.Background(), message); err != nil {
			logrus.Errorf("error occured on web socket sending message: %s", err.Error())
			return
		}
	}
}
