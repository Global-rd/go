package main

import (
	"fmt"
	"full-webservice/config"
	"os"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
