package model

// Teacher undefined
type Teacher struct {
	TeacherJobId string `json:"teacherJobId"`
	TeacherName  string `json:"teacherName"`
	TeacherSex   string `json:"teacherSex"`
	CourseId     int64  `json:"courseId"`
}

// TableName 表名称
func (*Teacher) TableName() string {
	return "teacher"
}
