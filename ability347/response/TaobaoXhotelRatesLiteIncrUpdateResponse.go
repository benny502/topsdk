package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelRatesLiteIncrUpdateResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   result
	*/
	Result domain.TaobaoXhotelRatesLiteIncrUpdateResultSet `json:"result,omitempty" `
}
