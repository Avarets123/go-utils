package pkg

import "sync"

type HandlerF = func(int) int

func PoolWokers(in chan int, workersCount int, f HandlerF) <-chan int {

	out := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		for range workersCount {
			wg.Add(1)
			go func() {
				worker(in, out, f)
				wg.Done()
			}()
		}

		wg.Wait()
		close(out)

	}()

	return out
}

func worker(in <-chan int, out chan<- int, f HandlerF) {
	for v := range in {
		out <- f(v)
	}
}
