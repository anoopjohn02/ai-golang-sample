package web

import (
	"context"
	"log"
	"net/http"

	"github.com/anoopjohn02/ai-golang-sample/internal/commons"
	"github.com/anoopjohn02/ai-golang-sample/internal/service"
	"github.com/gin-gonic/gin"
)

type RestService struct {
	engine *gin.Engine
	server *http.Server
	stop   chan string
}

func (r *RestService) Start() {
	go func() {
		if err := r.server.ListenAndServe(); err != nil {
			log.Printf("Error with http server: %v", err)
		}
	}()
	<-r.stop
}

func (r *RestService) Stop() {
	var stopMsg = "stop"
	r.stop <- stopMsg
	err := r.server.Shutdown(context.Background())
	if err != nil {
		log.Printf("Error with http server: %v", err)
	}
}

func NewService(ctx *commons.AIContext, doc *service.DocumentService) *RestService {
	engine := gin.Default()
	setupRoutes(engine, ctx, doc)

	return &RestService{
		engine: engine,
		server: &http.Server{Addr: ":" + "8080", Handler: engine},
		stop:   make(chan string),
	}
}

func setupRoutes(engine *gin.Engine, ctx *commons.AIContext, doc *service.DocumentService) {
	engine.GET("/api/ping", ping)
	engine.Static("/ui", "./static")
	engine.StaticFile("/", "./static/index.html")
	v1 := engine.Group("/v1")
	{
		chats := v1.Group("/chat")
		{
			chatController := NewChatController(ctx, doc)
			chats.POST("/stream", chatController.StreamChat)
		}
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}
