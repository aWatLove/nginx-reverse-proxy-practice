package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) hello(c *gin.Context) {
	//name := c.Param("name")
	//hello := getHelloString()
	hello := "hello world!"
	c.JSON(http.StatusOK, hello)
}

func getHelloString(name string) string {
	return fmt.Sprintf("hello %s", name)
}
