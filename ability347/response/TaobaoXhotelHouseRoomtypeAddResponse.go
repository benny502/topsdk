package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelHouseRoomtypeAddResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   房型信息
	*/
	Xroomtype domain.TaobaoXhotelHouseRoomtypeAddXRoomType `json:"xroomtype,omitempty" `
}
