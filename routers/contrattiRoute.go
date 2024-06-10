package routes

import (
	controller "root/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ContrattiRoute(db *gorm.DB,router *gin.Engine) {
	userGroup := router.Group("server/contratti/api/method")
  userGroup.GET("/get", controller.GetContratti(db))
	userGroup.POST("/post", controller.CreateContratti(db))
	userGroup.DELETE("/delete/:id", controller.DeleteContratti(db))
	userGroup.PUT("/put/:id", controller.UpdateContratti(db))
}
