package domain


type TaobaoXhotelPromotionQueryTimeRange struct {
    /*
        开始时间     */
    StartTime  *string `json:"start_time,omitempty" `

    /*
        结束时间     */
    EndTime  *string `json:"end_time,omitempty" `

}

func (s *TaobaoXhotelPromotionQueryTimeRange) SetStartTime(v string) *TaobaoXhotelPromotionQueryTimeRange {
    s.StartTime = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryTimeRange) SetEndTime(v string) *TaobaoXhotelPromotionQueryTimeRange {
    s.EndTime = &v
    return s
}
