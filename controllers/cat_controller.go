package controllers

import (
	"cats/models"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateBreed(breed string) bool {
	client := resty.New()
	resp, err := client.R().Get("https://api.thecatapi.com/v1/breeds")

	if err != nil {
		return false
	}

	var breeds []map[string]interface{}
	err = json.Unmarshal(resp.Body(), &breeds)
	if err != nil {
		return false
	}

	for _, b := range breeds {
		if b["name"] == breed {
			return true
		}
	}

	return false
}

func CreateCat(c *gin.Context) {
	var input models.CreateCatInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !ValidateBreed(input.Breed) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breed"})
		return
	}

	cat := models.Cat{Name: input.Name, YearsOfExperience: input.YearsOfExperience, Breed: input.Breed, Salary: input.Salary}
	models.DB.Create(&cat)

	c.JSON(http.StatusOK, gin.H{"data": cat})
}

func FindCats(c *gin.Context) {
	var cats []models.Cat
	models.DB.Find(&cats)

	c.JSON(http.StatusOK, gin.H{"data": cats})
}

func FindCat(c *gin.Context) {
	var cat models.Cat

	if err := models.DB.Where("id = ?", c.Param("id")).First(&cat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cat})
}

func UpdateCat(c *gin.Context) {
	var cat models.Cat
	if err := models.DB.Where("id = ?", c.Param("id")).First(&cat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateCatInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&cat).Update("salary", input.Salary)
	c.JSON(http.StatusOK, gin.H{"data": cat})
}

func DeleteCat(c *gin.Context) {
	var cat models.Cat
	if err := models.DB.Where("id = ?", c.Param("id")).First(&cat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&cat)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
