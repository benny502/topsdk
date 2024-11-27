package domain


type TaobaoXhotelPromotionCreateResultSet struct {
    /*
        促销活动ID，创建活动会返回，更新不返回     */
    PromotionRuleId  *int64 `json:"promotion_rule_id,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        是否成功，true代表成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        日志id     */
    TraceId  *string `json:"trace_id,omitempty" `

}

func (s *TaobaoXhotelPromotionCreateResultSet) SetPromotionRuleId(v int64) *TaobaoXhotelPromotionCreateResultSet {
    s.PromotionRuleId = &v
    return s
}
func (s *TaobaoXhotelPromotionCreateResultSet) SetErrorCode(v string) *TaobaoXhotelPromotionCreateResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelPromotionCreateResultSet) SetErrorMsg(v string) *TaobaoXhotelPromotionCreateResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelPromotionCreateResultSet) SetSuccess(v bool) *TaobaoXhotelPromotionCreateResultSet {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelPromotionCreateResultSet) SetTraceId(v string) *TaobaoXhotelPromotionCreateResultSet {
    s.TraceId = &v
    return s
}
