echo "删除二进制文件"
rm dbproxy

# echo "结束容器 dbproxy"
# docker stop dbproxy

# echo "删除容器 dbproxy"
# docker rm dbproxy

echo "删除镜像 fileserver/dbproxy"
docker image rm fileserver/dbproxy