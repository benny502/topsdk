package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelRoomGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   房间信息
	*/
	Room domain.TaobaoXhotelRoomGetFirstResult `json:"room,omitempty" `
}
