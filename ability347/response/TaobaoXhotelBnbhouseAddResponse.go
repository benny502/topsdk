package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelBnbhouseAddResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   系统自动生成
	*/
	Results []domain.TaobaoXhotelBnbhouseAddXHotel `json:"results,omitempty" `
}
