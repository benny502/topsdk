package request

import (
        "topsdk/ability347/domain"
        "topsdk/util"
    )

type TaobaoXhotelBnbownerAddRequest struct {
    /*
        添加房东信息的对象     */
    AddOwnerParam  *domain.TaobaoXhotelBnbownerAddAddOwnerParam `json:"add_owner_param,omitempty" required:"false" `
}

func (s *TaobaoXhotelBnbownerAddRequest) SetAddOwnerParam(v domain.TaobaoXhotelBnbownerAddAddOwnerParam) *TaobaoXhotelBnbownerAddRequest {
    s.AddOwnerParam = &v
    return s
}

func (req *TaobaoXhotelBnbownerAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AddOwnerParam != nil) {
        paramMap["add_owner_param"] = util.ConvertStruct(*req.AddOwnerParam)
    }
    return paramMap
}

func (req *TaobaoXhotelBnbownerAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}