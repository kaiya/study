FROM centos
RUN sed -i 's/enabled=0/enabled=1/g' /etc/yum.repos.d/CentOS-Base.repo && yum clean all && yum makecache && yum update -y && yum install -y binutils vim gdb
RUN curl -o '/tmp/go.tar.gz' 'https://dl.google.com/go/go1.17.2.linux-amd64.tar.gz' && tar -zxf /tmp/go.tar.gz -C /usr/lib/ && ln -s /usr/lib/go/bin/* /usr/local/bin/
RUN go install github.com/go-delve/delve/cmd/dlv@latest && ln -s $(go env GOPATH)/bin/dlv /usr/local/bin/dlv
