package server

import (
	"log/slog"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sam-chadnaldo/chadnaldo-website-backend/internal/app/server/middleware"
)

type Handler struct{
	log *slog.Logger
}

func NewHandler( log *slog.Logger ) *Handler {
	return &Handler{
		log: log,
	}
}

func (c *Handler) InitRouts() *gin.Engine{
	router := gin.New()
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	   }))

	router.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
			c.Header("Access-Control-Expose-Headers", "Content-Length")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Max-Age", "43200")
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// health checker
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	v1 := router.Group("/v1")
	protected := v1.Group("/")

	protected.Use(middleware.UserIdentity())


	return router
}