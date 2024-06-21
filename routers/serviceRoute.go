package routes

import (
	controller "root/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ServiceRoute(db *gorm.DB, router *gin.Engine) {
	userGroup := router.Group("server/api/service/method")
	userGroup.GET("/ricerca/get", controller.GetRicerca(db))
	userGroup.GET("/file/get", controller.GetFile(db))
}
