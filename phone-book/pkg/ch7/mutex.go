package ch7

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex
var m2 sync.RWMutex

var wg sync.WaitGroup

func MutexUse() {
	wg.Add(2)

	go func() {
		m.Lock()
		fmt.Println("Locked by first goroutine")
		time.Sleep(2 * time.Second)
		fmt.Println("Unlocked by first goroutine")
		m.Unlock()
		wg.Done()
	}()

	go func() {
		m.Lock()
		fmt.Println("Locked by second goroutine")
		time.Sleep(2 * time.Second)
		fmt.Println("Unlocked by second goroutine")
		m.Unlock()
		wg.Done()

	}()

	wg.Wait()

}

func RWMutexUse() {

	go func() {
		m2.RLock()
		print("Read lock")
		time.Sleep(3 * time.Second)
		m2.RUnlock()
		print("Read unlock")
	}()

	time.Sleep(time.Second)

	go func() {
		m2.Lock()
		print("Write lock")
		time.Sleep(1 * time.Second)
		m2.Unlock()
		print("Write unlock")
	}()

	time.Sleep(time.Second)
	go func() {
		m2.RLock()
		print("Read 2 lock")
		time.Sleep(3 * time.Second)
		m2.RUnlock()
		print("Read 2 unlock")
	}()

	<-time.NewTimer(10 * time.Second).C

}

func print(msg string) {
	fmt.Println(msg)
}
