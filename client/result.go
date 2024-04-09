package client

import (
	"context"
	commonApi "github.com/oldma3095/1712634983/protos/common/api"
	"net/http"
)

func (client *Clients) PushResultToMaster(data *commonApi.Result, predictErr error) {
	req := &commonApi.ResultReq{}
	if predictErr != nil {
		req.Code = http.StatusInternalServerError
		req.Msg = predictErr.Error()
	} else {
		req.Data = data
	}

	res, err := client.apiServiceClient.Result(context.Background(), req)
	if err != nil {
		client.log.Error(err.Error())
		return
	}
	if res.Code != 0 && res.Code != http.StatusOK {
		client.log.Error(res.Msg)
	}
}
