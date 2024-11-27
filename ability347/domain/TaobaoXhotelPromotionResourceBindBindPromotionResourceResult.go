package domain


type TaobaoXhotelPromotionResourceBindBindPromotionResourceResult struct {
    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误描述     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        是否成功，True or False，全部失败返回失败，部分/全部成功返回成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        日志id     */
    TraceId  *string `json:"trace_id,omitempty" `

    /*
        绑定详情     */
    Data  *TaobaoXhotelPromotionResourceBindBindData `json:"data,omitempty" `

}

func (s *TaobaoXhotelPromotionResourceBindBindPromotionResourceResult) SetErrorCode(v string) *TaobaoXhotelPromotionResourceBindBindPromotionResourceResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindPromotionResourceResult) SetErrorMsg(v string) *TaobaoXhotelPromotionResourceBindBindPromotionResourceResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindPromotionResourceResult) SetSuccess(v bool) *TaobaoXhotelPromotionResourceBindBindPromotionResourceResult {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindPromotionResourceResult) SetTraceId(v string) *TaobaoXhotelPromotionResourceBindBindPromotionResourceResult {
    s.TraceId = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindPromotionResourceResult) SetData(v TaobaoXhotelPromotionResourceBindBindData) *TaobaoXhotelPromotionResourceBindBindPromotionResourceResult {
    s.Data = &v
    return s
}
