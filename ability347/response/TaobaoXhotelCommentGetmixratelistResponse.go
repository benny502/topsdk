package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelCommentGetmixratelistResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   接口返回model
	*/
	Result domain.TaobaoXhotelCommentGetmixratelistResult `json:"result,omitempty" `
}
