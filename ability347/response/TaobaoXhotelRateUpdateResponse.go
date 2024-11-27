package response

import (
)

type TaobaoXhotelRateUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        酒店商品ID-酒店RPid
    */
    GidAndRpid  string `json:"gid_and_rpid,omitempty" `
}
