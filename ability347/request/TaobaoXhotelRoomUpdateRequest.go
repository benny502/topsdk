package request

type TaobaoXhotelRoomUpdateRequest struct {
	/*
	   废弃，使用out_rid     */
	Gid *int64 `json:"gid,omitempty" required:"false" `
	/*
	   卖家房型ID     */
	OutRid *string `json:"out_rid,omitempty" required:"false" `
	/*
	   系统商，一般不填写，使用须申请     */
	Vendor *string `json:"vendor,omitempty" required:"false" `
	/*
	   废弃，宝贝名称展示在店铺里     */
	Title *string `json:"title,omitempty" required:"false" `
	/*
	   废弃，房型购买须知展示在PC购物路径     */
	Guide *string `json:"guide,omitempty" required:"false" `
	/*
	   废弃，宝贝描述展示在宝贝详情页面     */
	Desc *string `json:"desc,omitempty" required:"false" `
	/*
	   废弃，宝贝图片，没有默认使用标准酒店房型图片     */
	Pic *[]byte `json:"pic,omitempty" required:"false" `
	/*
	   废弃，房型是否提供发票     */
	HasReceipt *bool `json:"has_receipt,omitempty" required:"false" `
	/*
	   废弃，房型发票类型。A,B。分别代表： A:酒店住宿发票,B:其他     */
	ReceiptType *string `json:"receipt_type,omitempty" required:"false" `
	/*
	   废弃，房型发票类型为其他时的发票描述,不能超过30个字     */
	ReceiptOtherTypeDesc *string `json:"receipt_other_type_desc,omitempty" required:"false" `
	/*
	   废弃，房型发票说明，不能超过100个字     */
	ReceiptInfo *string `json:"receipt_info,omitempty" required:"false" `
	/*
	   房型共享库存日历。quota物理库存；al_quota保留房库存；sp_quota超预定库存。其中保留房库存和超预定库存需要向运营申请权限示例：[{"date":2011-01-28,"quota":10,"al_quota":2,"sp_quota":3}]     */
	Inventory *string `json:"inventory,omitempty" required:"false" `
	/*
	   房型库存开关。 1，开；2，关     */
	RoomSwitchCal *string `json:"room_switch_cal,omitempty" required:"false" `
	/*
	   超预定库存截止时间     */
	SuperbookEndTime *string `json:"superbook_end_time,omitempty" required:"false" `
	/*
	   超预定库存开始时间     */
	SuperbookStartTime *string `json:"superbook_start_time,omitempty" required:"false" `
	/*
	   保留房库存截止时间     */
	AllotmentEndTime *string `json:"allotment_end_time,omitempty" required:"false" `
	/*
	   保留房库存截止时间     */
	AllotmentStartTime *string `json:"allotment_start_time,omitempty" required:"false" `
	/*
	   宝贝状态,1上架。     */
	Status *int64 `json:"status,omitempty" required:"false" `
}

func (s *TaobaoXhotelRoomUpdateRequest) SetGid(v int64) *TaobaoXhotelRoomUpdateRequest {
	s.Gid = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetOutRid(v string) *TaobaoXhotelRoomUpdateRequest {
	s.OutRid = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetVendor(v string) *TaobaoXhotelRoomUpdateRequest {
	s.Vendor = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetTitle(v string) *TaobaoXhotelRoomUpdateRequest {
	s.Title = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetGuide(v string) *TaobaoXhotelRoomUpdateRequest {
	s.Guide = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetDesc(v string) *TaobaoXhotelRoomUpdateRequest {
	s.Desc = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetPic(v []byte) *TaobaoXhotelRoomUpdateRequest {
	s.Pic = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetHasReceipt(v bool) *TaobaoXhotelRoomUpdateRequest {
	s.HasReceipt = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetReceiptType(v string) *TaobaoXhotelRoomUpdateRequest {
	s.ReceiptType = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetReceiptOtherTypeDesc(v string) *TaobaoXhotelRoomUpdateRequest {
	s.ReceiptOtherTypeDesc = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetReceiptInfo(v string) *TaobaoXhotelRoomUpdateRequest {
	s.ReceiptInfo = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetInventory(v string) *TaobaoXhotelRoomUpdateRequest {
	s.Inventory = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetRoomSwitchCal(v string) *TaobaoXhotelRoomUpdateRequest {
	s.RoomSwitchCal = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetSuperbookEndTime(v string) *TaobaoXhotelRoomUpdateRequest {
	s.SuperbookEndTime = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetSuperbookStartTime(v string) *TaobaoXhotelRoomUpdateRequest {
	s.SuperbookStartTime = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetAllotmentEndTime(v string) *TaobaoXhotelRoomUpdateRequest {
	s.AllotmentEndTime = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetAllotmentStartTime(v string) *TaobaoXhotelRoomUpdateRequest {
	s.AllotmentStartTime = &v
	return s
}
func (s *TaobaoXhotelRoomUpdateRequest) SetStatus(v int64) *TaobaoXhotelRoomUpdateRequest {
	s.Status = &v
	return s
}

func (req *TaobaoXhotelRoomUpdateRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Gid != nil {
		paramMap["gid"] = *req.Gid
	}
	if req.OutRid != nil {
		paramMap["out_rid"] = *req.OutRid
	}
	if req.Vendor != nil {
		paramMap["vendor"] = *req.Vendor
	}
	if req.Title != nil {
		paramMap["title"] = *req.Title
	}
	if req.Guide != nil {
		paramMap["guide"] = *req.Guide
	}
	if req.Desc != nil {
		paramMap["desc"] = *req.Desc
	}
	if req.HasReceipt != nil {
		paramMap["has_receipt"] = *req.HasReceipt
	}
	if req.ReceiptType != nil {
		paramMap["receipt_type"] = *req.ReceiptType
	}
	if req.ReceiptOtherTypeDesc != nil {
		paramMap["receipt_other_type_desc"] = *req.ReceiptOtherTypeDesc
	}
	if req.ReceiptInfo != nil {
		paramMap["receipt_info"] = *req.ReceiptInfo
	}
	if req.Inventory != nil {
		paramMap["inventory"] = *req.Inventory
	}
	if req.RoomSwitchCal != nil {
		paramMap["room_switch_cal"] = *req.RoomSwitchCal
	}
	if req.SuperbookEndTime != nil {
		paramMap["superbook_end_time"] = *req.SuperbookEndTime
	}
	if req.SuperbookStartTime != nil {
		paramMap["superbook_start_time"] = *req.SuperbookStartTime
	}
	if req.AllotmentEndTime != nil {
		paramMap["allotment_end_time"] = *req.AllotmentEndTime
	}
	if req.AllotmentStartTime != nil {
		paramMap["allotment_start_time"] = *req.AllotmentStartTime
	}
	if req.Status != nil {
		paramMap["status"] = *req.Status
	}
	return paramMap
}

func (req *TaobaoXhotelRoomUpdateRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	if req.Pic != nil {
		fileMap["pic"] = *req.Pic
	}
	return fileMap
}
