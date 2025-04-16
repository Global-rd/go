package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/utils"
	"net/http"
	"time"
)

type Uselessfact struct {
	UselessFact string `json:"useless_fact"`
}

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	uselessfact := utils.Getuselessfact(context.Background())

	select {
	case fetched := <-uselessfact:
		log.Println("Time elapsed: ", time.Since(start))
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fetched)
		return
	case <-time.After(time.Millisecond * 30):
		log.Println("Time elapsed: ", time.Since(start))
		w.WriteHeader(http.StatusRequestTimeout)
		fmt.Fprintf(w, "request timed out")
	}

}
