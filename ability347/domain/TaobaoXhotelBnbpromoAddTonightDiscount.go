package domain


type TaobaoXhotelBnbpromoAddTonightDiscount struct {
    /*
        活动折扣     */
    InvestmentNumber  *string `json:"investment_number,omitempty" `

    /*
        起始时间     */
    StartTime  *string `json:"start_time,omitempty" `

}

func (s *TaobaoXhotelBnbpromoAddTonightDiscount) SetInvestmentNumber(v string) *TaobaoXhotelBnbpromoAddTonightDiscount {
    s.InvestmentNumber = &v
    return s
}
func (s *TaobaoXhotelBnbpromoAddTonightDiscount) SetStartTime(v string) *TaobaoXhotelBnbpromoAddTonightDiscount {
    s.StartTime = &v
    return s
}
