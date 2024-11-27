package domain


type TaobaoXhotelPromotionResourceBindBindPromotionResourceParam struct {
    /*
        促销规则Id     */
    PromotionRuleId  *int64 `json:"promotion_rule_id,omitempty" `

    /*
        活动绑定资源，最多绑定200个 rate_plan_code + room_code 组合；     */
    Rooms  *[]TaobaoXhotelPromotionResourceBindResourceRoom `json:"rooms,omitempty" `

    /*
        供应商酒店code     */
    HotelCode  *string `json:"hotel_code,omitempty" `

}

func (s *TaobaoXhotelPromotionResourceBindBindPromotionResourceParam) SetPromotionRuleId(v int64) *TaobaoXhotelPromotionResourceBindBindPromotionResourceParam {
    s.PromotionRuleId = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindPromotionResourceParam) SetRooms(v []TaobaoXhotelPromotionResourceBindResourceRoom) *TaobaoXhotelPromotionResourceBindBindPromotionResourceParam {
    s.Rooms = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindPromotionResourceParam) SetHotelCode(v string) *TaobaoXhotelPromotionResourceBindBindPromotionResourceParam {
    s.HotelCode = &v
    return s
}
