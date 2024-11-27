package request


type TaobaoXhotelBnbcommonAddRequest struct {
    /*
        参数     */
    Param  *string `json:"param,omitempty" required:"false" `
    /*
        业务场景     */
    Scene  *string `json:"scene,omitempty" required:"false" `
}

func (s *TaobaoXhotelBnbcommonAddRequest) SetParam(v string) *TaobaoXhotelBnbcommonAddRequest {
    s.Param = &v
    return s
}
func (s *TaobaoXhotelBnbcommonAddRequest) SetScene(v string) *TaobaoXhotelBnbcommonAddRequest {
    s.Scene = &v
    return s
}

func (req *TaobaoXhotelBnbcommonAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Param != nil) {
        paramMap["param"] = *req.Param
    }
    if(req.Scene != nil) {
        paramMap["scene"] = *req.Scene
    }
    return paramMap
}

func (req *TaobaoXhotelBnbcommonAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}