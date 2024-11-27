package domain


type TaobaoXhotelBaseinfoRoomGetRatepPlan struct {
    /*
        系统商     */
    Vendor  *string `json:"vendor,omitempty" `

    /*
        房价名称     */
    Name  *string `json:"name,omitempty" `

    /*
        ratePlanCode     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" `

    /*
        1：开启2：关闭。     */
    Status  *int64 `json:"status,omitempty" `

}

func (s *TaobaoXhotelBaseinfoRoomGetRatepPlan) SetVendor(v string) *TaobaoXhotelBaseinfoRoomGetRatepPlan {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelBaseinfoRoomGetRatepPlan) SetName(v string) *TaobaoXhotelBaseinfoRoomGetRatepPlan {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelBaseinfoRoomGetRatepPlan) SetRatePlanCode(v string) *TaobaoXhotelBaseinfoRoomGetRatepPlan {
    s.RatePlanCode = &v
    return s
}
func (s *TaobaoXhotelBaseinfoRoomGetRatepPlan) SetStatus(v int64) *TaobaoXhotelBaseinfoRoomGetRatepPlan {
    s.Status = &v
    return s
}
