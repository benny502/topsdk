package request

import (
        "topsdk/ability347/domain"
        "topsdk/util"
    )

type TaobaoXhotelBnbpromoAddRequest struct {
    /*
        营销类型     */
    PromoInfo  *domain.TaobaoXhotelBnbpromoAddPromoInfo `json:"promo_info" required:"true" `
}

func (s *TaobaoXhotelBnbpromoAddRequest) SetPromoInfo(v domain.TaobaoXhotelBnbpromoAddPromoInfo) *TaobaoXhotelBnbpromoAddRequest {
    s.PromoInfo = &v
    return s
}

func (req *TaobaoXhotelBnbpromoAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PromoInfo != nil) {
        paramMap["promo_info"] = util.ConvertStruct(*req.PromoInfo)
    }
    return paramMap
}

func (req *TaobaoXhotelBnbpromoAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}