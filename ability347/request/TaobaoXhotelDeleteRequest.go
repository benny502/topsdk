package request


type TaobaoXhotelDeleteRequest struct {
    /*
        酒店id，传参方式  hid   或者   outer_id+vendor     */
    Hid  *int64 `json:"hid,omitempty" required:"false" `
    /*
        酒店vendor defalutValue��taobao    */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        酒店编码     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
}

func (s *TaobaoXhotelDeleteRequest) SetHid(v int64) *TaobaoXhotelDeleteRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelDeleteRequest) SetVendor(v string) *TaobaoXhotelDeleteRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelDeleteRequest) SetOuterId(v string) *TaobaoXhotelDeleteRequest {
    s.OuterId = &v
    return s
}

func (req *TaobaoXhotelDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Hid != nil) {
        paramMap["hid"] = *req.Hid
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    return paramMap
}

func (req *TaobaoXhotelDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}