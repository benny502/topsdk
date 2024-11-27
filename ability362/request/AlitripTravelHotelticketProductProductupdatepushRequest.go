package request

import (
	"github.com/benny502/topsdk/ability362/domain"
	"github.com/benny502/topsdk/util"
)

type AlitripTravelHotelticketProductProductupdatepushRequest struct {
	/*
	   系统商分配给飞猪卖家的账号     */
	AccessKey *string `json:"access_key" required:"true" `
	/*
	   变更的产品信息     */
	ProductUpdates *[]domain.AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO `json:"product_updates" required:"true" `
}

func (s *AlitripTravelHotelticketProductProductupdatepushRequest) SetAccessKey(v string) *AlitripTravelHotelticketProductProductupdatepushRequest {
	s.AccessKey = &v
	return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushRequest) SetProductUpdates(v []domain.AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO) *AlitripTravelHotelticketProductProductupdatepushRequest {
	s.ProductUpdates = &v
	return s
}

func (req *AlitripTravelHotelticketProductProductupdatepushRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.AccessKey != nil {
		paramMap["access_key"] = *req.AccessKey
	}
	if req.ProductUpdates != nil {
		paramMap["product_updates"] = util.ConvertStructList(*req.ProductUpdates)
	}
	return paramMap
}

func (req *AlitripTravelHotelticketProductProductupdatepushRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
