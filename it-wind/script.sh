#!/bin/sh

echo "FROM centos" > Dockerfile

#yum change aliyun.repo
#yum install
echo "RUN sed -i 's/enabled=0/enabled=1/g' /etc/yum.repos.d/CentOS-Base.repo && \
yum clean all && \
yum makecache && \
yum update -y && \
yum install -y binutils vim gdb" >> Dockerfile

#go1.17.2
echo "RUN curl -o '/tmp/go.tar.gz' 'https://dl.google.com/go/go1.17.2.linux-amd64.tar.gz' && \
tar -zxf /tmp/go.tar.gz -C /usr/lib/ && ln -s /usr/lib/go/bin/* /usr/local/bin/" >> Dockerfile

#dlv latest
echo "RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
ln -s \$(go env GOPATH)/bin/dlv /usr/local/bin/dlv" >> Dockerfile

docker stop godev
docker rm godev
docker rmi godev

docker build -t godev .
v_path=$(go env GOPATH)
if [ -z "$(go env GOPATH)" ]
then
    v_path=~/go
fi
docker run -dit --name godev -v $v_path:/root/go godev bash
