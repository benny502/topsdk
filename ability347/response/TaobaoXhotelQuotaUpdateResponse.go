package response

import (
)

type TaobaoXhotelQuotaUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        更新失败补充描述消息
    */
    WarnMessage  string `json:"warn_message,omitempty" `
    /*
        errorCode
    */
    BizErrorCode  string `json:"biz_error_code,omitempty" `
    /*
        更新失败错误信息
    */
    BizErrorMsg  string `json:"biz_error_msg,omitempty" `
}
