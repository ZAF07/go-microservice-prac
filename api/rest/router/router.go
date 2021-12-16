package router

import (
	"github.com/ZAF07/go-microservice-prac/api/rest/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
	// Set CORS config
	r.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	}))

	// Initialise a new route group
	hello := r.Group("v1")
	{
		// Initialise and invoke controller
		helloAPI := new(controller.GreetAPI)
		hello.GET("/hi", helloAPI.SayHello)
	}
	return r
}