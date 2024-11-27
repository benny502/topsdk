package response

import (
    "topsdk/ability347/domain"
)

type TaobaoXhotelXitemQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回结果
    */
    Result  domain.TaobaoXhotelXitemQueryResultSet `json:"result,omitempty" `
}
