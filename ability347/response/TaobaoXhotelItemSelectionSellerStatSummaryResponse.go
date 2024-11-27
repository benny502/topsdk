package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelItemSelectionSellerStatSummaryResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回参数
	*/
	Result domain.TaobaoXhotelItemSelectionSellerStatSummaryHsfResult `json:"result,omitempty" `
}
