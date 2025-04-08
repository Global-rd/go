package routes

import "net/http"

func Attachmhain(m *http.ServeMux) {
	m.HandleFunc("/", HelloWorld)
}

func AttachIntFetcher(m *http.ServeMux, ch <-chan int) {
	m.HandleFunc("/", IntFetcher(ch))
}

func AttachPrimeFetcher(m *http.ServeMux, ch <-chan int) {
	m.HandleFunc("/", PrimeFetcher(ch))
}

func AttachTimeoutFetcher(m *http.ServeMux) {
	m.HandleFunc("/", TimeoutFetcher)
}
