package model

import "time"

// Application undefined
type Application struct {
	ApplicationId             int64     `json:"applicationId"`             // 申请ID
	ApplicationTypeId         int64     `json:"applicationTypeId"`         // Sick病假Personal事假AnnualLeave年假Bereavement丧假Other其他
	ApplicationContent        string    `json:"applicationContent"`        // 申请表内容
	ApplicationCreatedTime    time.Time `json:"applicationCreatedTime"`    // 申请表 创建的时间
	ApplicationReviewedTime   time.Time `json:"applicationReviewedTime"`   // 申请表 审核完成的时间
	ApplicationStartTime      time.Time `json:"applicationStartTime"`      // 申请表 请假开始时间
	ApplicationEndTime        time.Time `json:"applicationEndTime"`        // 申请表 请假结束的时间
	ApplicationStatus         string    `json:"applicationStatus"`         // pending(待审核)approved(审核通过)rejected(审核不通过)
	ApplicationReply          string    `json:"applicationReply"`          // 审核回复
	ApplicationApplicantJobId string    `json:"applicationApplicantJobId"` // 申请人
	ApplicationReviewerJobId  string    `json:"applicationReviewerJobId"`  // 审批人
}

// Applicationtype undefined
type Applicationtype struct {
	ApplicationTypeId   int64  `json:"applicationTypeId"`
	ApplicationTypeName string `json:"applicationTypeName"`
}

type GetApplicationList struct {
	ApplicationId             int64           `json:"applicationId"`             // 申请ID
	ApplicationTypeId         int64           `json:"applicationTypeId"`         // Sick病假Personal事假AnnualLeave年假Bereavement丧假Other其他
	ApplicationContent        string          `json:"applicationContent"`        // 申请表内容
	ApplicationCreatedTime    time.Time       `json:"applicationCreatedTime"`    // 申请表 创建的时间
	ApplicationReviewedTime   time.Time       `json:"applicationReviewedTime"`   // 申请表 审核完成的时间
	ApplicationStartTime      time.Time       `json:"applicationStartTime"`      // 申请表 请假开始时间
	ApplicationEndTime        time.Time       `json:"applicationEndTime"`        // 申请表 请假结束的时间
	ApplicationStatus         string          `json:"applicationStatus"`         // pending(待审核)approved(审核通过)rejected(审核不通过)
	ApplicationReply          string          `json:"applicationReply"`          // 审核回复
	ApplicationApplicantJobId string          `json:"applicationApplicantJobId"` // 申请人
	ApplicationReviewerJobId  string          `json:"applicationReviewerJobId"`  // 审批人
	Applicationtype           Applicationtype `json:"applicationType" gorm:"table:applicationtype;foreignKey:ApplicationTypeId;references:ApplicationTypeId"`
	Student                   Student         `json:"student" gorm:"table:student;foreignKey:ApplicationApplicantJobId;references:StudentJobId"`
}

// TableName 表名称
func (*Application) TableName() string {
	return "application"
}

// TableName 表名称
func (*Applicationtype) TableName() string {
	return "applicationtype"
}

// TableName 表名称
func (*GetApplicationList) TableName() string {
	return "application"
}
