package rpc

import (
	"context"
	"errors"
	"fmt"
	"g09-to-do-list/common"
	"g09-to-do-list/plugin/rpccaller"
	goservice "github.com/200Lab-Education/go-sdk"
)

type itemService struct {
	appService goservice.ServiceContext
}

func NewItemService(appService goservice.ServiceContext) *itemService {
	return &itemService{appService: appService}
}

func (s *itemService) GetItemLikes(ctx context.Context, ids []int) (map[int]int, error) {
	type requestBody struct {
		Ids []int `json:"ids"`
	}

	logger := s.appService.Logger("rpc-call-api")

	rpc := s.appService.MustGet(common.PluginRpcApi).(rpccaller.Rpc)
	client := rpc.GetRestyClient()
	var response struct {
		Data map[int]int `json:"data"`
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody{Ids: ids}).
		SetResult(&response).
		Post(fmt.Sprintf("%s/%s", rpc.GetServiceUrl(), "v1/rpc/get_item_likes"))

	if err != nil {
		logger.Errorln(err)
		return nil, err
	}

	if !resp.IsSuccess() {
		//log.Println(resp.RawResponse)
		logger.Errorln(resp.RawResponse)
		return nil, errors.New("cannot call api get item likes")
	}

	return response.Data, nil

}
