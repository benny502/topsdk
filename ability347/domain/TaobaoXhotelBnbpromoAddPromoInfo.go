package domain


type TaobaoXhotelBnbpromoAddPromoInfo struct {
    /*
        今夜特惠     */
    TonightDiscount  *TaobaoXhotelBnbpromoAddTonightDiscount `json:"tonight_discount,omitempty" `

    /*
        连住优惠     */
    LongOrderInfo  *TaobaoXhotelBnbpromoAddLongOrderInfo `json:"long_order_info,omitempty" `

    /*
        早定优惠     */
    EarlyBookingInfo  *TaobaoXhotelBnbpromoAddEarlyBookingInfo `json:"early_booking_info,omitempty" `

    /*
        天天特惠     */
    DailyBookingInfo  *TaobaoXhotelBnbpromoAddDailyBookingInfo `json:"daily_booking_info,omitempty" `

    /*
        民宿优惠     */
    GeneralBookingInfo  *TaobaoXhotelBnbpromoAddGeneralBookingInfo `json:"general_booking_info,omitempty" `

}

func (s *TaobaoXhotelBnbpromoAddPromoInfo) SetTonightDiscount(v TaobaoXhotelBnbpromoAddTonightDiscount) *TaobaoXhotelBnbpromoAddPromoInfo {
    s.TonightDiscount = &v
    return s
}
func (s *TaobaoXhotelBnbpromoAddPromoInfo) SetLongOrderInfo(v TaobaoXhotelBnbpromoAddLongOrderInfo) *TaobaoXhotelBnbpromoAddPromoInfo {
    s.LongOrderInfo = &v
    return s
}
func (s *TaobaoXhotelBnbpromoAddPromoInfo) SetEarlyBookingInfo(v TaobaoXhotelBnbpromoAddEarlyBookingInfo) *TaobaoXhotelBnbpromoAddPromoInfo {
    s.EarlyBookingInfo = &v
    return s
}
func (s *TaobaoXhotelBnbpromoAddPromoInfo) SetDailyBookingInfo(v TaobaoXhotelBnbpromoAddDailyBookingInfo) *TaobaoXhotelBnbpromoAddPromoInfo {
    s.DailyBookingInfo = &v
    return s
}
func (s *TaobaoXhotelBnbpromoAddPromoInfo) SetGeneralBookingInfo(v TaobaoXhotelBnbpromoAddGeneralBookingInfo) *TaobaoXhotelBnbpromoAddPromoInfo {
    s.GeneralBookingInfo = &v
    return s
}
