package domain


type TaobaoXhotelPromotionQueryDateRange struct {
    /*
        开始入住时间     */
    Start  *string `json:"start,omitempty" `

    /*
        截止入住时间     */
    End  *string `json:"end,omitempty" `

}

func (s *TaobaoXhotelPromotionQueryDateRange) SetStart(v string) *TaobaoXhotelPromotionQueryDateRange {
    s.Start = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryDateRange) SetEnd(v string) *TaobaoXhotelPromotionQueryDateRange {
    s.End = &v
    return s
}
