package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/itsrkator/event-booking/controllers"
	"github.com/itsrkator/event-booking/middlewares"
)

func RegisterRoutes(app *gin.Engine) {
	app.GET("/events", controllers.GetEvents)
	app.GET("/events/:id", controllers.GetEvent)

	protected := app.Group("/")
	protected.Use(middlewares.AuthMiddleware)

	protected.POST("/events", controllers.CreateEvent)
	protected.PUT("/events/:id", controllers.UpdateEvent)
	protected.DELETE("/events/:id", controllers.DeleteEvent)

	// protected.GET("/events/:id/")
	protected.POST("/events/:id/register", controllers.RegisterForEvent)
	protected.DELETE("/events/:id/register", controllers.CancelRegistration)

	app.POST("/signup", controllers.Signup)
	app.POST("/login", controllers.Login)
}
