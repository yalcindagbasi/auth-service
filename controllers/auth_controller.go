package controllers

import (
	"auth-service/database"
	"auth-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcı kaydedilemedi"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Kayıt başarılı"})
}

func Login(c *gin.Context) {
	var user models.User
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}

	result := database.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	if user.Password != input.Password { // TODO: Şifreyi hashleyip karşılaştırmalıyız!
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Şifre hatalı"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Giriş başarılı"})
}
