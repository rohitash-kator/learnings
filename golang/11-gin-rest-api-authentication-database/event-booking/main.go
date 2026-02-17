package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsrkator/event-booking/db"
	"github.com/itsrkator/event-booking/routes"
)

func main() {
	app := gin.Default()
	db.InitDb()

	routes.RegisterRoutes(app)
	app.Run(":3000") // localhost:3000
}
