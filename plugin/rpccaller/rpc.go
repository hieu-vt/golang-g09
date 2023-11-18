package rpccaller

import "github.com/go-resty/resty/v2"

type Rpc interface {
	GetServiceUrl() string
	GetRestyClient() *resty.Client
}
