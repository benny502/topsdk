package domain


type TaobaoXhotelBnbpromoGetRateinfos struct {
    /*
        外部房源id     */
    OutRid  *string `json:"out_rid,omitempty" `

    /*
        外部rpcode     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" `

}

func (s *TaobaoXhotelBnbpromoGetRateinfos) SetOutRid(v string) *TaobaoXhotelBnbpromoGetRateinfos {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelBnbpromoGetRateinfos) SetRatePlanCode(v string) *TaobaoXhotelBnbpromoGetRateinfos {
    s.RatePlanCode = &v
    return s
}
