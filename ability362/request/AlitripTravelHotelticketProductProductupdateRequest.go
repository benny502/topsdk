package request

import (
        "topsdk/ability362/domain"
        "topsdk/util"
    )

type AlitripTravelHotelticketProductProductupdateRequest struct {
    /*
        系统商分配给飞猪卖家的账号     */
    AccessKey  *string `json:"access_key" required:"true" `
    /*
        变更的产品信息     */
    ProductUpdates  *[]domain.AlitripTravelHotelticketProductProductupdateProductUpdateDTO `json:"product_updates" required:"true" `
}

func (s *AlitripTravelHotelticketProductProductupdateRequest) SetAccessKey(v string) *AlitripTravelHotelticketProductProductupdateRequest {
    s.AccessKey = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdateRequest) SetProductUpdates(v []domain.AlitripTravelHotelticketProductProductupdateProductUpdateDTO) *AlitripTravelHotelticketProductProductupdateRequest {
    s.ProductUpdates = &v
    return s
}

func (req *AlitripTravelHotelticketProductProductupdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AccessKey != nil) {
        paramMap["access_key"] = *req.AccessKey
    }
    if(req.ProductUpdates != nil) {
        paramMap["product_updates"] = util.ConvertStructList(*req.ProductUpdates)
    }
    return paramMap
}

func (req *AlitripTravelHotelticketProductProductupdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}