package db

import (
	"database/sql"
	"file-store-server/db/mysql"
	"log"
)

// OnFileUploadFinished 文件上传完成后保存meta
func OnFileUploadFinished(fileSha1 string, fileName string, fileSize int64, fileAddr string) bool {
	stmt, err := mysql.DBConn().Prepare("insert ignore into tbl_file (`file_sha1`, `file_name`, `file_size`, `file_addr`, `status`) value(?,?,?,?,1)")
	if err != nil {
		log.Printf("Failed to prepare statement: %s\n", err)
		return false
	}
	defer stmt.Close()

	res, err := stmt.Exec(fileSha1, fileName, fileSize, fileAddr)
	if err != nil {
		log.Printf("Failed to execute statement: %s\n", err)
		return false
	}

	if rf, err := res.RowsAffected(); err == nil {
		if rf <= 0 {
			log.Printf("File with SHA1:%s has been uploaded before", fileSha1)
		}
		return true
	}
	return false

}

// TableFile 表文件
type TableFile struct {
	FileSha1 string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

// GetFileMeta 从mysql获取文件元信息
func GetFileMeta(filesha1 string) (*TableFile, error) {
	stmt, err := mysql.DBConn().Prepare(`select file_sha1, file_name, file_size, file_addr from tbl_file where file_sha1=? and status=1`)
	if err != nil {
		log.Printf("Prepare failed: %s\n", err)
		return nil, err
	}
	defer stmt.Close()

	tFile := TableFile{}
	err = stmt.QueryRow(filesha1).Scan(&tFile.FileSha1, &tFile.FileName, &tFile.FileSize, &tFile.FileAddr)
	if err != nil {
		log.Printf("Qurey failed: %s\n", err)
		return nil, err
	}
	return &tFile, nil
}
