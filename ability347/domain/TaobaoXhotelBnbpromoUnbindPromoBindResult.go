package domain


type TaobaoXhotelBnbpromoUnbindPromoBindResult struct {
    /*
        成功或失败     */
    Success  *bool `json:"success,omitempty" `

    /*
        失败信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        外部rid     */
    OutRid  *string `json:"out_rid,omitempty" `

    /*
        外部rp     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" `

}

func (s *TaobaoXhotelBnbpromoUnbindPromoBindResult) SetSuccess(v bool) *TaobaoXhotelBnbpromoUnbindPromoBindResult {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUnbindPromoBindResult) SetErrorMsg(v string) *TaobaoXhotelBnbpromoUnbindPromoBindResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUnbindPromoBindResult) SetOutRid(v string) *TaobaoXhotelBnbpromoUnbindPromoBindResult {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUnbindPromoBindResult) SetRatePlanCode(v string) *TaobaoXhotelBnbpromoUnbindPromoBindResult {
    s.RatePlanCode = &v
    return s
}
