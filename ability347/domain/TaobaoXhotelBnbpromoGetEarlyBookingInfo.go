package domain


type TaobaoXhotelBnbpromoGetEarlyBookingInfo struct {
    /*
        提前预定天数,时间范围为1-60天     */
    MinPreBookingDays  *int64 `json:"min_pre_booking_days,omitempty" `

    /*
        固定折扣百分比     */
    InvestmentNumber  *int64 `json:"investment_number,omitempty" `

    /*
        连住天数     */
    MinContinuityStay  *int64 `json:"min_continuity_stay,omitempty" `

}

func (s *TaobaoXhotelBnbpromoGetEarlyBookingInfo) SetMinPreBookingDays(v int64) *TaobaoXhotelBnbpromoGetEarlyBookingInfo {
    s.MinPreBookingDays = &v
    return s
}
func (s *TaobaoXhotelBnbpromoGetEarlyBookingInfo) SetInvestmentNumber(v int64) *TaobaoXhotelBnbpromoGetEarlyBookingInfo {
    s.InvestmentNumber = &v
    return s
}
func (s *TaobaoXhotelBnbpromoGetEarlyBookingInfo) SetMinContinuityStay(v int64) *TaobaoXhotelBnbpromoGetEarlyBookingInfo {
    s.MinContinuityStay = &v
    return s
}
