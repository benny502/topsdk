package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelDeleteResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   删除结果
	*/
	Result domain.TaobaoXhotelDeleteResultSet `json:"result,omitempty" `
}
