package client

import (
	"context"
	"github.com/oldma3095/1712634983/cache"
	commonApi "github.com/oldma3095/1712634983/protos/common/api"
	"net/http"
)

func (client *Clients) PushResultToMaster(data cache.NiuNiuResult, predictErr error) {
	req := &commonApi.ResultReq{}
	if predictErr != nil {
		req.Code = http.StatusInternalServerError
		req.Msg = predictErr.Error()
	} else {
		req.Data = &commonApi.Result{
			Id:            data.Id,
			SaveTime:      data.SaveTime,
			Image:         data.Image,
			RawImage:      data.RawImage,
			Video:         data.Video,
			VideoSize:     data.VideoSize,
			Flag:          data.Flag,
			Banker:        data.Banker,
			Player1:       data.Player1,
			Player2:       data.Player2,
			Player3:       data.Player3,
			Other:         data.Other,
			SimpleFlag:    data.SimpleFlag,
			SimpleBanker:  data.SimpleBanker,
			SimplePlayer1: data.SimplePlayer1,
			SimplePlayer2: data.SimplePlayer2,
			SimplePlayer3: data.SimplePlayer3,
			SimpleOther:   data.SimpleOther,
		}
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
