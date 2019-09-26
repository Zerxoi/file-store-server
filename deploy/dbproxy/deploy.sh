echo "编译生成二进制文件"
GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o dbproxy $GOPATH/src/file-store-server/service/dbproxy

echo "创建镜像 dbproxy"
docker build -t dbproxy .

echo "启动容器 dbproxy"
docker run -d --name=dbproxy dbproxy