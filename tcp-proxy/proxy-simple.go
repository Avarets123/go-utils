package tcpproxy

import (
	"fmt"
	"io"
	"log"
	"net"
)

type SimpleProxy struct {
	address    string
	listenPort int
}

func NewSimpleProxy(address string, listenPort int) *SimpleProxy {
	return &SimpleProxy{
		address:    address,
		listenPort: listenPort,
	}
}

func (p *SimpleProxy) handleConn(src net.Conn) {

	dst, err := net.Dial("tcp", p.address)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		_, err := io.Copy(dst, src)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	_, err = io.Copy(src, dst)
	if err != nil {
		log.Fatalln(err)
	}

}

func (p *SimpleProxy) RunProxy() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", p.listenPort))
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Proxy listening port is: %d", p.listenPort)
	log.Printf("Proxing address is: %s \n", p.address)

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go p.handleConn(con)

	}

}
