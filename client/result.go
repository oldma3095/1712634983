package client

import (
	"context"
	"fmt"
	commonApi "github.com/oldma3095/1712634983/protos/common/api"
	"go_poker/model"
	"net/http"
	"strings"
)

func (client *Clients) PushResultToMaster(domain string, result model.Result, predictErr error) {
	req := &commonApi.ResultReq{}
	if predictErr != nil {
		req.Code = http.StatusInternalServerError
		req.Msg = predictErr.Error()
	} else {
		data := &commonApi.Result{
			Id:        uint64(result.ID),
			SaveTime:  result.SaveTime,
			VideoSize: result.VideoSize,
			Flag:      []string{},
			Banker:    []string{},
			Player1:   []string{},
			Player2:   []string{},
			Player3:   []string{},
			Other:     []string{},
		}
		if result.Video != "" {
			data.Video = fmt.Sprintf("%s/%s", domain, result.Video)
		}
		if result.Image != "" {
			data.Image = fmt.Sprintf("%s/%s", domain, result.Image)
		}
		if result.RawImage != "" {
			data.RawImage = fmt.Sprintf("%s/%s", domain, result.RawImage)
		}
		if result.Flag != "" {
			data.Flag = strings.Split(result.Flag, ",")
		}
		if result.Banker != "" {
			data.Banker = strings.Split(result.Banker, ",")
		}
		if result.Player1 != "" {
			data.Player1 = strings.Split(result.Player1, ",")
		}
		if result.Player2 != "" {
			data.Player2 = strings.Split(result.Player2, ",")
		}
		if result.Player3 != "" {
			data.Player3 = strings.Split(result.Player3, ",")
		}
		if result.Other != "" {
			data.Other = strings.Split(result.Other, ",")
		}
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
