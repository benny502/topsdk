package request


type TaobaoXhotelMultiplerateDeleteRequest struct {
    /*
        渠道，和推送房价所使用的渠道保持一致     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        商家价格政策编码     */
    RateplanCode  *string `json:"rateplan_code" required:"true" `
    /*
        商家房型编码     */
    OutRid  *string `json:"out_rid" required:"true" `
    /*
        连住天数     */
    Occupancy  *int64 `json:"occupancy,omitempty" required:"false" `
    /*
        入住人数     */
    Lengthofstay  *int64 `json:"lengthofstay,omitempty" required:"false" `
}

func (s *TaobaoXhotelMultiplerateDeleteRequest) SetVendor(v string) *TaobaoXhotelMultiplerateDeleteRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelMultiplerateDeleteRequest) SetRateplanCode(v string) *TaobaoXhotelMultiplerateDeleteRequest {
    s.RateplanCode = &v
    return s
}
func (s *TaobaoXhotelMultiplerateDeleteRequest) SetOutRid(v string) *TaobaoXhotelMultiplerateDeleteRequest {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelMultiplerateDeleteRequest) SetOccupancy(v int64) *TaobaoXhotelMultiplerateDeleteRequest {
    s.Occupancy = &v
    return s
}
func (s *TaobaoXhotelMultiplerateDeleteRequest) SetLengthofstay(v int64) *TaobaoXhotelMultiplerateDeleteRequest {
    s.Lengthofstay = &v
    return s
}

func (req *TaobaoXhotelMultiplerateDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.RateplanCode != nil) {
        paramMap["rateplan_code"] = *req.RateplanCode
    }
    if(req.OutRid != nil) {
        paramMap["out_rid"] = *req.OutRid
    }
    if(req.Occupancy != nil) {
        paramMap["occupancy"] = *req.Occupancy
    }
    if(req.Lengthofstay != nil) {
        paramMap["lengthofstay"] = *req.Lengthofstay
    }
    return paramMap
}

func (req *TaobaoXhotelMultiplerateDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}