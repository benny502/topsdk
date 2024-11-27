package request


type TaobaoXhotelRateGetRequest struct {
    /*
        gid酒店商品id     */
    Gid  *int64 `json:"gid,omitempty" required:"false" `
    /*
        酒店RPID     */
    Rpid  *int64 `json:"rpid,omitempty" required:"false" `
    /*
        用于标示该宝贝的售卖渠道信息，允许同一个卖家酒店房型在淘宝系统发布多个售卖渠道的宝贝的价格。     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        卖家自己系统的Code，简称RateCode     */
    RateplanCode  *string `json:"rateplan_code,omitempty" required:"false" `
    /*
        卖家房型ID, 这是卖家自己系统中的房型ID 注意：需要按照规则组合     */
    OutRid  *string `json:"out_rid,omitempty" required:"false" `
    /*
        RateID     */
    RateId  *int64 `json:"rate_id,omitempty" required:"false" `
}

func (s *TaobaoXhotelRateGetRequest) SetGid(v int64) *TaobaoXhotelRateGetRequest {
    s.Gid = &v
    return s
}
func (s *TaobaoXhotelRateGetRequest) SetRpid(v int64) *TaobaoXhotelRateGetRequest {
    s.Rpid = &v
    return s
}
func (s *TaobaoXhotelRateGetRequest) SetVendor(v string) *TaobaoXhotelRateGetRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelRateGetRequest) SetRateplanCode(v string) *TaobaoXhotelRateGetRequest {
    s.RateplanCode = &v
    return s
}
func (s *TaobaoXhotelRateGetRequest) SetOutRid(v string) *TaobaoXhotelRateGetRequest {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelRateGetRequest) SetRateId(v int64) *TaobaoXhotelRateGetRequest {
    s.RateId = &v
    return s
}

func (req *TaobaoXhotelRateGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Gid != nil) {
        paramMap["gid"] = *req.Gid
    }
    if(req.Rpid != nil) {
        paramMap["rpid"] = *req.Rpid
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.RateplanCode != nil) {
        paramMap["rateplan_code"] = *req.RateplanCode
    }
    if(req.OutRid != nil) {
        paramMap["out_rid"] = *req.OutRid
    }
    if(req.RateId != nil) {
        paramMap["rate_id"] = *req.RateId
    }
    return paramMap
}

func (req *TaobaoXhotelRateGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}