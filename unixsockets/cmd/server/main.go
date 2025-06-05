package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	sockPath := "/tmp/test.socket"

	_, err := os.Stat(sockPath)
	if err == nil {
		os.Remove(sockPath)
	}

	fmt.Println("UNIX socket server running on path: ", sockPath)
	s, err := net.Listen("unix", sockPath)
	if err != nil {
		panic(err)
	}

	for {
		c, _ := s.Accept()
		fmt.Println("New client connected")
		go handleConn(c)

	}

}

func handleConn(con net.Conn) {

	b := make([]byte, 1024)

	for {
		n, _ := con.Read(b)

		fmt.Println(string(b))

		if strings.TrimSpace(string(b[:n])) == "STOP" {
			fmt.Println("DISCONNECT")
			con.Close()
			return
		}

		con.Write([]byte("Success"))

	}

}
