package domain


type TaobaoXhotelPromotionResourceBindBindData struct {
    /*
        促销规则Id     */
    PromotionRuleId  *int64 `json:"promotion_rule_id,omitempty" `

    /*
        供应商酒店code     */
    HotelCode  *string `json:"hotel_code,omitempty" `

    /*
        绑定成功结果，部分成功的时候有值     */
    SuccessRooms  *[]TaobaoXhotelPromotionResourceBindBindSuccItem `json:"success_rooms,omitempty" `

    /*
        绑定失败结果，部分成功的时候有值     */
    FailedRooms  *[]TaobaoXhotelPromotionResourceBindBindFailItem `json:"failed_rooms,omitempty" `

}

func (s *TaobaoXhotelPromotionResourceBindBindData) SetPromotionRuleId(v int64) *TaobaoXhotelPromotionResourceBindBindData {
    s.PromotionRuleId = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindData) SetHotelCode(v string) *TaobaoXhotelPromotionResourceBindBindData {
    s.HotelCode = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindData) SetSuccessRooms(v []TaobaoXhotelPromotionResourceBindBindSuccItem) *TaobaoXhotelPromotionResourceBindBindData {
    s.SuccessRooms = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindData) SetFailedRooms(v []TaobaoXhotelPromotionResourceBindBindFailItem) *TaobaoXhotelPromotionResourceBindBindData {
    s.FailedRooms = &v
    return s
}
