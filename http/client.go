package http

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Client struct {
	client *http.Client
	logger *log.Logger
}

func (c *Client) AuthPocket(addr, consumerKey string) error {
	req, err := http.NewRequest(http.MethodGet, "http://localhost"+addr+"/auth", nil)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var r io.Reader = res.Body
	r = io.TeeReader(r, os.Stderr)

	ioutil.ReadAll(r)

	return nil
}

func NewClient(logger *log.Logger) *Client {
	return &Client{
		client: &http.Client{},
		logger: logger,
	}
}
