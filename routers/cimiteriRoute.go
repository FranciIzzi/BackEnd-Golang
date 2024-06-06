package routes

import (
	controller "root/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CimiteriRoute(db *gorm.DB,router *gin.Engine) {
	userGroup := router.Group("server/cimiteri/api/method")
  userGroup.GET("/get", controller.GetCimiteri(db))
	userGroup.POST("/post", controller.CreateCimitero(db))
	userGroup.DELETE("/delete/:id", controller.DeleteCimitero(db))
	userGroup.PUT("/put/:id", controller.UpdateCimitero(db))
}
