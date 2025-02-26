package main

type BTree[T any] struct {
	value T
	Left  *BTree[T]
	Right *BTree[T]
}

func (b BTree[T]) GetData() T {
	return b.value
}

func (b *BTree[T]) SetData(value T) {
	b.value = value
}
