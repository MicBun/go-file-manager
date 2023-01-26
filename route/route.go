package route

import (
	"github.com/MicBun/go-file-manager/controllers"
	"github.com/MicBun/go-file-manager/middleware"
	"github.com/MicBun/go-file-manager/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/login", controllers.Login)
	r.POST("/resetUserDatabase", utils.ResetUserDatabase)
	r.GET("/download", controllers.DownloadFile)

	fileRoutes := r.Group("/file")
	fileRoutes.Use(middleware.JwtAuthMiddleware())
	fileRoutes.GET("/list", controllers.ListFile)
	fileRoutes.POST("/upload", controllers.UploadFile)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
