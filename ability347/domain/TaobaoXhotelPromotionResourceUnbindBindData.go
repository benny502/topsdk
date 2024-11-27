package domain


type TaobaoXhotelPromotionResourceUnbindBindData struct {
    /*
        促销规则Id     */
    PromotionRuleId  *int64 `json:"promotion_rule_id,omitempty" `

    /*
        供应商酒店code     */
    HotelCode  *string `json:"hotel_code,omitempty" `

    /*
        解绑成功结果，部分成功的时候有值     */
    SuccessRooms  *[]TaobaoXhotelPromotionResourceUnbindUnBindSuccItem `json:"success_rooms,omitempty" `

    /*
        解绑失败结果，部分成功的时候有值     */
    FailedRooms  *[]TaobaoXhotelPromotionResourceUnbindUnBindFailItem `json:"failed_rooms,omitempty" `

}

func (s *TaobaoXhotelPromotionResourceUnbindBindData) SetPromotionRuleId(v int64) *TaobaoXhotelPromotionResourceUnbindBindData {
    s.PromotionRuleId = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindBindData) SetHotelCode(v string) *TaobaoXhotelPromotionResourceUnbindBindData {
    s.HotelCode = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindBindData) SetSuccessRooms(v []TaobaoXhotelPromotionResourceUnbindUnBindSuccItem) *TaobaoXhotelPromotionResourceUnbindBindData {
    s.SuccessRooms = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindBindData) SetFailedRooms(v []TaobaoXhotelPromotionResourceUnbindUnBindFailItem) *TaobaoXhotelPromotionResourceUnbindBindData {
    s.FailedRooms = &v
    return s
}
