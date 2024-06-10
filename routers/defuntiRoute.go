package routes

import (
	controller "root/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DefuntiRoute(db *gorm.DB, router *gin.Engine) {
	userGroup := router.Group("server/defunti/api/method")
	userGroup.GET("/get", controller.GetDefunti(db))
	userGroup.POST("/post", controller.CreateDefunto(db))
	userGroup.DELETE("/delete/:id", controller.DeleteDefunto(db))
	userGroup.PUT("/put/:id", controller.UpdateDefunto(db))
}
