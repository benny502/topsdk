package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelBnbroomtypeAddResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   房源信息
	*/
	Xroomtype domain.TaobaoXhotelBnbroomtypeAddXRoomType `json:"xroomtype,omitempty" `
}
