package db

import (
	"file-store-server/db/mysql"
	"log"
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
