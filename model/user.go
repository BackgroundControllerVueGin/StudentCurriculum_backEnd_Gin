package model

import "time"

// User undefined
type User struct {
	UserId           int64     `json:"userId" gorm:"user_id"`
	UserJobId        string    `json:"userJobId" gorm:"user_job_id"`
	UserPassword     string    `json:"userPassword" gorm:"user_password"`
	UserCreateOfTime time.Time `json:"userCreateOfTime" gorm:"user_create_of_time"`
	UserRoleId       int64     `json:"userRoleId" gorm:"user_role_id"`
	UserAvatar       string    `json:"userAvatar" gorm:"user_avatar"`
}

type UserRoleString struct {
	UserId           int64     `json:"userId" gorm:"user_id"`
	UserJobId        string    `json:"userJobId" gorm:"user_job_id"`
	UserPassword     string    `json:"userPassword" gorm:"user_password"`
	UserCreateOfTime time.Time `json:"userCreateOfTime" gorm:"user_create_of_time"`
	UserRoleId       string    `json:"userRoleId" gorm:"user_role_id"`
	UserAvatar       string    `json:"userAvatar" gorm:"user_avatar"`
}

type UserAndRole struct {
	UserId           int64     `json:"userId" gorm:"user_id"`
	UserJobId        string    `json:"userJobId" gorm:"user_job_id"`
	UserPassword     string    `json:"userPassword" gorm:"user_password"`
	UserCreateOfTime time.Time `json:"userCreateOfTime" gorm:"user_create_of_time"`
	UserRoleId       int64     `json:"userRoleId" gorm:"user_role_id"`
	UserAvatar       string    `json:"userAvatar" gorm:"user_avatar"`
	RoleName         string    `json:"roleName" gorm:"role_name"`
}

// TableName 表名称
func (*User) TableName() string {
	return "user"
}
