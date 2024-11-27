package domain


type TaobaoXhotelBnbpromoBindPromoBindResult struct {
    /*
        活动是否报名成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        活动失败原因     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        外部rid     */
    OutRid  *string `json:"out_rid,omitempty" `

    /*
        外部rp     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" `

}

func (s *TaobaoXhotelBnbpromoBindPromoBindResult) SetSuccess(v bool) *TaobaoXhotelBnbpromoBindPromoBindResult {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelBnbpromoBindPromoBindResult) SetErrorMsg(v string) *TaobaoXhotelBnbpromoBindPromoBindResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelBnbpromoBindPromoBindResult) SetOutRid(v string) *TaobaoXhotelBnbpromoBindPromoBindResult {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelBnbpromoBindPromoBindResult) SetRatePlanCode(v string) *TaobaoXhotelBnbpromoBindPromoBindResult {
    s.RatePlanCode = &v
    return s
}
