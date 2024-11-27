package response

import (
)

type TaobaoXhotelBnbhouseDeleteResponse struct {

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
        是否出错
    */
    ErrorMsg  string `json:"error_msg,omitempty" `
}
