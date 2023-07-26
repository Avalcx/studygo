package auth

import (
	"project01/controller/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthGroup(g *gin.RouterGroup) {
	g.POST("/login", auth.Login)
	g.GET("/logout", auth.Logout)
}
