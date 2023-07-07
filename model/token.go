package model

import "github.com/dgrijalva/jwt-go"

type UserToken struct {
	No         string `json:"no"`
	UserJobId  string `json:"user_job_id"`
	Password   string `json:"password"`
	UserRoleId int64  `json:"userRoleId"`
	Token      string `json:"token"`
	jwt.StandardClaims
}
