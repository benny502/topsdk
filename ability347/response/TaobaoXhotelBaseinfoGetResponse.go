package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelBaseinfoGetResponse struct {

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
	Result domain.TaobaoXhotelBaseinfoGetResultSet `json:"result,omitempty" `
}
