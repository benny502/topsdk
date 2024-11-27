package request

import (
	"github.com/benny502/topsdk/ability347/domain"
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelPromotionQueryRequest struct {
	/*
	   查询活动请求     */
	QueryPromotionParam *domain.TaobaoXhotelPromotionQueryQueryPromotionParam `json:"query_promotion_param" required:"true" `
}

func (s *TaobaoXhotelPromotionQueryRequest) SetQueryPromotionParam(v domain.TaobaoXhotelPromotionQueryQueryPromotionParam) *TaobaoXhotelPromotionQueryRequest {
	s.QueryPromotionParam = &v
	return s
}

func (req *TaobaoXhotelPromotionQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryPromotionParam != nil {
		paramMap["query_promotion_param"] = util.ConvertStruct(*req.QueryPromotionParam)
	}
	return paramMap
}

func (req *TaobaoXhotelPromotionQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
