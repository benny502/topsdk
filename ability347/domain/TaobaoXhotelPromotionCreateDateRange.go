package domain


type TaobaoXhotelPromotionCreateDateRange struct {
    /*
        开始日期，不传默认今天     */
    Start  *string `json:"start,omitempty" `

    /*
        结束日期，不传默认今天+10年     */
    End  *string `json:"end,omitempty" `

}

func (s *TaobaoXhotelPromotionCreateDateRange) SetStart(v string) *TaobaoXhotelPromotionCreateDateRange {
    s.Start = &v
    return s
}
func (s *TaobaoXhotelPromotionCreateDateRange) SetEnd(v string) *TaobaoXhotelPromotionCreateDateRange {
    s.End = &v
    return s
}
