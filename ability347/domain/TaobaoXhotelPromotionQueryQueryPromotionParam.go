package domain


type TaobaoXhotelPromotionQueryQueryPromotionParam struct {
    /*
        促销规则Id     */
    PromotionRuleId  *int64 `json:"promotion_rule_id,omitempty" `

    /*
        页码 defalutValue:1    */
    PageNo  *int64 `json:"page_no,omitempty" `

    /*
        每页大小，最大50 defalutValue:50    */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        酒店code     */
    HotelCode  *string `json:"hotel_code,omitempty" `

}

func (s *TaobaoXhotelPromotionQueryQueryPromotionParam) SetPromotionRuleId(v int64) *TaobaoXhotelPromotionQueryQueryPromotionParam {
    s.PromotionRuleId = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryQueryPromotionParam) SetPageNo(v int64) *TaobaoXhotelPromotionQueryQueryPromotionParam {
    s.PageNo = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryQueryPromotionParam) SetPageSize(v int64) *TaobaoXhotelPromotionQueryQueryPromotionParam {
    s.PageSize = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryQueryPromotionParam) SetHotelCode(v string) *TaobaoXhotelPromotionQueryQueryPromotionParam {
    s.HotelCode = &v
    return s
}
