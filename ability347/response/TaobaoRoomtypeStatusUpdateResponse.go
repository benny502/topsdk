package response

import (
)

type TaobaoRoomtypeStatusUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        错误
    */
    ErrorMsg  string `json:"error_msg,omitempty" `
}
