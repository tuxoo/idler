package server

import (
	"context"
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/config"
	"github.com/eugene-krivtsov/idler/internal/transport/ws"
	"github.com/gorilla/websocket"
	"net/http"
)

type WSServer struct {
	upgrader *websocket.Upgrader

	wsHandler *http.Handler
	wsServer  *http.Server
	hub       *ws.Hub
}

func NewWSServer(cfg *config.Config, wsHandler http.Handler, hub *ws.Hub) *WSServer {
	wsServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.WS.Port),
		Handler: wsHandler,
	}

	return &WSServer{
		wsServer: wsServer,
		hub:      hub,
	}
}

func (s *WSServer) Run() error {
	return s.wsServer.ListenAndServe()
}

func (s *WSServer) Shutdown(ctx context.Context) error {
	return s.wsServer.Shutdown(ctx)
}
