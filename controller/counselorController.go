package controller

import (
	"StudentCurriculum_backEnd_Gin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateApplication(ctx *gin.Context) {
	var jsoninfo model.Application
	if err := ctx.ShouldBindJSON(&jsoninfo); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "error",
		})
		return
	}
	var application model.Application
	result := db.Model(&application).Where("application_id = ?", jsoninfo.ApplicationId).Updates(model.Application{ApplicationReply: jsoninfo.ApplicationReply, ApplicationStatus: jsoninfo.ApplicationStatus})
	if result.Error != nil || result.RowsAffected == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "提交申请失败",
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 200,
			"msg":  "审核完成",
			"data": nil,
		})
	}
}

func GetApplicationList(ctx *gin.Context) {
	fmt.Println("GetApplicationList is running")
	json := make(map[string]interface{})
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "json传输异常",
			"data": nil,
		})
	}
	userJobId := json["userJobId"].(string)
	var applications []model.GetApplicationList
	result := db.Select("*").Preload("Applicationtype").Preload("Student").
		Table("application").
		Joins("INNER JOIN applicationtype ON application.application_type_id = applicationtype.application_type_id").
		Joins("INNER JOIN student ON application.application_applicant_job_id = student.student_job_id").
		Where("application_reviewer_job_id = ?", userJobId).
		Find(&applications)
	fmt.Println(applications)
	if result.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "获取列表失败",
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "获取该老师的申请表",
			"data": applications,
		})
	}
}
