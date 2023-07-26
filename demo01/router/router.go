package router

import (
	"project01/middlerware"
	"project01/router/auth"
	"project01/router/hello"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	// 注册全局拦截器
	r.Use(middlerware.JWTAuthMiddleware())
	// 注册路由组
	authGroup := r.Group("/api/auth")
	auth.RegisterAuthGroup(authGroup)

	helloGroup := r.Group("/api/hello")
	hello.RegisterHelloGroup(helloGroup)
}
