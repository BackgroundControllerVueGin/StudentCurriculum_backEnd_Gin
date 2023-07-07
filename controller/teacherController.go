package controller

import (
	"StudentCurriculum_backEnd_Gin/model"
	"StudentCurriculum_backEnd_Gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RemoveStudent(ctx *gin.Context) {
	fmt.Println("RemoveStudent is running")
	json := make(map[string]interface{})
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "json传输异常",
			"data": nil,
		})
	}
	studentJobId := json["studentJobId"]
	var student model.Student
	result := db.Where("student_job_id = ?", studentJobId).Delete(&student)
	if result.RowsAffected == 0 || result.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "该学生不存在",
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "学生删除成功",
			"data": nil,
		})
	}
}

func AddStudent(ctx *gin.Context) {
	var jsoninfo model.Student
	if err := ctx.ShouldBindJSON(&jsoninfo); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "error",
		})
		return
	}
	result_search := db.First("student_job_id = ?", jsoninfo.StudentJobId)
	if result_search.RowsAffected == 0 {
		result_add := db.Create(&jsoninfo)
		if result_add.Error == nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 200,
				"msg":  "学生添加成功",
				"data": nil,
			})
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 400,
				"msg":  "添加失败",
				"data": nil,
			})
		}
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "该学生已存在",
			"data": nil,
		})
	}
}

func GetStudentList(ctx *gin.Context) {
	fmt.Println("GetStudentList is running")
	json := make(map[string]interface{})
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "json传输异常",
			"data": nil,
		})
	}
	//currentPageNum, _ := utils.InterfaceToInt(json["currentPageNum"])
	//pageSizeNum, _ := utils.InterfaceToInt(json["pageSizeNum"])
	var studentList []model.GetStudentList
	total := map[string]interface{}{}
	//var studentClass []map[string]interface{}
	countResult := db.Select("count(*) as total").Table("student").
		Joins("JOIN class ON student.student_class_id = class.class_id").
		Find(&total)
	total_number := utils.GetInterfaceToInt(total["total"])
	currentPageNum := utils.GetInterfaceToInt(json["currentPageNum"])
	pageSizeNum := utils.GetInterfaceToInt(json["pageSizeNum"])
	fmt.Printf("%d %d %d", total_number, currentPageNum, pageSizeNum)
	result := db.Select("*").Table("student").Preload("StudentClass").Offset(currentPageNum - 1*pageSizeNum).Limit(pageSizeNum).
		Joins("JOIN class ON student.student_class_id = class.class_id").
		Find(&studentList)
	if result.Error != nil && countResult.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "获取学生列表失败",
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "获取学生列表",
			"data": gin.H{
				"total":       total["total"],
				"studentList": studentList,
			},
		})
	}
}

func GetStudentData(ctx *gin.Context) {
	fmt.Println("GetStudentData is running")
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
	var student []model.Student
	result := db.First(&student, studentJobId)
	if result.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "获取学生信息失败",
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "获取学生信息成功",
			"data": gin.H{
				"studentList": student,
			},
		})
	}

}
