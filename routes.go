package main

import (
	"github.com/gin-gonic/gin"
	"slw.app/ginessential/controller"
	"slw.app/ginessential/middleware"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r
}
