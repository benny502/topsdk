package domain


type TaobaoXhotelItemSelectionSellerStatExposureSellerStatExposureElementList struct {
    /*
        日期     */
    StatDate  *string `json:"stat_date,omitempty" `

    /*
        曝光率     */
    ExposedPercent  *string `json:"exposed_percent,omitempty" `

    /*
        shid维度访问量     */
    ShidTotalAmount  *string `json:"shid_total_amount,omitempty" `

    /*
        对应商品曝光数量     */
    ExposedAmount  *string `json:"exposed_amount,omitempty" `

}

func (s *TaobaoXhotelItemSelectionSellerStatExposureSellerStatExposureElementList) SetStatDate(v string) *TaobaoXhotelItemSelectionSellerStatExposureSellerStatExposureElementList {
    s.StatDate = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatExposureSellerStatExposureElementList) SetExposedPercent(v string) *TaobaoXhotelItemSelectionSellerStatExposureSellerStatExposureElementList {
    s.ExposedPercent = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatExposureSellerStatExposureElementList) SetShidTotalAmount(v string) *TaobaoXhotelItemSelectionSellerStatExposureSellerStatExposureElementList {
    s.ShidTotalAmount = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatExposureSellerStatExposureElementList) SetExposedAmount(v string) *TaobaoXhotelItemSelectionSellerStatExposureSellerStatExposureElementList {
    s.ExposedAmount = &v
    return s
}
