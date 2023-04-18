package client

type ClientOptions struct{}

type Client struct{}

func New(opt *ClientOptions) *Client {
	return &Client{}
}
