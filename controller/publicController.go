package controller

import (
	"StudentCurriculum_backEnd_Gin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserName(ctx *gin.Context) {
	fmt.Println("GetUserName is running")
	json := make(map[string]interface{})
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "json传输异常",
			"data": nil,
		})
	}
	userJobId := json["userJobId"]
	fmt.Println(userJobId)
	var dbData []map[string]interface{}
	result := db.Table("user").Select("*").
		Joins("left join teacher on teacher.teacher_job_id = user_job_id ").
		Joins("left join student on student.student_job_id = user_job_id ").
		Joins("left join dean on dean.dean_job_id = user_job_id ").
		Joins("left join counselor on counselor.counselor_job_id = user_job_id ").
		Where("user_job_id=?", userJobId).Order("user.user_role_id").
		Limit(1).Scan(&dbData)
	if dbData[0]["teacher_name"] != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户信息获取成功",
			"data": dbData[0]["teacher_name"],
		})
	} else if dbData[0]["student_name"] != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户信息获取成功",
			"data": dbData[0]["student_name"],
		})
	} else if dbData[0]["counselor_name"] != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户信息获取成功",
			"data": dbData[0]["counselor_name"],
		})
	} else if dbData[0]["dean_name"] != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户信息获取成功",
			"data": dbData[0]["dean_name"],
		})
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "用户信息获取失败",
			"data": nil,
		})
	}
	fmt.Println(result)
}

func UploadIconFile(ctx *gin.Context) {
	fmt.Println("UploadIconFile is running")
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
	userFileList := json["fileList"].(string)
	var user model.User
	db.First(&user, "user_job_id", userJobId)
	fmt.Println(user)
	result := db.Model(&user).Where("user_job_id = ?", userJobId).Updates(model.User{UserJobId: userJobId, UserAvatar: userFileList})
	if result.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  result.Error,
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 200,
			"msg":  "更新头像成功",
			"data": map[string]interface{}{
				"userJobId":  user.UserJobId,
				"userAvatar": user.UserAvatar,
			},
		})
	}
}

func GetMajor(ctx *gin.Context) {
	fmt.Println("GetMajor is running")
	json := make(map[string]interface{})
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "json传输异常",
			"data": nil,
		})
	}
	var dbData []map[string]interface{}
	classYear := json["classYear"]
	result := db.Table("class").Select("DISTINCT class_major").Where("class_year = ?", classYear).
		Order("class_major").
		Find(&dbData)
	fmt.Println(dbData)
	if result.RowsAffected != 0 {
		var data []string
		for _, value := range dbData {
			class_major, ok := value["class_major"].(string)
			if ok {
				data = append(data, class_major)
			}
		}
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 200,
			"msg":  "获取所有专业",
			"data": gin.H{
				"majorList": data,
			},
		})
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "获取不到数据",
			"data": gin.H{
				"majorList": nil,
			},
		})
	}
}

func GetClassName(ctx *gin.Context) {
	fmt.Println("GetClassName is running")
	json := make(map[string]interface{})
	err := ctx.BindJSON(&json)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "json传输异常",
			"data": nil,
		})
	}
	var dbData []map[string]interface{}
	classYear := json["classYear"]
	classMajor := json["classMajor"]
	result := db.Table("class").Select("DISTINCT class_name").Where("class_year = ?", classYear).
		Where("class_major=?", classMajor).
		Find(&dbData)
	fmt.Println(dbData)
	if result.RowsAffected != 0 {
		var data []string
		for _, value := range dbData {
			classNameList, ok := value["class_name"].(string)
			if ok {
				data = append(data, classNameList)
			}
		}
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 200,
			"msg":  "获取到班级名列表",
			"data": gin.H{
				"classNameList": data,
			},
		})
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "获取不到数据",
			"data": gin.H{
				"classNameList": nil,
			},
		})
	}
}

func GetCourseRecordOfWeek(ctx *gin.Context) {
	fmt.Println("GetCourseRecordOfWeek is running")
	var jsoninfo model.PostCourseRecordOfWeek
	if err := ctx.ShouldBindJSON(&jsoninfo); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 405,
			"msg":  "error",
		})
		return
	}
	fmt.Printf("%v \n", jsoninfo)

	var getCourseRecordOfWeek model.GetCourseRecordOfWeek

	result := db.Select("*").
		Preload("Courserecord").
		Preload("Courserecord.Course").
		Table("courserecord as csr").
		Joins("JOIN course as cs ON cs.course_id = csr.courserecord_course_id").
		Joins("JOIN class as c ON c.class_id = cs.course_class_id").
		Joins("JOIN coursetable as cst ON cst.coursetable_class_id = c.class_id").
		Where("cst.coursetable_school_year = ?", jsoninfo.CourseTableYear).
		Where("courserecord_number_of_week = ?", jsoninfo.NumberOfWeek).
		Where("class_name = ?", jsoninfo.ClassName).
		Find(&getCourseRecordOfWeek)

	if result.Error != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 201,
			"msg":  "获取列表失败",
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查找到该班级课程记录",
			"data": getCourseRecordOfWeek,
		})
	}
}
