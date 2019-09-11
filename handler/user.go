package handler

import (
	"file-store-server/db"
	"file-store-server/util"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	pwdSalt = "!@#$%^&*()"
)

// SignupHandler 处理用户注册请求
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.Write(data)
		return
	}

	r.ParseForm()
	username := r.FormValue("username")
	passwd := r.FormValue("password")

	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("Invalid parameter."))
		return
	}

	encPasswd := util.Sha1([]byte(passwd + pwdSalt))
	ok := db.UserSignup(username, encPasswd)
	if ok {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}

// SignInHandler 登录接口
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	// 1.验证用户名及密码
	r.ParseForm()
	username := r.FormValue("username")
	passwd := r.FormValue("password")
	encpwd := util.Sha1([]byte(passwd + pwdSalt))
	pwdChecked := db.UserSignIn(username, encpwd)
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}

	// 2.生成访问凭证
	token := GenToken(username)
	ok := db.UpdateToken(username, token)
	if !ok {
		w.Write([]byte("FAILED"))
		return
	}

	// 3.重定向到首页
	// fmt.Println("http://" + r.Host + "/static/view/home.html")
	// w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "http://" + r.Host + "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	w.Write(resp.JSONBytes())
}

// GenToken 产生40位的token
func GenToken(username string) string {
	// 40 bits (username + timestamp + tokenSalt) + timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	return util.MD5([]byte(username+ts+"_tokensalt")) + ts[:8]
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
