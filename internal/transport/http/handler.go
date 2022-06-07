package http

import (
	"fmt"
	"github.com/eugene-krivtsov/idler/docs"
	"github.com/eugene-krivtsov/idler/internal/config"
	"github.com/eugene-krivtsov/idler/internal/service"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

type Handler struct {
	userService         service.Users
	tokenManager        auth.TokenManager
	conversationService service.Conversations
}

func NewHandler(userService service.Users, tokenManager auth.TokenManager, dialogService service.Conversations) *Handler {
	return &Handler{
		userService:         userService,
		tokenManager:        tokenManager,
		conversationService: dialogService,
	}
}

func (h *Handler) Init(cfg config.HTTPConfig) *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.Default(),
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("api/v1/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		h.initUserRoutes(api)
		h.initConversationRoutes(api)
	}
}
