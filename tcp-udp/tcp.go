package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	go CreateServer()
	time.Sleep(time.Second)

	ConnecClient()

}

func CreateServer() {
	cAddrr := "localhost:2222"

	l, err := net.Listen("tcp4", cAddrr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server listen tcp connections")
	b := make([]byte, 1024)

	for {
		con, err := l.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println("New client connected!")

		for {
			n, _ := con.Read(b)
			fmt.Println("S side: >>", string(b))
			if strings.TrimSpace(string(b[0:n])) == "STOP" {
				fmt.Println("Disconnect")
				con.Close()
				return
			}

			con.Write(b)
		}

	}

}

func ConnecClient() {
	cAddrr := "localhost:2222"

	c, err := net.Dial("tcp4", cAddrr)
	if err != nil {
		panic(err)
	}

	for {
		r := bufio.NewReader(os.Stdin)

		text, _ := r.ReadString('\n')
		fmt.Fprintf(c, "Client txt: %s\n", text)
		fmt.Println("")
		if strings.TrimSpace(text) == "STOP" {
			c.Close()
			fmt.Println("Disconnect!")
			return
		}

		rs, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Println("c Side: >> ", rs)

	}
}
