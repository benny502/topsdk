package response

import (
)

type TaobaoXhotelRoomUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        gid酒店商品id
    */
    Gid  int64 `json:"gid,omitempty" `
}
