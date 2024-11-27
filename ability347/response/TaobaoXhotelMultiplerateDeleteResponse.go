package response

import (
    "topsdk/ability347/domain"
)

type TaobaoXhotelMultiplerateDeleteResponse struct {

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
    Result  domain.TaobaoXhotelMultiplerateDeleteResultSet `json:"result,omitempty" `
}
