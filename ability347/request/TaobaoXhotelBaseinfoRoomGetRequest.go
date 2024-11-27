package request


type TaobaoXhotelBaseinfoRoomGetRequest struct {
    /*
        卖家系统中的酒店ID。     */
    OutHid  *string `json:"out_hid" required:"true" `
    /*
        用于标示该酒店发布的渠道信息     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        是否需要房价基本信息（false为不需要），默认为需要 defalutValue��true    */
    IsNeedRatePlan  *bool `json:"is_need_rate_plan,omitempty" required:"false" `
}

func (s *TaobaoXhotelBaseinfoRoomGetRequest) SetOutHid(v string) *TaobaoXhotelBaseinfoRoomGetRequest {
    s.OutHid = &v
    return s
}
func (s *TaobaoXhotelBaseinfoRoomGetRequest) SetVendor(v string) *TaobaoXhotelBaseinfoRoomGetRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelBaseinfoRoomGetRequest) SetIsNeedRatePlan(v bool) *TaobaoXhotelBaseinfoRoomGetRequest {
    s.IsNeedRatePlan = &v
    return s
}

func (req *TaobaoXhotelBaseinfoRoomGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OutHid != nil) {
        paramMap["out_hid"] = *req.OutHid
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.IsNeedRatePlan != nil) {
        paramMap["is_need_rate_plan"] = *req.IsNeedRatePlan
    }
    return paramMap
}

func (req *TaobaoXhotelBaseinfoRoomGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}