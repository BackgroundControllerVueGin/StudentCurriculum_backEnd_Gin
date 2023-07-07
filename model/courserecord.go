package model

import "time"

// Courserecord undefined
type Courserecord struct {
	CourserecordId           int64     `json:"courserecordId" gorm:"column:courserecord_id"`
	CourserecordTableId      int64     `json:"courserecordTableId" gorm:"column:courserecord_table_id"`
	CourserecordNumberOfWeek int64     `json:"courserecordNumberOfWeek" gorm:"column:courserecord_number_of_week"`
	CourserecordDayOfWeek    int8      `json:"courserecordDayOfWeek" gorm:"column:courserecord_day_of_week"`
	CourserecordStartTime    time.Time `json:"courserecordStartTime" gorm:"column:courserecord_start_time"`
	CourserecordCourseId     int64     `json:"courserecordCourseId" gorm:"column:courserecord_course_id"`
	CourserecordAddress      string    `json:"courserecordAddress" gorm:"column:courserecord_address"`
}

// TableName 表名称
func (*Courserecord) TableName() string {
	return "courserecord"
}
