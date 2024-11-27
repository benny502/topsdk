package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelPromotionQueryResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   查询结果
	*/
	Result domain.TaobaoXhotelPromotionQueryQueryPromotionResult `json:"result,omitempty" `
}
