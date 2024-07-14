package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

const (
	PORT     = ":3000"
	PROTOCOL = "tcp4"
)

func main() {
	listener, err := net.Listen(PROTOCOL, PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Printf("Listening on %s", PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("err acception conn %s", err.Error())
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		// read client request data
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("request: %s", bytes)

		// prepend prefix and send as response
		line := fmt.Sprintf("%s %s", "prefix", bytes)
		fmt.Printf("response: %s", line)
		conn.Write([]byte(line))
	}
}
