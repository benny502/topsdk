package domain


type TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceResult struct {
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
        解绑详情，部分成功的时候有值     */
    Data  *TaobaoXhotelPromotionResourceUnbindBindData `json:"data,omitempty" `

}

func (s *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceResult) SetErrorCode(v string) *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceResult) SetErrorMsg(v string) *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceResult) SetSuccess(v bool) *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceResult {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceResult) SetTraceId(v string) *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceResult {
    s.TraceId = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceResult) SetData(v TaobaoXhotelPromotionResourceUnbindBindData) *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceResult {
    s.Data = &v
    return s
}
