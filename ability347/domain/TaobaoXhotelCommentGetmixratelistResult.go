package domain


type TaobaoXhotelCommentGetmixratelistResult struct {
    /*
        状态码     */
    HttpStatusCode  *int64 `json:"http_status_code,omitempty" `

    /*
        模型     */
    Model  *TaobaoXhotelCommentGetmixratelistGetMixRateListResult `json:"model,omitempty" `

    /*
        msgCode     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        msgInfo     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoXhotelCommentGetmixratelistResult) SetHttpStatusCode(v int64) *TaobaoXhotelCommentGetmixratelistResult {
    s.HttpStatusCode = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistResult) SetModel(v TaobaoXhotelCommentGetmixratelistGetMixRateListResult) *TaobaoXhotelCommentGetmixratelistResult {
    s.Model = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistResult) SetMsgCode(v string) *TaobaoXhotelCommentGetmixratelistResult {
    s.MsgCode = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistResult) SetMsgInfo(v string) *TaobaoXhotelCommentGetmixratelistResult {
    s.MsgInfo = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistResult) SetSuccess(v bool) *TaobaoXhotelCommentGetmixratelistResult {
    s.Success = &v
    return s
}
