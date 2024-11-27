package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelBnbreviewAddResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   调用返回结果
	*/
	Result domain.TaobaoXhotelBnbreviewAddBnbResult `json:"result,omitempty" `
}
