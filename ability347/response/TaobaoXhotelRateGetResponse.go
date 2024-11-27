package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelRateGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   rate
	*/
	Rate domain.TaobaoXhotelRateGetRate `json:"rate,omitempty" `
}
