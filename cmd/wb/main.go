package main

import (
	"log"
	"os"

	"github.com/himetani/workbook/cmd"
)

var (
	logger = log.New(os.Stdout, "Info: ", log.LstdFlags)
	addr   = ":8080"
)

func main() {
	cmd.Execute()
}

/*
func main() {
	var (
		svrStartUp  sync.WaitGroup
		authCode    sync.WaitGroup
		consumerKey = os.Getenv("WB_CONSUMER_KEY")
		addr        = ":8080"
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
}
*/
