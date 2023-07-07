package model

import (
	"time"
)

// Course undefined
type Course struct {
	CourseId           int64     `json:"courseId" gorm:"column:course_id"`
	CourseName         string    `json:"courseName" gorm:"column:course_name"`
	CourseCreditshours string    `json:"courseCreditshours" gorm:"column:course_creditshours"`
	CourseCredit       string    `json:"courseCredit" gorm:"column:course_credit"`
	CourseDeanJobId    string    `json:"courseDeanJobId" gorm:"column:course_dean_job_id"`
	CourseStartDate    time.Time `json:"courseStartDate" gorm:"column:course_start_date"`
	CourseEndDate      time.Time `json:"courseEndDate" gorm:"column:course_end_date"`
	CourseTeacherJobId string    `json:"courseTeacherJobId" gorm:"column:course_teacher_job_id"`
	CourseClassId      string    `json:"courseClassId" gorm:"column:course_class_id"`
}

// TableName 表名称
func (*Course) TableName() string {
	return "course"
}
