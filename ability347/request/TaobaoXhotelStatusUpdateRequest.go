package request


type TaobaoXhotelStatusUpdateRequest struct {
    /*
        飞猪酒店id     */
    Hid  *int64 `json:"hid,omitempty" required:"false" `
    /*
        对接系统商，不填默认taobao defalutValue��taobao    */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        外部酒店id     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        （0在线，-1删除， -2停售）     */
    Status  *int64 `json:"status,omitempty" required:"false" `
}

func (s *TaobaoXhotelStatusUpdateRequest) SetHid(v int64) *TaobaoXhotelStatusUpdateRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelStatusUpdateRequest) SetVendor(v string) *TaobaoXhotelStatusUpdateRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelStatusUpdateRequest) SetOuterId(v string) *TaobaoXhotelStatusUpdateRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelStatusUpdateRequest) SetStatus(v int64) *TaobaoXhotelStatusUpdateRequest {
    s.Status = &v
    return s
}

func (req *TaobaoXhotelStatusUpdateRequest) ToMap() map[string]interface{} {
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
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    return paramMap
}

func (req *TaobaoXhotelStatusUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}