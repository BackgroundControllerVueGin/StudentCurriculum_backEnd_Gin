package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestApi(ctx *gin.Context) {
	fmt.Println(db_err)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "测试成功",
	})
}
