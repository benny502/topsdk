package domain


type TaobaoXhotelItemSelectionSellerStatExposureModule struct {
    /*
        返回结果     */
    SellerStatExposureElementList  *[]TaobaoXhotelItemSelectionSellerStatExposureSellerStatExposureElementList `json:"seller_stat_exposure_element_list,omitempty" `

}

func (s *TaobaoXhotelItemSelectionSellerStatExposureModule) SetSellerStatExposureElementList(v []TaobaoXhotelItemSelectionSellerStatExposureSellerStatExposureElementList) *TaobaoXhotelItemSelectionSellerStatExposureModule {
    s.SellerStatExposureElementList = &v
    return s
}
