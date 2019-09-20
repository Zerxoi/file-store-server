package handler

import (
	"database/sql"
	"encoding/json"
	"file-store-server/common"
	"file-store-server/config"
	"file-store-server/db"
	"file-store-server/meta"
	"file-store-server/mq"
	"file-store-server/store/oss"
	"file-store-server/util"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadHandler 处理文件伤上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传 HTML 页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, fmt.Sprintf("Internal server error: %s", err))
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 接受文件流及存储到本地目录
		f, header, err := r.FormFile("file")
		if err != nil {
			log.Printf("Failed to get data: %s\n", err)
			return
		}
		defer f.Close()

		fileMeta := meta.FileMeta{
			FileName: header.Filename,
			Location: "/tmp/" + header.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		nf, err := os.Create(fileMeta.Location)
		if err != nil {
			log.Printf("Failed to create file: %s\n", err)
			return
		}
		defer nf.Close()

		fileMeta.FileSize, err = io.Copy(nf, f)
		if err != nil {
			log.Printf("Failed to save data into file: %s\n", err)
			return
		}

		nf.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(nf)
		ossPath := "oss/" + header.Filename

		data := mq.TransferData{
			FileSha1:      fileMeta.FileSha1,
			CurLocation:   fileMeta.Location,
			DestLocation:  ossPath,
			DestStoreType: common.StoreOSS,
		}
		pubData, err := json.Marshal(data)
		if err != nil {
			log.Println("UploadHandler:", err)
			return
		}
		ok := mq.Publish(config.TransExchangeName, config.TransOSSRoutingKey, pubData)
		log.Println("Producer publish", string(pubData), "to", config.TransExchangeName)
		if !ok {
			// TODO: 加入重试发送逻辑
		}
		//err = oss.Bucket().PutObject(ossPath, nf)
		//if err != nil {
		//	log.Println("UploadHandler:", err)
		//	w.Write([]byte("Upload failed"))
		//	return
		//}

		//fileMeta.Location = ossPath

		ok = meta.UpdateFileMetaDB(fileMeta)
		if !ok {
			log.Println("Failed to update file meta DB")
			return
		}
		// log.Println("文件数据库更新")

		err = r.ParseForm()
		if err != nil {
			return
		}

		// 什么时候给了请求 username 参数??????????????
		username := r.FormValue("username")
		ok = db.OnUserFileUploadFinished(username, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize)
		if !ok {
			w.Write([]byte("Upload Failed"))
			return
		}
		// log.Println("用户文件数据库更新")

		http.Redirect(w, r, "/static/view/home.html", http.StatusFound)
	}
}

// GETUploadHandler 获取文件上传页面
func GETUploadHandler(c *gin.Context) {
	// 返回上传 HTML 页面
	c.Redirect(http.StatusFound, "/static/view/index.html")
}

// POSTUploadHandler 处理文件上传
func POSTUploadHandler(c *gin.Context) {
	// 接受文件流及存储到本地目录
	f, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("POSTUploadHandler", err)
		return
	}
	defer f.Close()

	fileMeta := meta.FileMeta{
		FileName: header.Filename,
		Location: "/tmp/" + header.Filename,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	nf, err := os.Create(fileMeta.Location)
	if err != nil {
		log.Println("POSTUploadHandler", err)
		return
	}
	defer nf.Close()

	fileMeta.FileSize, err = io.Copy(nf, f)
	if err != nil {
		log.Println("POSTUploadHandler", err)
		return
	}

	nf.Seek(0, 0)
	fileMeta.FileSha1 = util.FileSha1(nf)
	ossPath := "oss/" + header.Filename

	data := mq.TransferData{
		FileSha1:      fileMeta.FileSha1,
		CurLocation:   fileMeta.Location,
		DestLocation:  ossPath,
		DestStoreType: common.StoreOSS,
	}
	pubData, err := json.Marshal(data)
	if err != nil {
		log.Println("POSTUploadHandler", err)
		return
	}
	ok := mq.Publish(config.TransExchangeName, config.TransOSSRoutingKey, pubData)
	log.Println("Producer publish", string(pubData), "to", config.TransExchangeName)
	if !ok {
		// TODO: 加入重试发送逻辑
	}

	ok = meta.UpdateFileMetaDB(fileMeta)
	if !ok {
		log.Println("Failed to update file meta DB")
		resp := util.RespMsg{
			Code: -3,
			Msg:  "更新文件列表失败",
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	username := c.Request.FormValue("username")
	ok = db.OnUserFileUploadFinished(username, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize)
	if !ok {
		resp := util.RespMsg{
			Code: -3,
			Msg:  "更新用户文件列表失败",
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := util.RespMsg{
		Code: 0,
		Msg:  "文件" + fileMeta.FileName + "上传成功",
		Data: nil,
	}
	c.JSON(http.StatusOK, resp)
}

// GetFileMetaHandler 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filesha1 := r.Form["filesha1"][0]

	// fileMeta := meta.GetFileMeta(filesha1)
	fileMeta, err := meta.GetFuncMetaDB(filesha1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Get file meta from mysql failed: %s\n", err)
		return
	}
	data, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("JSON marshal failed: %s\n", err)
		return
	}
	w.Write(data)
}

// POSTGetFileMetaHandler 文件元数据处理程序
func POSTGetFileMetaHandler(c *gin.Context) {
	filesha1 := c.Request.FormValue("filesha1")
	fileMeta, err := meta.GetFuncMetaDB(filesha1)
	if err != nil {
		log.Println("POSTGetFileMetaHandler", err)
		resp := util.RespMsg{
			Code: -3,
			Msg:  "从文件列表获取元信息失败",
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	data, err := json.Marshal(fileMeta)
	if err != nil {
		log.Println("POSTGetFileMetaHandler", err)
		resp := util.RespMsg{
			Code: -4,
			Msg:  "文件元信息Marshal失败",
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := util.RespMsg{
		Code: 0,
		Msg:  "文件元信息获取成功",
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}

// FileQueryHandler 查询批量的文件元信息
func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	limitCnt, _ := strconv.Atoi(r.Form.Get("limit"))
	username := r.FormValue("username")
	// fMetas := meta.GetLastFileMetas(limitCnt)
	userFiles := db.QueryUserFile(username, limitCnt)
	if userFiles == nil {
		log.Println("FileQueryHandler: not Found")
		return
	}
	fmt.Println(userFiles)
	// resp := util.RespMsg{
	// 	Code: 0,
	//	Msg:  "OK",
	//	Data: userFiles,
	// }
	jsbytes, err := json.Marshal(userFiles)
	if err != nil {
		log.Println("FileQueryHandler:", err)
		return
	}
	w.Write(jsbytes)
}

// POSTFileQueryHandler 获取文件队列的处理器
func POSTFileQueryHandler(c *gin.Context) {
	limitCnt, _ := strconv.Atoi(c.Request.FormValue("limit"))
	username := c.Request.FormValue("username")
	// fMetas := meta.GetLastFileMetas(limitCnt)
	userFiles := db.QueryUserFile(username, limitCnt)
	if userFiles == nil {
		log.Println("FileQueryHandler: not Found")
		resp := util.RespMsg{
			Code: -3,
			Msg:  "没有查询到用户的文件列表",
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	fmt.Println(userFiles)

	jsbytes, err := json.Marshal(userFiles)
	if err != nil {
		log.Println("POSTFileQueryHandler:", err)
		return
	}
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: string(jsbytes),
	}
	c.JSON(http.StatusOK, resp)
}

// DownloadHandler 下载文件
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fsha1 := r.Form.Get("filesha1")
	fm := meta.GetFileMeta(fsha1)

	f, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("File open failed: %s\n", err)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	_, err = ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("File read failed: %s\n", err)
		return
	}

	w.Header().Set("Conten-Type", "application/octect-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fm.FileName+"\"")

	w.Write(data)
}

// GETDownloadHandler 文件下载处理器
func GETDownloadHandler(c *gin.Context) {
	fsha1 := c.Request.FormValue("filesha1")
	fm := meta.GetFileMeta(fsha1)

	// TODO: 添加OSS文件的下载

	f, err := os.Open(fm.Location)
	if err != nil {
		log.Printf("File open failed: %s\n", err)
		resp := util.RespMsg{
			Code: -1,
			Msg:  "未找到目标文件",
			Data: nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	defer f.Close()

	c.FileAttachment(fm.Location, fm.FileName)
}

// FileMetaUpdateHandler 文件元信息更新借口(重命名)
func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filesha1")
	newFileName := r.Form.Get("filename")

	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	curlFileMeta := meta.GetFileMeta(fileSha1)
	curlFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curlFileMeta)

	data, err := json.Marshal(curlFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("FileMeta update failed: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

/* // TODO: 重命名处理器
func POSTRenameHandler(c *gin.Context) {
	fileSha1 := c.Request.FormValue("filesha1")
	newFileName := c.Request.FormValue("filename")

} */

// FileDeleteHandler 删除文件及其元信息
func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()

	fileSha1 := r.Form.Get("filesha1")
	fMeta := meta.GetFileMeta(fileSha1)
	err := os.Remove(fMeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("File remove failed: %s\n", err)
	}
	meta.RemovFileMeta(fileSha1)
	w.WriteHeader(http.StatusOK)

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

	_, err = db.GetFileMeta(filesha1)
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

	ok := db.OnUserFileUploadFinished(username, filesha1, filename, int64(filesize))
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

// POSTFastUploadHandler 文件快传处理函数
func POSTFastUploadHandler(c *gin.Context) {
	// TODO: 改进,去掉多余的参数,将快传放入普通传输的内部
	username := c.Request.FormValue("username")
	filesha1 := c.Request.FormValue("filesha1")
	filename := c.Request.FormValue("filename")
	filesize, _ := strconv.Atoi(c.Request.FormValue("filesize"))

	_, err := db.GetFileMeta(filesha1)
	if err != nil {
		resp := util.RespMsg{
			Code: -1,
			Msg:  "秒传失败",
		}
		// 服务器内部出错
		log.Println("POSTFastUploadHandler:", err)
		c.JSON(http.StatusOK, resp)
		return
	}

	ok := db.OnUserFileUploadFinished(username, filesha1, filename, int64(filesize))
	if ok {
		resp := util.RespMsg{
			Code: 0,
			Msg:  "秒传成功",
		}
		c.JSON(http.StatusOK, resp)
	} else {
		resp := util.RespMsg{
			Code: -2,
			Msg:  "秒传失败,请稍后重试",
		}
		c.JSON(http.StatusOK, resp)
	}
}

// DownloadURLHandler 生成文件的下载地址
func DownloadURLHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filesha1 := r.FormValue("filesha1")
	row, _ := db.GetFileMeta(filesha1)
	var signedURL string
	if row.FileAddr.Valid {
		signedURL = oss.DownloadURL(string(row.FileAddr.String))
		if signedURL == "" {
			log.Println("DownloadURLHandler: Failed to get file's address")
			return
		}
		w.Write([]byte(signedURL))
	}
	log.Println("File's address is null")
}

// POSTDownloadURLHandler 获取文件URL地址
func POSTDownloadURLHandler(c *gin.Context) {
	filesha1 := c.Request.FormValue("filesha1")
	row, err := db.GetFileMeta(filesha1)
	if err != nil {
		log.Println("PUTDownloadURLHandler:", err)
		resp := util.RespMsg{
			Code: -1,
			Msg:  "未找到该文件",
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	var signedURL string
	fmt.Println(row.FileAddr.Valid, row.FileAddr.String)
	if row.FileAddr.Valid {
		signedURL = oss.DownloadURL(string(row.FileAddr.String))
		c.Writer.Write([]byte(signedURL))
		c.Data(http.StatusOK, "text/plain", []byte(signedURL))
		return
	}
	resp := util.RespMsg{
		Code: -1,
		Msg:  "文件路径为空",
	}
	c.JSON(http.StatusOK, resp)
}
