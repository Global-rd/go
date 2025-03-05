package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"golang.org/x/exp/rand"
)

const (
	RANDCHARS     = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvxxyz"
	RANDOM_LENGHT = 10
)

func RandomSelect() string {
	buff := strings.Builder{}
	for {
		if buff.Len() >= RANDOM_LENGHT {
			break
		}
		randomIndex := rand.Intn(len(RANDCHARS))
		randomCharacter := []byte(RANDCHARS)[randomIndex]
		buff.WriteByte(randomCharacter)
	}
	return buff.String()
}

func main() {
	apiserver := &http.Server{
		Addr: ":5000",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", RandomSelect())
	})
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()
		defer func() {
			log.Println("API shut down...")
		}()
		if err := apiserver.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("API shutdown error: %v", err)
		}
	}()
	log.Println("API is listening on port 5000...")
	if err := apiserver.ListenAndServe(); err != nil {
		log.Fatalln("Error at serving API endpoints: ", err)
	}
}
