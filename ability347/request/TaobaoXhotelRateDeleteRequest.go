package request


type TaobaoXhotelRateDeleteRequest struct {
    /*
        rateId     */
    RateId  *int64 `json:"rate_id,omitempty" required:"false" `
    /*
        ratepland标识     */
    RpId  *int64 `json:"rp_id,omitempty" required:"false" `
    /*
        系统商，一般不用填写，使用须申请     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        商家价格政策编码     */
    RateplanCode  *string `json:"rateplan_code,omitempty" required:"false" `
    /*
        商家房型ID     */
    OutRid  *string `json:"out_rid,omitempty" required:"false" `
    /*
        房型id     */
    Gid  *string `json:"gid,omitempty" required:"false" `
}

func (s *TaobaoXhotelRateDeleteRequest) SetRateId(v int64) *TaobaoXhotelRateDeleteRequest {
    s.RateId = &v
    return s
}
func (s *TaobaoXhotelRateDeleteRequest) SetRpId(v int64) *TaobaoXhotelRateDeleteRequest {
    s.RpId = &v
    return s
}
func (s *TaobaoXhotelRateDeleteRequest) SetVendor(v string) *TaobaoXhotelRateDeleteRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelRateDeleteRequest) SetRateplanCode(v string) *TaobaoXhotelRateDeleteRequest {
    s.RateplanCode = &v
    return s
}
func (s *TaobaoXhotelRateDeleteRequest) SetOutRid(v string) *TaobaoXhotelRateDeleteRequest {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelRateDeleteRequest) SetGid(v string) *TaobaoXhotelRateDeleteRequest {
    s.Gid = &v
    return s
}

func (req *TaobaoXhotelRateDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RateId != nil) {
        paramMap["rate_id"] = *req.RateId
    }
    if(req.RpId != nil) {
        paramMap["rp_id"] = *req.RpId
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
    if(req.Gid != nil) {
        paramMap["gid"] = *req.Gid
    }
    return paramMap
}

func (req *TaobaoXhotelRateDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}