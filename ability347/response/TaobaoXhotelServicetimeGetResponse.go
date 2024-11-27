package response

import (
    "topsdk/ability347/domain"
)

type TaobaoXhotelServicetimeGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        result
    */
    Result  domain.TaobaoXhotelServicetimeGetResultSet `json:"result,omitempty" `
}
