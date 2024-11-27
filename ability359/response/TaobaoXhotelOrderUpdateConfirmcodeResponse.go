package response

import (
)

type TaobaoXhotelOrderUpdateConfirmcodeResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        错误描述
    */
    ErrorMsg  string `json:"error_msg,omitempty" `
    /*
        操作结果,成功返回success
    */
    Result  string `json:"result,omitempty" `
    /*
        是否操作成功
    */
    Error  bool `json:"error,omitempty" `
}
