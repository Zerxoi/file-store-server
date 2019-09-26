package api

import (
	"context"
	dbProto "file-store-server/service/dbproxy/proto"
	"file-store-server/service/dbproxy/rpc"
	"file-store-server/service/transfer/proto"
	"file-store-server/util"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"

	"github.com/gin-gonic/gin"
)

var transferCli proto.TransferService

func init() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"172.17.0.1:8500",
		}
	})

	service := micro.NewService(micro.Registry(reg))

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

	filename := header.Filename
	fileaddr := "/tmp/" + filename

	nf, err := os.Create(fileaddr)
	if err != nil {
		msg := "创建文件失败"
		log.Println(msg)
		resp.Code = -2
		resp.Msg = msg
		return
	}
	defer nf.Close()

	filesize, err := io.Copy(nf, f)
	if err != nil {
		msg := "文件复制失败"
		log.Println(msg)
		resp.Code = -3
		resp.Msg = msg
		return
	}

	nf.Seek(0, 0)
	filesha1 := util.FileSha1(nf)

	_, err = rpc.Client().UploadFile(context.TODO(), &dbProto.ReqUploadFile{
		FileSha1: filesha1,
		FileName: filename,
		FileSize: filesize,
		FileAddr: fileaddr,
	})
	if err != nil {
		log.Println(err)
		resp.Code = -4
		resp.Msg = err.Error()
		return
	}

	username := c.Request.FormValue("username")
	_, err = rpc.Client().UploadUserFile(context.TODO(), &dbProto.ReqUploadUserFile{
		Username: username,
		FileSha1: filesha1,
		FileSize: filesize,
		FileName: filename,
	})
	if err != nil {
		log.Println(err)
		resp.Code = -5
		resp.Msg = err.Error()
		return
	}

	// 调用文件传输微服务
	_, err = transferCli.Transfer(context.TODO(), &proto.ReqTrans{
		FileSha1:     filesha1,
		CurLocation:  fileaddr,
		DestLocation: "oss/" + header.Filename,
	})

	if err != nil {
		log.Println(err)
		resp.Code = -6
		resp.Msg = err.Error()
		return
	}

	resp.Code = 0
	resp.Msg = "文件更新成功"
}
