package response

import (
)

type TaobaoXhotelRateAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        酒店商品id-酒店rpID
    */
    GidAndRpid  string `json:"gid_and_rpid,omitempty" `
    /*
        results
    */
    Results  []string `json:"results,omitempty" `
    /*
        warnMessage
    */
    WarnMessage  string `json:"warn_message,omitempty" `
}
