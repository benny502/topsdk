package request

import (
	"github.com/benny502/topsdk/ability359/domain"
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelOrderUpdateConfirmcodeRequest struct {
	/*
	   系统入参     */
	Param *domain.TaobaoXhotelOrderUpdateConfirmcodeUpdateOrderConfirmCodeParam `json:"param,omitempty" required:"false" `
}

func (s *TaobaoXhotelOrderUpdateConfirmcodeRequest) SetParam(v domain.TaobaoXhotelOrderUpdateConfirmcodeUpdateOrderConfirmCodeParam) *TaobaoXhotelOrderUpdateConfirmcodeRequest {
	s.Param = &v
	return s
}

func (req *TaobaoXhotelOrderUpdateConfirmcodeRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Param != nil {
		paramMap["param"] = util.ConvertStruct(*req.Param)
	}
	return paramMap
}

func (req *TaobaoXhotelOrderUpdateConfirmcodeRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
