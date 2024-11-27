package request


type TaobaoXhotelServicetimeGetRequest struct {
    /*
        酒店id     */
    Hid  *int64 `json:"hid" required:"true" `
}

func (s *TaobaoXhotelServicetimeGetRequest) SetHid(v int64) *TaobaoXhotelServicetimeGetRequest {
    s.Hid = &v
    return s
}

func (req *TaobaoXhotelServicetimeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Hid != nil) {
        paramMap["hid"] = *req.Hid
    }
    return paramMap
}

func (req *TaobaoXhotelServicetimeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}