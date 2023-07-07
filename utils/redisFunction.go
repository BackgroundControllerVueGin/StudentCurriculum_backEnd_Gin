package utils

import (
	"StudentCurriculum_backEnd_Gin/common"
	"StudentCurriculum_backEnd_Gin/model"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

var rdb = common.GetRedis()
var ctx = context.Background()
var expires = 1800 * time.Second

func JsonSetRedis(tokenStruct model.UserToken, key string) bool {
	fmt.Println(rdb)
	data, err := json.Marshal(tokenStruct)
	if err != nil {
		panic(err)
		return false
	} else {
		err = rdb.Set(ctx, key, data, expires).Err()
		if err != nil {
			panic(err)
			return false
		}
		return true
	}
}

func JsonGetRedis(key string) model.UserToken {
	// 从 Redis 中获取字符串并反序列化为结构体
	var result model.UserToken
	data, err := rdb.Get(ctx, key).Bytes()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}
	return result
}
