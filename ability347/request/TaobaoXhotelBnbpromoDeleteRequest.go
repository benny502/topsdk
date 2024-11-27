package request


type TaobaoXhotelBnbpromoDeleteRequest struct {
    /*
        外部活动code     */
    OuterActivityCode  *string `json:"outer_activity_code,omitempty" required:"false" `
}

func (s *TaobaoXhotelBnbpromoDeleteRequest) SetOuterActivityCode(v string) *TaobaoXhotelBnbpromoDeleteRequest {
    s.OuterActivityCode = &v
    return s
}

func (req *TaobaoXhotelBnbpromoDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OuterActivityCode != nil) {
        paramMap["outer_activity_code"] = *req.OuterActivityCode
    }
    return paramMap
}

func (req *TaobaoXhotelBnbpromoDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}