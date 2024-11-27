package request

import (
	"github.com/benny502/topsdk/ability347/domain"
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelPromotionResourceUnbindRequest struct {
	/*
	   资源解绑请求     */
	UnBindPromotionResourceParam *domain.TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceParam `json:"un_bind_promotion_resource_param" required:"true" `
}

func (s *TaobaoXhotelPromotionResourceUnbindRequest) SetUnBindPromotionResourceParam(v domain.TaobaoXhotelPromotionResourceUnbindUnBindPromotionResourceParam) *TaobaoXhotelPromotionResourceUnbindRequest {
	s.UnBindPromotionResourceParam = &v
	return s
}

func (req *TaobaoXhotelPromotionResourceUnbindRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.UnBindPromotionResourceParam != nil {
		paramMap["un_bind_promotion_resource_param"] = util.ConvertStruct(*req.UnBindPromotionResourceParam)
	}
	return paramMap
}

func (req *TaobaoXhotelPromotionResourceUnbindRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
