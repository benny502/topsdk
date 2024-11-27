package domain


type TaobaoXhotelPromotionCreateTimeRange struct {
    /*
        开始时间,今夜甩卖 start_time 最早开始时间 14:00,限时特惠活动预订时间限当天，start_time最早开始时间 00:00:00     */
    StartTime  *string `json:"start_time,omitempty" `

    /*
        今夜甩卖支持end_time 如果到第二天的凌晨，则需+24h，如第二天凌晨4点，即传28:00:00，最大到31:00:00；限时特惠活动预订时间限当天end_time最大为23:59:59。     */
    EndTime  *string `json:"end_time,omitempty" `

}

func (s *TaobaoXhotelPromotionCreateTimeRange) SetStartTime(v string) *TaobaoXhotelPromotionCreateTimeRange {
    s.StartTime = &v
    return s
}
func (s *TaobaoXhotelPromotionCreateTimeRange) SetEndTime(v string) *TaobaoXhotelPromotionCreateTimeRange {
    s.EndTime = &v
    return s
}
