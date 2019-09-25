package rpc

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"file-store-server/service/dbproxy/mysql"
	"file-store-server/service/dbproxy/proto"
	"fmt"
	"log"
	"sort"
	"time"
)

// DBProxy DBProxyService 接口的实现
type DBProxy struct{}

// UploadFile 文件元信息上传
func (db *DBProxy) UploadFile(c context.Context, req *proto.ReqUploadFile, resp *proto.RespUploadFile) error {
	stmt, err := mysql.DBConn().Prepare("insert ignore into tbl_file (`file_sha1`, `file_name`, `file_size`, `file_addr`, `create_at`, `update_at`, `status`) value(?,?,?,?,?,?,1)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	timeNow := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(req.FileSha1, req.FileName, req.FileSize, req.FileAddr, timeNow, timeNow)
	if err != nil {
		log.Printf("文件元信息上传SQL语句执行失败")
		return err
	}

	if rf, err := res.RowsAffected(); err == nil {
		if rf <= 0 {
			log.Printf("SHA1为%s的文件已存在", req.FileSha1)
		}
		resp.Code = 0
		resp.Message = "文件元信息上传成功"
		return nil
	}
	log.Println("文件元信息上传SQL语句执行未生效")
	return err
}

// FileMeta 获取文件元信息
func (db *DBProxy) FileMeta(c context.Context, req *proto.ReqFileMeta, resp *proto.RespFileMeta) error {
	stmt, err := mysql.DBConn().Prepare(`select file_sha1, file_name, file_size, file_addr from tbl_file where file_sha1=? and status=1`)
	if err != nil {
		log.Println("获取文件元信息SQL语句准备失败")
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(req.FileSha1).Scan(&resp.FileSha1, &resp.FileName, &resp.FileSize, &resp.FileAddr)
	if err != nil {
		log.Printf("获取文件元信息SQL语句查询失败")
		return err
	}
	resp.Code = 0
	resp.Message = "获取文件元信息成功"
	return nil
}

// UpdateFileLocation 更新文件位置
func (db *DBProxy) UpdateFileLocation(c context.Context, req *proto.ReqUpdateFileLocation, resp *proto.RespUpdateFileLocation) error {
	stmt, err := mysql.DBConn().Prepare(
		"update tbl_file set`file_addr`=?, `update_at`=? where  `file_sha1`=? limit 1")
	if err != nil {
		log.Println("更新文件位置SQL语句准备失败")
		return err
	}
	defer stmt.Close()

	ret, err := stmt.Exec(req.FileAddr, time.Now().Format("2006-01-02 15:04:05"), req.FileSha1)
	if err != nil {
		log.Println("更新文件位置SQL语句执行失败")
		return err
	}
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			log.Printf("文件地址在修改之前已经是 %s\n", req.FileAddr)
		}
		resp.Code = 0
		resp.Message = "更新文件位置成功"
		return nil
	}
	log.Println("更新文件位置SQL执行但未生效")
	return err
}

// UploadUserFile 上传用户文件元信息
func (db *DBProxy) UploadUserFile(c context.Context, req *proto.ReqUploadUserFile, resp *proto.RespUploadUserFile) error {
	stmt, err := mysql.DBConn().Prepare("insert ignore into tbl_user_file (`user_name`, `file_sha1`, `file_name`, `file_size`, `upload_at`) values(?,?,?,?,?)")
	if err != nil {
		log.Println("上传用户文件元信息SQL语句准备失败")
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(req.Username, req.FileSha1, req.FileName, req.FileSize, time.Now())
	if err != nil {
		log.Println("上传用户文件元信息SQL语句执行失败")
		return err
	}
	if rf, err := res.RowsAffected(); err == nil {
		if rf <= 0 {
			log.Printf("SHA1为%s的文件已存在", req.FileSha1)
		}
		resp.Code = 0
		resp.Message = "用户文件元信息上传成功"
		return nil
	}
	log.Println("上传用户文件元信息SQL语句执行未生效")
	return err
}

// UserFile 用户文件表结构体
type UserFile struct {
	FileSha1   string
	FileSize   int64
	FileName   string
	UploadAt   string
	LastUpdate string
}

// QueryUserFile 查询用户文件
func (db *DBProxy) QueryUserFile(c context.Context, req *proto.ReqQueryUserFile, resp *proto.RespQueryUserFile) error {
	stmt, err := mysql.DBConn().Prepare("select `file_sha1`, `file_size`, `file_name`, `upload_at`, `last_update` from tbl_user_file where user_name=? limit ?")
	if err != nil {
		log.Println("查询用户文件SQL语句准备失败")
		return err
	}
	defer stmt.Close()

	// limit 在home.html 的 updateFileList函数的 form data中为15
	rows, err := stmt.Query(req.Username, req.Limit)
	if err != nil {
		log.Println("查询用户文件SQL语句查询失败")
		return err
	}
	var userFiles []UserFile
	for rows.Next() {
		uf := UserFile{}
		err = rows.Scan(&uf.FileSha1, &uf.FileSize, &uf.FileName, &uf.UploadAt, &uf.LastUpdate)
		if err != nil {
			log.Println("查询用户文件 Scan失败")
			return err
		}
		userFiles = append(userFiles, uf)
	}
	if userFiles == nil {
		resp.Code = 1
		resp.Message = " 该用户还没有文件"
		return nil
	}

	sort.Sort(ByUploadTime(userFiles))
	jsbyte, err := json.Marshal(userFiles)
	if err != nil {
		log.Println(err)
		return err
	}
	resp.Code = 0
	resp.Message = " 用户文件查询成功"
	resp.UserFileByte = jsbyte
	return nil
}

// SignupUser 用户注册
func (db *DBProxy) SignupUser(c context.Context, req *proto.ReqSignupUser, resp *proto.RespSignupUser) error {
	stmt, err := mysql.DBConn().Prepare(
		"insert ignore into tbl_user(`user_name`,`user_pwd`) values (?,?)")
	if err != nil {
		log.Println("用户注册SQL语句准备失败")
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(req.Username, req.Passwd)
	if err != nil {
		log.Printf("用户注册SQL语句执行失败")
		return err
	}
	if rowsAffected, err := res.RowsAffected(); err == nil {
		if rowsAffected <= 0 {
			log.Println("该用户已注册")
			return errors.New("该用户已注册")
		}
		resp.Code = 0
		resp.Message = "用户注册成功"
		return nil
	}
	fmt.Println("用户注册失败")
	return errors.New("用户注册失败")
}

// SigninUser 用户登录
func (db *DBProxy) SigninUser(c context.Context, req *proto.ReqSigninUser, resp *proto.RespSigninUser) error {
	stmt, err := mysql.DBConn().Prepare(`select user_pwd from tbl_user where user_name=? limit 1`)
	if err != nil {
		log.Println("用户登录SQL语句准备失败")
		return err
	}
	defer stmt.Close()

	encpasswd := ""
	err = stmt.QueryRow(req.Username).Scan(&encpasswd)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("该用户未注册")
			return errors.New("该用户未注册")
		}
		log.Println("用户登录SQL语句查询失败")
		return err
	}
	if req.Encpwd == encpasswd {
		resp.Code = 0
		resp.Message = "用户登录成功"
		return nil
	}
	log.Println("用户密码输入错误")
	return errors.New("用户密码输入错误")
}

// UserInfo 获取用户元信息
func (db *DBProxy) UserInfo(c context.Context, req *proto.ReqUserInfo, resp *proto.RespUserInfo) error {
	stmt, err := mysql.DBConn().Prepare(`select user_name, email, phone, email_validated, phone_validated, signup_at, last_active, status from tbl_user where user_name=? limit 1`)
	if err != nil {
		log.Println("获取用户元信息SQL语句准备失败")
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(req.Username).Scan(&resp.Username, &resp.Email, &resp.Phone, &resp.EmailValidated, &resp.PhoneValidated, &resp.SignupAt, &resp.LastActive, &resp.Status)
	if err != nil {
		log.Println("获取用户元信息SQL语句查询失败")
		return err
	}
	resp.Code = 0
	resp.Message = "获取用户元信息成功"
	return nil
}

// Token 从数据库获取Token
func (db *DBProxy) Token(c context.Context, req *proto.ReqToken, resp *proto.RespToken) error {
	stmt, err := mysql.DBConn().Prepare(`select user_token from tbl_user_token where user_name=? limit 1`)
	if err != nil {
		log.Println("从数据库获取Token SQL语句准备失败")
		return err
	}
	defer stmt.Close()

	var dbtoken string
	err = stmt.QueryRow(req.Username).Scan(&dbtoken)
	if err != nil {
		log.Println("从数据库获取Token SQL语句查询失败")
		return err
	}
	resp.Code = 0
	resp.Message = "从数据库获取Token成功"
	resp.Token = dbtoken
	return nil
}

// UpdateToken 更新Token
func (db *DBProxy) UpdateToken(c context.Context, req *proto.ReqUpdateToken, resp *proto.RespUpdateToken) error {
	stmt, err := mysql.DBConn().Prepare("replace into tbl_user_token (`user_name`, `user_token`) values(?,?)")
	if err != nil {
		log.Println("更新Token SQL语句准备失败")
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(req.Username, req.Token)
	if err != nil {
		log.Println("更新Token SQL语句执行失败")
		return err
	}
	if rowsAffected, err := res.RowsAffected(); err == nil {
		if rowsAffected <= 0 {
			log.Println("SQL语句Token更新未生效")
			return errors.New("SQL语句Token更新未生效")
		}
		resp.Code = 0
		resp.Message = "Token更新完成"
		return nil
	}
	fmt.Println("Token更新失败")
	return err
}
