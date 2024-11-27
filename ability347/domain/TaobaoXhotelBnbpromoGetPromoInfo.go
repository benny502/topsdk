package domain


type TaobaoXhotelBnbpromoGetPromoInfo struct {
    /*
        天天特惠优惠     */
    DailyBookingInfo  *TaobaoXhotelBnbpromoGetDailyBookingInfo `json:"daily_booking_info,omitempty" `

    /*
        早定优惠     */
    EarlyBookingInfo  *TaobaoXhotelBnbpromoGetEarlyBookingInfo `json:"early_booking_info,omitempty" `

    /*
        连住优惠     */
    LongOrderInfo  *TaobaoXhotelBnbpromoGetLongOrderInfo `json:"long_order_info,omitempty" `

}

func (s *TaobaoXhotelBnbpromoGetPromoInfo) SetDailyBookingInfo(v TaobaoXhotelBnbpromoGetDailyBookingInfo) *TaobaoXhotelBnbpromoGetPromoInfo {
    s.DailyBookingInfo = &v
    return s
}
func (s *TaobaoXhotelBnbpromoGetPromoInfo) SetEarlyBookingInfo(v TaobaoXhotelBnbpromoGetEarlyBookingInfo) *TaobaoXhotelBnbpromoGetPromoInfo {
    s.EarlyBookingInfo = &v
    return s
}
func (s *TaobaoXhotelBnbpromoGetPromoInfo) SetLongOrderInfo(v TaobaoXhotelBnbpromoGetLongOrderInfo) *TaobaoXhotelBnbpromoGetPromoInfo {
    s.LongOrderInfo = &v
    return s
}
