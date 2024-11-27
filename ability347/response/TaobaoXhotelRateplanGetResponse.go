package response

import (
    "topsdk/ability347/domain"
)

type TaobaoXhotelRateplanGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        rateplan
    */
    Rateplan  domain.TaobaoXhotelRateplanGetRatePlan `json:"rateplan,omitempty" `
}
