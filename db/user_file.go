package db

import (
	"database/sql"
	"file-store-server/db/mysql"
	"file-store-server/util"
	"log"
	"net/http"
	"strconv"
	"time"
)

// UserFile 用户文件表结构体
type UserFile struct {
	FileSha1   string
	FileSize   int64
	FileName   string
	UploadAt   string
	LastUpdate string
}

// OnUserFileUploadFinished 更新用户文件数据库
func OnUserFileUploadFinished(username, filesha1, filename string, filesize int64) bool {
	stmt, err := mysql.DBConn().Prepare("insert ignore into tbl_user_file (`user_name`, `file_sha1`, `file_name`, `file_size`, `upload_at`) values(?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, filesha1, filename, filesize, time.Now())
	if err != nil {
		return false
	}
	return true
}

// QueryUserFile 查询用户文件
func QueryUserFile(username string, limit int) []UserFile {
	stmt, err := mysql.DBConn().Prepare("select `file_sha1`, `file_size`, `file_name`, `upload_at`, `last_update` from tbl_user_file where user_name=? limit ?")
	if err != nil {
		log.Println("QueryUserFile:", err)
		return nil
	}
	defer stmt.Close()

	// limit 在home.html 的 updateFileList函数的 form data中为15
	rows, err := stmt.Query(username, limit)
	if err != nil {
		log.Println("QueryUserFile:", err)
		return nil
	}
	var userFiles []UserFile
	for rows.Next() {
		uf := UserFile{}
		err = rows.Scan(&uf.FileSha1, &uf.FileSize, &uf.FileName, &uf.UploadAt, &uf.LastUpdate)
		if err != nil {
			log.Println("QueryUserFile:", err)
			return nil
		}
		userFiles = append(userFiles, uf)
	}
	return userFiles
}

// TryFastUploadHandler 尝试秒传接口
func TryFastUploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("TryFastUploadHandler:", err)
		return
	}
	username := r.FormValue("username")
	filesha1 := r.FormValue("filesha1")
	filename := r.FormValue("filename")
	filesize, _ := strconv.Atoi(r.FormValue("filesize"))

	/*
		stmt, err := mysql.DBConn().Prepare("select file_sha1, file_size, file_name from tbl_file where file_sha1=? limit 1")
		if err != nil {
			log.Println("TryFastUploadHandler:", err)
			return
		}
		defer stmt.Close()

		row := stmt.QueryRow(filesha1)
	*/

	_, err = GetFileMeta(filesha1)
	if err != nil {
		if err == sql.ErrNoRows { // 数据库中没有该数据
			resp := util.RespMsg{
				Code: -1,
				Msg:  "秒传失败",
			}
			w.Write(resp.JSONBytes())
			return
		}
		// 服务器内部出错
		log.Println("TryFastUploadHandler:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	ok := OnUserFileUploadFinished(username, filesha1, filename, int64(filesize))
	if ok {
		resp := util.RespMsg{
			Code: 0,
			Msg:  "秒传成功",
		}
		w.Write(resp.JSONBytes())
	} else {
		resp := util.RespMsg{
			Code: -2,
			Msg:  "秒传失败,请稍后重试",
		}
		w.Write(resp.JSONBytes())
	}
}
