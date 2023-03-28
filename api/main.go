package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/jiaruling/StreamMediaDevelopment/api/handlers"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handlers.CreateUser)      // 创建用户
	router.POST("/user/:username", handlers.Login) // 用户登录
	// router.GET("/user/:username/videos", handlers.)             // 用户资源列表
	// router.GET("/user/:username/videos/:vid-id", handlers.)     // 单个用户资源
	// router.DDELETE("/user/:username/videos/:vid-id", handlers.) // 删除用户资源
	// router.POST("/user/:username/videos/:vid-id", handlers.) // 删除用户资源
	// router.GET("/videos/:vid-id", handlers.) // 删除用户资源
	// router.DDELETE("/videos/:vid-id", handlers.) // 删除用户资源
	return router
}

func main() {
	router := RegisterHandlers()
	log.Fatal(http.ListenAndServe(":8000", router))
}
