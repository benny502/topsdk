package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelGetEntityByTagResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   出参
	*/
	TagQueryResult domain.TaobaoXhotelGetEntityByTagTagQueryResult `json:"tag_query_result,omitempty" `
}
