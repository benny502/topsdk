package domain


type TaobaoXhotelItemSelectionSellerStatHotshidResult struct {
    /*
        错误码     */
    Code  *string `json:"code,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        返回结构     */
    Module  *TaobaoXhotelItemSelectionSellerStatHotshidModule `json:"module,omitempty" `

    /*
        错误信息     */
    Message  *string `json:"message,omitempty" `

}

func (s *TaobaoXhotelItemSelectionSellerStatHotshidResult) SetCode(v string) *TaobaoXhotelItemSelectionSellerStatHotshidResult {
    s.Code = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatHotshidResult) SetSuccess(v bool) *TaobaoXhotelItemSelectionSellerStatHotshidResult {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatHotshidResult) SetModule(v TaobaoXhotelItemSelectionSellerStatHotshidModule) *TaobaoXhotelItemSelectionSellerStatHotshidResult {
    s.Module = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatHotshidResult) SetMessage(v string) *TaobaoXhotelItemSelectionSellerStatHotshidResult {
    s.Message = &v
    return s
}
