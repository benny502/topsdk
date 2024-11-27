package response

import (
)

type TaobaoXhotelRateplanUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        修改的rp id
    */
    Rpid  int64 `json:"rpid,omitempty" `
}
