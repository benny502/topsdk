package request


type TaobaoXhotelBnbownerDeleteRequest struct {
    /*
        对接系统商名称：可为空不要乱填，需要申请后使用，默认taobao defalutValue��taobao    */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        房东Id，系统商outer_id     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
}

func (s *TaobaoXhotelBnbownerDeleteRequest) SetVendor(v string) *TaobaoXhotelBnbownerDeleteRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelBnbownerDeleteRequest) SetOuterId(v string) *TaobaoXhotelBnbownerDeleteRequest {
    s.OuterId = &v
    return s
}

func (req *TaobaoXhotelBnbownerDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    return paramMap
}

func (req *TaobaoXhotelBnbownerDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}