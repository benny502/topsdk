package domain


type TaobaoXhotelBnbreviewReplyBnbResult struct {
    /*
        状态，成功true，失败false     */
    Success  *bool `json:"success,omitempty" `

    /*
        响应吗     */
    ResultCode  *string `json:"result_code,omitempty" `

    /*
        响应信息     */
    ResultMsg  *string `json:"result_msg,omitempty" `

}

func (s *TaobaoXhotelBnbreviewReplyBnbResult) SetSuccess(v bool) *TaobaoXhotelBnbreviewReplyBnbResult {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelBnbreviewReplyBnbResult) SetResultCode(v string) *TaobaoXhotelBnbreviewReplyBnbResult {
    s.ResultCode = &v
    return s
}
func (s *TaobaoXhotelBnbreviewReplyBnbResult) SetResultMsg(v string) *TaobaoXhotelBnbreviewReplyBnbResult {
    s.ResultMsg = &v
    return s
}
