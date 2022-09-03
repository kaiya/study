package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	Message       = "Pong"
	StopCharacter = "\r\n\r\n"
)

func main2() {

	//gin.Default()
	//engine := gin.Default()
	//engine.GET("/ping", func(context *gin.Context) {
	//	param := context.Param("key")
	//	// H is a shortcut for map[string]interface{}
	//	context.JSON(200, gin.H{
	//		"message":"pong",
	//		"key": param,
	//	})
	//	err := engine.Run()
	//	if err != nil {
	//		fmt.Errorf(err.Error())
	//		return
	//	}
	//})
}

//func maina() {
//
//	fmt.Println("main.go running")
//	// goroutine.Sleep("hello")
//	// goroutine.Sleep("world")
//
//	//port := 3333
//	//SocketServer(port)
//
//	//ginServer()
//
//	/*
//	resp, err := http.DefaultClient.Get("https://kaiyai.com")
//	if err != nil {
//		log.Fatalln("err")
//	}
//	fmt.Println(resp)
//	 */
//}
func SocketClient(port int) {
	listen, err := net.Listen("tcp6", ":"+strconv.Itoa(port))
	if err != nil {
		return
	}
	defer listen.Close()
}

func SocketServer(port int) {
	listen, err := net.Listen("tcp4", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Socket listen port %d failed, %s", port, err)
	}
	defer listen.Close()
	log.Printf("Begin listen port: %d", port)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)
ILOOP:
	for {
		n, err := r.Read(buf)
		data := string(buf[:n])
		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			log.Println("Receive: ", data)
			if isTransportOver(data) {
				break ILOOP
			}
		default:
			log.Fatalf("Receive data failed:%s", err)
			return
		}
	}
	w.Write([]byte(Message))
	w.Flush()
	log.Printf("Send: %s", Message)
}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, StopCharacter)
	return
}
