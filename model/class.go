package model

// Class undefined
type Class struct {
	ClassId             string `json:"classId" gorm:"class_id"`
	ClassName           string `json:"className" gorm:"class_name"`
	ClassYear           string `json:"classYear" gorm:"class_year"`
	ClassMajor          string `json:"classMajor" gorm:"class_major"`
	ClassCounselorJobId string `json:"classCounselorJobId" gorm:"class_counselor_job_id"`
}

// TableName 表名称
func (*Class) TableName() string {
	return "class"
}
