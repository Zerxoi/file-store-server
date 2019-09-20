package handler

import (
	"file-store-server/db"
	"file-store-server/util"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	pwdSalt = "!@#$%^&*()"
)

// GETSignupHandler 返回处理页面
func GETSignupHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

// POSTSignupHandler 处理注册POST请求
func POSTSignupHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	if len(username) < 3 || len(passwd) < 5 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Invalid parameter",
			"code": -1,
		})
		return
	}

	encPasswd := util.Sha1([]byte(passwd + pwdSalt))
	ok := db.UserSignup(username, encPasswd)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Signup successed",
			"code": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Signup successed",
			"code": -2,
		})
	}
}

// GETSigninHandler 响应登录页面
func GETSigninHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signin.html")
}

// POSTSigninHandler 登c陆处理
func POSTSigninHandler(c *gin.Context) {
	// 1.验证用户名及密码
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")
	encpwd := util.Sha1([]byte(passwd + pwdSalt))
	pwdChecked := db.UserSignIn(username, encpwd)
	if !pwdChecked {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "用户名或密码不正确",
			"code": -1,
		})
		return
	}

	// 2.生成访问凭证
	token := GenToken(username)
	ok := db.UpdateToken(username, token)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Token 更新失败",
			"code": -2,
		})
		return
	}

	resp := util.RespMsg{
		Code: 0,
		Msg:  "登录成功",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	c.JSON(http.StatusOK, resp)
}

// GenToken 产生40位的token
func GenToken(username string) string {
	// 40 bits (username + timestamp + tokenSalt) + timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	return util.MD5([]byte(username+"_tokensalt")) + ts[:8]
}

// POSTUserInfoHandler 请求用户信息
func POSTUserInfoHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	fmt.Println(username)
	userinfo, err := db.GetUserInfo(username)
	if err != nil {
		log.Println("POSTUserInfoHandler:", err)
		resp := util.RespMsg{
			Code: -4,
			Msg:  "数据获取用户信息出错",
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	// 4.组装并且相应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: userinfo,
	}
	c.JSON(http.StatusOK, resp)
}

// UserInfoHandler 查询用户信息
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	/*
		在拦截器中实现
		// 1.解析请求数据
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			return
		}
		username := r.FormValue("username")
		token := r.FormValue("token")

		// 2.验证token是否有效
		ok := db.IsTokenValid(username, token)
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			log.Println("Token is not right")
			return
		}
	*/

	// 3.查询用户信息
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	username := r.FormValue("username")

	userinfo, err := db.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println(err)
		return
	}

	// 4.组装并且相应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: userinfo,
	}
	w.Write(resp.JSONBytes())
}
