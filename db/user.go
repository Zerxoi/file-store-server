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
