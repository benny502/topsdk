package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelXitemDeleteResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果
	*/
	Result domain.TaobaoXhotelXitemDeleteResultSet `json:"result,omitempty" `
}
