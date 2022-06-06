package ws

import (
	"github.com/eugene-krivtsov/idler/internal/config"
	"github.com/eugene-krivtsov/idler/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Handler struct {
	Hub            *Hub
	Upgrader       *websocket.Upgrader
	MessageService service.Messages
}

func NewHandler(cfg config.WSConfig, hub *Hub, messageService service.Messages) *Handler {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  cfg.ReadBufferSize,
		WriteBufferSize: cfg.ReadBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return &Handler{
		Hub:            hub,
		Upgrader:       upgrader,
		MessageService: messageService,
	}
}

func (h *Handler) Init() http.Handler {
	handler := gin.New()

	handler.GET("/conversation", func(c *gin.Context) {
		conn, err := h.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		client := NewClient(conn, h.Hub, h.MessageService)
		client.HandleMessage()
	})

	return handler
}
