package pkg

import (
	"context"
	"sync"
	"time"
)

func Procces1(i int) int {
	time.Sleep(10 * time.Millisecond)

	count := 0

	for range 10000000 {
		count++
	}

	return i * i
}

func FanIn(ctx context.Context, chans []chan int, f func(int) int) <-chan int {
	out := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for _, ch := range chans {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case <-ctx.Done():
						return
					case v, ok := <-ch:
						if !ok {
							return
						}
						out <- f(v)
					}
				}
			}()
		}

		wg.Wait()
		close(out)

	}()

	return out

}
