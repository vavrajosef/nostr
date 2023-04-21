package client

import "net/http"

func New() *Client {
	return &Client{
		client: http.DefaultClient,
	}
}

type Client struct {
	client *http.Client
}
