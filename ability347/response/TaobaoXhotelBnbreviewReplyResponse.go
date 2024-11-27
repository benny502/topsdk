package response

import (
    "topsdk/ability347/domain"
)

type TaobaoXhotelBnbreviewReplyResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        调用返回结果
    */
    Result  domain.TaobaoXhotelBnbreviewReplyBnbResult `json:"result,omitempty" `
}
