package handler

import (
	redisconn "file-store-server/cache/redis"
	"file-store-server/db"
	"file-store-server/util"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

// MultipartUploadInfo 分块上传文件信息
type MultipartUploadInfo struct {
	FileSha1   string
	FileSize   int
	UploadID   string
	ChunkSize  int
	ChunkCount int
}

// InitialMultipartUploadHandler 初始化分块上传
func InitialMultipartUploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 解析用户请求参数
	err := r.ParseForm()
	if err != nil {
		log.Println("InitialMultipartUploadHandler:", err)
		return
	}
	username := r.FormValue("username")
	filename := r.FormValue("filename")
	filesha1 := r.FormValue("filesha1")
	filesize, err := strconv.Atoi(r.FormValue("filesize"))
	if err != nil {
		w.Write(util.NewRespMsg(-1, "Params invalid", nil).JSONBytes())
		log.Println("InitialMultipartUploadHandler:", err)
		return
	}

	// 在唯一数据库中存在, 使用秒传
	if _, err := db.GetFileMeta(filesha1); err == nil {
		db.OnUserFileUploadFinished(username, filesha1, filename, int64(filesize))
		return
	}

	// 如果文件不大可以不使用分块上传

	// 2.获得redis的一个连接
	rConn := redisconn.RedisPool().Get()
	defer rConn.Close()

	// 3.生成分块上传的初始化信息
	upInfo := MultipartUploadInfo{
		FileSha1:   filesha1,
		FileSize:   filesize,
		UploadID:   username + fmt.Sprintf("%x", time.Now().UnixNano()),
		ChunkSize:  5 * 2 << 20,
		ChunkCount: int(math.Ceil(float64(filesize) / (5 * 2 << 20))),
	}

	// 4.将初始化信息写入到redis缓存
	rConn.Do("HSET", "MP_"+upInfo.UploadID, "chunkcount", upInfo.ChunkCount)
	rConn.Do("HSET", "MP_"+upInfo.UploadID, "filesha1", upInfo.FileSha1)
	rConn.Do("HSET", "MP_"+upInfo.UploadID, "filesize", upInfo.FileSize)

	// 5.将响应初始化数据返回到客户端
	w.Write(util.NewRespMsg(0, "OK", upInfo).JSONBytes())
}

// UploadPartHandler 上传文件分块
func UploadPartHandler(w http.ResponseWriter, r *http.Request) {
	// 1.解析用户请求参数
	err := r.ParseForm()
	if err != nil {
		log.Println("UploadPartHandler:", err)
	}
	// username := r.FormValue("username")
	uploadid := r.FormValue("uploadid")
	chunkIndex := r.FormValue("index")

	// 2.获得redis连接池中的一个连接
	rConn := redisconn.RedisPool().Get()
	defer rConn.Close()

	// 3.获得文件句柄,用于存储分块内容
	fpath := "/home/zerxoi/Desktop/" + uploadid
	err = os.MkdirAll(fpath, 0744)
	if err != nil {
		log.Println("UploadPartHandler:", err)
		return
	}
	f, err := os.Create(fpath + "/" + chunkIndex)
	if err != nil {
		w.Write(util.NewRespMsg(-1, "Upload part failed", nil).JSONBytes())
		log.Println("UploadPartHandler:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 2<<20)
	for {
		n, err := r.Body.Read(buf)
		// 可以在此加上一个hash校验,验证文件完整性
		f.Write(buf[:n])
		if err != nil {
			break
		}
	}

	// 4.更新redis缓存状态
	rConn.Do("HSET", "MP_"+uploadid, "chkidx_"+chunkIndex, 1)

	// 5.返回处理结果到客户端
	w.Write(util.NewRespMsg(0, "OK", nil).JSONBytes())
}

// CompleteUploadHandler 通知上传合并接口
func CompleteUploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1.解析请求参数
	err := r.ParseForm()
	if err != nil {
		log.Println("CompleteUploadHandler:", err)
		return
	}
	uploadid := r.FormValue("uploadid")
	filesha1 := r.FormValue("filesha1")
	filesize, err := strconv.Atoi(r.FormValue("filesize"))
	if err != nil {
		log.Println("CompleteUploadHandler:", err)
		return
	}
	filename := r.FormValue("filename")
	username := r.FormValue("username")

	// 2.获得redis连接
	rConn := redisconn.RedisPool().Get()
	defer rConn.Close()

	// 3.通过uploadID查询redis判断是否完成所有分块的上传
	data, err := redis.Values(rConn.Do("HGETALL", "MP_"+uploadid))
	if err != nil {
		w.Write(util.NewRespMsg(-1, "complete upload failed", nil).JSONBytes())
		log.Println("CompleteUploadHandler:", err)
		return
	}
	totalCount := 0
	chunkCount := 0
	for i := 0; i < len(data); i += 2 {
		k := string(data[i].([]byte))
		v := string(data[i+1].([]byte))
		if k == "chunkcount" {
			totalCount, _ = strconv.Atoi(v)
		} else if strings.HasPrefix(k, "chkidx_") && v == "1" {
			chunkCount++
		}

	}
	if totalCount != chunkCount {
		w.Write(util.NewRespMsg(-1, "Invalid request", nil).JSONBytes())
		return
	}

	// 4.TODO:合并分块

	fmt.Println("moring")
	// 5.更新唯一文件表和用户文件表
	db.OnFileUploadFinished(filesha1, filename, int64(filesize), "")
	db.OnUserFileUploadFinished(username, filesha1, filename, int64(filesize))

	// 6.向客户端响应处理结果
	w.Write(util.NewRespMsg(0, "OK", nil).JSONBytes())
}
