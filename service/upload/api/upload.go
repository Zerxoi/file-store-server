package api

import (
	"encoding/json"
	"file-store-server/common"
	"file-store-server/config"
	"file-store-server/db"
	"file-store-server/meta"
	"file-store-server/mq"
	"file-store-server/util"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// GETUploadHandler 获取文件上传页面
func GETUploadHandler(c *gin.Context) {
	// 返回上传 HTML 页面
	c.Redirect(http.StatusFound, "/static/view/upload.html")
}

// POSTUploadHandler 处理文件上传
func POSTUploadHandler(c *gin.Context) {
	resp := util.RespMsg{}
	defer func() {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST,OPTIONS")
		if resp.Code < 0 {
			// 上传失败
			log.Println(resp.Msg)
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			// 上传成功
			c.JSON(http.StatusOK, resp)
		}
	}()
	// 接受文件流及存储到本地目录
	f, header, err := c.Request.FormFile("file")
	if err != nil {
		msg := "Form Data 获取文件失败"
		log.Println(msg)
		resp.Code = -1
		resp.Msg = msg
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
		msg := "创建文件失败"
		log.Println(msg)
		resp.Code = -2
		resp.Msg = msg
		return
	}
	defer nf.Close()

	fileMeta.FileSize, err = io.Copy(nf, f)
	if err != nil {
		msg := "文件复制失败"
		log.Println(msg)
		resp.Code = -3
		resp.Msg = msg
		return
	}

	nf.Seek(0, 0)
	fileMeta.FileSha1 = util.FileSha1(nf)

	ok := meta.UpdateFileMetaDB(fileMeta)
	if !ok {
		msg := "数据库元信息更新失败"
		log.Println(msg)
		resp.Code = -4
		resp.Msg = msg
		return
	}

	username := c.Request.FormValue("username")
	ok = db.OnUserFileUploadFinished(username, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize)
	if !ok {
		msg := "用户文件列表更新失败"
		log.Println(msg)
		resp.Code = -5
		resp.Msg = msg
		return
	}

	data := mq.TransferData{
		FileSha1:      fileMeta.FileSha1,
		CurLocation:   fileMeta.Location,
		DestLocation:  "oss/" + header.Filename,
		DestStoreType: common.StoreOSS,
	}
	pubData, err := json.Marshal(data)
	if err != nil {
		msg := "JOSN Marshal 失败"
		log.Println(msg)
		resp.Code = -6
		resp.Msg = msg
		return
	}

	ok = mq.Publish(config.TransExchangeName, config.TransOSSRoutingKey, pubData)
	log.Println("Producer publish", string(pubData), "to", config.TransExchangeName)
	if !ok {
		// TODO: 加入重试发送逻辑
		msg := "文件写入OSS失败"
		log.Println(msg)
		resp.Code = -7
		resp.Msg = msg
		return
	}

	resp.Code = 0
	resp.Msg = "文件更新成功"
}
