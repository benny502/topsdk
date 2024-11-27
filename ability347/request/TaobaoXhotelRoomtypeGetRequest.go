package request


type TaobaoXhotelRoomtypeGetRequest struct {
    /*
        废弃，使用商家房型ID     */
    Rid  *int64 `json:"rid,omitempty" required:"false" `
    /*
        商家房型ID     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        系统商，一般不填写，使用须申请     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
}

func (s *TaobaoXhotelRoomtypeGetRequest) SetRid(v int64) *TaobaoXhotelRoomtypeGetRequest {
    s.Rid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetRequest) SetOuterId(v string) *TaobaoXhotelRoomtypeGetRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetRequest) SetVendor(v string) *TaobaoXhotelRoomtypeGetRequest {
    s.Vendor = &v
    return s
}

func (req *TaobaoXhotelRoomtypeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Rid != nil) {
        paramMap["rid"] = *req.Rid
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    return paramMap
}

func (req *TaobaoXhotelRoomtypeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}