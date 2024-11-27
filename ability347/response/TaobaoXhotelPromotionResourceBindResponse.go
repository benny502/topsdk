package response

import (
    "topsdk/ability347/domain"
)

type TaobaoXhotelPromotionResourceBindResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        绑定结果
    */
    Result  domain.TaobaoXhotelPromotionResourceBindBindPromotionResourceResult `json:"result,omitempty" `
}
