package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/itsrkator/event-booking/controllers"
)

func RegisterRoutes(app *gin.Engine) {
	app.GET("/events", controllers.GetEvents)
	app.POST("/events", controllers.CreateEvent)
	app.GET("/events/:id", controllers.GetEvent)
	app.PUT("/events/:id", controllers.UpdateEvent)
	app.DELETE("/events/:id", controllers.DeleteEvent)

	app.POST("/signup", controllers.Signup)
	app.POST("/login", controllers.Login)
}
