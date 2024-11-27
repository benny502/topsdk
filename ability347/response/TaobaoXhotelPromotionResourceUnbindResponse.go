package response

import (
    "topsdk/ability347/domain"
)

type TaobaoXhotelPromotionResourceUnbindResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        解绑结果
    */
    Result  domain.TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceResult `json:"result,omitempty" `
}
