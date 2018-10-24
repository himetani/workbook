package pocket

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	AccessToken string
	client      *http.Client
	consumerKey string
	RedirectURL string
	RequestCode string
	logger      *log.Logger
	Username    string
}

func newRequest(path, jsonstr string) (*http.Request, error) {
	req, err := http.NewRequest("POST", "https://getpocket.com/"+path, bytes.NewBuffer([]byte(jsonstr)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Accept", "application/json")

	return req, nil
}

func (c *Client) GetRequestCode() error {
	path := "/v3/oauth/request"
	jsonstr := "{\"consumer_key\": \"" + c.consumerKey + "\", \"redirect_uri\": \"" + c.RedirectURL + "\"}"
	c.logger.Printf("Get Request Code, path: %s, body: %s", path, jsonstr)

	req, err := newRequest(path, jsonstr)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var r io.Reader = resp.Body
	//r = io.TeeReader(r, os.Stderr)
	//fmt.Println("")

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code was not 200. got=%d", resp.StatusCode)
	}

	var requestCodeResponse requestCodeResponse
	if err := json.NewDecoder(r).Decode(&requestCodeResponse); err != nil {
		return err
	}

	c.RequestCode = requestCodeResponse.Code
	return nil
}

func (c *Client) GetAccessToken() error {
	c.logger.Printf("Get AccessToken, consumerKey: %s, requestCode %s\n", c.consumerKey, c.RequestCode)

	if c.consumerKey == "" || c.RequestCode == "" {
		return errors.New("consumerKey or requestCode is empty")
	}

	jsonstr := "{\"consumer_key\": \"" + c.consumerKey + "\", \"code\": \"" + c.RequestCode + "\"}"

	req, err := newRequest("/v3/oauth/authorize", jsonstr)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code was not 200. got=%d", resp.StatusCode)
	}

	var r io.Reader = resp.Body
	//r = io.TeeReader(r, os.Stderr)

	var accessTokenResponse accessTokenResponse
	if err := json.NewDecoder(r).Decode(&accessTokenResponse); err != nil {
		return err
	}

	c.AccessToken = accessTokenResponse.AccessToken
	c.Username = accessTokenResponse.Username
	return nil
}

func NewClient(redirectURL, consumerKey string, logger *log.Logger) *Client {
	return &Client{
		client:      &http.Client{},
		consumerKey: consumerKey,
		RedirectURL: redirectURL,
		logger:      logger,
	}
}

type requestCodeResponse struct {
	Code string `json:"code"`
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}
