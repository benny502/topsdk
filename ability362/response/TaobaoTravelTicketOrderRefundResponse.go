package response

import (
)

type TaobaoTravelTicketOrderRefundResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        系统自动生成
    */
    IsSuccess  bool `json:"is_success,omitempty" `
}
