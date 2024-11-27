package response

import (
)

type TaobaoXhotelRateplanAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        生成的rp id
    */
    Rpid  int64 `json:"rpid,omitempty" `
}
