package main

import (
	"fmt"
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

	wg.Add(1)
	go func() {

		wg.Wait()

		time.Sleep(5 * time.Second)
		fmt.Println("hoge")
		if err := client.AuthPocket(addr, consumerKey); err != nil {
			logger.Panic(err)
		}
	}()
	srv.Serve(addr, &wg)
}
