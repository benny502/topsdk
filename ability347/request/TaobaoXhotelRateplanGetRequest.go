package request


type TaobaoXhotelRateplanGetRequest struct {
    /*
        废弃，使用rateplan_code     */
    Rpid  *int64 `json:"rpid,omitempty" required:"false" `
    /*
        系统商，一般不填写，使用须申请     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        卖家自己系统的Code，简称RateCode     */
    RateplanCode  *string `json:"rateplan_code,omitempty" required:"false" `
}

func (s *TaobaoXhotelRateplanGetRequest) SetRpid(v int64) *TaobaoXhotelRateplanGetRequest {
    s.Rpid = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRequest) SetVendor(v string) *TaobaoXhotelRateplanGetRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRequest) SetRateplanCode(v string) *TaobaoXhotelRateplanGetRequest {
    s.RateplanCode = &v
    return s
}

func (req *TaobaoXhotelRateplanGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Rpid != nil) {
        paramMap["rpid"] = *req.Rpid
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.RateplanCode != nil) {
        paramMap["rateplan_code"] = *req.RateplanCode
    }
    return paramMap
}

func (req *TaobaoXhotelRateplanGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}