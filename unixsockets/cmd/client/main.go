package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	sockPath := "/tmp/test.socket"

	con, err := net.Dial("unix", sockPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to socket!")
	data := make([]byte, 500)
	for {

		r := bufio.NewReader(os.Stdin)

		n, _ := r.Read(data)
		con.Write(data[:n])
		n2, _ := con.Read(data)

		fmt.Println(string(data[:n2]))

	}

}
