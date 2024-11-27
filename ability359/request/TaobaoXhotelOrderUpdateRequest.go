package request


type TaobaoXhotelOrderUpdateRequest struct {
    /*
        订单号     */
    Tid  *int64 `json:"tid" required:"true" `
    /*
        操作的类型：1.确认无房（取消预订，710发送短信提醒买家申请退款）2.确认预订 3.入住 4.离店 5.noshow 6.关单     */
    OptType  *int64 `json:"opt_type" required:"true" `
    /*
        是否把代理直签的订单同步到酒店，Y为同步，N不同步     */
    SyncToHotel  *string `json:"sync_to_hotel,omitempty" required:"false" `
    /*
        退款费用 defalutValue��0    */
    RefundFee  *int64 `json:"refund_fee,omitempty" required:"false" `
    /*
        取消类型，6 代表的是用户取消，reasonType=7代表的是小二协商 defalutValue��0    */
    ReasonType  *int64 `json:"reason_type,omitempty" required:"false" `
    /*
        开票金额 defalutValue��-1    */
    InvoiceAmount  *int64 `json:"invoice_amount,omitempty" required:"false" `
    /*
        外部订单号（酒店确认号）     */
    ConfirmCode  *string `json:"confirm_code,omitempty" required:"false" `
    /*
        错误码 defalutValue��0    */
    HotelReverseReasonCode  *int64 `json:"hotel_reverse_reason_code,omitempty" required:"false" `
    /*
        拒单原因描述     */
    HotelReverseReasonDesc  *string `json:"hotel_reverse_reason_desc,omitempty" required:"false" `
    /*
        补充细节     */
    HotelReverseReasonDetail  *string `json:"hotel_reverse_reason_detail,omitempty" required:"false" `
}

func (s *TaobaoXhotelOrderUpdateRequest) SetTid(v int64) *TaobaoXhotelOrderUpdateRequest {
    s.Tid = &v
    return s
}
func (s *TaobaoXhotelOrderUpdateRequest) SetOptType(v int64) *TaobaoXhotelOrderUpdateRequest {
    s.OptType = &v
    return s
}
func (s *TaobaoXhotelOrderUpdateRequest) SetSyncToHotel(v string) *TaobaoXhotelOrderUpdateRequest {
    s.SyncToHotel = &v
    return s
}
func (s *TaobaoXhotelOrderUpdateRequest) SetRefundFee(v int64) *TaobaoXhotelOrderUpdateRequest {
    s.RefundFee = &v
    return s
}
func (s *TaobaoXhotelOrderUpdateRequest) SetReasonType(v int64) *TaobaoXhotelOrderUpdateRequest {
    s.ReasonType = &v
    return s
}
func (s *TaobaoXhotelOrderUpdateRequest) SetInvoiceAmount(v int64) *TaobaoXhotelOrderUpdateRequest {
    s.InvoiceAmount = &v
    return s
}
func (s *TaobaoXhotelOrderUpdateRequest) SetConfirmCode(v string) *TaobaoXhotelOrderUpdateRequest {
    s.ConfirmCode = &v
    return s
}
func (s *TaobaoXhotelOrderUpdateRequest) SetHotelReverseReasonCode(v int64) *TaobaoXhotelOrderUpdateRequest {
    s.HotelReverseReasonCode = &v
    return s
}
func (s *TaobaoXhotelOrderUpdateRequest) SetHotelReverseReasonDesc(v string) *TaobaoXhotelOrderUpdateRequest {
    s.HotelReverseReasonDesc = &v
    return s
}
func (s *TaobaoXhotelOrderUpdateRequest) SetHotelReverseReasonDetail(v string) *TaobaoXhotelOrderUpdateRequest {
    s.HotelReverseReasonDetail = &v
    return s
}

func (req *TaobaoXhotelOrderUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    if(req.OptType != nil) {
        paramMap["opt_type"] = *req.OptType
    }
    if(req.SyncToHotel != nil) {
        paramMap["sync_to_hotel"] = *req.SyncToHotel
    }
    if(req.RefundFee != nil) {
        paramMap["refund_fee"] = *req.RefundFee
    }
    if(req.ReasonType != nil) {
        paramMap["reason_type"] = *req.ReasonType
    }
    if(req.InvoiceAmount != nil) {
        paramMap["invoice_amount"] = *req.InvoiceAmount
    }
    if(req.ConfirmCode != nil) {
        paramMap["confirm_code"] = *req.ConfirmCode
    }
    if(req.HotelReverseReasonCode != nil) {
        paramMap["hotel_reverse_reason_code"] = *req.HotelReverseReasonCode
    }
    if(req.HotelReverseReasonDesc != nil) {
        paramMap["hotel_reverse_reason_desc"] = *req.HotelReverseReasonDesc
    }
    if(req.HotelReverseReasonDetail != nil) {
        paramMap["hotel_reverse_reason_detail"] = *req.HotelReverseReasonDetail
    }
    return paramMap
}

func (req *TaobaoXhotelOrderUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}