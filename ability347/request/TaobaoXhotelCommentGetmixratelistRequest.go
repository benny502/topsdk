package request

import (
        "topsdk/ability347/domain"
        "topsdk/util"
    )

type TaobaoXhotelCommentGetmixratelistRequest struct {
    /*
        1     */
    ParamGetMixRateListParam  *domain.TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam `json:"param_get_mix_rate_list_param,omitempty" required:"false" `
}

func (s *TaobaoXhotelCommentGetmixratelistRequest) SetParamGetMixRateListParam(v domain.TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam) *TaobaoXhotelCommentGetmixratelistRequest {
    s.ParamGetMixRateListParam = &v
    return s
}

func (req *TaobaoXhotelCommentGetmixratelistRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ParamGetMixRateListParam != nil) {
        paramMap["param_get_mix_rate_list_param"] = util.ConvertStruct(*req.ParamGetMixRateListParam)
    }
    return paramMap
}

func (req *TaobaoXhotelCommentGetmixratelistRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}