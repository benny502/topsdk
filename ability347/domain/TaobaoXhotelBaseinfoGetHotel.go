package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelBaseinfoGetHotel struct {
	/*
	   酒店ID     */
	Hid *int64 `json:"hid,omitempty" `

	/*
	   酒店修改备注     */
	Remark *string `json:"remark,omitempty" `

	/*
	   酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886     */
	Tel *string `json:"tel,omitempty" `

	/*
	   酒店设施。json格式     */
	HotelFacilities *string `json:"hotel_facilities,omitempty" `

	/*
	   酒店类型     */
	Type *string `json:"type,omitempty" `

	/*
	   扩展信息     */
	Ext *string `json:"ext,omitempty" `

	/*
	   城市编码     */
	City *int64 `json:"city,omitempty" `

	/*
	   邮编     */
	PostalCode *string `json:"postal_code,omitempty" `

	/*
	   楼层信息     */
	Floors *string `json:"floors,omitempty" `

	/*
	   卖家名称     */
	SellerNick *string `json:"seller_nick,omitempty" `

	/*
	   酒店中文描述     */
	Description *string `json:"description,omitempty" `

	/*
	   省份编码     */
	Province *int64 `json:"province,omitempty" `

	/*
	   经度     */
	Longitude *string `json:"longitude,omitempty" `

	/*
	   匹配是否人工确认     */
	DataConfirm *int64 `json:"data_confirm,omitempty" `

	/*
	   房间数     */
	Rooms *int64 `json:"rooms,omitempty" `

	/*
	   货币类型（编码,字母编码）     */
	CurrencyCodeName *string `json:"currency_code_name,omitempty" `

	/*
	   酒店状态     */
	Status *int64 `json:"status,omitempty" `

	/*
	   酒店英文描述     */
	EnDesc *string `json:"en_desc,omitempty" `

	/*
	   domestic=0时，固定China； domestic=1时，是海外国家编码值     */
	Country *string `json:"country,omitempty" `

	/*
	   酒店入住政策     */
	HotelPolicies *string `json:"hotel_policies,omitempty" `

	/*
	   纬度     */
	Latitude *string `json:"latitude,omitempty" `

	/*
	   操作人信息     */
	OperXiaoerName *string `json:"oper_xiaoer_name,omitempty" `

	/*
	   酒店外部ID     */
	OuterId *string `json:"outer_id,omitempty" `

	/*
	   酒店修改时间     */
	GmtModified *util.LocalTime `json:"gmt_modified,omitempty" `

	/*
	   酒店下架类型     */
	DownShelfType *int64 `json:"down_shelf_type,omitempty" `

	/*
	   酒店英文地址     */
	EnAddr *string `json:"en_addr,omitempty" `

	/*
	   标准酒店ID     */
	Shid *int64 `json:"shid,omitempty" `

	/*
	   商圈     */
	Business *string `json:"business,omitempty" `

	/*
	   酒店曾用名     */
	UsedName *string `json:"used_name,omitempty" `

	/*
	   酒店图片信息     */
	Pics *string `json:"pics,omitempty" `

	/*
	   房间设施     */
	RoomFacilities *string `json:"room_facilities,omitempty" `

	/*
	   酒店名     */
	Name *string `json:"name,omitempty" `

	/*
	   坐标类型，现在支持：G : Google:B : 百度;A : 高德;M : Mapbar;L : 灵图     */
	PositionType *string `json:"position_type,omitempty" `

	/*
	   酒店支付结算类型     */
	BillingProcessType *int64 `json:"billing_process_type,omitempty" `

	/*
	   地区编码     */
	District *int64 `json:"district,omitempty" `

	/*
	   酒店名称(英文)     */
	NameE *string `json:"name_e,omitempty" `

	/*
	   酒店创建时间     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   酒店的销售渠道     */
	Vendor *string `json:"vendor,omitempty" `

	/*
	   扩展信息     */
	Extend *string `json:"extend,omitempty" `

	/*
	   酒店星级     */
	Star *string `json:"star,omitempty" `

	/*
	   预订须知     */
	BookingNotice *string `json:"booking_notice,omitempty" `

	/*
	   是否国外。     */
	Domestic *int64 `json:"domestic,omitempty" `

	/*
	   装修时间     */
	DecorateTime *string `json:"decorate_time,omitempty" `

	/*
	   来源     */
	Source *int64 `json:"source,omitempty" `

	/*
	   酒店地址     */
	Address *string `json:"address,omitempty" `

	/*
	   卖家ID     */
	SellerId *int64 `json:"seller_id,omitempty" `

	/*
	   匹配状态     */
	MatchStatus *int64 `json:"match_status,omitempty" `

	/*
	   酒店服务     */
	Service *string `json:"service,omitempty" `

	/*
	   品牌     */
	Brand *string `json:"brand,omitempty" `

	/*
	   判断该酒店是不是对应的卖家直营     */
	KzzyTag *int64 `json:"kzzy_tag,omitempty" `

	/*
	   开业时间     */
	OpeningTime *string `json:"opening_time,omitempty" `
}

func (s *TaobaoXhotelBaseinfoGetHotel) SetHid(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.Hid = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetRemark(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Remark = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetTel(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Tel = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetHotelFacilities(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.HotelFacilities = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetType(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Type = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetExt(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Ext = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetCity(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.City = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetPostalCode(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.PostalCode = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetFloors(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Floors = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetSellerNick(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.SellerNick = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetDescription(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Description = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetProvince(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.Province = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetLongitude(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Longitude = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetDataConfirm(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.DataConfirm = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetRooms(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.Rooms = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetCurrencyCodeName(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.CurrencyCodeName = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetStatus(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.Status = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetEnDesc(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.EnDesc = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetCountry(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Country = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetHotelPolicies(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.HotelPolicies = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetLatitude(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Latitude = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetOperXiaoerName(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.OperXiaoerName = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetOuterId(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.OuterId = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetGmtModified(v util.LocalTime) *TaobaoXhotelBaseinfoGetHotel {
	s.GmtModified = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetDownShelfType(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.DownShelfType = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetEnAddr(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.EnAddr = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetShid(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.Shid = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetBusiness(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Business = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetUsedName(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.UsedName = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetPics(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Pics = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetRoomFacilities(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.RoomFacilities = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetName(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Name = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetPositionType(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.PositionType = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetBillingProcessType(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.BillingProcessType = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetDistrict(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.District = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetNameE(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.NameE = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetGmtCreate(v util.LocalTime) *TaobaoXhotelBaseinfoGetHotel {
	s.GmtCreate = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetVendor(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Vendor = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetExtend(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Extend = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetStar(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Star = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetBookingNotice(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.BookingNotice = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetDomestic(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.Domestic = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetDecorateTime(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.DecorateTime = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetSource(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.Source = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetAddress(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Address = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetSellerId(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.SellerId = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetMatchStatus(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.MatchStatus = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetService(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Service = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetBrand(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.Brand = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetKzzyTag(v int64) *TaobaoXhotelBaseinfoGetHotel {
	s.KzzyTag = &v
	return s
}
func (s *TaobaoXhotelBaseinfoGetHotel) SetOpeningTime(v string) *TaobaoXhotelBaseinfoGetHotel {
	s.OpeningTime = &v
	return s
}
