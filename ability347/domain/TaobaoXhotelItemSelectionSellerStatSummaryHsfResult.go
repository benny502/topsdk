package domain


type TaobaoXhotelItemSelectionSellerStatSummaryHsfResult struct {
    /*
        错误码     */
    Code  *string `json:"code,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        返回结果     */
    Module  *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult `json:"module,omitempty" `

    /*
        接口信息     */
    Message  *string `json:"message,omitempty" `

}

func (s *TaobaoXhotelItemSelectionSellerStatSummaryHsfResult) SetCode(v string) *TaobaoXhotelItemSelectionSellerStatSummaryHsfResult {
    s.Code = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummaryHsfResult) SetSuccess(v bool) *TaobaoXhotelItemSelectionSellerStatSummaryHsfResult {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummaryHsfResult) SetModule(v TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) *TaobaoXhotelItemSelectionSellerStatSummaryHsfResult {
    s.Module = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummaryHsfResult) SetMessage(v string) *TaobaoXhotelItemSelectionSellerStatSummaryHsfResult {
    s.Message = &v
    return s
}
