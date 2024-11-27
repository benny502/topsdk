package request


type TaobaoXhotelGetEntityByTagRequest struct {
    /*
        标签     */
    Tag  *int64 `json:"tag,omitempty" required:"false" `
    /*
        查询token，填入上一页查询的返回结果，只能按顺序单线程查询     */
    TokenStr  *string `json:"token_str,omitempty" required:"false" `
}

func (s *TaobaoXhotelGetEntityByTagRequest) SetTag(v int64) *TaobaoXhotelGetEntityByTagRequest {
    s.Tag = &v
    return s
}
func (s *TaobaoXhotelGetEntityByTagRequest) SetTokenStr(v string) *TaobaoXhotelGetEntityByTagRequest {
    s.TokenStr = &v
    return s
}

func (req *TaobaoXhotelGetEntityByTagRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Tag != nil) {
        paramMap["tag"] = *req.Tag
    }
    if(req.TokenStr != nil) {
        paramMap["token_str"] = *req.TokenStr
    }
    return paramMap
}

func (req *TaobaoXhotelGetEntityByTagRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}