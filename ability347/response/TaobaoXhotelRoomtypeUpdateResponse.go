package response

import (
    "topsdk/ability347/domain"
)

type TaobaoXhotelRoomtypeUpdateResponse struct {

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
    Xroomtype  domain.TaobaoXhotelRoomtypeUpdateXRoomType `json:"xroomtype,omitempty" `
}
