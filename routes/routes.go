package routes

import (
	"cats/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	r.POST("/cats", controllers.CreateCat)
	r.GET("/cats", controllers.FindCats)
	r.GET("/cats/:id", controllers.FindCat)
	r.PUT("/cats/:id", controllers.UpdateCat)
	r.DELETE("/cats/:id", controllers.DeleteCat)

	r.POST("/missions", controllers.CreateMission)
	r.GET("/missions", controllers.FindMissions)
	r.GET("/missions/:id", controllers.FindMission)
	r.PUT("/missions/:id", controllers.UpdateMission)
	r.DELETE("/missions/:id", controllers.DeleteMission)

	return r
}
