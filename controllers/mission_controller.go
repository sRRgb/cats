package controllers

import (
	"cats/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMission(c *gin.Context) {
	var input models.CreateMissionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	mission := models.Mission{CatID: input.CatID, Complete: input.Complete}
	models.DB.Create(&mission)

	for _, t := range input.Targets {
		target := models.Target{Name: t.Name, Country: t.Country, Notes: t.Notes, Complete: t.Complete, MissionID: mission.ID}
		models.DB.Create(&target)
	}

	models.DB.Preload("Targets").First(&mission)
	c.JSON(http.StatusOK, gin.H{"data": mission})
}

func FindMissions(c *gin.Context) {
	var missions []models.Mission
	models.DB.Preload("Targets").Find(&missions)

	c.JSON(http.StatusOK, gin.H{"data": missions})
}

func FindMission(c *gin.Context) {
	var mission models.Mission
	if err := models.DB.Preload("Targets").Where("id = ?", c.Param("id")).First(&mission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mission})
}

func UpdateMission(c *gin.Context) {
	var mission models.Mission
	if err := models.DB.Where("id = ?", c.Param("id")).First(&mission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.CreateMissionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	models.DB.Model(&mission).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": mission})
}

func DeleteMission(c *gin.Context) {
	var mission models.Mission
	if err := models.DB.Where("id = ?", c.Param("id")).First(&mission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})

		return
	}

	if mission.CatID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mission is assigned to a cat"})

		return
	}

	models.DB.Delete(&mission)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
