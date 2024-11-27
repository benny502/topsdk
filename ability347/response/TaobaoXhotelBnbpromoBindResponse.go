package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelBnbpromoBindResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   营销绑定返回对象
	*/
	Module domain.TaobaoXhotelBnbpromoBindPromoBindResult `json:"module,omitempty" `
}
