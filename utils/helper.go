package utils

import (
	"github.com/MicBun/go-file-manager/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// ResetUserDatabase godoc
// @Summary Reset user database
// @Description Reset user database
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /resetUserDatabase [post]
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
