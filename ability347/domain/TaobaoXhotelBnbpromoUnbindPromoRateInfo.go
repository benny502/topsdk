package domain


type TaobaoXhotelBnbpromoUnbindPromoRateInfo struct {
    /*
        外部rp     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" `

    /*
        外部rid     */
    OutRid  *string `json:"out_rid,omitempty" `

}

func (s *TaobaoXhotelBnbpromoUnbindPromoRateInfo) SetRatePlanCode(v string) *TaobaoXhotelBnbpromoUnbindPromoRateInfo {
    s.RatePlanCode = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUnbindPromoRateInfo) SetOutRid(v string) *TaobaoXhotelBnbpromoUnbindPromoRateInfo {
    s.OutRid = &v
    return s
}
