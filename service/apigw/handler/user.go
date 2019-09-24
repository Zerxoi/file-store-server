package handler

import (
	"context"
	"file-store-server/config"
	"file-store-server/db"
	userProto "file-store-server/service/account/proto"
	upProto "file-store-server/service/upload/proto"
	"file-store-server/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
)

var (
	userCli userProto.UserService
	upCli   upProto.UploadService
)

func init() {
	service := micro.NewService()
	// 初始化,解析命令行参数
	service.Init()

	// 初始化一个rpcClient
	userCli = userProto.NewUserService("go.micro.service.user", service.Client())
	upCli = upProto.NewUploadService("go.micro.service.upload", service.Client())
}

// GETSignupHandler 返回处理页面
func GETSignupHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

// POSTSignupHandler 处理注册POST请求
func POSTSignupHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")
	resp, err := userCli.Signup(context.TODO(), &userProto.ReqSignup{
		Username: username,
		Password: passwd,
	})

	if err != nil {
		log.Println(err)
		c.Error(err)
	}

	c.JSON(http.StatusOK, resp)
}

// GETSigninHandler 转跳至用户登录页面
func GETSigninHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signin.html")
}

// POSTSigninHandler 登陆处理
func POSTSigninHandler(c *gin.Context) {
	// 1.验证用户名及密码
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")
	encpwd := util.Sha1([]byte(passwd + config.PwdSalt))
	signinResp, err := userCli.Signin(context.TODO(), &userProto.ReqSignin{
		Username: username,
		Password: encpwd,
	})
	if err != nil {
		log.Println(err)
		c.Error(err) // 响应内容为空
		return
	}
	upResp, err := upCli.UploadEntry(context.TODO(), &upProto.ReqEntry{})
	if err != nil {
		log.Println(err)
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, struct {
		Token       string
		UploadEntry string
	}{
		signinResp.Token,
		upResp.Entry,
	})
}

// POSTUserInfoHandler 请求用户信息
func POSTUserInfoHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	log.Println("UserName:", username)
	resp, err := userCli.UserInfo(context.TODO(), &userProto.ReqUserInfo{
		Username: username,
	})
	if err != nil {
		log.Println(err)
		c.Error(err) // 响应内容为空
		return
	}

	c.JSON(http.StatusOK, resp)
}

// POSTFileQueryHandler 获取文件队列的处理器
func POSTFileQueryHandler(c *gin.Context) {
	limitCnt, _ := strconv.Atoi(c.Request.FormValue("limit"))
	username := c.Request.FormValue("username")
	resp, err := userCli.UserFiles(context.TODO(), &userProto.ReqUserFile{
		Username: username,
		Limit:    int32(limitCnt),
	})
	if err != nil {
		log.Println(err)
		c.Error(err) // 响应内容为空
		return
	}
	c.JSON(http.StatusOK, resp)
}

// HTTPInterceptor HTTP 请求拦截器
func HTTPInterceptor(c *gin.Context) {
	username := c.Request.FormValue("username")
	token := c.Request.FormValue("token")
	ok := db.IsTokenValid(username, token)
	if !ok {
		c.Abort()
		log.Println("Token is not validated")
		resp := util.RespMsg{
			Code: -3,
			Msg:  "Token 无效",
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	c.Next()
}
