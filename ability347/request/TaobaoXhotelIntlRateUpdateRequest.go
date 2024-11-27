package request

import (
        "topsdk/ability347/domain"
        "topsdk/util"
    )

type TaobaoXhotelIntlRateUpdateRequest struct {
    /*
        rate更新参数     */
    UpdateRateParam  *domain.TaobaoXhotelIntlRateUpdateUpdateRateParam `json:"update_rate_param,omitempty" required:"false" `
}

func (s *TaobaoXhotelIntlRateUpdateRequest) SetUpdateRateParam(v domain.TaobaoXhotelIntlRateUpdateUpdateRateParam) *TaobaoXhotelIntlRateUpdateRequest {
    s.UpdateRateParam = &v
    return s
}

func (req *TaobaoXhotelIntlRateUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.UpdateRateParam != nil) {
        paramMap["update_rate_param"] = util.ConvertStruct(*req.UpdateRateParam)
    }
    return paramMap
}

func (req *TaobaoXhotelIntlRateUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}