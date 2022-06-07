package http

import (
	"fmt"
	"github.com/eugene-krivtsov/idler/docs"
	"github.com/eugene-krivtsov/idler/internal/config"
	"github.com/eugene-krivtsov/idler/internal/service"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

type Handler struct {
	userService   service.Users
	tokenManager  auth.TokenManager
	dialogService service.Conversations
}

func NewHandler(userService service.Users, tokenManager auth.TokenManager, dialogService service.Conversations) *Handler {
	return &Handler{
		userService:   userService,
		tokenManager:  tokenManager,
		dialogService: dialogService,
	}
}

func (h *Handler) Init(cfg config.HTTPConfig) *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.Default(),
		//func(c *gin.Context) {
		//	c.Header("Access-Control-Allow-Origin", "*")
		//	c.Header("Access-Control-Allow-Credentials", "true")
		//	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		//
		//	if c.Request.Method == "OPTIONS" {
		//		c.AbortWithStatus(204)
		//		return
		//	}
		//
		//	c.Next()
		//},
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	h.initUserRoutes(router)
	h.initConversationRoutes(router)

	return router
}
