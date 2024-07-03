package main

import (
	"cats/models"
	"cats/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func main() {
	models.ConnectDatabase()

	r := gin.New()

	//Logger middleware
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	r.Use(ginlogrus.Logger(logger), gin.Recovery())

	routes.SetupRouter(r)
	r.Run()
}
