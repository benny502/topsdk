package domain


type TaobaoXhotelBnbpromoAddDailyBookingInfo struct {
    /*
        活动折扣     */
    InvestmentNumber  *int64 `json:"investment_number,omitempty" `

    /*
        生效星期，星期一星期二生效就填1，2     */
    ValidWeeks  *[]int32 `json:"valid_weeks,omitempty" `

}

func (s *TaobaoXhotelBnbpromoAddDailyBookingInfo) SetInvestmentNumber(v int64) *TaobaoXhotelBnbpromoAddDailyBookingInfo {
    s.InvestmentNumber = &v
    return s
}
func (s *TaobaoXhotelBnbpromoAddDailyBookingInfo) SetValidWeeks(v []int32) *TaobaoXhotelBnbpromoAddDailyBookingInfo {
    s.ValidWeeks = &v
    return s
}
