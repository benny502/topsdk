package domain


type TaobaoXhotelPromotionQueryQueryPromotionResult struct {
    /*
        日志id     */
    TraceId  *string `json:"trace_id,omitempty" `

    /*
        请求传promotion_rule_id，total_count 指酒店code总量；请求传hotel_code，total_count是指promotion_rule_id总量；请求传promotion_rule_id + hotel_code，total_count = 1。     */
    TotalCount  *int64 `json:"total_count,omitempty" `

    /*
        直连营销     */
    DirectPromotions  *[]TaobaoXhotelPromotionQueryDirectPromotion `json:"direct_promotions,omitempty" `

    /*
        错误描述     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

}

func (s *TaobaoXhotelPromotionQueryQueryPromotionResult) SetTraceId(v string) *TaobaoXhotelPromotionQueryQueryPromotionResult {
    s.TraceId = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryQueryPromotionResult) SetTotalCount(v int64) *TaobaoXhotelPromotionQueryQueryPromotionResult {
    s.TotalCount = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryQueryPromotionResult) SetDirectPromotions(v []TaobaoXhotelPromotionQueryDirectPromotion) *TaobaoXhotelPromotionQueryQueryPromotionResult {
    s.DirectPromotions = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryQueryPromotionResult) SetErrorMsg(v string) *TaobaoXhotelPromotionQueryQueryPromotionResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryQueryPromotionResult) SetSuccess(v bool) *TaobaoXhotelPromotionQueryQueryPromotionResult {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryQueryPromotionResult) SetErrorCode(v string) *TaobaoXhotelPromotionQueryQueryPromotionResult {
    s.ErrorCode = &v
    return s
}
