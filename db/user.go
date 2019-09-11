package db

import (
	"file-store-server/db/mysql"
	"fmt"
	"log"
	"time"
)

// UserInfo 用户信息
type UserInfo struct {
	Username       string
	Email          string
	Phone          string
	EmailValidated bool
	PhoneValidated bool
	SignupAt       string
	LastActve      string
	Status         int
}

// UserSignup 通过用户名及密码完成user表的注册操作
func UserSignup(username string, passwd string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"insert ignore into tbl_user(`user_name`,`user_pwd`) values (?,?)")
	if err != nil {
		log.Printf("Failed to insert: %s\n", err)
		return false
	}
	defer stmt.Close()
	res, err := stmt.Exec(username, passwd)
	if err != nil {
		log.Printf("Failed to execute: %s\n", err)
		return false
	}
	rowsAffected, err := res.RowsAffected()
	if err == nil && rowsAffected > 0 {
		return true
	}
	return false
}

// UserSignIn 判断密码是否相等
func UserSignIn(username string, encpwd string) bool {
	stmt, err := mysql.DBConn().Prepare(`select * from tbl_user where user_name=? limit 1`)
	if err != nil {
		log.Println(err)
		return false
	}

	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		log.Println(err)
		return false
	} else if rows == nil {
		log.Printf("Username not found %s: %s\n", username, err)
		return false
	}

	pRows := mysql.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
		return true
	}
	return false
}

// UpdateToken 刷新用户登录的token
func UpdateToken(username, token string) bool {
	stmt, err := mysql.DBConn().Prepare("replace into tbl_user_token (`user_name`, `user_token`) values(?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	res, err := stmt.Exec(username, token)
	if err != nil {
		log.Println(err)
		return false
	}
	if rows, err := res.RowsAffected(); rows > 0 && err == nil {
		return true
	}
	return false
}

// IsTokenValid 检验token是否有效
func IsTokenValid(username, token string) bool {
	// 检查token时效性
	var st int64
	_, err := fmt.Sscanf(token[32:], "%x", &st)
	if err != nil {
		log.Println(err)
		return false
	}
	if time.Now().Unix()-int64(st) > 86400 {
		log.Println("Token out of date")
		return false
	}

	// 验证token是否相等
	stmt, err := mysql.DBConn().Prepare(`select user_token from tbl_user_token where user_name=? limit 1`)
	if err != nil {
		log.Println(err)
		return false
	}

	defer stmt.Close()

	var dbtoken string
	err = stmt.QueryRow(username).Scan(&dbtoken)
	if dbtoken[:32] == token[:32] {
		return true
	}

	return false
}

// GetUserInfo 返回用户信息
func GetUserInfo(username string) (UserInfo, error) {
	userinfo := UserInfo{}
	stmt, err := mysql.DBConn().Prepare(`select user_name, signup_at from tbl_user where user_name=? limit 1`)
	if err != nil {
		log.Println(err)
		return userinfo, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&userinfo.Username, &userinfo.SignupAt)
	if err != nil {
		log.Println(err)
		return userinfo, err
	}

	return userinfo, nil
}

func returnBool(v int) bool {
	if v == 0 {
		return false
	}
	return true
}
