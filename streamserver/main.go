package main

import (
	"log"
	"net/http"

	"github.com/jiaruling/StreamMediaDevelopment/streamserver/handlers"
	"github.com/jiaruling/StreamMediaDevelopment/streamserver/limiter"
	"github.com/jiaruling/StreamMediaDevelopment/streamserver/response"
	"github.com/julienschmidt/httprouter"
)

type middlewareHandler struct {
	r *httprouter.Router
	l *limiter.ConnLimiter
}

func NewMiddlewareHandler(r *httprouter.Router, cc int) http.Handler {
	return middlewareHandler{r: r, l: limiter.NewConnLimiter(cc)}
}

func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		response.SendErrorResponse(w, http.StatusTooManyRequests, "too many requests")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", handlers.HealthCheck)
	router.GET("/videos/:vid-id", handlers.StreamHandler)
	router.POST("/upload/:vid-id", handlers.UploadHandler)
	return router
}

func main() {
	router := RegisterHandlers()
	mh := NewMiddlewareHandler(router, 2)
	log.Fatal(http.ListenAndServe(":8003", mh))
}
