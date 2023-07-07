package controller

import (
	"StudentCurriculum_backEnd_Gin/model"
	"StudentCurriculum_backEnd_Gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func UserLogin(ctx *gin.Context) {
	fmt.Println("User_login is running")
	var jsoninfo model.User
	if err := ctx.ShouldBindJSON(&jsoninfo); err != nil {
		fmt.Printf("%s and %s\n", jsoninfo.UserJobId, jsoninfo.UserPassword)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "error",
		})
		return
	}
	var user model.UserAndRole
	result := db.Table("user").Select("*").Joins("left join role on role.role_id = user_role_id ").Where("user_job_id=?", jsoninfo.UserJobId).Scan(&user)
	fmt.Println(user)
	if result.Error != nil {
		if result.RowsAffected != 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg":  result.Error,
				"code": 502,
			})
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg":  "登录失败，找不到该账号",
				"code": 502,
			})
		}
	} else {

		bcrypt_password := jsoninfo.UserPassword
		err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(bcrypt_password))
		if err == nil {
			user.UserPassword = jsoninfo.UserPassword
			var usertoken model.UserToken
			usertoken.No = strconv.Itoa(int(user.UserId))
			usertoken.UserJobId = user.UserJobId
			usertoken.Password = user.UserPassword
			usertoken.UserRoleId = user.UserRoleId
			token := utils.GenerateToken(&usertoken)
			usertoken.Token = token
			flag := utils.JsonSetRedis(usertoken, usertoken.UserJobId)
			if flag {
				ctx.JSON(http.StatusOK, gin.H{
					"code":        200,
					"msg":         "登录成功！",
					"avatar":      user.UserAvatar,
					"authorities": user.RoleName,
					"token":       token,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code":        201,
					"msg":         "token存入失败！",
					"avatar":      user.UserAvatar,
					"authorities": user.RoleName,
					"token":       nil,
				})
			}
		} else {
			fmt.Println(err)
		}
	}
}

func UserRegister(ctx *gin.Context) {
	utils.UserDbAdd(db, ctx)
}
