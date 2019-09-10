package handler

import (
	"file-store-server/db"
	"file-store-server/util"
	"io/ioutil"
	"log"
	"net/http"
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
