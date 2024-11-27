package domain


type TaobaoXhotelBaseinfoRoomGetRoomType struct {
    /*
        房价列表     */
    RatePlanList  *[]TaobaoXhotelBaseinfoRoomGetRatepPlan `json:"rate_plan_list,omitempty" `

    /*
        系统商     */
    Vendor  *string `json:"vendor,omitempty" `

    /*
        房型名称     */
    Name  *string `json:"name,omitempty" `

    /*
        房型状态。0:正常，-1:删除，-2:停售     */
    Status  *int64 `json:"status,omitempty" `

    /*
        outerId     */
    OuterId  *string `json:"outer_id,omitempty" `

}

func (s *TaobaoXhotelBaseinfoRoomGetRoomType) SetRatePlanList(v []TaobaoXhotelBaseinfoRoomGetRatepPlan) *TaobaoXhotelBaseinfoRoomGetRoomType {
    s.RatePlanList = &v
    return s
}
func (s *TaobaoXhotelBaseinfoRoomGetRoomType) SetVendor(v string) *TaobaoXhotelBaseinfoRoomGetRoomType {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelBaseinfoRoomGetRoomType) SetName(v string) *TaobaoXhotelBaseinfoRoomGetRoomType {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelBaseinfoRoomGetRoomType) SetStatus(v int64) *TaobaoXhotelBaseinfoRoomGetRoomType {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelBaseinfoRoomGetRoomType) SetOuterId(v string) *TaobaoXhotelBaseinfoRoomGetRoomType {
    s.OuterId = &v
    return s
}
