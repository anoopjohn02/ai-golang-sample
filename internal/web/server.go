package web

import (
	"context"
	"log"
	"net/http"

	"github.com/anoopjohn02/ai-golang-sample/internal/commons"
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

func NewService(deviceContext *commons.DeviceContext) *RestService {
	engine := gin.Default()
	setupRoutes(engine, deviceContext)

	return &RestService{
		engine: engine,
		server: &http.Server{Addr: ":" + "8080", Handler: engine},
		stop:   make(chan string),
	}
}

func setupRoutes(engine *gin.Engine, deviceContext *commons.DeviceContext) {
	engine.GET("/api/ping", ping)
	engine.Static("/static", "./static")
	engine.StaticFile("/", "./static/index.html")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}
