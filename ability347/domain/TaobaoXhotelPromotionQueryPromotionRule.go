package domain


type TaobaoXhotelPromotionQueryPromotionRule struct {
    /*
        连住天数     */
    LengthOfStay  *int64 `json:"length_of_stay,omitempty" `

    /*
        预定日期     */
    BookDate  *TaobaoXhotelPromotionQueryDateRange `json:"book_date,omitempty" `

    /*
        入住时间的适用星期     */
    StayWeekdays  *[]int32 `json:"stay_weekdays,omitempty" `

    /*
        提前预订天数     */
    MinAdvanceDay  *int64 `json:"min_advance_day,omitempty" `

    /*
        预定时间     */
    BookTime  *TaobaoXhotelPromotionQueryTimeRange `json:"book_time,omitempty" `

    /*
        不生效日期     */
    UnStayDate  *[]TaobaoXhotelPromotionQueryDateRange `json:"un_stay_date,omitempty" `

    /*
        1 Fixed:固定金额,单位分，1000;0Percentage: 折扣比例，20,表示优惠20%，比如100元，优惠20%后为 80元。     */
    DiscountValue  *int64 `json:"discount_value,omitempty" `

    /*
        促销限制的入住时间范围     */
    StayDate  *TaobaoXhotelPromotionQueryDateRange `json:"stay_date,omitempty" `

}

func (s *TaobaoXhotelPromotionQueryPromotionRule) SetLengthOfStay(v int64) *TaobaoXhotelPromotionQueryPromotionRule {
    s.LengthOfStay = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryPromotionRule) SetBookDate(v TaobaoXhotelPromotionQueryDateRange) *TaobaoXhotelPromotionQueryPromotionRule {
    s.BookDate = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryPromotionRule) SetStayWeekdays(v []int32) *TaobaoXhotelPromotionQueryPromotionRule {
    s.StayWeekdays = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryPromotionRule) SetMinAdvanceDay(v int64) *TaobaoXhotelPromotionQueryPromotionRule {
    s.MinAdvanceDay = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryPromotionRule) SetBookTime(v TaobaoXhotelPromotionQueryTimeRange) *TaobaoXhotelPromotionQueryPromotionRule {
    s.BookTime = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryPromotionRule) SetUnStayDate(v []TaobaoXhotelPromotionQueryDateRange) *TaobaoXhotelPromotionQueryPromotionRule {
    s.UnStayDate = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryPromotionRule) SetDiscountValue(v int64) *TaobaoXhotelPromotionQueryPromotionRule {
    s.DiscountValue = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryPromotionRule) SetStayDate(v TaobaoXhotelPromotionQueryDateRange) *TaobaoXhotelPromotionQueryPromotionRule {
    s.StayDate = &v
    return s
}
