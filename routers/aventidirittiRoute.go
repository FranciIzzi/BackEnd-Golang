package routes

import (
	controller "root/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AventiDirittiRoute(db *gorm.DB, router *gin.Engine) {
	userGroup := router.Group("server/aventi-diritti/api/method")
	userGroup.GET("/get", controller.GetAventiDiritti(db))
	userGroup.POST("/post", controller.CreateAventiDiritti(db))
	userGroup.DELETE("/delete/:id", controller.DeleteAventiDiritti(db))
	userGroup.PUT("/put/:id", controller.UpdateAventiDiritti(db))
}
