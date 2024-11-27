package request


type TaobaoXhotelBnbpromoGetRequest struct {
    /*
        外部活动code     */
    OuterActivityCode  *string `json:"outer_activity_code,omitempty" required:"false" `
}

func (s *TaobaoXhotelBnbpromoGetRequest) SetOuterActivityCode(v string) *TaobaoXhotelBnbpromoGetRequest {
    s.OuterActivityCode = &v
    return s
}

func (req *TaobaoXhotelBnbpromoGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OuterActivityCode != nil) {
        paramMap["outer_activity_code"] = *req.OuterActivityCode
    }
    return paramMap
}

func (req *TaobaoXhotelBnbpromoGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}