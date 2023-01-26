package controllers

import (
	"github.com/MicBun/go-file-manager/models"
	jwtauth "github.com/MicBun/go-file-manager/utils/jwtAuth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ResetUserDatabase(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user := models.NewUser()
	err := user.ResetUserDatabase(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Reset user database successfully",
	})
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	userInterface := models.NewUser()
	err = userInterface.Login(db, &user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, _ := jwtauth.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
