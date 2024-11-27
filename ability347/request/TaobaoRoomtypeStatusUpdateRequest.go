package request


type TaobaoRoomtypeStatusUpdateRequest struct {
    /*
        对接系统商，不填默认taobao defalutValue��taobao    */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        卖家房型id     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        飞猪房型id     */
    Rid  *int64 `json:"rid,omitempty" required:"false" `
    /*
        0在线，-1删除， -2停售）     */
    Status  *int64 `json:"status,omitempty" required:"false" `
}

func (s *TaobaoRoomtypeStatusUpdateRequest) SetVendor(v string) *TaobaoRoomtypeStatusUpdateRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoRoomtypeStatusUpdateRequest) SetOuterId(v string) *TaobaoRoomtypeStatusUpdateRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoRoomtypeStatusUpdateRequest) SetRid(v int64) *TaobaoRoomtypeStatusUpdateRequest {
    s.Rid = &v
    return s
}
func (s *TaobaoRoomtypeStatusUpdateRequest) SetStatus(v int64) *TaobaoRoomtypeStatusUpdateRequest {
    s.Status = &v
    return s
}

func (req *TaobaoRoomtypeStatusUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Rid != nil) {
        paramMap["rid"] = *req.Rid
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    return paramMap
}

func (req *TaobaoRoomtypeStatusUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}