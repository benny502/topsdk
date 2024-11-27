package request


type AlitripTravelHotelticketOrderRefundRequest struct {
    /*
        系统商订单号     */
    OrderId  *string `json:"order_id,omitempty" required:"false" `
    /*
        退款失败原因     */
    FailMsg  *string `json:"fail_msg,omitempty" required:"false" `
    /*
        退款结果状态 1:退款成功  2:退款失败     */
    Status  *int64 `json:"status" required:"true" `
    /*
        飞猪订单号     */
    FliggyOrderId  *string `json:"fliggy_order_id" required:"true" `
}

func (s *AlitripTravelHotelticketOrderRefundRequest) SetOrderId(v string) *AlitripTravelHotelticketOrderRefundRequest {
    s.OrderId = &v
    return s
}
func (s *AlitripTravelHotelticketOrderRefundRequest) SetFailMsg(v string) *AlitripTravelHotelticketOrderRefundRequest {
    s.FailMsg = &v
    return s
}
func (s *AlitripTravelHotelticketOrderRefundRequest) SetStatus(v int64) *AlitripTravelHotelticketOrderRefundRequest {
    s.Status = &v
    return s
}
func (s *AlitripTravelHotelticketOrderRefundRequest) SetFliggyOrderId(v string) *AlitripTravelHotelticketOrderRefundRequest {
    s.FliggyOrderId = &v
    return s
}

func (req *AlitripTravelHotelticketOrderRefundRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.FailMsg != nil) {
        paramMap["fail_msg"] = *req.FailMsg
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.FliggyOrderId != nil) {
        paramMap["fliggy_order_id"] = *req.FliggyOrderId
    }
    return paramMap
}

func (req *AlitripTravelHotelticketOrderRefundRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}