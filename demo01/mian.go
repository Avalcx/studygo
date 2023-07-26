package main

import (
	_ "project01/config"
	"project01/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.RegisterRouter(r)
	r.Run()
}
