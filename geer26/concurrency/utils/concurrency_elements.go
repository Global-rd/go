package utils

import "context"

func Producer[T any, K any](ctx context.Context, done <-chan K, fn func() T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			default:
				out <- fn()
			}
		}
	}()
	return out
}

func Limiter[T any, K any](ctx context.Context, done <-chan K, in <-chan T, limit int) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for i := 0; i < limit; i++ {
			select {
			case <-done:
				return
			case out <- <-in:
			}
		}
	}()
	return out
}

func Filter[T any, K any](ctx context.Context, done <-chan K, in <-chan T, filterfunc func(T) bool) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case v := <-in:
				if filterfunc(v) {
					out <- v
				}
			}
		}
	}()
	return out
}
