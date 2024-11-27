package domain


type TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult struct {
    /*
        曝光率     */
    ExposedPercent  *string `json:"exposed_percent,omitempty" `

    /*
        supplier参数     */
    SupplierParam  *string `json:"supplier_param,omitempty" `

    /*
        标准酒店维度曝光总数     */
    ShidTotalAmount  *string `json:"shid_total_amount,omitempty" `

    /*
        hid参数     */
    HidParam  *string `json:"hid_param,omitempty" `

    /*
        rate最低分     */
    MinRateScore  *string `json:"min_rate_score,omitempty" `

    /*
        不可售情况     */
    UnsaleReseasonInfo  *string `json:"unsale_reseason_info,omitempty" `

    /*
        rate最高分     */
    MaxRateScore  *string `json:"max_rate_score,omitempty" `

    /*
        选品情况     */
    SelectionMessageInfo  *string `json:"selection_message_info,omitempty" `

    /*
        rate平均分     */
    AvgRateScore  *string `json:"avg_rate_score,omitempty" `

    /*
        日期     */
    DateParam  *string `json:"date_param,omitempty" `

    /*
        商品总数     */
    TotalAmount  *string `json:"total_amount,omitempty" `

    /*
        vendor参数     */
    VendorParam  *string `json:"vendor_param,omitempty" `

    /*
        曝光总数     */
    ExposedAmount  *string `json:"exposed_amount,omitempty" `

    /*
        选品情况     */
    SelectionMessageInfoJson  *string `json:"selection_message_info_json,omitempty" `

    /*
        不可售情况     */
    UnsaleReasonInfoJson  *string `json:"unsale_reason_info_json,omitempty" `

    /*
        sellerId参数     */
    SellerIdParam  *string `json:"seller_id_param,omitempty" `

    /*
        可售商品数量     */
    CanSaleAmount  *string `json:"can_sale_amount,omitempty" `

    /*
        选品保留商品数量     */
    SelectedAmount  *string `json:"selected_amount,omitempty" `

}

func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetExposedPercent(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.ExposedPercent = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetSupplierParam(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.SupplierParam = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetShidTotalAmount(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.ShidTotalAmount = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetHidParam(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.HidParam = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetMinRateScore(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.MinRateScore = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetUnsaleReseasonInfo(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.UnsaleReseasonInfo = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetMaxRateScore(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.MaxRateScore = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetSelectionMessageInfo(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.SelectionMessageInfo = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetAvgRateScore(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.AvgRateScore = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetDateParam(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.DateParam = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetTotalAmount(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.TotalAmount = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetVendorParam(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.VendorParam = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetExposedAmount(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.ExposedAmount = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetSelectionMessageInfoJson(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.SelectionMessageInfoJson = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetUnsaleReasonInfoJson(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.UnsaleReasonInfoJson = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetSellerIdParam(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.SellerIdParam = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetCanSaleAmount(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.CanSaleAmount = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult) SetSelectedAmount(v string) *TaobaoXhotelItemSelectionSellerStatSummarySellerStatSummaryResult {
    s.SelectedAmount = &v
    return s
}
