package db

import (
	"file-store-server/db/mysql"
	"log"
)

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
