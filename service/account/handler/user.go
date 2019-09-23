package handler

import (
	"context"
	"encoding/json"
	"file-store-server/common"
	"file-store-server/config"
	"file-store-server/db"
	"file-store-server/handler"
	"file-store-server/service/account/proto"
	"file-store-server/util"
	"fmt"
	"log"
)

// User UserService 接口的实现
type User struct{}

// Signup 处理用户注册请求
func (u *User) Signup(ctx context.Context, req *proto.ReqSignup, resp *proto.RespSignup) error {
	username := req.Username
	passwd := req.Password
	// 参数简单校验
	if len(username) < 3 || len(passwd) < 5 {
		log.Println("注册参数无效")
		return fmt.Errorf("注册参数无效")
	}
	encPasswd := util.Sha1([]byte(passwd + config.PwdSalt))
	ok := db.UserSignup(username, encPasswd)
	if ok {
		resp.Code = common.StatusOK
		resp.Message = "注册成功"
		return nil
	}
	log.Println("注册失败")
	return fmt.Errorf("注册失败")
}

// Signin 处理用户登录
func (u *User) Signin(c context.Context, req *proto.ReqSignin, resp *proto.RespSignin) error {
	username := req.Username
	encpwd := req.Password

	pwdChecked := db.UserSignIn(username, encpwd)
	if !pwdChecked {
		log.Println("密码错误")
		return fmt.Errorf("密码错误")
	}

	// 2.生成访问凭证
	token := handler.GenToken(username)
	ok := db.UpdateToken(username, token)
	if !ok {
		log.Println("Token 更新失败")
		return fmt.Errorf("Token 更新失败")
	}

	resp.Code = common.StatusOK
	resp.Message = "用户登录成功"
	resp.Token = token
	return nil
}

// UserInfo 用户信息获取
func (u *User) UserInfo(c context.Context, req *proto.ReqUserInfo, resp *proto.RespUserInfo) error {
	username := req.Username
	userinfo := db.UserInfo{}

	userinfo, err := db.GetUserInfo(username)
	if err != nil {
		log.Println("获取用户元信息出错")
		return fmt.Errorf("获取用户元信息出错")
	}

	resp.Code = common.StatusOK
	resp.Message = "用户信息获取成功"
	resp.Username = userinfo.Username
	resp.Email = userinfo.Email
	resp.Phone = userinfo.Phone
	resp.SignupAt = userinfo.SignupAt
	resp.LastActveAt = userinfo.LastActve
	resp.Status = int32(userinfo.Status)
	return nil
}

// UserFiles 获取用户文件
func (u *User) UserFiles(c context.Context, req *proto.ReqUserFile, resp *proto.RespUserFile) error {
	limitCnt := req.Limit
	username := req.Username

	userFiles := db.QueryUserFile(username, int(limitCnt))
	if userFiles == nil {
		log.Println("该用户不存在或者还没有文件")
		return fmt.Errorf("该用户不存在或者还没有文件")
	}

	jsbytes, err := json.Marshal(userFiles)
	fmt.Println(jsbytes)
	if err != nil {
		log.Println("UserFiles: JSON Marshal 失败")
		return fmt.Errorf("UserFiles: JSON Marshal 失败")
	}
	resp.Code = common.StatusOK
	resp.Message = "用户文件获取成功"
	resp.FileData = string(jsbytes)
	return nil
}
