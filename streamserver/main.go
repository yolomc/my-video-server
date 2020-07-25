package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid", streamHandler)
	router.POST("/upload/:vid", uploadHandler)
	router.GET("/testpage", testPageHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8081", NewMiddlerWareHandler(r, 2))
}
