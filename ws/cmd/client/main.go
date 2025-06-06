package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var server = "localhost:2223"
var TIMESWAIT = 0
var TIMESWAITMAX = 5
var in = bufio.NewReader(os.Stdin)

func main() {

	singalsCH := make(chan os.Signal, 1)
	signal.Notify(singalsCH, os.Interrupt)

	input := make(chan string, 1)
	go getInput(input)

	URL := url.URL{
		Scheme: "ws",
		Host:   server,
		Path:   "ws",
	}

	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Received msg: ", string(msg))

		}
	}()

	for {
		select {
		case <-time.After(4 * time.Second):
			fmt.Println("Please give input ", TIMESWAIT)
			TIMESWAIT++
			if TIMESWAIT < TIMESWAITMAX {
				syscall.Kill(os.Getpid(), syscall.SIGINT)
			}
		case <-done:
			return
		case t := <-input:
			err = c.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				fmt.Println(err)
				return
			}
			TIMESWAIT = 0
			go getInput(input)
		case <-singalsCH:
			err = c.WriteMessage(1, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println(err)
				return
			}
			select {
			case <-done:
			case <-time.After(2 * time.Second):
			}
			return
		}

	}

}

func getInput(input chan string) {
	result, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	input <- result

}
