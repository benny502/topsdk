package request


type TaobaoXhotelBnbhouseDeleteRequest struct {
    /*
        门店Id，传参方式为hid或outer_id+vendor     */
    Hid  *int64 `json:"hid,omitempty" required:"false" `
    /*
        对接系统商名称：可为空不要乱填，需要申请后使用 defalutValue��taobao    */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        门店Id，系统商outer_id     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
}

func (s *TaobaoXhotelBnbhouseDeleteRequest) SetHid(v int64) *TaobaoXhotelBnbhouseDeleteRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelBnbhouseDeleteRequest) SetVendor(v string) *TaobaoXhotelBnbhouseDeleteRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelBnbhouseDeleteRequest) SetOuterId(v string) *TaobaoXhotelBnbhouseDeleteRequest {
    s.OuterId = &v
    return s
}

func (req *TaobaoXhotelBnbhouseDeleteRequest) ToMap() map[string]interface{} {
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

func (req *TaobaoXhotelBnbhouseDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}