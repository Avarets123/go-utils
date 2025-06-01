package ch7

import (
	"context"
	"fmt"
	"time"
)

func UseCtx() {
	F1(3 * time.Second)
	F2(3*time.Second, 2*time.Second)
	F3(2*time.Second, 1*time.Second)

}

func F1(t time.Duration) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("F1() Done", c1.Err())
		return
	case r := <-time.After(t):
		fmt.Println("F1(): ", r)
	}
}

func F2(t, t2 time.Duration) {
	c2, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("F2() : ", c2.Err())
		return
	case r := <-time.After(t2):
		fmt.Println("f2(): ", r)
	}

}

func F3(t, t2 time.Duration) {
	deadline := time.Now().Add(t)
	c3, cancel := context.WithDeadline(context.TODO(), deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("F3(): ", c3.Err())
	case r := <-time.After(t2):
		fmt.Println("F3(): ", r)
	}

}
