package utils

import (
	"StudentCurriculum_backEnd_Gin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func UserDataAdd(jsoninfo model.User) model.User {
	jsoninfo.UserPassword, _ = HashPassword(jsoninfo.UserPassword)
	jsoninfo.UserCreateOfTime = time.Now()
	return jsoninfo
}

func SearchUserIsExist(jsoninfo model.User, db *gorm.DB) bool {
	result := db.Where("user_job_id =?", jsoninfo.UserJobId).First(&jsoninfo)
	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func UserDbAdd(db *gorm.DB, ctx *gin.Context) {
	var jsoninfo model.UserRoleString
	if err := ctx.ShouldBindJSON(&jsoninfo); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "error",
		})
		return
	}
	fmt.Println(jsoninfo)
	var user model.User
	user.UserId = jsoninfo.UserId
	user.UserJobId = jsoninfo.UserJobId
	user.UserPassword, _ = HashPassword(jsoninfo.UserPassword)
	user.UserCreateOfTime = jsoninfo.UserCreateOfTime
	user.UserRoleId = StringToInt64(jsoninfo.UserRoleId)
	user.UserAvatar = jsoninfo.UserAvatar

	if user.UserRoleId == 1 {
		fmt.Println("1")
		var dean model.Dean
		result := db.Where("dean_job_id = ?", user.UserJobId).First(&dean)
		if result.RowsAffected != 0 {
			if SearchUserIsExist(user, db) != false {
				UserDataAdd(user)
				result := db.Create(&user)
				fmt.Println(result)
				ctx.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "注册成功",
					"data": nil,
				})
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"code": 201,
					"msg":  "该用户已注册",
					"data": nil,
				})
			}
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 201,
				"msg":  "未找到该用户信息",
				"data": nil,
			})
		}
	} else if user.UserRoleId == 2 {
		fmt.Println("2")
		var Teacher model.Teacher
		result := db.Where("teacher_job_id = ?", user.UserJobId).First(&Teacher)
		if result.RowsAffected != 0 {
			if SearchUserIsExist(user, db) != false {
				UserDataAdd(user)
				result := db.Create(&user)
				fmt.Println(result)
				ctx.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "注册成功",
					"data": nil,
				})
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"code": 201,
					"msg":  "该用户已注册",
					"data": nil,
				})
			}
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 201,
				"msg":  "未找到该用户信息",
				"data": nil,
			})
		}
	} else if user.UserRoleId == 3 {
		fmt.Println("3")
		var Student model.Student
		result := db.Where("student_job_id = ?", user.UserJobId).First(&Student)
		if result.RowsAffected != 0 {
			if SearchUserIsExist(user, db) != false {
				UserDataAdd(user)
				result := db.Create(&user)
				fmt.Println(result)
				ctx.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "注册成功",
					"data": nil,
				})
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"code": 201,
					"msg":  "该用户已注册",
					"data": nil,
				})
			}
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 201,
				"msg":  "未找到该用户信息",
				"data": nil,
			})
		}
	} else if user.UserRoleId == 4 {
		fmt.Println("4")
		var Counselor model.Counselor
		result := db.Where("counselor_job_id = ?", user.UserJobId).First(&Counselor)
		if result.RowsAffected != 0 {
			if SearchUserIsExist(user, db) != false {
				UserDataAdd(user)
				result := db.Create(&user)
				fmt.Println(result)
				ctx.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "注册成功",
					"data": nil,
				})
			} else {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"code": 201,
					"msg":  "该用户已注册",
					"data": nil,
				})
			}
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 201,
				"msg":  "未找到该用户信息",
				"data": nil,
			})
		}
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "未找到该用户信息",
			"data": nil,
		})
	}

}
