package request

import (
	"github.com/benny502/topsdk/ability347/domain"
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelBnbpromoUnbindRequest struct {
	/*
	   营销活动code     */
	ActivityCode *string `json:"activity_code,omitempty" required:"false" `
	/*
	   营销     */
	RateInfos *[]domain.TaobaoXhotelBnbpromoUnbindPromoRateInfo `json:"rate_infos,omitempty" required:"false" `
}

func (s *TaobaoXhotelBnbpromoUnbindRequest) SetActivityCode(v string) *TaobaoXhotelBnbpromoUnbindRequest {
	s.ActivityCode = &v
	return s
}
func (s *TaobaoXhotelBnbpromoUnbindRequest) SetRateInfos(v []domain.TaobaoXhotelBnbpromoUnbindPromoRateInfo) *TaobaoXhotelBnbpromoUnbindRequest {
	s.RateInfos = &v
	return s
}

func (req *TaobaoXhotelBnbpromoUnbindRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.ActivityCode != nil {
		paramMap["activity_code"] = *req.ActivityCode
	}
	if req.RateInfos != nil {
		paramMap["rate_infos"] = util.ConvertStructList(*req.RateInfos)
	}
	return paramMap
}

func (req *TaobaoXhotelBnbpromoUnbindRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
