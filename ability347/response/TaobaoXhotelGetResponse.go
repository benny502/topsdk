package response

import (
    "topsdk/ability347/domain"
)

type TaobaoXhotelGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        查询得到的hotel
    */
    Xhotel  domain.TaobaoXhotelGetFirstResult `json:"xhotel,omitempty" `
}
