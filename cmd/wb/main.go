package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/himetani/workbook/http"
)

var (
	logger = log.New(os.Stdout, "Info: ", log.LstdFlags)
	addr   = ":8080"
)

func main() {
	var (
		wg          sync.WaitGroup
		consumerKey = "dummy"
		addr        = ":8080"
	)

	srv := http.NewServer(logger)
	client := http.NewClient(logger)
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		wg.Wait()

		time.Sleep(3 * time.Second)

		if err := client.AuthPocket(addr, consumerKey); err != nil {
			logger.Panic(err)
		}
		cancel()
	}()

	srv.Serve(addr, &wg, ctx)
}
