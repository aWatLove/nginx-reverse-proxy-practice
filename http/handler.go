package http

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/api/hello", h.hello)

	return router
}
