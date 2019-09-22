package handler

import (
	"context"
	"file-store-server/common"
	"file-store-server/config"
	"file-store-server/db"
	"file-store-server/service/account/proto"
	"file-store-server/util"
	"fmt"
	"log"
)

// User UserService 接口的实现
type User struct{}

// Signup 处理用户注册请求
func (u *User) Signup(ctx context.Context, req *proto.ReqSignup, res *proto.RespSignup) error {
	username := req.Username
	passwd := req.Password
	// 参数简单校验
	if len(username) < 3 || len(passwd) < 5 {
		log.Println("注册参数无效")
		res.Code = common.StatusParamInvalid
		res.Message = "注册参数无效"
		return fmt.Errorf("注册参数无效")
	}
	encPasswd := util.Sha1([]byte(passwd + config.PwdSalt))
	ok := db.UserSignup(username, encPasswd)
	if ok {
		res.Code = common.StatusOK
		res.Message = "注册成功"
		return nil
	}
	log.Println("注册失败")
	res.Code = common.StatusRegisterFailed
	res.Message = "注册失败"
	return fmt.Errorf("注册失败")
}
