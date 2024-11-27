package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelHouseAddResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   酒店信息
	*/
	Xhotel domain.TaobaoXhotelHouseAddXHotel `json:"xhotel,omitempty" `
}
