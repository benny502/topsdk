package domain


type TaobaoXhotelBaseinfoGetRatePlan struct {
    /*
        系统商，一般不填写，使用须申请     */
    Vendor  *string `json:"vendor,omitempty" `

    /*
        房价名称     */
    Name  *string `json:"name,omitempty" `

    /*
        卖家自己系统的Code，简称RateCode     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" `

    /*
        阿里房价id     */
    RatePlanId  *int64 `json:"rate_plan_id,omitempty" `

    /*
        1：开启2：关闭。     */
    Status  *int64 `json:"status,omitempty" `

}

func (s *TaobaoXhotelBaseinfoGetRatePlan) SetVendor(v string) *TaobaoXhotelBaseinfoGetRatePlan {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRatePlan) SetName(v string) *TaobaoXhotelBaseinfoGetRatePlan {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRatePlan) SetRatePlanCode(v string) *TaobaoXhotelBaseinfoGetRatePlan {
    s.RatePlanCode = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRatePlan) SetRatePlanId(v int64) *TaobaoXhotelBaseinfoGetRatePlan {
    s.RatePlanId = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRatePlan) SetStatus(v int64) *TaobaoXhotelBaseinfoGetRatePlan {
    s.Status = &v
    return s
}
