package pkg

import "context"

func FanOut(ctx context.Context, in chan int, chansCount int) []chan int {
	chans := make([]chan int, chansCount)

	for i := range chansCount {
		chans[i] = pipeline(ctx, in)
	}

	return chans

}

func pipeline(ctx context.Context, in chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}

	}()

	return out

}
