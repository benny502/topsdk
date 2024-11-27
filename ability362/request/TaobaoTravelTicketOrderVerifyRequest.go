package request

import (
	"github.com/benny502/topsdk/ability362/domain"
	"github.com/benny502/topsdk/util"
)

type TaobaoTravelTicketOrderVerifyRequest struct {
	/*
	   核销次数     */
	CheckNum *int64 `json:"check_num,omitempty" required:"false" `
	/*
	   下单订单ID     */
	OrderId *int64 `json:"order_id,omitempty" required:"false" `
	/*
	   门票取消数量     */
	ReturnNum *int64 `json:"return_num,omitempty" required:"false" `
	/*
	   门票总共允许核销次数     */
	TotalNum *int64 `json:"total_num,omitempty" required:"false" `
	/*
	   外部订单ID     */
	OutOrderId *string `json:"out_order_id,omitempty" required:"false" `
	/*
	   (新接入使用voucher_infos)用户短信会收到的确认号     */
	ConfirmCode *string `json:"confirm_code,omitempty" required:"false" `
	/*
	   使用凭证信息     */
	VoucherInfos *[]domain.TaobaoTravelTicketOrderVerifyVoucherInfoDto `json:"voucher_infos,omitempty" required:"false" `
	/*
	   供应商核销回调类型：0表示使用本次核销数量（常规），1表示使用总核销数量（已使用+本次）     */
	WriteOffType *int64 `json:"write_off_type,omitempty" required:"false" `
}

func (s *TaobaoTravelTicketOrderVerifyRequest) SetCheckNum(v int64) *TaobaoTravelTicketOrderVerifyRequest {
	s.CheckNum = &v
	return s
}
func (s *TaobaoTravelTicketOrderVerifyRequest) SetOrderId(v int64) *TaobaoTravelTicketOrderVerifyRequest {
	s.OrderId = &v
	return s
}
func (s *TaobaoTravelTicketOrderVerifyRequest) SetReturnNum(v int64) *TaobaoTravelTicketOrderVerifyRequest {
	s.ReturnNum = &v
	return s
}
func (s *TaobaoTravelTicketOrderVerifyRequest) SetTotalNum(v int64) *TaobaoTravelTicketOrderVerifyRequest {
	s.TotalNum = &v
	return s
}
func (s *TaobaoTravelTicketOrderVerifyRequest) SetOutOrderId(v string) *TaobaoTravelTicketOrderVerifyRequest {
	s.OutOrderId = &v
	return s
}
func (s *TaobaoTravelTicketOrderVerifyRequest) SetConfirmCode(v string) *TaobaoTravelTicketOrderVerifyRequest {
	s.ConfirmCode = &v
	return s
}
func (s *TaobaoTravelTicketOrderVerifyRequest) SetVoucherInfos(v []domain.TaobaoTravelTicketOrderVerifyVoucherInfoDto) *TaobaoTravelTicketOrderVerifyRequest {
	s.VoucherInfos = &v
	return s
}
func (s *TaobaoTravelTicketOrderVerifyRequest) SetWriteOffType(v int64) *TaobaoTravelTicketOrderVerifyRequest {
	s.WriteOffType = &v
	return s
}

func (req *TaobaoTravelTicketOrderVerifyRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.CheckNum != nil {
		paramMap["check_num"] = *req.CheckNum
	}
	if req.OrderId != nil {
		paramMap["order_id"] = *req.OrderId
	}
	if req.ReturnNum != nil {
		paramMap["return_num"] = *req.ReturnNum
	}
	if req.TotalNum != nil {
		paramMap["total_num"] = *req.TotalNum
	}
	if req.OutOrderId != nil {
		paramMap["out_order_id"] = *req.OutOrderId
	}
	if req.ConfirmCode != nil {
		paramMap["confirm_code"] = *req.ConfirmCode
	}
	if req.VoucherInfos != nil {
		paramMap["voucher_infos"] = util.ConvertStructList(*req.VoucherInfos)
	}
	if req.WriteOffType != nil {
		paramMap["write_off_type"] = *req.WriteOffType
	}
	return paramMap
}

func (req *TaobaoTravelTicketOrderVerifyRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
