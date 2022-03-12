package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
)

type recvData struct {
	len  int
	data []byte
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Printf("Usage: %s ip port", args[0])
	}
	ip := args[1]
	port, err := strconv.Atoi(args[2])
	fmt.Printf("try to conn %s:%d\n", ip, port)
	if err != nil {
		fmt.Println("err:", err)
	}
	conn, err := net.Dial("tcp", args[1]+":"+args[2])
	if err != nil {
		fmt.Println("err:", err)
	}
	// 接收到连接
	defer conn.Close()
	recvChan := make(chan recvData, 10)
	fmt.Println("conn to target, begin to transfor data")
	go recv(conn, recvChan)
	go handleRecvData(recvChan)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	for {
		snd := make([]byte, 0)
		fmt.Print("input:")
		fmt.Scan(&snd)
		fmt.Print("sending ", snd)
		send(conn, snd)
		fmt.Println(" done")

		// send(conn, []byte("wait client send\n"))

	}
}

func send(conn net.Conn, data []byte) {
	// data := <-dataChan
	n, err := conn.Write(data)
	if err != nil {
		fmt.Println("err:", err)

	}
	if n != len(data) {
		fmt.Printf("just send %d bytes data, but should send %d bytes data\n", n, len(data))
	}
}

func recv(conn net.Conn, recvChan chan<- recvData) {
	defer func() {
		fmt.Println("接收数据的协程关闭")
	}()
	tmp := make([]byte, 1024)
	for {
		n, err := conn.Read(tmp)

		if err != nil {
			fmt.Println("err:", err)

			if err == io.EOF {
				// recvChan <- recvData{
				// 	len:  -1,
				// 	data: []byte("EOF"),
				// }
				return
			}
			recvChan <- recvData{
				len:  0,
				data: nil,
			}
		}
		recvChan <- recvData{
			len:  n,
			data: tmp,
		}
		fmt.Println("read once data")
	}

}

func handleRecvData(recvChan <-chan recvData) {
	defer func() {
		fmt.Println("处理数据的协程关闭")
	}()
	for {
		fmt.Println("test")
		select {
		case recv, ok := <-recvChan:
			if !ok {
				return
			}
			if recv.len == 0 && recv.data == nil {
				fmt.Println("recv data err1, continue")
				return
			} else if recv.data == nil && recv.len != 0 {
				fmt.Println("recv data err2, continue")
				continue
			} else if recv.len == 0 && recv.data != nil {
				// 判断是不是EOF

				fmt.Println("recv data err3, continue")
				return
			}
			fmt.Println("\nrecv:", recv.data[0:recv.len-1])

		}

	}
}
