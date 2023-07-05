package main

import (
	"os"
	"github.com/jacker1342/gpio_project/controller"
	"github.com/jacker1342/gpio_project/entity"
	"github.com/jacker1342/gpio_project/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Remove("./GPIO.db")
	// entity.MGDatabase()
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// User Routes
			protected.GET("/users", controller.ListUsers)
			protected.GET("/user/:id", controller.GetUser)
			protected.PATCH("/users", controller.UpdateUser)
			protected.DELETE("/users/:id", controller.DeleteUser)


			
			
		}
	}

	// users Routes
	r.POST("/users", controller.CreateUser)
	// r.PATCH("/users", controller.UpdateUser)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS"{
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}