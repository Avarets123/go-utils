package ch7

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ChPractic() {

	ch := make(chan int, 1)

	ch <- 10
	close(ch)

	fmt.Println(<-ch)

	_, ok := <-ch
	fmt.Println("is closed", !ok)

	for i := 5; i > 0; i-- {
		fmt.Println(<-ch)
	}

	ch2 := make(chan bool, 3)

	ch2 <- true
	ch2 <- false
	ch2 <- true
	// close(ch2)

	for i := range ch2 {
		if !i {
			break
		}
		fmt.Println(i)
	}

}

func SelectPractic() {
	wg := sync.WaitGroup{}
	nch := make(chan int)
	ech := make(chan bool)
	defer close(nch)
	defer close(ech)
	wg.Add(1)
	go func() {
		nch <- 12
		gen(0, 14, nch, ech)
		wg.Done()
	}()

	for v := range nch {
		fmt.Println(v)
		if v == 10 {
			ech <- true
			break
		}
	}

	wg.Wait()

}

func gen(min, max int, numbersCh chan<- int, end chan bool) {
	for {
		select {
		case numbersCh <- rand.Intn(max-min) + min:

		case <-end:
			fmt.Println("ENded")
			return

		case <-time.After(4 * time.Second):
			{
				fmt.Println("TIME AFTER")
				return
			}
		}
	}

}

func TimeoutGorutine1() {
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "OK"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}

	ch2 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "OK"
	}()

	select {
	case res := <-ch2:
		fmt.Println(res)
	case <-time.After(4 * time.Second):
		fmt.Println("Timeout ch2")
	}

}

func TimeoutGorutine(d time.Duration, result chan<- bool) {
	temp := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		defer close(temp)
	}()

	select {
	case <-temp:
		result <- false
	case <-time.After(d):
		result <- true
	}

}

func ChWithDefault() {
	count := 10
	ch := make(chan int, 5)
	defer close(ch)

	for i := range count {

		select {
		case ch <- i * i:
			fmt.Println("Process ", i)
		default:
			fmt.Println("Channel unavailable")
		}

	}

	for {
		select {
		case res := <-ch:
			fmt.Println("Result ", res)
		default:
			fmt.Println("Empty")
			return
		}
	}

}

func NilChannel(wg *sync.WaitGroup) {

	ch := make(chan int)
	sum := 0

	t := time.NewTimer(time.Second)

	send := func(c chan int) {
		for {
			c <- rand.Intn(10)
		}
	}

	add := func(ch chan int) {
		for {

			select {
			case input := <-ch:
				sum += input
			case <-t.C:
				fmt.Println("Timer end!")
				fmt.Println(sum)
				ch = nil
				wg.Done()
			}
		}
	}

	go add(ch)
	go send(ch)

}
