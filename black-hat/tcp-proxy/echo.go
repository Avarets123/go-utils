package tcpproxy

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func echo2(con net.Conn) {
	defer con.Close()
	b := make([]byte, 512)

	for {
		size, err := con.Read(b[0:])

		if err == io.EOF {
			log.Println("Client disconnected!")
		}
		if err != nil {
			log.Println(err)
			break
		}

		log.Printf("Received %d bytes: %s \n", size, string(b))

		if _, err := con.Write(b[:size]); err != nil {
			log.Fatal("Unable to write data")
		}

	}

}

func echo(con net.Conn) {
	io.Copy(con, con)
	con.Close()
}

func echo3(con net.Conn) {

	reader := bufio.NewReader(con)
	// reader := bufio.NewReaderSize(con, 5)

	s, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Readed bytes %v, %s", len(s), s)

	log.Print("Writing data!")

	writer := bufio.NewWriter(con)

	writer.WriteString(s)

	// writer.Flush()

}

func RunEchoServer() {

	listener, _ := net.Listen("tcp", ":12345")

	fmt.Println("Listening port 12345")

	for {
		con, err := listener.Accept()

		fmt.Println("Received connection")

		if err != nil {
			log.Fatalln(err)
		}

		go echo3(con)

	}

}
