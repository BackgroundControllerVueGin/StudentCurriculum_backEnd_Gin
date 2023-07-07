package model

import "time"

type PostCourseRecordOfWeek struct {
	ClassYear       string `json:"classYear" gorm:"class_year"`
	ClassMajor      string `json:"classMajor" gorm:"class_major"`
	ClassName       string `json:"className" gorm:"class_name"`
	CourseTableYear string `json:"courseTableYear" gorm:"course_table_year"`
	NumberOfWeek    string `json:"numberOfWeek" gorm:"number_of_week"`
}

//
//type GetCourseRecordOfWeek struct {
//	Class GetClass
//}

type GetCourseRecordOfWeek struct {
	ClassId                  string            `json:"classID" gorm:"column:class_id"`
	ClassYear                string            `json:"classYear" gorm:"column:class_year"`
	ClassMajor               string            `json:"classMajor" gorm:"column:class_major"`
	ClassName                string            `json:"className" gorm:"column:class_name"`
	CoursetableSchoolYear    string            `json:"courseTableYear" gorm:"column:coursetable_school_year"`
	CourserecordNumberOfWeek string            `json:"numberOfWeek" gorm:"column:courserecord_number_of_week"`
	CourseID                 int64             `json:"-" gorm:"column:course_id;table:course"`
	Courserecord             []GetCourserecord `json:"courseRecordList"  gorm:"table:courserecord;foreignKey:CourserecordCourseId;references:CourseID;"`
	//Course                   Course          `json:"courseList" gorm:"table:course;foreignKey:CourseClassId;references:ClassId"`
}

type GetCourserecord struct {
	CourserecordId           int64  `json:"courserecordId" gorm:"column:courserecord_id"`
	CourserecordTableId      int64  `json:"courserecordTableId" gorm:"column:courserecord_table_id"`
	CourserecordNumberOfWeek int64  `json:"courserecordNumberOfWeek" gorm:"column:courserecord_number_of_week"`
	CourserecordDayOfWeek    int8   `json:"courserecordDayOfWeek" gorm:"column:courserecord_day_of_week"`
	CourserecordStartTime    string `json:"courserecordStartTime" gorm:"column:courserecord_start_time"`
	CourserecordCourseId     int64  `json:"courserecordCourseId" gorm:"column:courserecord_course_id;" `
	CourserecordAddress      string `json:"courserecordAddress" gorm:"column:courserecord_address"`
	Course                   Course `json:"courseList" gorm:"table:course;foreignKey:CourserecordCourseId;references:CourseId"`
}

// Course undefined
type GetCourse struct {
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
func (*GetCourse) TableName() string {
	return "course"
}

//	func (GetCourseRecordOfWeek) ChangeGetCourseRecordOfWeekName() string {
//		return "class"
//	}

func (this *GetCourserecord) TableName() string {
	return "courserecord"
}

type Results struct {
	Class        Class
	Courserecord Courserecord
	Course       Course
	Coursetable  Coursetable
}
