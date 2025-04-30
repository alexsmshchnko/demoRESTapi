package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() (err error) {
	listener, err := net.Listen("tcp4", ":8081")
	if err != nil {
		return fmt.Errorf("net.Listen start error: %w", err)
	}
	defer func() {
		fmt.Printf("listeren.Close err: %v", listener.Close())
	}()

	fmt.Printf("server started: %s %s\n", listener.Addr().Network(), listener.Addr().String())

	// ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	for {
		// select {
		// case <-ctx.Done():
		// 	fmt.Printf("signal interrupt received\n")
		// 	return
		// default:
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("listener.Accept error: %v\n", err)
			continue
		}

		go handleConnection(conn)
		// }

	}

	// return
}

func handleConnection(conn net.Conn) {
	defer func(addr string) {
		fmt.Printf("connection %s close: error %v\n", addr, conn.Close())
	}(conn.RemoteAddr().String())
	fmt.Printf("client connected: %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		var buf []byte
		message, err := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		if err != nil {
			fmt.Printf("client disconnected: %s\n", conn.RemoteAddr())
			return
		}

		fmt.Printf("new message received from %s : %s\n", conn.RemoteAddr(), message)

		if message == "exit" {
			buf = []byte("bye.\n")
			fmt.Printf("buf: %s\n", string(buf))
			conn.Write(buf)
			return
		}

		buf = []byte(fmt.Sprintf("your message: %s\n", message))
		fmt.Printf("buf: %s\n", string(buf))
		conn.Write(buf)
	}
}
