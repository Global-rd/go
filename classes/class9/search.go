package main

type FirstFilterString FirstFilter[string]

type FirstFilter[K any] func(K) bool

func First[T any](values []T, filter FirstFilter[T]) (result T, value bool) {
	for _, value := range values {
		if ok := filter(value); ok {
			return value, ok
		}
	}

	return result, false
}

func FirstString(values []string, filter FirstFilterString) (string, bool) {
	for _, value := range values {
		if ok := filter(value); ok {
			return value, ok
		}
	}

	return "", false
}

func FirstInt(values []int, filter func(int) bool) (int, bool) {
	for _, value := range values {
		if ok := filter(value); ok {
			return value, ok
		}
	}

	return 0, false
}
