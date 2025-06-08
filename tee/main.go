package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	TeeDynamic()
}

func TeeDynamic() {
	in := make(chan int)

	go func() {
		for i := range 100 {
			in <- i
		}
		defer close(in)
	}()

	chs := TeeDynamicProccess(context.Background(), in, 5)

	wg := sync.WaitGroup{}
	wg.Add(len(chs))
	for i, ch := range chs {
		go func() {
			defer wg.Done()
			for v := range ch {
				fmt.Println("Ch", i, ":", v)
				time.Sleep(time.Millisecond * 300)
			}
		}()
	}

	wg.Wait()

}

func TeeDynamicProccess(ctx context.Context, in chan int, chsCount int) []chan int {

	chs := make([]chan int, chsCount)

	for i := range chsCount {
		chs[i] = make(chan int)
	}

	go func() {
		for _, ch := range chs {
			defer close(ch)
		}

		wgRoot := sync.WaitGroup{}
		wgRoot.Add(1)

		go func() {
			defer wgRoot.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-in:
					if !ok {
						return
					}

					wgChild := sync.WaitGroup{}
					wgChild.Add(len(chs))
					for _, ch := range chs {
						go func() {
							defer wgChild.Done()
							ch <- v
						}()

					}
					wgChild.Wait()

				}

			}
		}()

		wgRoot.Wait()

	}()

	return chs

}

func Tee() {

	in := make(chan int)
	go func() {
		defer close(in)
		for i := range 100 {
			in <- i
		}
	}()

	ctx := context.Background()

	ch1, ch2 := TeeProccess(ctx, in)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range ch1 {
			fmt.Println("Ch1: ", v)
			time.Sleep(time.Millisecond * 300)

		}
	}()

	go func() {
		defer wg.Done()
		for v := range ch2 {
			fmt.Println("Ch2: ", v)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	wg.Wait()

}

func TeeProccess(ctx context.Context, in <-chan int) (<-chan int, <-chan int) {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {

		defer close(ch1)
		defer close(ch2)

		for {

			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				wg := sync.WaitGroup{}
				wg.Add(2)
				go func() {
					defer wg.Done()
					ch1 <- v
				}()

				go func() {
					defer wg.Done()
					ch2 <- v
				}()
				wg.Wait()

			case <-ctx.Done():
				return
			}

		}

	}()

	return ch1, ch2

}

func TeeOldVersion() {

	in := make(chan int)

	go func() {
		defer close(in)
		for i := range 10 {
			in <- i
		}
	}()

	ch1, ch2 := teeProccessOld(context.Background(), in)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range ch1 {
			fmt.Println("Ch1: ", v)
		}

	}()

	go func() {
		defer wg.Done()
		for v := range ch2 {
			fmt.Println("Ch2: ", v)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	wg.Wait()

}

func teeProccessOld(ctx context.Context, in chan int) (chan int, chan int) {

	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		defer close(ch1)
		defer close(ch2)

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				res := v
				var ch1, ch2 = ch1, ch2
				for range 2 {
					select {
					case ch1 <- res:
						ch1 = nil
					case ch2 <- res:
						ch2 = nil

					}
				}

			}

		}
	}()

	return ch1, ch2

}

func Proccess(i int) int {
	return i * i
}
