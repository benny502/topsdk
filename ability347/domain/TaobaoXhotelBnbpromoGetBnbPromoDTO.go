package domain


type TaobaoXhotelBnbpromoGetBnbPromoDTO struct {
    /*
        参与活动的rates     */
    RateInfos  *[]TaobaoXhotelBnbpromoGetRateinfos `json:"rate_infos,omitempty" `

    /*
        外部活动code     */
    OuterActivityCode  *string `json:"outer_activity_code,omitempty" `

    /*
        优惠信息     */
    PromoInfo  *TaobaoXhotelBnbpromoGetPromoInfo `json:"promo_info,omitempty" `

}

func (s *TaobaoXhotelBnbpromoGetBnbPromoDTO) SetRateInfos(v []TaobaoXhotelBnbpromoGetRateinfos) *TaobaoXhotelBnbpromoGetBnbPromoDTO {
    s.RateInfos = &v
    return s
}
func (s *TaobaoXhotelBnbpromoGetBnbPromoDTO) SetOuterActivityCode(v string) *TaobaoXhotelBnbpromoGetBnbPromoDTO {
    s.OuterActivityCode = &v
    return s
}
func (s *TaobaoXhotelBnbpromoGetBnbPromoDTO) SetPromoInfo(v TaobaoXhotelBnbpromoGetPromoInfo) *TaobaoXhotelBnbpromoGetBnbPromoDTO {
    s.PromoInfo = &v
    return s
}
