package channels

import "sync"

func FanInChannels[T any](inChannels []chan T) chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	for _, in := range inChannels {
		wg.Add(1)
		go func(in <-chan T) {
			defer wg.Done()
			for v := range in {
				out <- v
			}
		}(in)
	}
	go func() { wg.Wait(); close(out) }()

	return out
}
