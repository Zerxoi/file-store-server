package db

import (
	"database/sql"
	"file-store-server/db/mysql"
	"fmt"
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
			log.Printf("SHA1为%s的文件已存在", fileSha1)
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

// UpdateFileLocation : 更新文件的存储地址(如文件被转移了)
func UpdateFileLocation(filehash string, fileaddr string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"update tbl_file set`file_addr`=? where  `file_sha1`=? limit 1")
	if err != nil {
		fmt.Println("预编译sql失败, err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(fileaddr, filehash)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("文件地址在修改之前已经是 %s\n", fileaddr)
		}
		return true
	}
	log.Println(err)
	return false
}
