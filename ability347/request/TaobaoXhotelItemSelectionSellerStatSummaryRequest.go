package request


type TaobaoXhotelItemSelectionSellerStatSummaryRequest struct {
    /*
        日期  默认为昨天     */
    Date  *string `json:"date,omitempty" required:"false" `
    /*
        hid  默认为all     */
    Hid  *string `json:"hid,omitempty" required:"false" `
    /*
        vendor 默认为all     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        supplier 默认为all     */
    Supplier  *string `json:"supplier,omitempty" required:"false" `
    /*
        外部酒店编码     */
    OutHid  *string `json:"out_hid,omitempty" required:"false" `
}

func (s *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetDate(v string) *TaobaoXhotelItemSelectionSellerStatSummaryRequest {
    s.Date = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetHid(v string) *TaobaoXhotelItemSelectionSellerStatSummaryRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetVendor(v string) *TaobaoXhotelItemSelectionSellerStatSummaryRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetSupplier(v string) *TaobaoXhotelItemSelectionSellerStatSummaryRequest {
    s.Supplier = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatSummaryRequest) SetOutHid(v string) *TaobaoXhotelItemSelectionSellerStatSummaryRequest {
    s.OutHid = &v
    return s
}

func (req *TaobaoXhotelItemSelectionSellerStatSummaryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Date != nil) {
        paramMap["date"] = *req.Date
    }
    if(req.Hid != nil) {
        paramMap["hid"] = *req.Hid
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.Supplier != nil) {
        paramMap["supplier"] = *req.Supplier
    }
    if(req.OutHid != nil) {
        paramMap["out_hid"] = *req.OutHid
    }
    return paramMap
}

func (req *TaobaoXhotelItemSelectionSellerStatSummaryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}