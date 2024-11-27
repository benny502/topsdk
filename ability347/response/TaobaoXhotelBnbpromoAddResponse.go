package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelBnbpromoAddResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   营销添加返回对象
	*/
	Module domain.TaobaoXhotelBnbpromoAddPromoCode `json:"module,omitempty" `
}
