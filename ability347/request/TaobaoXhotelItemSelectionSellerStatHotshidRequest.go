package request


type TaobaoXhotelItemSelectionSellerStatHotshidRequest struct {
    /*
        日期  默认为昨天     */
    Date  *string `json:"date,omitempty" required:"false" `
    /*
        酒店id  默认all     */
    Hid  *string `json:"hid,omitempty" required:"false" `
    /*
        vendor  默认all     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        supplier  默认all     */
    Supplier  *string `json:"supplier,omitempty" required:"false" `
    /*
        酒店编码     */
    OutHid  *string `json:"out_hid,omitempty" required:"false" `
}

func (s *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetDate(v string) *TaobaoXhotelItemSelectionSellerStatHotshidRequest {
    s.Date = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetHid(v string) *TaobaoXhotelItemSelectionSellerStatHotshidRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetVendor(v string) *TaobaoXhotelItemSelectionSellerStatHotshidRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetSupplier(v string) *TaobaoXhotelItemSelectionSellerStatHotshidRequest {
    s.Supplier = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatHotshidRequest) SetOutHid(v string) *TaobaoXhotelItemSelectionSellerStatHotshidRequest {
    s.OutHid = &v
    return s
}

func (req *TaobaoXhotelItemSelectionSellerStatHotshidRequest) ToMap() map[string]interface{} {
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

func (req *TaobaoXhotelItemSelectionSellerStatHotshidRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}