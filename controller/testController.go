package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestApi(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "测试成功",
	})
}
