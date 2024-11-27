package domain


type TaobaoXhotelItemSelectionSellerStatHotshidModule struct {
    /*
        热销标准酒店中卖家覆盖的数量     */
    CoveredHidAmount  *string `json:"covered_hid_amount,omitempty" `

    /*
        热销标准酒店中卖家可售的酒店数量     */
    CanSaleHidAmount  *string `json:"can_sale_hid_amount,omitempty" `

}

func (s *TaobaoXhotelItemSelectionSellerStatHotshidModule) SetCoveredHidAmount(v string) *TaobaoXhotelItemSelectionSellerStatHotshidModule {
    s.CoveredHidAmount = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatHotshidModule) SetCanSaleHidAmount(v string) *TaobaoXhotelItemSelectionSellerStatHotshidModule {
    s.CanSaleHidAmount = &v
    return s
}
