package request

import (
	"github.com/benny502/topsdk/ability347/domain"
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelPromotionCreateRequest struct {
	/*
	   促销活动对象     */
	CreatePromotionParam *domain.TaobaoXhotelPromotionCreateCreatePromotionParam `json:"create_promotion_param" required:"true" `
}

func (s *TaobaoXhotelPromotionCreateRequest) SetCreatePromotionParam(v domain.TaobaoXhotelPromotionCreateCreatePromotionParam) *TaobaoXhotelPromotionCreateRequest {
	s.CreatePromotionParam = &v
	return s
}

func (req *TaobaoXhotelPromotionCreateRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.CreatePromotionParam != nil {
		paramMap["create_promotion_param"] = util.ConvertStruct(*req.CreatePromotionParam)
	}
	return paramMap
}

func (req *TaobaoXhotelPromotionCreateRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
