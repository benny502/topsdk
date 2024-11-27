package response

import (
    "topsdk/ability347/domain"
)

type TaobaoXhotelBnbownerAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        查询结果集
    */
    Result  domain.TaobaoXhotelBnbownerAddResultSet `json:"result,omitempty" `
}
