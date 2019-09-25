package handler

import (
	"context"
	"errors"
	"file-store-server/common"
	"file-store-server/config"
	"file-store-server/service/account/proto"
	dbProto "file-store-server/service/dbproxy/proto"
	"file-store-server/service/dbproxy/rpc"
	"file-store-server/util"
	"fmt"
	"log"
	"time"
)

// User UserServiceHandler 接口的实现
type User struct{}

// Signup 处理用户注册请求
func (u *User) Signup(ctx context.Context, req *proto.ReqSignup, resp *proto.RespSignup) error {
	username := req.Username
	passwd := req.Password
	// 参数简单校验
	if len(username) < 3 || len(passwd) < 5 {
		log.Println("注册参数无效")
		return errors.New("注册参数无效")
	}
	encPasswd := util.Sha1([]byte(passwd + config.PwdSalt))

	dbresp, err := rpc.Client().SignupUser(context.TODO(), &dbProto.ReqSignupUser{
		Username: username,
		Passwd:   encPasswd,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(dbresp.Code, dbresp.Message)
	resp.Code = common.StatusOK
	resp.Message = "注册成功"
	return nil
}

// GenToken 产生40位的token
func GenToken(username string) string {
	// 40 bits (username + timestamp + tokenSalt) + timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	return util.MD5([]byte(username+"_tokensalt")) + ts[:8]
}

// Signin 处理用户登录
func (u *User) Signin(c context.Context, req *proto.ReqSignin, resp *proto.RespSignin) error {
	username := req.Username
	encpwd := req.Password

	_, err := rpc.Client().SigninUser(context.TODO(), &dbProto.ReqSigninUser{
		Username: username,
		Encpwd:   encpwd,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	// 2.生成访问凭证
	token := GenToken(username)

	_, err = rpc.Client().SigninUser(context.TODO(), &dbProto.ReqSigninUser{
		Username: req.Username,
		Encpwd:   req.Password,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = rpc.Client().UpdateToken(context.TODO(), &dbProto.ReqUpdateToken{
		Username: username,
		Token:    token,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	resp.Code = common.StatusOK
	resp.Message = "用户登录成功"
	resp.Token = token
	return nil
}

// UserInfo 用户信息获取
func (u *User) UserInfo(c context.Context, req *proto.ReqUserInfo, resp *proto.RespUserInfo) error {
	username := req.Username

	userInfoResp, err := rpc.Client().UserInfo(context.TODO(), &dbProto.ReqUserInfo{
		Username: username,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	resp.Code = common.StatusOK
	resp.Message = "用户信息获取成功"
	resp.Username = username
	resp.Email = userInfoResp.Email
	resp.Phone = userInfoResp.Phone
	resp.SignupAt = userInfoResp.SignupAt
	resp.LastActveAt = userInfoResp.LastActive
	resp.Status = userInfoResp.Status
	return nil
}

// UserFiles 获取用户文件
func (u *User) UserFiles(c context.Context, req *proto.ReqUserFile, resp *proto.RespUserFile) error {
	limitCnt := req.Limit
	username := req.Username

	userFilesResp, err := rpc.Client().QueryUserFile(context.TODO(), &dbProto.ReqQueryUserFile{
		Username: username,
		Limit:    limitCnt,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	if userFilesResp.Code == 1 {
		resp.Code = common.StatusOK
		resp.Message = userFilesResp.Message
		return nil
	}

	resp.Code = common.StatusOK
	resp.Message = "用户文件获取成功"
	resp.FileData = string(userFilesResp.UserFileByte)
	return nil
}
