package model

// Dean undefined
type Dean struct {
	DeanJobId string `json:"deanJobId"`
	DeanName  string `json:"deanName"`
}

// TableName 表名称
func (*Dean) TableName() string {
	return "dean"
}
