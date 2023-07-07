package controller

import (
	"StudentCurriculum_backEnd_Gin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetConselorData(ctx *gin.Context) {
	fmt.Println("GetConselorData is running")
	json := make(map[string]interface{})
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "json传输异常",
			"data": nil,
		})
	}
	studentJobId := json["studentJobId"].(string)
	var counselor []model.Counselor
	result := db.Joins("JOIN class ON counselor.counselor_job_id = class.class_counselor_job_id").
		Joins("JOIN student ON student.student_class_id = class.class_id").
		Where("student_job_id = ?", studentJobId).First(&counselor)
	if result.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "获取信息失败",
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": gin.H{
				"studentList": counselor,
			},
		})
	}

}

func SubmitApplication(ctx *gin.Context) {
	var jsoninfo model.Application
	if err := ctx.ShouldBindJSON(&jsoninfo); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "error",
		})
		return
	}
	result_add := db.Create(&jsoninfo)
	if result_add.Error == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 200,
			"msg":  "提交申请成功",
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "提交申请失败",
			"data": nil,
		})
	}
}
