package domain


type TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceParam struct {
    /*
        促销规则Id     */
    PromotionRuleId  *int64 `json:"promotion_rule_id,omitempty" `

    /*
        活动解绑资源     */
    Rooms  *[]TaobaoXhotelPromotionResourceUnbindResourceRoom `json:"rooms,omitempty" `

    /*
        供应商酒店code     */
    HotelCode  *string `json:"hotel_code,omitempty" `

}

func (s *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceParam) SetPromotionRuleId(v int64) *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceParam {
    s.PromotionRuleId = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceParam) SetRooms(v []TaobaoXhotelPromotionResourceUnbindResourceRoom) *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceParam {
    s.Rooms = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceParam) SetHotelCode(v string) *TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceParam {
    s.HotelCode = &v
    return s
}
