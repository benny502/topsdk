package request

import (
        "topsdk/util"
    )

type TaobaoXhotelOrderSearchRequest struct {
    /*
        酒店订单oids列表，多个oid用英文逗号隔开，一次不超过20个。     */
    OrderIds  *string `json:"order_ids,omitempty" required:"false" `
    /*
        酒店订单tids列表，多个tid用英文逗号隔开，一次不超过20个。oids和tids都传的情况下默认使用tids     */
    OrderTids  *string `json:"order_tids,omitempty" required:"false" `
    /*
        订单创建时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss     */
    CreatedStart  *util.LocalTime `json:"created_start" required:"true" `
    /*
        订单创建时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于created_start或者间隔不能大于30     */
    CreatedEnd  *util.LocalTime `json:"created_end" required:"true" `
    /*
        分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。页面大小为20 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        系统商标识     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        外部订单号out_oids列表，多个oid用英文逗号隔开，一次不超过20个。     */
    OutOids  *string `json:"out_oids,omitempty" required:"false" `
    /*
        入住时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss     */
    CheckinDateStart  *util.LocalTime `json:"checkin_date_start,omitempty" required:"false" `
    /*
        入住时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于checkin_date_start或者间隔不能大于30     */
    CheckinDateEnd  *util.LocalTime `json:"checkin_date_end,omitempty" required:"false" `
    /*
        离店时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss     */
    CheckoutDateStart  *util.LocalTime `json:"checkout_date_start,omitempty" required:"false" `
    /*
        离店时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于checkout_date_start或者间隔不能大于30     */
    CheckoutDateEnd  *util.LocalTime `json:"checkout_date_end,omitempty" required:"false" `
    /*
        订单状态（可发多个，逗号隔开）     */
    TradeStatus  *string `json:"trade_status,omitempty" required:"false" `
    /*
        订单类型（true为直连，false为非直连）     */
    Direct  *bool `json:"direct,omitempty" required:"false" `
}

func (s *TaobaoXhotelOrderSearchRequest) SetOrderIds(v string) *TaobaoXhotelOrderSearchRequest {
    s.OrderIds = &v
    return s
}
func (s *TaobaoXhotelOrderSearchRequest) SetOrderTids(v string) *TaobaoXhotelOrderSearchRequest {
    s.OrderTids = &v
    return s
}
func (s *TaobaoXhotelOrderSearchRequest) SetCreatedStart(v util.LocalTime) *TaobaoXhotelOrderSearchRequest {
    s.CreatedStart = &v
    return s
}
func (s *TaobaoXhotelOrderSearchRequest) SetCreatedEnd(v util.LocalTime) *TaobaoXhotelOrderSearchRequest {
    s.CreatedEnd = &v
    return s
}
func (s *TaobaoXhotelOrderSearchRequest) SetPageNo(v int64) *TaobaoXhotelOrderSearchRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoXhotelOrderSearchRequest) SetVendor(v string) *TaobaoXhotelOrderSearchRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelOrderSearchRequest) SetOutOids(v string) *TaobaoXhotelOrderSearchRequest {
    s.OutOids = &v
    return s
}
func (s *TaobaoXhotelOrderSearchRequest) SetCheckinDateStart(v util.LocalTime) *TaobaoXhotelOrderSearchRequest {
    s.CheckinDateStart = &v
    return s
}
func (s *TaobaoXhotelOrderSearchRequest) SetCheckinDateEnd(v util.LocalTime) *TaobaoXhotelOrderSearchRequest {
    s.CheckinDateEnd = &v
    return s
}
func (s *TaobaoXhotelOrderSearchRequest) SetCheckoutDateStart(v util.LocalTime) *TaobaoXhotelOrderSearchRequest {
    s.CheckoutDateStart = &v
    return s
}
func (s *TaobaoXhotelOrderSearchRequest) SetCheckoutDateEnd(v util.LocalTime) *TaobaoXhotelOrderSearchRequest {
    s.CheckoutDateEnd = &v
    return s
}
func (s *TaobaoXhotelOrderSearchRequest) SetTradeStatus(v string) *TaobaoXhotelOrderSearchRequest {
    s.TradeStatus = &v
    return s
}
func (s *TaobaoXhotelOrderSearchRequest) SetDirect(v bool) *TaobaoXhotelOrderSearchRequest {
    s.Direct = &v
    return s
}

func (req *TaobaoXhotelOrderSearchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderIds != nil) {
        paramMap["order_ids"] = *req.OrderIds
    }
    if(req.OrderTids != nil) {
        paramMap["order_tids"] = *req.OrderTids
    }
    if(req.CreatedStart != nil) {
        paramMap["created_start"] = *req.CreatedStart
    }
    if(req.CreatedEnd != nil) {
        paramMap["created_end"] = *req.CreatedEnd
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.OutOids != nil) {
        paramMap["out_oids"] = *req.OutOids
    }
    if(req.CheckinDateStart != nil) {
        paramMap["checkin_date_start"] = *req.CheckinDateStart
    }
    if(req.CheckinDateEnd != nil) {
        paramMap["checkin_date_end"] = *req.CheckinDateEnd
    }
    if(req.CheckoutDateStart != nil) {
        paramMap["checkout_date_start"] = *req.CheckoutDateStart
    }
    if(req.CheckoutDateEnd != nil) {
        paramMap["checkout_date_end"] = *req.CheckoutDateEnd
    }
    if(req.TradeStatus != nil) {
        paramMap["trade_status"] = *req.TradeStatus
    }
    if(req.Direct != nil) {
        paramMap["direct"] = *req.Direct
    }
    return paramMap
}

func (req *TaobaoXhotelOrderSearchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}