echo "编译生成二进制文件"
GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o apigw $GOPATH/src/file-store-server/service/apigw

echo "创建镜像 fileserver/apigw"
docker build -t fileserver/apigw .

echo "启动容器 apigw"
docker run -d --name=apigw -p 8000:8000 fileserver/apigw