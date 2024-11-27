package request


type TaobaoXhotelItemSelectionSellerStatExposureRequest struct {
    /*
        日期 默认为昨天     */
    Date  *string `json:"date,omitempty" required:"false" `
    /*
        hid  默认为all     */
    Hid  *string `json:"hid,omitempty" required:"false" `
    /*
        默认为all     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        默认为all     */
    Supplier  *string `json:"supplier,omitempty" required:"false" `
    /*
        酒店编码     */
    OutHid  *string `json:"out_hid,omitempty" required:"false" `
}

func (s *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetDate(v string) *TaobaoXhotelItemSelectionSellerStatExposureRequest {
    s.Date = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetHid(v string) *TaobaoXhotelItemSelectionSellerStatExposureRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetVendor(v string) *TaobaoXhotelItemSelectionSellerStatExposureRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetSupplier(v string) *TaobaoXhotelItemSelectionSellerStatExposureRequest {
    s.Supplier = &v
    return s
}
func (s *TaobaoXhotelItemSelectionSellerStatExposureRequest) SetOutHid(v string) *TaobaoXhotelItemSelectionSellerStatExposureRequest {
    s.OutHid = &v
    return s
}

func (req *TaobaoXhotelItemSelectionSellerStatExposureRequest) ToMap() map[string]interface{} {
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

func (req *TaobaoXhotelItemSelectionSellerStatExposureRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}