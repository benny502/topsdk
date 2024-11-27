package response

import (
)

type TaobaoXhotelBnbroomtypeDeleteResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否出错
    */
    Error  bool `json:"error,omitempty" `
    /*
        错误信息
    */
    ErrorMsg  string `json:"error_msg,omitempty" `
}
