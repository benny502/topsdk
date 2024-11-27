package request


type TaobaoXhotelRoomGetRequest struct {
    /*
        卖家渠道 如果gid为空，那么out_rid和vendor都不能为空。 支持通过gid或者通过out_rid和vendor来获取商品     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        外部房型id 如果gid为空，那么out_rid和vendor都不能为空 支持通过gid或者通过out_rid和vendor来获取商品     */
    OutRid  *string `json:"out_rid,omitempty" required:"false" `
    /*
        废弃     */
    Gid  *int64 `json:"gid,omitempty" required:"false" `
}

func (s *TaobaoXhotelRoomGetRequest) SetVendor(v string) *TaobaoXhotelRoomGetRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelRoomGetRequest) SetOutRid(v string) *TaobaoXhotelRoomGetRequest {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelRoomGetRequest) SetGid(v int64) *TaobaoXhotelRoomGetRequest {
    s.Gid = &v
    return s
}

func (req *TaobaoXhotelRoomGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.OutRid != nil) {
        paramMap["out_rid"] = *req.OutRid
    }
    if(req.Gid != nil) {
        paramMap["gid"] = *req.Gid
    }
    return paramMap
}

func (req *TaobaoXhotelRoomGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}