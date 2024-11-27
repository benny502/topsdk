package response

import (
)

type TaobaoXhotelOrderUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回提示信息
    */
    Result  string `json:"result,omitempty" `
}
