package response

import (
)

type TaobaoTravelTicketOrderVerifyResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        成功状态true or false
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
