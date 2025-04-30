package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	conn, err := net.Dial("tcp4", ":8081")
	if err != nil {
		fmt.Println(err)
	}

	var str string

	serverReader := bufio.NewReader(conn)
	serverWriter := bufio.NewWriter(conn)
	serverWriter.WriteString("hi\n")
	serverWriter.Flush()
	// conn.Write([]byte("hi\n"))
	// fmt.Println(serverReader. serverReader.Size(), serverReader.Buffered(), str)
	// bts, _ := serverReader.ReadBytes(10)
	// bts = bts[:len(bts)-1]
	// fmt.Println(bts)

	str, _ = serverReader.ReadString('\n')
	str = strings.TrimSpace(str)
	fmt.Println(serverReader.Size(), serverReader.Buffered(), str)
	time.Sleep(2 * time.Second)

	// serverWriter.WriteString("aboba\n")
	// serverWriter.Flush()
	// str, _ = serverReader.ReadString('\n')
	// str = strings.TrimSpace(str)
	// fmt.Println(serverReader.Size(), serverReader.Buffered(), str)
	// time.Sleep(2 * time.Second)

	// serverWriter.WriteString("exit\n")
	// serverWriter.Flush()
	// str, _ = serverReader.ReadString('\n')
	// str = strings.TrimSpace(str)
	// fmt.Println(serverReader.Size(), serverReader.Buffered(), str)
	// time.Sleep(2 * time.Second)

	// serverWriter.WriteString("elsagate\n")
	// serverWriter.Flush()

	// str, err = serverReader.ReadString('\n')
	// if errors.Is(err, io.EOF) {
	// 	fmt.Println(err)
	// }
	// str = strings.TrimSpace(str)
	// fmt.Println(serverReader.Size(), serverReader.Buffered(), str)
	// time.Sleep(2 * time.Second)

	conn.Close()
}
