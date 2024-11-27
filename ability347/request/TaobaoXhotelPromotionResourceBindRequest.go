package request

import (
	"github.com/benny502/topsdk/ability347/domain"
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelPromotionResourceBindRequest struct {
	/*
	   资源绑定请求     */
	BindPromotionResourceParam *domain.TaobaoXhotelPromotionResourceBindBindPromotionResourceParam `json:"bind_promotion_resource_param,omitempty" required:"false" `
}

func (s *TaobaoXhotelPromotionResourceBindRequest) SetBindPromotionResourceParam(v domain.TaobaoXhotelPromotionResourceBindBindPromotionResourceParam) *TaobaoXhotelPromotionResourceBindRequest {
	s.BindPromotionResourceParam = &v
	return s
}

func (req *TaobaoXhotelPromotionResourceBindRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.BindPromotionResourceParam != nil {
		paramMap["bind_promotion_resource_param"] = util.ConvertStruct(*req.BindPromotionResourceParam)
	}
	return paramMap
}

func (req *TaobaoXhotelPromotionResourceBindRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
