package domain


type TaobaoXhotelBnbpromoUpdateRateInfo struct {
    /*
        外部价格计划code     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" `

    /*
        外部房型id     */
    OutRid  *string `json:"out_rid,omitempty" `

}

func (s *TaobaoXhotelBnbpromoUpdateRateInfo) SetRatePlanCode(v string) *TaobaoXhotelBnbpromoUpdateRateInfo {
    s.RatePlanCode = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUpdateRateInfo) SetOutRid(v string) *TaobaoXhotelBnbpromoUpdateRateInfo {
    s.OutRid = &v
    return s
}
