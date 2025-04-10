package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"patterns/channels"
	"patterns/client"
	"patterns/config"
	"patterns/model"
	"patterns/server"
)

func parseFlags() (bool, int) {
	var isServer bool
	var requestCount int
	flag.BoolVar(&isServer, "server", false, "run as server")
	flag.IntVar(&requestCount, "requests", 0, "number of requests to send")
	flag.Parse()
	return isServer, requestCount
}

func runServer(logger *slog.Logger, cfg *config.ServerCfg) {
	router := server.NewDummyRouter(logger)
	dummyServer, err := server.NewServer(router, cfg, server.WithLogger(logger))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	if err = dummyServer.Start(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func sendRequests(logger *slog.Logger, count int, cfg *config.ServerCfg) []chan model.Result {
	requests := client.NewRequest(logger, cfg)

	resultChannels := make([]chan model.Result, 0)
	for range count {
		resultChan := make(chan model.Result)
		resultChannels = append(resultChannels, resultChan)
		go requests.CallResourceWithTimeout(resultChan)
	}
	return resultChannels
}

func runClient(logger *slog.Logger, count int, cfg *config.ServerCfg) {
	resultChannels := sendRequests(logger, count, cfg)
	mergedChannel := channels.FanInChannels[model.Result](resultChannels)

	for result := range mergedChannel {
		fmt.Printf("call to resource %s ended with result: %s\n", result.ResourceId, result.Status)
	}
}

func main() {
	logger := slog.Default()
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	isServer, requestCount := parseFlags()

	if isServer {
		runServer(logger, cfg)
	} else {
		if requestCount <= 0 || requestCount > config.MaxRequests {
			logger.Error(fmt.Sprintf("invalid request count: %d, must be between 1 and %d ", requestCount, config.MaxRequests))
			os.Exit(1)
		}

		runClient(logger, requestCount, cfg)
	}
}
