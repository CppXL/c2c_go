package contrcli

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

func ShowAllCli(Connlist []net.Conn) {
	for idx, conn := range Connlist {
		log.Printf("No %d Conn %v\n", idx+1, conn.RemoteAddr())
	}
}
func Snd2Cli(conn net.Conn) {
	var code int
	fmt.Print("输入命令:")
	fmt.Scan(&code)
	conn.Write([]byte(strconv.Itoa(code)))
	cmdrt := make([]byte, 0)
	br := bufio.NewReader(conn)
	for {
		data, err := br.ReadByte()
		if err == io.EOF {
			break
		}

		cmdrt = append(cmdrt, data)
		log.Printf("Client执行结果:%v", cmdrt)
	}

}
