package routes

import (
	"StudentCurriculum_backEnd_Gin/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(gin_server *gin.Engine) *gin.Engine {
	gin_server.GET("/test", controller.TestApi)
	gin_server.POST("/login", controller.UserLogin)
	gin_server.POST("/register", controller.UserRegister)

	gin_server.POST("/public/getUsername", controller.GetUserName)
	gin_server.POST("/public/upload", controller.UploadIconFile)
	gin_server.POST("/public/getMajor", controller.GetMajor)
	gin_server.POST("/public/getClassName", controller.GetClassName)
	gin_server.POST("/public/getCourseRecordOfWeek", controller.GetCourseRecordOfWeek)

	teacherMid := gin_server.Group("/teacher")
	teacherMid.POST("/removeStudent", teacherMiddleware, controller.RemoveStudent)
	teacherMid.POST("/addStudent", teacherMiddleware, controller.AddStudent)
	teacherMid.POST("/getStudentList", teacherMiddleware, controller.GetStudentList)
	teacherMid.POST("/getStudent", teacherMiddleware, controller.GetStudentData)

	studentMid := gin_server.Group("/student")
	studentMid.POST("/getConselor", studentMiddleware, controller.GetConselorData)
	studentMid.POST("/submitApplication", studentMiddleware, controller.SubmitApplication)

	counselorMid := gin_server.Group("/counselor")
	counselorMid.POST("/updateApplication", counselorMiddleware, controller.UpdateApplication)
	counselorMid.POST("/getApplicationList", counselorMiddleware, controller.GetApplicationList)

	return gin_server
}
