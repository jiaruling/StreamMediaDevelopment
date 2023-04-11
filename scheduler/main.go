package main

import (
	"log"
	"net/http"

	"github.com/jiaruling/StreamMediaDevelopment/scheduler/handlers"
	"github.com/jiaruling/StreamMediaDevelopment/scheduler/taskrunner"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", handlers.HealthCheck)
	router.DELETE("/video-delete-record/:vid-id", handlers.VidDelRecHandler) // 创建用户
	return router
}

func main() {
	taskrunner.Start()
	router := RegisterHandlers()
	log.Fatal(http.ListenAndServe(":8002", router))
}
