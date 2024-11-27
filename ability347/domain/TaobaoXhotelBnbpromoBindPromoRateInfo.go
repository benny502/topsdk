package domain


type TaobaoXhotelBnbpromoBindPromoRateInfo struct {
    /*
        外部rp     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" `

    /*
        外部rid     */
    OutRid  *string `json:"out_rid,omitempty" `

}

func (s *TaobaoXhotelBnbpromoBindPromoRateInfo) SetRatePlanCode(v string) *TaobaoXhotelBnbpromoBindPromoRateInfo {
    s.RatePlanCode = &v
    return s
}
func (s *TaobaoXhotelBnbpromoBindPromoRateInfo) SetOutRid(v string) *TaobaoXhotelBnbpromoBindPromoRateInfo {
    s.OutRid = &v
    return s
}
