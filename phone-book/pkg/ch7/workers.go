package ch7

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Client struct {
	Id      int
	Integer int
}

type Result struct {
	Job    Client
	Square int
}

func RunWorkers() {
	size := runtime.GOMAXPROCS(0)
	clients := make(chan Client, size)
	data := make(chan Result, size)
	// finish := make(chan bool)
	wg := sync.WaitGroup{}

	go createReqs(15, clients)

	workersCount := 10
	wg.Add(workersCount)

	for range workersCount {
		go func() {
			worker(clients, data)
			wg.Done()
		}()
	}

	go func() {
		handleReqs(data)
		// finish <- true

	}()

	wg.Wait()

	// fmt.Println("Finish: ", <-finish)

}

func worker(clients <-chan Client, data chan<- Result) {
	for client := range clients {
		square := client.Integer * client.Integer
		result := Result{client, square}
		data <- result
		time.Sleep(time.Second)
	}
}

func handleReqs(dataCH <-chan Result) {
	for data := range dataCH {
		fmt.Printf("%+v \n", data)

	}

}

func createReqs(n int, clients chan<- Client) {
	defer close(clients)
	for v := range n {
		clients <- Client{v, v}
	}
}
