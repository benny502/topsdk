package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelRoomtypeConflictDataResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   查询结果集
	*/
	Result domain.TaobaoXhotelRoomtypeConflictDataResultSet `json:"result,omitempty" `
}
