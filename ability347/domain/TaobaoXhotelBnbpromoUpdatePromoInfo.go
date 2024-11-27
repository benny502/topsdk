package domain


type TaobaoXhotelBnbpromoUpdatePromoInfo struct {
    /*
        连住优惠的入参     */
    LongOrderInfo  *TaobaoXhotelBnbpromoUpdateLongOrderInfo `json:"long_order_info,omitempty" `

    /*
        早定优惠的入参     */
    EarlyBookingInfo  *TaobaoXhotelBnbpromoUpdateEarlyBookingInfo `json:"early_booking_info,omitempty" `

    /*
        天天特惠的入参     */
    DailyBookingInfo  *TaobaoXhotelBnbpromoUpdateDailyBookingInfo `json:"daily_booking_info,omitempty" `

}

func (s *TaobaoXhotelBnbpromoUpdatePromoInfo) SetLongOrderInfo(v TaobaoXhotelBnbpromoUpdateLongOrderInfo) *TaobaoXhotelBnbpromoUpdatePromoInfo {
    s.LongOrderInfo = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUpdatePromoInfo) SetEarlyBookingInfo(v TaobaoXhotelBnbpromoUpdateEarlyBookingInfo) *TaobaoXhotelBnbpromoUpdatePromoInfo {
    s.EarlyBookingInfo = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUpdatePromoInfo) SetDailyBookingInfo(v TaobaoXhotelBnbpromoUpdateDailyBookingInfo) *TaobaoXhotelBnbpromoUpdatePromoInfo {
    s.DailyBookingInfo = &v
    return s
}
