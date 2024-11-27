package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelOrderSearchXHotelOrder struct {
	/*
	   酒店id     */
	Hid *int64 `json:"hid,omitempty" `

	/*
	   合作方订单号,最长250个字符     */
	OutOid *string `json:"out_oid,omitempty" `

	/*
	   离店时间     */
	CheckoutDate *util.LocalTime `json:"checkout_date,omitempty" `

	/*
	   入住人信息     */
	Guests *[]TaobaoXhotelOrderSearchXOrderGuest `json:"guests,omitempty" `

	/*
	   RatePlan 中的 rpid     */
	Rpid *int64 `json:"rpid,omitempty" `

	/*
	   支付类型 可选值 1：预付 5：前台面付     */
	Type *int64 `json:"type,omitempty" `

	/*
	   订单创建时间     */
	Created *util.LocalTime `json:"created,omitempty" `

	/*
	   酒店订单id     */
	Oid *int64 `json:"oid,omitempty" `

	/*
	   tid     */
	Tid *int64 `json:"tid,omitempty" `

	/*
	   买家最晚到达时间     */
	ArriveLate *util.LocalTime `json:"arrive_late,omitempty" `

	/*
	   联系人姓名     */
	ContactName *string `json:"contact_name,omitempty" `

	/*
	   入住时间     */
	CheckinDate *util.LocalTime `json:"checkin_date,omitempty" `

	/*
	   天数     */
	Nights *int64 `json:"nights,omitempty" `

	/*
	   卖家淘宝账号     */
	SellerNick *string `json:"seller_nick,omitempty" `

	/*
	   房间数     */
	RoomNumber *int64 `json:"room_number,omitempty" `

	/*
	   实付款（分）     */
	Payment *int64 `json:"payment,omitempty" `

	/*
	   联系人电话     */
	ContactPhone *string `json:"contact_phone,omitempty" `

	/*
	   买家最早到达时间     */
	ArriveEarly *util.LocalTime `json:"arrive_early,omitempty" `

	/*
	   付款时间     */
	PayTime *util.LocalTime `json:"pay_time,omitempty" `

	/*
	   交易状态。 WAIT_BUYER_PAY:预订中/等待买家付款, WAIT_SELLER_SEND_GOODS:预订中/等待卖家发货(确认), TRADE_CLOSED:结束/预订失败/交易关闭, TRADE_FINISHED:结束/交易成功, TRADE_NO_CREATE_PAY:结束/预订失败/没有创建支付宝交易, TRADE_CLOSED_BY_TAOBAO:结束/预订失败/预订被卖家关闭, TRADE_SUCCESS:交易中/预订成功/卖家已确认, TRADE_CHECKIN:交易中/预定成功/买家入住, TRADE_CHECKOUT:交易中/预定成功/买家离店, TRADE_SETTLEING:交易中/预定成功/结账中, TRADE_SETTLE_SUCCESS:结束/预定成功/结账成功     */
	TradeStatus *string `json:"trade_status,omitempty" `

	/*
	   订单修改时间     */
	Modified *util.LocalTime `json:"modified,omitempty" `

	/*
	   买家淘宝账号     */
	BuyerNick *string `json:"buyer_nick,omitempty" `

	/*
	   买家留言     */
	Message *string `json:"message,omitempty" `

	/*
	   房型id     */
	Rid *int64 `json:"rid,omitempty" `

	/*
	   结束时间     */
	EndTime *util.LocalTime `json:"end_time,omitempty" `

	/*
	   总房价（分）     */
	TotalRoomPrice *int64 `json:"total_room_price,omitempty" `

	/*
	   商品id     */
	Gid *int64 `json:"gid,omitempty" `

	/*
	   下单时每间夜的价格（分）     */
	Prices *[]int64 `json:"prices,omitempty" `

	/*
	   支付宝交易号，28位字符     */
	AlipayTradeNo *string `json:"alipay_trade_no,omitempty" `

	/*
	   预付订单使用，1未发货，2已发货，3已收货，4已经退货，8还未创建物流订单     */
	LogisticsStatus *string `json:"logistics_status,omitempty" `

	/*
	   预付订单使用，1买家已经申请退款，等待卖家同意，2卖家已经同意退款，等待买家退货，3买家已经退货，等待卖家确认收货，4退款关闭，5退款成功，6卖家拒绝退款，9没有申请退款     */
	RefundStatus *string `json:"refund_status,omitempty" `
}

func (s *TaobaoXhotelOrderSearchXHotelOrder) SetHid(v int64) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Hid = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetOutOid(v string) *TaobaoXhotelOrderSearchXHotelOrder {
	s.OutOid = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetCheckoutDate(v util.LocalTime) *TaobaoXhotelOrderSearchXHotelOrder {
	s.CheckoutDate = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetGuests(v []TaobaoXhotelOrderSearchXOrderGuest) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Guests = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetRpid(v int64) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Rpid = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetType(v int64) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Type = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetCreated(v util.LocalTime) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Created = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetOid(v int64) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Oid = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetTid(v int64) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Tid = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetArriveLate(v util.LocalTime) *TaobaoXhotelOrderSearchXHotelOrder {
	s.ArriveLate = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetContactName(v string) *TaobaoXhotelOrderSearchXHotelOrder {
	s.ContactName = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetCheckinDate(v util.LocalTime) *TaobaoXhotelOrderSearchXHotelOrder {
	s.CheckinDate = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetNights(v int64) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Nights = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetSellerNick(v string) *TaobaoXhotelOrderSearchXHotelOrder {
	s.SellerNick = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetRoomNumber(v int64) *TaobaoXhotelOrderSearchXHotelOrder {
	s.RoomNumber = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetPayment(v int64) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Payment = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetContactPhone(v string) *TaobaoXhotelOrderSearchXHotelOrder {
	s.ContactPhone = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetArriveEarly(v util.LocalTime) *TaobaoXhotelOrderSearchXHotelOrder {
	s.ArriveEarly = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetPayTime(v util.LocalTime) *TaobaoXhotelOrderSearchXHotelOrder {
	s.PayTime = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetTradeStatus(v string) *TaobaoXhotelOrderSearchXHotelOrder {
	s.TradeStatus = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetModified(v util.LocalTime) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Modified = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetBuyerNick(v string) *TaobaoXhotelOrderSearchXHotelOrder {
	s.BuyerNick = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetMessage(v string) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Message = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetRid(v int64) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Rid = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetEndTime(v util.LocalTime) *TaobaoXhotelOrderSearchXHotelOrder {
	s.EndTime = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetTotalRoomPrice(v int64) *TaobaoXhotelOrderSearchXHotelOrder {
	s.TotalRoomPrice = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetGid(v int64) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Gid = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetPrices(v []int64) *TaobaoXhotelOrderSearchXHotelOrder {
	s.Prices = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetAlipayTradeNo(v string) *TaobaoXhotelOrderSearchXHotelOrder {
	s.AlipayTradeNo = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetLogisticsStatus(v string) *TaobaoXhotelOrderSearchXHotelOrder {
	s.LogisticsStatus = &v
	return s
}
func (s *TaobaoXhotelOrderSearchXHotelOrder) SetRefundStatus(v string) *TaobaoXhotelOrderSearchXHotelOrder {
	s.RefundStatus = &v
	return s
}
