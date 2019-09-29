echo "删除二进制文件"
rm apigw

# echo "结束容器 apigw"
# docker stop apigw

# echo "删除容器 apigw"
# docker rm apigw

echo "删除镜像 fileserver/apigw"
docker image rm fileserver/apigw