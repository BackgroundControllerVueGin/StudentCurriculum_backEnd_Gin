package model

import "time"

// Coursetable undefined
type Coursetable struct {
	CoursetableId         int64     `json:"coursetableId" gorm:"coursetable_id"`
	CoursetableStartTime  time.Time `json:"coursetableStartTime" gorm:"coursetable_start_time"`
	CoursetableClassId    string    `json:"coursetableClassId" gorm:"coursetable_class_id"`
	CoursetableSchoolYear string    `json:"coursetableSchoolYear" gorm:"coursetable_school_year"`
	CoursetableMajor      string    `json:"coursetableMajor" gorm:"coursetable_major"`
	CousrsetableDeanJobId string    `json:"cousrsetableDeanJobId" gorm:"cousrsetable_dean_job_id"`
}

// TableName 表名称
func (*Coursetable) TableName() string {
	return "coursetable"
}
