package request

import (
        "topsdk/ability362/domain"
        "topsdk/util"
    )

type AlitripTravelHotelticketOrderCreateRequest struct {
    /*
        扩展参数 支持的key: ["hotelName" : "酒店名称", "roomName" : "房型名称", "productName" : "产品名称",  "desc" : "描述"] value字符长度不超过100     */
    ExtendParams  *string `json:"extend_params,omitempty" required:"false" `
    /*
        系统商订单号     */
    OrderId  *string `json:"order_id,omitempty" required:"false" `
    /*
        创单出票失败原因信息     */
    FailMsg  *string `json:"fail_msg,omitempty" required:"false" `
    /*
        凭证信息 无数据时传空集合     */
    Vouchers  *[]domain.AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO `json:"vouchers,omitempty" required:"false" `
    /*
        创单结果状态 1：创单出票成功， 2：创单出票失败     */
    Status  *int64 `json:"status" required:"true" `
    /*
        飞猪订单号     */
    FliggyOrderId  *string `json:"fliggy_order_id" required:"true" `
}

func (s *AlitripTravelHotelticketOrderCreateRequest) SetExtendParams(v string) *AlitripTravelHotelticketOrderCreateRequest {
    s.ExtendParams = &v
    return s
}
func (s *AlitripTravelHotelticketOrderCreateRequest) SetOrderId(v string) *AlitripTravelHotelticketOrderCreateRequest {
    s.OrderId = &v
    return s
}
func (s *AlitripTravelHotelticketOrderCreateRequest) SetFailMsg(v string) *AlitripTravelHotelticketOrderCreateRequest {
    s.FailMsg = &v
    return s
}
func (s *AlitripTravelHotelticketOrderCreateRequest) SetVouchers(v []domain.AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO) *AlitripTravelHotelticketOrderCreateRequest {
    s.Vouchers = &v
    return s
}
func (s *AlitripTravelHotelticketOrderCreateRequest) SetStatus(v int64) *AlitripTravelHotelticketOrderCreateRequest {
    s.Status = &v
    return s
}
func (s *AlitripTravelHotelticketOrderCreateRequest) SetFliggyOrderId(v string) *AlitripTravelHotelticketOrderCreateRequest {
    s.FliggyOrderId = &v
    return s
}

func (req *AlitripTravelHotelticketOrderCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ExtendParams != nil) {
        paramMap["extend_params"] = *req.ExtendParams
    }
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.FailMsg != nil) {
        paramMap["fail_msg"] = *req.FailMsg
    }
    if(req.Vouchers != nil) {
        paramMap["vouchers"] = util.ConvertStructList(*req.Vouchers)
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.FliggyOrderId != nil) {
        paramMap["fliggy_order_id"] = *req.FliggyOrderId
    }
    return paramMap
}

func (req *AlitripTravelHotelticketOrderCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}