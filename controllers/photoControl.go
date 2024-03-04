package controllers

import (
	"net/http"
	"restapi/database"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	var photos models.PhotoModel
	c.BindJSON(&photos)
	database.DB.Create(&photos)
	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func GetPhotos(c *gin.Context) {
	var photos []models.PhotoModel
	database.DB.Find(&photos)
	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func GetPhoto(c *gin.Context) {
	var photos models.PhotoModel
	id := c.Params.ByName("photo_id")
	database.DB.Where("id = ?", id).First(&photos)
	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func UpdatePhoto(c *gin.Context) {
	var photos models.PhotoModel
	c.BindJSON(&photos)
	database.DB.Save(&photos)
	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func DeletePhoto(c *gin.Context) {
	var photos models.PhotoModel
	id := c.Params.ByName("photo_id")
	database.DB.Where("id = ?", id).Delete(&photos)
	c.JSON(http.StatusOK, gin.H{"data": "Photo deleted"})
}
