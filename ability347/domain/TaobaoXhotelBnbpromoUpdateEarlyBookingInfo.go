package domain


type TaobaoXhotelBnbpromoUpdateEarlyBookingInfo struct {
    /*
        最少提前预定天数，数字范围限定在1-60     */
    MinPreBookingDays  *int64 `json:"min_pre_booking_days,omitempty" `

    /*
        折扣比例，填30就意味着原价的30%，也就是打3折。数字范围限定在10-95之间     */
    InvestmentNumber  *int64 `json:"investment_number,omitempty" `

    /*
        连住天数,可为空，范围1-60     */
    MinContinuityStay  *int64 `json:"min_continuity_stay,omitempty" `

}

func (s *TaobaoXhotelBnbpromoUpdateEarlyBookingInfo) SetMinPreBookingDays(v int64) *TaobaoXhotelBnbpromoUpdateEarlyBookingInfo {
    s.MinPreBookingDays = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUpdateEarlyBookingInfo) SetInvestmentNumber(v int64) *TaobaoXhotelBnbpromoUpdateEarlyBookingInfo {
    s.InvestmentNumber = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUpdateEarlyBookingInfo) SetMinContinuityStay(v int64) *TaobaoXhotelBnbpromoUpdateEarlyBookingInfo {
    s.MinContinuityStay = &v
    return s
}
