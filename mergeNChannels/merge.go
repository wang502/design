package mergeNChannels

import (
	"sync"
)

func asChan(vs ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range vs {
			out <- v
		}
		close(out)
	}()
	return out
}

func merge(chans ...chan int) <-chan int {
	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chans))
		for _, c := range chans {
			go func(c chan int) {
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}

		wg.Wait()
		close(out)
	}()
	return out
}
