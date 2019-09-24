package api

import (
	"context"
	"file-store-server/db"
	"file-store-server/meta"
	"file-store-server/service/transfer/proto"
	"file-store-server/util"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/micro/go-micro"

	"github.com/gin-gonic/gin"
)

var transferCli proto.TransferService

func init() {
	service := micro.NewService()

	service.Init()

	transferCli = proto.NewTransferService("go.micro.service.transfer", service.Client())
}

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

	// 调用文件传输微服务
	transferResp, err := transferCli.Transfer(context.TODO(), &proto.ReqTrans{
		FileSha1:     fileMeta.FileSha1,
		CurLocation:  fileMeta.Location,
		DestLocation: "oss/" + header.Filename,
	})

	if err != nil {
		log.Println(transferResp.Message)
		resp.Code = -6
		resp.Msg = transferResp.Message
		return
	}

	resp.Code = 0
	resp.Msg = "文件更新成功"
}
