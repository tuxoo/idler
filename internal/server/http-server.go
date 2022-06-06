package server

import (
	"context"
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/config"
	"net/http"
)

type HTTPServer struct {
	httpServer *http.Server
}

func NewHTTPServer(cfg *config.Config, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		httpServer: &http.Server{
			Addr:           fmt.Sprintf(":%s", cfg.HTTP.Port),
			Handler:        handler,
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 28,
		},
	}
}

func (s *HTTPServer) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
