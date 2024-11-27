package domain


type TaobaoXhotelItemSelectionSellerStatExposureResult struct {
    /*
        错误码     */
    Code  *string `json:"code,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        返回结果     */
    Module  *TaobaoXhotelItemSelectionSellerStatExposureModule `json:"module,omitempty" `

    /*
        接口信息     */
    Message  *string `json:"message,omitempty" `

}

func (s *TaobaoXhotelItemSelectionSellerStatExposureResult) SetCode(v string) *TaobaoXhotelItemSelectionSellerStatExposureResult {
    s.Code = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatExposureResult) SetSuccess(v bool) *TaobaoXhotelItemSelectionSellerStatExposureResult {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatExposureResult) SetModule(v TaobaoXhotelItemSelectionSellerStatExposureModule) *TaobaoXhotelItemSelectionSellerStatExposureResult {
    s.Module = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatExposureResult) SetMessage(v string) *TaobaoXhotelItemSelectionSellerStatExposureResult {
    s.Message = &v
    return s
}
