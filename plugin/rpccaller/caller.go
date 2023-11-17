package rpccaller

import (
	"flag"
	"github.com/go-resty/resty/v2"
)

type rpcCaller struct {
	prefix     string
	client     *resty.Client
	serviceURL string
}

func NewRpcCaller(prefix string) *rpcCaller {
	return &rpcCaller{prefix: prefix}
}

func (c *rpcCaller) GetPrefix() string {
	return c.prefix
}

func (c *rpcCaller) Get() interface{} {
	return c
}

func (c *rpcCaller) Name() string {
	return c.prefix
}

func (c *rpcCaller) InitFlags() {
	flag.StringVar(&c.serviceURL, "item-service-url", "http://localhost:3001", "URL of item service")
}

func (c *rpcCaller) Configure() error {
	c.client = resty.New()
	return nil
}

func (c *rpcCaller) Run() error {
	return nil
}

func (c *rpcCaller) Stop() <-chan bool {
	ch := make(chan bool)
	go func() {
		ch <- true
	}()
	return ch
}
