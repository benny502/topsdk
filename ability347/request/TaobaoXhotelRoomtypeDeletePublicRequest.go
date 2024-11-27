package request


type TaobaoXhotelRoomtypeDeletePublicRequest struct {
    /*
        房型rid ，传参方式：rid    或者   outer_id+vendor     */
    Rid  *int64 `json:"rid,omitempty" required:"false" `
    /*
        vendor     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        外部房型ID     */
    OuterRid  *string `json:"outer_rid,omitempty" required:"false" `
    /*
        具体操作人，比如酒店帐号、小二名称等     */
    Operator  *string `json:"operator" required:"true" `
}

func (s *TaobaoXhotelRoomtypeDeletePublicRequest) SetRid(v int64) *TaobaoXhotelRoomtypeDeletePublicRequest {
    s.Rid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeDeletePublicRequest) SetVendor(v string) *TaobaoXhotelRoomtypeDeletePublicRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelRoomtypeDeletePublicRequest) SetOuterRid(v string) *TaobaoXhotelRoomtypeDeletePublicRequest {
    s.OuterRid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeDeletePublicRequest) SetOperator(v string) *TaobaoXhotelRoomtypeDeletePublicRequest {
    s.Operator = &v
    return s
}

func (req *TaobaoXhotelRoomtypeDeletePublicRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Rid != nil) {
        paramMap["rid"] = *req.Rid
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.OuterRid != nil) {
        paramMap["outer_rid"] = *req.OuterRid
    }
    if(req.Operator != nil) {
        paramMap["operator"] = *req.Operator
    }
    return paramMap
}

func (req *TaobaoXhotelRoomtypeDeletePublicRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}