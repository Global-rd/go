package main

type Cache[T any] interface {
	GetData() T
	SetData(T)
}
