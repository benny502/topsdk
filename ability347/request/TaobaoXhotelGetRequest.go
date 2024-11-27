package request


type TaobaoXhotelGetRequest struct {
    /*
        废弃，请使用outer_id     */
    Hid  *int64 `json:"hid,omitempty" required:"false" `
    /*
        卖家系统中的酒店ID。     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        系统商，一般不用填写，使用须申请     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        是否需要在售状态(默认false不需要) defalutValue��false    */
    NeedSaleInfo  *bool `json:"need_sale_info,omitempty" required:"false" `
}

func (s *TaobaoXhotelGetRequest) SetHid(v int64) *TaobaoXhotelGetRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelGetRequest) SetOuterId(v string) *TaobaoXhotelGetRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelGetRequest) SetVendor(v string) *TaobaoXhotelGetRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelGetRequest) SetNeedSaleInfo(v bool) *TaobaoXhotelGetRequest {
    s.NeedSaleInfo = &v
    return s
}

func (req *TaobaoXhotelGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Hid != nil) {
        paramMap["hid"] = *req.Hid
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.NeedSaleInfo != nil) {
        paramMap["need_sale_info"] = *req.NeedSaleInfo
    }
    return paramMap
}

func (req *TaobaoXhotelGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}