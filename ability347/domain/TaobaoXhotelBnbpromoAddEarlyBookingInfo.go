package domain


type TaobaoXhotelBnbpromoAddEarlyBookingInfo struct {
    /*
        活动折扣     */
    InvestmentNumber  *int64 `json:"investment_number,omitempty" `

    /*
        早定天数     */
    MinPreBookingDays  *int64 `json:"min_pre_booking_days,omitempty" `

}

func (s *TaobaoXhotelBnbpromoAddEarlyBookingInfo) SetInvestmentNumber(v int64) *TaobaoXhotelBnbpromoAddEarlyBookingInfo {
    s.InvestmentNumber = &v
    return s
}
func (s *TaobaoXhotelBnbpromoAddEarlyBookingInfo) SetMinPreBookingDays(v int64) *TaobaoXhotelBnbpromoAddEarlyBookingInfo {
    s.MinPreBookingDays = &v
    return s
}
