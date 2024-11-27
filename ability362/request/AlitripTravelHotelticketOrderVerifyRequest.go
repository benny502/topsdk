package request

import (
        "topsdk/ability362/domain"
        "topsdk/util"
    )

type AlitripTravelHotelticketOrderVerifyRequest struct {
    /*
        扩展参数     */
    ExtendParams  *string `json:"extend_params,omitempty" required:"false" `
    /*
        系统商订单号     */
    OrderId  *string `json:"order_id,omitempty" required:"false" `
    /*
        凭证信息     */
    Vouchers  *[]domain.AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO `json:"vouchers,omitempty" required:"false" `
    /*
        飞猪订单号     */
    FliggyOrderId  *string `json:"fliggy_order_id" required:"true" `
}

func (s *AlitripTravelHotelticketOrderVerifyRequest) SetExtendParams(v string) *AlitripTravelHotelticketOrderVerifyRequest {
    s.ExtendParams = &v
    return s
}
func (s *AlitripTravelHotelticketOrderVerifyRequest) SetOrderId(v string) *AlitripTravelHotelticketOrderVerifyRequest {
    s.OrderId = &v
    return s
}
func (s *AlitripTravelHotelticketOrderVerifyRequest) SetVouchers(v []domain.AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO) *AlitripTravelHotelticketOrderVerifyRequest {
    s.Vouchers = &v
    return s
}
func (s *AlitripTravelHotelticketOrderVerifyRequest) SetFliggyOrderId(v string) *AlitripTravelHotelticketOrderVerifyRequest {
    s.FliggyOrderId = &v
    return s
}

func (req *AlitripTravelHotelticketOrderVerifyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ExtendParams != nil) {
        paramMap["extend_params"] = *req.ExtendParams
    }
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.Vouchers != nil) {
        paramMap["vouchers"] = util.ConvertStructList(*req.Vouchers)
    }
    if(req.FliggyOrderId != nil) {
        paramMap["fliggy_order_id"] = *req.FliggyOrderId
    }
    return paramMap
}

func (req *AlitripTravelHotelticketOrderVerifyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}