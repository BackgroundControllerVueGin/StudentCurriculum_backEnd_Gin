package model

// Counselor undefined
type Counselor struct {
	CounselorJobId string `json:"counselorJobId"`
	CounselorName  string `json:"counselorName"`
	CounselorSex   string `json:"counselorSex"`
}

// TableName 表名称
func (*Counselor) TableName() string {
	return "counselor"
}
