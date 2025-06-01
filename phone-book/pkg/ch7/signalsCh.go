package ch7

import (
	"fmt"
	"sync"
	"time"
)

func SignalsChannel() {
	a, b, c, d := make(chan any), make(chan any), make(chan any), make(chan any)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		D(d)
		wg.Done()
	}()

	go A(a, b)

	go func() {
		D(d)
		wg.Done()
	}()

	go B(b, c)

	go func() {
		D(d)
		wg.Done()
	}()

	go C(c, d)

	close(a)

	wg.Wait()
}

func A(a, b chan any) {
	<-a
	fmt.Println("A()")
	time.Sleep(2 * time.Second)
	close(b)
}
func B(a, b chan any) {
	<-a
	fmt.Println("B()")
	time.Sleep(2 * time.Second)
	close(b)
}
func C(a, b chan any) {
	<-a
	fmt.Println("C()")
	time.Sleep(2 * time.Second)
	close(b)
}
func D(a chan any) {
	<-a
	fmt.Println("D()")
	time.Sleep(2 * time.Second)
}
