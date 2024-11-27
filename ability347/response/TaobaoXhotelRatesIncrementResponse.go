package response

import (
)

type TaobaoXhotelRatesIncrementResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        gid和rpid组合数组
gid_rpid
    */
    GidAndRpids  []string `json:"gid_and_rpids,omitempty" `
}
