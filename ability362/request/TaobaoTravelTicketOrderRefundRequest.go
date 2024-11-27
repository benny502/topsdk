package request


type TaobaoTravelTicketOrderRefundRequest struct {
    /*
        退票数量     */
    RefundNum  *int64 `json:"refund_num,omitempty" required:"false" `
    /*
        下单时订单ID     */
    OrderId  *int64 `json:"order_id" required:"true" `
    /*
        退票结果；1: 退票成功；2: 退票失败     */
    RefundStatus  *int64 `json:"refund_status" required:"true" `
    /*
        退票失败理由     */
    RefundFailureReason  *string `json:"refund_failure_reason,omitempty" required:"false" `
    /*
        退票的批次号     */
    RefundBatchNo  *string `json:"refund_batch_no,omitempty" required:"false" `
    /*
        退款类型：1：全额退；3：部分退；默认为1     */
    RefundType  *int64 `json:"refund_type,omitempty" required:"false" `
    /*
        飞猪退款流水号，部分退必填     */
    FliggyRefundFlowId  *string `json:"fliggy_refund_flow_id,omitempty" required:"false" `
}

func (s *TaobaoTravelTicketOrderRefundRequest) SetRefundNum(v int64) *TaobaoTravelTicketOrderRefundRequest {
    s.RefundNum = &v
    return s
}
func (s *TaobaoTravelTicketOrderRefundRequest) SetOrderId(v int64) *TaobaoTravelTicketOrderRefundRequest {
    s.OrderId = &v
    return s
}
func (s *TaobaoTravelTicketOrderRefundRequest) SetRefundStatus(v int64) *TaobaoTravelTicketOrderRefundRequest {
    s.RefundStatus = &v
    return s
}
func (s *TaobaoTravelTicketOrderRefundRequest) SetRefundFailureReason(v string) *TaobaoTravelTicketOrderRefundRequest {
    s.RefundFailureReason = &v
    return s
}
func (s *TaobaoTravelTicketOrderRefundRequest) SetRefundBatchNo(v string) *TaobaoTravelTicketOrderRefundRequest {
    s.RefundBatchNo = &v
    return s
}
func (s *TaobaoTravelTicketOrderRefundRequest) SetRefundType(v int64) *TaobaoTravelTicketOrderRefundRequest {
    s.RefundType = &v
    return s
}
func (s *TaobaoTravelTicketOrderRefundRequest) SetFliggyRefundFlowId(v string) *TaobaoTravelTicketOrderRefundRequest {
    s.FliggyRefundFlowId = &v
    return s
}

func (req *TaobaoTravelTicketOrderRefundRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefundNum != nil) {
        paramMap["refund_num"] = *req.RefundNum
    }
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.RefundStatus != nil) {
        paramMap["refund_status"] = *req.RefundStatus
    }
    if(req.RefundFailureReason != nil) {
        paramMap["refund_failure_reason"] = *req.RefundFailureReason
    }
    if(req.RefundBatchNo != nil) {
        paramMap["refund_batch_no"] = *req.RefundBatchNo
    }
    if(req.RefundType != nil) {
        paramMap["refund_type"] = *req.RefundType
    }
    if(req.FliggyRefundFlowId != nil) {
        paramMap["fliggy_refund_flow_id"] = *req.FliggyRefundFlowId
    }
    return paramMap
}

func (req *TaobaoTravelTicketOrderRefundRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}