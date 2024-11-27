package response

import (
    "topsdk/ability347/domain"
)

type TaobaoXhotelPromotionCreateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        营销活动创建结果
    */
    Result  domain.TaobaoXhotelPromotionCreateResultSet `json:"result,omitempty" `
}
