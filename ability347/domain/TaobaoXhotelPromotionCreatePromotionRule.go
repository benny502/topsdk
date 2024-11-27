package domain


type TaobaoXhotelPromotionCreatePromotionRule struct {
    /*
        连住天数,仅【连住优惠】活动类型需要传值且必须大于等于1，其他活动传值不生效。     */
    LengthOfStay  *int64 `json:"length_of_stay,omitempty" `

    /*
        预定日期     */
    BookDate  *TaobaoXhotelPromotionCreateDateRange `json:"book_date,omitempty" `

    /*
        入住时间的适用星期，周一到周日，适用1，不适用是0, 不支持所有日期都设置0     */
    StayWeekdays  *[]int32 `json:"stay_weekdays,omitempty" `

    /*
        提前预订天数，仅【早订优惠】 时需要传值,需要大于等于1，其他活动传值不生效。     */
    MinAdvanceDay  *int64 `json:"min_advance_day,omitempty" `

    /*
        限时特惠、今夜甩卖（尾房）必传预订时间点，其他活动类型传了不生效。     */
    BookTime  *TaobaoXhotelPromotionCreateTimeRange `json:"book_time,omitempty" `

    /*
        不生效日期，支持传多段，最大限制10段     */
    UnStayDate  *[]TaobaoXhotelPromotionCreateDateRange `json:"un_stay_date,omitempty" `

    /*
        优惠值，与discount_type 使用。 优惠20%，discount_value = 20，discount_type = 1；立减20元，discount_value = 2000，discount_type = 2；表示立减20元，原价200元，优惠后的金额为180元。     */
    DiscountValue  *int64 `json:"discount_value,omitempty" `

    /*
        促销限制的入住时间范围     */
    StayDate  *TaobaoXhotelPromotionCreateDateRange `json:"stay_date,omitempty" `

}

func (s *TaobaoXhotelPromotionCreatePromotionRule) SetLengthOfStay(v int64) *TaobaoXhotelPromotionCreatePromotionRule {
    s.LengthOfStay = &v
    return s
}
func (s *TaobaoXhotelPromotionCreatePromotionRule) SetBookDate(v TaobaoXhotelPromotionCreateDateRange) *TaobaoXhotelPromotionCreatePromotionRule {
    s.BookDate = &v
    return s
}
func (s *TaobaoXhotelPromotionCreatePromotionRule) SetStayWeekdays(v []int32) *TaobaoXhotelPromotionCreatePromotionRule {
    s.StayWeekdays = &v
    return s
}
func (s *TaobaoXhotelPromotionCreatePromotionRule) SetMinAdvanceDay(v int64) *TaobaoXhotelPromotionCreatePromotionRule {
    s.MinAdvanceDay = &v
    return s
}
func (s *TaobaoXhotelPromotionCreatePromotionRule) SetBookTime(v TaobaoXhotelPromotionCreateTimeRange) *TaobaoXhotelPromotionCreatePromotionRule {
    s.BookTime = &v
    return s
}
func (s *TaobaoXhotelPromotionCreatePromotionRule) SetUnStayDate(v []TaobaoXhotelPromotionCreateDateRange) *TaobaoXhotelPromotionCreatePromotionRule {
    s.UnStayDate = &v
    return s
}
func (s *TaobaoXhotelPromotionCreatePromotionRule) SetDiscountValue(v int64) *TaobaoXhotelPromotionCreatePromotionRule {
    s.DiscountValue = &v
    return s
}
func (s *TaobaoXhotelPromotionCreatePromotionRule) SetStayDate(v TaobaoXhotelPromotionCreateDateRange) *TaobaoXhotelPromotionCreatePromotionRule {
    s.StayDate = &v
    return s
}
