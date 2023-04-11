package main

import (
	"net/http"

	"github.com/jiaruling/StreamMediaDevelopment/web/handlers"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", handlers.HomeHandler)

	router.POST("/", handlers.HomeHandler)

	router.GET("/userhome", handlers.UserHomeHandler)

	router.POST("/userhome", handlers.UserHomeHandler)
	// api 透传
	router.POST("/api", handlers.ApiHandler)
	// proxy 转发
	router.GET("/videos/:vid-id", handlers.ProxyVideoHandler)
	// proxy 转发
	router.POST("/upload/:vid-id", handlers.ProxyUploadHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./template"))

	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8000", r)
}
