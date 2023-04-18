package client

type Options struct{}

type Client struct{}

func New(opt *Options) *Client {
	return &Client{}
}
