package routes

import (
	"github.com/go-chi/chi/v5"
)

func Attachmhain(m *chi.Mux) {
	m.HandleFunc("/", HelloWorld)
}

func AttachIntFetcher(m *chi.Mux, ch <-chan int) {
	m.HandleFunc("/", IntFetcher(ch))
}

func AttachPrimeFetcher(m *chi.Mux, ch <-chan int) {
	m.HandleFunc("/", PrimeFetcher(ch))
}

func AttachTimeoutFetcher(m *chi.Mux) {
	m.HandleFunc("/", TimeoutFetcher)
}
