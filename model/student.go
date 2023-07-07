package model

import "time"

// Student undefined
type Student struct {
	StudentJobId     string    `json:"studentJobId"`
	StudentName      string    `json:"studentName"`
	StudentSex       string    `json:"studentSex"`
	StudentEnterYear time.Time `json:"studentEnterYear"`
	StudentMajor     string    `json:"studentMajor"`
	StudentClassId   string    `json:"studentClassId"`
}

type GetStudentList struct {
	StudentJobId     string    `json:"studentJobId"`
	StudentName      string    `json:"studentName"`
	StudentSex       string    `json:"studentSex"`
	StudentEnterYear time.Time `json:"studentEnterYear"`
	StudentMajor     string    `json:"studentMajor"`
	StudentClassId   string    `json:"studentClassId"`
	StudentClass     Class     `gorm:"table:class;foreignKey:StudentClassId;references:ClassId"`
}

// TableName 表名称
func (*Student) TableName() string {
	return "student"
}

// TableName 表名称
func (*GetStudentList) TableName() string {
	return "student"
}
