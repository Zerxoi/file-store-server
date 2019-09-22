package handler

import (
	"context"
	"file-store-server/service/account/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
)

var (
	userCli proto.UserService
)

func init() {
	service := micro.NewService()
	// 初始化,解析命令行参数
	service.Init()

	// 初始化一个rpcClient
	userCli = proto.NewUserService("go.micro.service.user", service.Client())
}

// GETSignupHandler 返回处理页面
func GETSignupHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

// POSTSignupHandler 处理注册POST请求
func POSTSignupHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")
	_, err := userCli.Signup(context.TODO(), &proto.ReqSignup{
		Username: username,
		Password: passwd,
	})

	if err != nil {
		log.Println("客户端服务远程调用失败")
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"code": resp.Code,
	// 	"msg":  resp.Message,
	// })
	c.Data(http.StatusOK, "text/plain", []byte("SUCCESS"))
}
