package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelRoomtypeGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   查询得到的RoomType
	*/
	Xroomtype domain.TaobaoXhotelRoomtypeGetXRoomType `json:"xroomtype,omitempty" `
}
