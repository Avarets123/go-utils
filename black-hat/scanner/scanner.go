package scanner

import (
	"encoding/json"
	"fmt"
	"net"
	"sort"
	"sync"
)

func simpleScanner() {

	openedPorts := []int{}
	portsCount := 65000
	wg := sync.WaitGroup{}

	wg.Add(portsCount)
	for i := 1; i <= portsCount; i++ {

		go func(p int) {

			defer wg.Done()

			address := fmt.Sprintf("scanme.nmap.org:%v", p)

			con, err := net.Dial("tcp", address)

			if err != nil {
				fmt.Println("Error in connetion ", err)
				return
			}

			openedPorts = append(openedPorts, p)

			b, _ := json.Marshal(con)
			fmt.Println(string(b))

			con.Close()
		}(i)
	}

	wg.Wait()

	fmt.Println(openedPorts)

}

func worker(address string, ports chan int, wg *sync.WaitGroup) {

	for v := range ports {
		wg.Done()

		address := fmt.Sprintf("%v:%v", address, v)

		con, err := net.Dial("tcp", address)

		if err != nil {
			continue
		}
		fmt.Println("Port on address opened:  ", address)
		con.Close()

	}

}

func AdvScanner(address string) {

	scanPortCount := 10000
	var wg sync.WaitGroup
	portsCh := make(chan int, 100)

	for i := 1; i <= 2000; i++ {
		go worker(address, portsCh, &wg)
	}

	for i := 1; i <= scanPortCount; i++ {
		wg.Add(1)
		portsCh <- i
	}

	wg.Wait()

}

func worker2(address string, ports, results chan int) {

	for port := range ports {

		con, err := scanPortOnAddress(address, port)

		if err != nil {
			results <- 0
			continue
		}
		fmt.Println("Port on address opened:  ", port)
		con.Close()

		results <- port

	}

}

func AdvScanner2(address string, portsCount int) {

	ports, results := make(chan int, 100), make(chan int)

	openedPorts := []int{}

	for i := 0; i <= 500; i++ {
		go worker2(address, ports, results)
	}

	go func() {
		for i := 1; i <= portsCount; i++ {
			ports <- i
		}
	}()

	for i := 1; i <= portsCount; i++ {
		res := <-results
		if res != 0 {
			openedPorts = append(openedPorts, res)
			continue
		}

	}

	close(ports)
	close(results)

	sort.Ints(openedPorts)

	fmt.Println("Opened ports is: ", openedPorts)

}

func scanPortOnAddress(address string, port int) (net.Conn, error) {
	return net.Dial("tcp", fmt.Sprintf("%v:%v", address, port))
}
