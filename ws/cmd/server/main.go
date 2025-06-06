package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader websocket.Upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	h := http.NewServeMux()

	addrr := "localhost:2223"

	h.HandleFunc("/ws", HandleWsConnection)

	s := http.Server{
		Addr:         addrr,
		Handler:      h,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server listen address: :2223")
	panic(s.ListenAndServe())

}

func HandleWsConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ws connect request")
	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("ERROR IN CONNECTING TO WS")
		fmt.Println(err)
		return
	}
	defer con.Close()

	for {
		mt, mb, err := con.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("Message type: ", mt, " and message: ", string(mb))
		err = con.WriteMessage(mt, mb)
		if err != nil {
			fmt.Println(err)
			break
		}

	}

}
