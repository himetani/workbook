package cmd

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/himetani/workbook/http"
	"github.com/himetani/workbook/pocket"
)

func RunAuth(consumerKey string, logger *log.Logger) error {
	var (
		svrStartUp sync.WaitGroup
		authCode   sync.WaitGroup
		addr       = ":8080"
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
			logger.Printf("%s\n", err.Error())
			return
		}

		fmt.Println("=> Access to https://getpocket.com/auth/authorize?request_token=" + client.RequestCode + "&redirect_uri=" + client.RedirectURL)

		authCode.Wait()

		cancel()
	}()

	srv.Serve(addr, &svrStartUp, &authCode, ctx)
	fmt.Printf("=> Username: %s, AccessToken %s\n", client.Username, client.AccessToken)
	authCode.Wait()

	return nil
}
