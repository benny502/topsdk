package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelRateDeleteResponse struct {

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
	Result domain.TaobaoXhotelRateDeleteResultSet `json:"result,omitempty" `
}
