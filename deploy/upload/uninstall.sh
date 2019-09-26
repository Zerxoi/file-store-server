echo "删除二进制文件"
rm upload

echo "结束容器 upload"
docker stop upload

echo "删除容器 upload"
docker rm upload

echo "删除镜像 fileserver/upload"
docker image rm fileserver/upload