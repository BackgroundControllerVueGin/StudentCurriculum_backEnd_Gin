package routes

import (
	"StudentCurriculum_backEnd_Gin/model"
	"StudentCurriculum_backEnd_Gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func studentMiddleware(ctx *gin.Context) {
	// 在这里实现学生权限过滤器的逻辑
	fmt.Println("studentMiddleware is running")
	token := ctx.Query("token")
	if token == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 404,
			"msg":  "缺少传入",
		})
	}
	var userToken model.UserToken
	var boolFlag bool
	userToken, boolFlag = utils.JwtVerify(token)
	if boolFlag {
		if userToken.UserRoleId == 3 {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(401, gin.H{
				"code": 201,
				"msg":  "权限验证失败"})
		}
	} else {
		ctx.AbortWithStatusJSON(401, gin.H{
			"code": 204,
			"msg":  "权限验证失败"})
	}
}

func teacherMiddleware(ctx *gin.Context) {
	fmt.Println("teacherMiddleware is running")
	// 在这里实现学生权限过滤器的逻辑
	token := ctx.Query("token")
	if token == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 404,
			"msg":  "缺少传入",
		})
	}
	var userToken model.UserToken
	var boolFlag bool
	userToken, boolFlag = utils.JwtVerify(token)
	if boolFlag {
		if userToken.UserRoleId == 2 {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(401, gin.H{
				"code": 201,
				"msg":  "权限验证失败"})
		}
	} else {
		ctx.AbortWithStatusJSON(401, gin.H{
			"code": 204,
			"msg":  "权限验证失败"})
	}
}

func counselorMiddleware(ctx *gin.Context) {
	// 在这里实现学生权限过滤器的逻辑
	fmt.Println("studentMiddleware is running")
	token := ctx.Query("token")
	if token == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 404,
			"msg":  "缺少传入",
		})
	}
	var userToken model.UserToken
	var boolFlag bool
	userToken, boolFlag = utils.JwtVerify(token)
	if boolFlag {
		if userToken.UserRoleId == 4 {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(401, gin.H{
				"code": 201,
				"msg":  "权限验证失败"})
		}
	} else {
		ctx.AbortWithStatusJSON(401, gin.H{
			"code": 204,
			"msg":  "权限验证失败"})
	}
}
