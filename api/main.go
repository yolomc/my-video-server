package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yolomc/my-video-server/api/gredis"
	"github.com/yolomc/my-video-server/api/middleware"
)

func RegisrerHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:username", Login)

	return router
}

func main() {
	gredis.Setup()
	r := RegisrerHandlers()
	http.ListenAndServe(":8080", middleware.NewMiddleWareHandler(r))
}
