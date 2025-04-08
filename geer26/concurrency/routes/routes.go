package routes

import "net/http"

func Attachmhain(m *http.ServeMux) {
	m.HandleFunc("/", HelloWorld)
}

func AttachIntFetcher(m *http.ServeMux) {
	m.HandleFunc("/", IntFetcher)
}

func AttachPrimeFetcher(m *http.ServeMux) {
	m.HandleFunc("/", PrimeFetcher)
}

func AttachTimeoutFetcher(m *http.ServeMux) {
	m.HandleFunc("/", TimeoutFetcher)
}
