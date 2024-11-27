package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelBnbpromoAddGeneralBookingInfo struct {
    /*
        活动折扣     */
    InvestmentNumber  *int64 `json:"investment_number,omitempty" `

    /*
        活动入住开始时间     */
    CheckInFrom  *util.LocalTime `json:"check_in_from,omitempty" `

    /*
        活动离店结束时间     */
    CheckOutTo  *util.LocalTime `json:"check_out_to,omitempty" `

}

func (s *TaobaoXhotelBnbpromoAddGeneralBookingInfo) SetInvestmentNumber(v int64) *TaobaoXhotelBnbpromoAddGeneralBookingInfo {
    s.InvestmentNumber = &v
    return s
}
func (s *TaobaoXhotelBnbpromoAddGeneralBookingInfo) SetCheckInFrom(v util.LocalTime) *TaobaoXhotelBnbpromoAddGeneralBookingInfo {
    s.CheckInFrom = &v
    return s
}
func (s *TaobaoXhotelBnbpromoAddGeneralBookingInfo) SetCheckOutTo(v util.LocalTime) *TaobaoXhotelBnbpromoAddGeneralBookingInfo {
    s.CheckOutTo = &v
    return s
}
