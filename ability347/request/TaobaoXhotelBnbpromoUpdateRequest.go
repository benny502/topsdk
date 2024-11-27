package request

import (
        "topsdk/ability347/domain"
        "topsdk/util"
    )

type TaobaoXhotelBnbpromoUpdateRequest struct {
    /*
        更新营销活动的入参     */
    UpdatePromoParam  *domain.TaobaoXhotelBnbpromoUpdateUpdatePromoParam `json:"update_promo_param" required:"true" `
}

func (s *TaobaoXhotelBnbpromoUpdateRequest) SetUpdatePromoParam(v domain.TaobaoXhotelBnbpromoUpdateUpdatePromoParam) *TaobaoXhotelBnbpromoUpdateRequest {
    s.UpdatePromoParam = &v
    return s
}

func (req *TaobaoXhotelBnbpromoUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.UpdatePromoParam != nil) {
        paramMap["update_promo_param"] = util.ConvertStruct(*req.UpdatePromoParam)
    }
    return paramMap
}

func (req *TaobaoXhotelBnbpromoUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}