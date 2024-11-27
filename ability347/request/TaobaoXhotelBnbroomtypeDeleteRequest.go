package request


type TaobaoXhotelBnbroomtypeDeleteRequest struct {
    /*
        房型Id，系统商outer_id     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        对接系统商名称：可为空不要乱填，需要申请后使用 defalutValue��taobao    */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        房型Id，传参方式为rid或outer_id+vendor     */
    Rid  *int64 `json:"rid,omitempty" required:"false" `
}

func (s *TaobaoXhotelBnbroomtypeDeleteRequest) SetOuterId(v string) *TaobaoXhotelBnbroomtypeDeleteRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeDeleteRequest) SetVendor(v string) *TaobaoXhotelBnbroomtypeDeleteRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeDeleteRequest) SetRid(v int64) *TaobaoXhotelBnbroomtypeDeleteRequest {
    s.Rid = &v
    return s
}

func (req *TaobaoXhotelBnbroomtypeDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.Rid != nil) {
        paramMap["rid"] = *req.Rid
    }
    return paramMap
}

func (req *TaobaoXhotelBnbroomtypeDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}