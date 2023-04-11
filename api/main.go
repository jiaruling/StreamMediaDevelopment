package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/jiaruling/StreamMediaDevelopment/api/handlers"
	"github.com/jiaruling/StreamMediaDevelopment/api/middleware"
	"github.com/jiaruling/StreamMediaDevelopment/api/session"
)

type middlewareHandler struct {
	r *httprouter.Router
}

func NewMiddlewareHandler(r *httprouter.Router) http.Handler {
	return middlewareHandler{r: r}
}

func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middleware.ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", handlers.HealthCheck)
	router.POST("/user", handlers.CreateUser)                             // 创建用户
	router.POST("/user/:username", handlers.Login)                        // 用户登录
	router.GET("/user/:username", handlers.GetUserInfo)                   // 获取用户信息
	router.POST("/user/:username/videos", handlers.AddNewVideo)           // 新增视频资源
	router.GET("/user/:username/videos", handlers.ListAllVideos)          // 用户资源列表
	router.DELETE("/user/:username/videos/:vid-id", handlers.DeleteVideo) // 删除用户资源
	router.POST("/videos/:vid-id/comments", handlers.PostComment)         // 新增评论
	router.GET("/videos/:vid-id/comments", handlers.ShowComments)         // 获取评列表
	return router
}

func main() {
	session.LoadSessionFromDB()
	router := RegisterHandlers()
	mh := NewMiddlewareHandler(router)
	log.Fatal(http.ListenAndServe(":8001", mh))
}
