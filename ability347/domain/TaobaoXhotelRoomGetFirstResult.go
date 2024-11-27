package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelRoomGetFirstResult struct {
	/*
	   发票类型为其他时的发票描述,不能超过30个汉字，60个字符     */
	ReceiptOtherTypeDesc *string `json:"receipt_other_type_desc,omitempty" `

	/*
	   发票类型。A,B。分别代表： A:酒店住宿发票,B:其他     */
	ReceiptType *string `json:"receipt_type,omitempty" `

	/*
	   酒店商品是否提供发票     */
	HasReceipt *bool `json:"has_receipt,omitempty" `

	/*
	   酒店商品图片Url。多个url用逗号隔开     */
	PicUrls *string `json:"pic_urls,omitempty" `

	/*
	   宝贝描述     */
	Desc *string `json:"desc,omitempty" `

	/*
	   购买须知     */
	Guide *string `json:"guide,omitempty" `

	/*
	   宝贝名称     */
	Title *string `json:"title,omitempty" `

	/*
	   rid房型id     */
	Rid *int64 `json:"rid,omitempty" `

	/*
	   hid酒店id     */
	Hid *int64 `json:"hid,omitempty" `

	/*
	   iid淘宝商品id     */
	Iid *int64 `json:"iid,omitempty" `

	/*
	   gid酒店商品id     */
	Gid *int64 `json:"gid,omitempty" `

	/*
	   发票说明，不能超过100个汉字,200个字符。     */
	ReceiptInfo *string `json:"receipt_info,omitempty" `

	/*
	   库存日历     */
	Inventory *string `json:"inventory,omitempty" `

	/*
	   宝贝状态。1：上架。2：下架。3：删除     */
	Status *int64 `json:"status,omitempty" `

	/*
	   橱窗推荐     */
	Recommend *bool `json:"recommend,omitempty" `

	/*
	   创建时间     */
	CreatedTime *util.LocalTime `json:"created_time,omitempty" `

	/*
	   修改时间     */
	ModifiedTime *util.LocalTime `json:"modified_time,omitempty" `

	/*
	   extend_info1     */
	ExtendInfo1 *string `json:"extend_info1,omitempty" `

	/*
	   extend_info2     */
	ExtendInfo2 *string `json:"extend_info2,omitempty" `

	/*
	   extend_info3     */
	ExtendInfo3 *string `json:"extend_info3,omitempty" `

	/*
	   卖家渠道     */
	Vendor *string `json:"vendor,omitempty" `

	/*
	   out_rid     */
	OutRid *string `json:"out_rid,omitempty" `

	/*
	   商品下架原因     */
	DownReason *string `json:"down_reason,omitempty" `

	/*
	   switchCalendar     */
	SwitchCalendar *string `json:"switch_calendar,omitempty" `
}

func (s *TaobaoXhotelRoomGetFirstResult) SetReceiptOtherTypeDesc(v string) *TaobaoXhotelRoomGetFirstResult {
	s.ReceiptOtherTypeDesc = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetReceiptType(v string) *TaobaoXhotelRoomGetFirstResult {
	s.ReceiptType = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetHasReceipt(v bool) *TaobaoXhotelRoomGetFirstResult {
	s.HasReceipt = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetPicUrls(v string) *TaobaoXhotelRoomGetFirstResult {
	s.PicUrls = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetDesc(v string) *TaobaoXhotelRoomGetFirstResult {
	s.Desc = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetGuide(v string) *TaobaoXhotelRoomGetFirstResult {
	s.Guide = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetTitle(v string) *TaobaoXhotelRoomGetFirstResult {
	s.Title = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetRid(v int64) *TaobaoXhotelRoomGetFirstResult {
	s.Rid = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetHid(v int64) *TaobaoXhotelRoomGetFirstResult {
	s.Hid = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetIid(v int64) *TaobaoXhotelRoomGetFirstResult {
	s.Iid = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetGid(v int64) *TaobaoXhotelRoomGetFirstResult {
	s.Gid = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetReceiptInfo(v string) *TaobaoXhotelRoomGetFirstResult {
	s.ReceiptInfo = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetInventory(v string) *TaobaoXhotelRoomGetFirstResult {
	s.Inventory = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetStatus(v int64) *TaobaoXhotelRoomGetFirstResult {
	s.Status = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetRecommend(v bool) *TaobaoXhotelRoomGetFirstResult {
	s.Recommend = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetCreatedTime(v util.LocalTime) *TaobaoXhotelRoomGetFirstResult {
	s.CreatedTime = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetModifiedTime(v util.LocalTime) *TaobaoXhotelRoomGetFirstResult {
	s.ModifiedTime = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetExtendInfo1(v string) *TaobaoXhotelRoomGetFirstResult {
	s.ExtendInfo1 = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetExtendInfo2(v string) *TaobaoXhotelRoomGetFirstResult {
	s.ExtendInfo2 = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetExtendInfo3(v string) *TaobaoXhotelRoomGetFirstResult {
	s.ExtendInfo3 = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetVendor(v string) *TaobaoXhotelRoomGetFirstResult {
	s.Vendor = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetOutRid(v string) *TaobaoXhotelRoomGetFirstResult {
	s.OutRid = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetDownReason(v string) *TaobaoXhotelRoomGetFirstResult {
	s.DownReason = &v
	return s
}
func (s *TaobaoXhotelRoomGetFirstResult) SetSwitchCalendar(v string) *TaobaoXhotelRoomGetFirstResult {
	s.SwitchCalendar = &v
	return s
}
