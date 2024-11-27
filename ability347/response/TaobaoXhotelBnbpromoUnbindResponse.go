package response

import (
    "topsdk/ability347/domain"
)

type TaobaoXhotelBnbpromoUnbindResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        营销解绑返回对象
    */
    Module  domain.TaobaoXhotelBnbpromoUnbindPromoBindResult `json:"module,omitempty" `
}
