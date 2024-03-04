package controllers

import (
	"net/http"
	"restapi/database"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

// CreateUser adalah handler untuk membuat register users
func CreateUser(c *gin.Context) {
	var users models.UserModel

	// Bind data JSON yang diterima dari permintaan ke struct UserModel
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan data users ke dalam basis data
	if err := database.DB.Create(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat users"})
		return
	}

	// Berhasil membuat users, kirimkan respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dibuat", "data": users})
}

func Login(c *gin.Context) {
	var users models.UserModel
	c.BindJSON(&users)
	var userDB models.UserModel
	database.DB.Where("email = ?", users.Email).First(&userDB)
	if userDB.Email == users.Email && userDB.Password == users.Password {
		c.JSON(http.StatusOK, gin.H{"data": userDB})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"data": "User not found"})
	}
}

func UpdateUser(c *gin.Context) {
	var users models.UserModel
	c.BindJSON(&users)
	database.DB.Save(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func DeleteUser(c *gin.Context) {
	var users models.UserModel
	id := c.Params.ByName("User_id")
	database.DB.Where("id = ?", id).Delete(&users)
	c.JSON(http.StatusOK, gin.H{"data": "User deleted"})
}
