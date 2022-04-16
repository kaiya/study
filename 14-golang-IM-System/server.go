package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	//在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	//消息广播的channel
	Message chan string
}

//创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//监听Message广播消息channel的goroutine 一旦有消息就发送给全部在线User
func (server *Server) ListenMessager() {
	for {
		msg := <-server.Message
		//将msg发送给全部在线User
		server.mapLock.Lock()
		for _, cli := range server.OnlineMap {
			cli.C <- msg
		}
		server.mapLock.Unlock()
	}
}

//广播消息的方法
func (server *Server) Brodcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	server.Message <- sendMsg
}

func (server *Server) Handler(conn net.Conn) {
	//当前连接的业务
	// fmt.Println("连接建立成功")
	user := NewUser(conn, server)

	user.Online()

	//监听用户是否活跃的channel
	isAlive := make(chan bool)

	//接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn read err:", err)
				return
			}
			//提取用户的消息 去除\n
			msg := string(buf[:n-1])
			//用户针对msg进行消息处理
			user.DoMessage(msg)
			//用户的任意消息 代表当前用户活跃
			isAlive <- true
		}
	}()
	//当前handler阻塞 方便演示
	for {
		select {
		case <-isAlive:
			//当前用户活跃 重置定时器
			//不做任何事情 为了激活select 更新下面的定时器(小技巧 isAlive的case一定要在定时器的case上面 这样如果活跃会直接执行下面的case 虽然不一定能进入case 但也实现了重置定时器)
		case <-time.After(time.Second * 300):
			//已经超时
			//将当前的User强制关闭
			user.SendMsg("你被T了!")
			//销毁用的资源
			close(user.C)
			//关闭连接
			conn.Close()
			//退出当前hanlder 的goroutine,用return或者runtime.Goexit()
			return //runtime.Goexit()
		}
	}
}

//启动服务器的方法
func (server *Server) Start() {

	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	//close listen socket
	defer listener.Close()
	//启动监听Message的goroutine
	go server.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accepet err: ", err)
			continue
		}
		//do handler
		go server.Handler(conn)
	}
}
