package domain


type TaobaoXhotelBnbpromoUpdateUpdatePromoParam struct {
    /*
        外部营销活动的code，最长40个字符     */
    OuterActivityCode  *string `json:"outer_activity_code,omitempty" `

    /*
        营销活动的具体参数对象，在每次添加更新的时候，long_order_info、early_booking_info、daily_booking_info 只能填1种类型，其他2种类型为空     */
    PromoInfo  *TaobaoXhotelBnbpromoUpdatePromoInfo `json:"promo_info,omitempty" `

    /*
        营销活动关联的价格计划和房型     */
    RateInfos  *[]TaobaoXhotelBnbpromoUpdateRateInfo `json:"rate_infos,omitempty" `

}

func (s *TaobaoXhotelBnbpromoUpdateUpdatePromoParam) SetOuterActivityCode(v string) *TaobaoXhotelBnbpromoUpdateUpdatePromoParam {
    s.OuterActivityCode = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUpdateUpdatePromoParam) SetPromoInfo(v TaobaoXhotelBnbpromoUpdatePromoInfo) *TaobaoXhotelBnbpromoUpdateUpdatePromoParam {
    s.PromoInfo = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUpdateUpdatePromoParam) SetRateInfos(v []TaobaoXhotelBnbpromoUpdateRateInfo) *TaobaoXhotelBnbpromoUpdateUpdatePromoParam {
    s.RateInfos = &v
    return s
}
