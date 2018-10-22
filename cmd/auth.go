package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/himetani/workbook/http"
	"github.com/himetani/workbook/pocket"
)

func RunAuth(consumerKey string) error {
	var (
		svrStartUp sync.WaitGroup
		authCode   sync.WaitGroup
		addr       = ":8080"
		logger     = log.New(os.Stdout, "Info: ", log.LstdFlags)
	)

	client := pocket.NewClient("http://localhost:8080/pocket/redirected", consumerKey, logger)
	srv := http.NewServer(client, logger)
	ctx, cancel := context.WithCancel(context.Background())

	svrStartUp.Add(1)
	authCode.Add(1)
	go func() {
		svrStartUp.Wait()

		time.Sleep(1 * time.Second)

		if err := client.GetRequestCode(); err != nil {
			logger.Panic(err)
		}

		fmt.Println("")
		fmt.Println("Access to https://getpocket.com/auth/authorize?request_token=" + client.RequestCode + "&redirect_uri=" + client.RedirectURL)

		authCode.Wait()

		cancel()
	}()

	srv.Serve(addr, &svrStartUp, &authCode, ctx)
	authCode.Wait()

	return nil
}
