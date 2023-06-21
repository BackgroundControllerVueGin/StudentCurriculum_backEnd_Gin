package routes

import (
	"StudentCurriculum_backEnd_Gin/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(gin_server *gin.Engine) *gin.Engine {
	gin_server.GET("/test", controller.TestApi)
	return gin_server
}
