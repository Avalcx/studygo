package hello

import (
	"project01/controller/hello"

	"github.com/gin-gonic/gin"
)

func RegisterHelloGroup(g *gin.RouterGroup) {
	g.GET("/say", hello.Hello)
}
