package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelGetFirstResult struct {
	/*
	   酒店ID     */
	Hid *int64 `json:"hid,omitempty" `

	/*
	   酒店状态：0: 正常;-2:停售；-1：删除     */
	Status *int64 `json:"status,omitempty" `

	/*
	   淘宝标准酒店信息     */
	SHotel *TaobaoXhotelGetShotel `json:"s_hotel,omitempty" `

	/*
	   未通过时的拒绝原因等。     */
	ErrorInfo *string `json:"error_info,omitempty" `

	/*
	   卖家自己系统的id     */
	OuterId *string `json:"outer_id,omitempty" `

	/*
	   酒店名称     */
	Name *string `json:"name,omitempty" `

	/*
	   0:国内;1:国外     */
	Domestic *int64 `json:"domestic,omitempty" `

	/*
	   国家编码     */
	Country *string `json:"country,omitempty" `

	/*
	   曾用名     */
	UsedName *string `json:"used_name,omitempty" `

	/*
	   省份编码     */
	Province *int64 `json:"province,omitempty" `

	/*
	   城市编码     */
	City *int64 `json:"city,omitempty" `

	/*
	   地区编码     */
	District *int64 `json:"district,omitempty" `

	/*
	   商圈信息     */
	Business *string `json:"business,omitempty" `

	/*
	   酒店地址     */
	Address *string `json:"address,omitempty" `

	/*
	   经度     */
	Longitude *string `json:"longitude,omitempty" `

	/*
	   纬度     */
	Latitude *string `json:"latitude,omitempty" `

	/*
	   坐标类型     */
	PositionType *string `json:"position_type,omitempty" `

	/*
	   酒店电话     */
	Tel *string `json:"tel,omitempty" `

	/*
	   扩展信息     */
	Extend *string `json:"extend,omitempty" `

	/*
	   此字段已废弃     */
	MatchStatus *int64 `json:"match_status,omitempty" `

	/*
	   创建时间     */
	CreatedTime *util.LocalTime `json:"created_time,omitempty" `

	/*
	   修改时间     */
	ModifiedTime *util.LocalTime `json:"modified_time,omitempty" `

	/*
	   匹配结果     */
	DataConfirmStr *string `json:"data_confirm_str,omitempty" `

	/*
	   逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡     */
	CreditCardTypes *string `json:"credit_card_types,omitempty" `

	/*
	   卖家酒店英文名称     */
	NameE *string `json:"name_e,omitempty" `

	/*
	   vendor     */
	Vendor *string `json:"vendor,omitempty" `

	/*
	   酒店维度特殊标签含义, json: {"pure-direct-hotel":0,"direct-manual-order-hotel":1,"ebk-direct-hotel":0,"non-direct-hotel":1,"allow-dingding-business-travel-hotel":1,"willing-dingding-bussiness-travel-hotel":0,"calendar-room-package-hotel":1,"dijiajiajia-hotel":0,"ebk-number-of-confirm-room-hotel":1} , key含义: pure-direct-hotel 表示纯直连酒店, direct-manual-order-hotel 和 ebk-direct-hotel 和 non-direct-hotel 这三个key对应value都是0 . direct-manual-order-hotel 表示 人工承接失败订单的酒店标签。如果某个直连酒店打了该标签，那么直连下单失败以后，允许人工承接订单，由人工跟进 . ebk-direct-hotel 表示 ebk直连酒店标。如果某个酒店打了该标签，那么这个酒店下允许通过ebk发布直连rp . non-direct-hotel 表示 卖家非直连酒店标签。如果某个酒店打了该标签，那么该酒店下单不会走直连交易。 allow-dingding-business-travel-hotel 表示 允许进入阿里商旅渠道（钉钉）售卖信用住的单体酒店 willing-dingding-bussiness-travel-hotel 表示 已签协议愿意加入阿里商旅渠道售卖信用住的单体酒店 . calendar-room-package-hotel 表示 酒店可以参加日历房套餐活动包括创建，修改，删除活动信息（高星集团GMV项目） dijiajiajia-hotel 表示 底价加价酒店权限标。只有打了该标的酒店才允许维护底价加价规则和包房rp . ebk-number-of-confirm-room-hotel 表示ebk确认订单，是否要输入外部确认号 . nonstandard-project-hotel 表示该酒店是否参加非标项目     */
	TagJson *string `json:"tag_json,omitempty" `

	/*
	   标识该酒店所走的 结算流程，如果是null 默认是 国内结算流程，否则是其他的，比如：海外信用住结算流程     */
	BillingProcessType *int64 `json:"billing_process_type,omitempty" `

	/*
	   货币类型（编码,字母编码）,hid 维度支持的币种信息,目前只能 add 时添加，不支持 update时更新,,如果DB中是null ，则默认是人民币 CNY     */
	CurrencyCodeName *string `json:"currency_code_name,omitempty" `

	/*
	   酒店对应的旺旺号     */
	AliNick *string `json:"ali_nick,omitempty" `

	/*
	   资源方房型设施     */
	StandardRoomFacilities *string `json:"standard_room_facilities,omitempty" `

	/*
	   资源方酒店服务     */
	StandardHotelService *string `json:"standard_hotel_service,omitempty" `

	/*
	   资源方酒店设施     */
	StandardHotelFacilities *string `json:"standard_hotel_facilities,omitempty" `

	/*
	   资源方预订须知     */
	StandardBookingNotice *string `json:"standard_booking_notice,omitempty" `

	/*
	   资源方娱乐设施     */
	StandardAmuseFacilities *string `json:"standard_amuse_facilities,omitempty" `

	/*
	   离线数据,该酒店在售1,不在售0,未知-1     */
	OnSale *int64 `json:"on_sale,omitempty" `

	/*
	   离线数据,该酒店热搜1,非热搜0,未知-1     */
	HotSearch *int64 `json:"hot_search,omitempty" `

	/*
	   离线数据,该酒店热卖1,非热卖0,未知-1     */
	HotSale *int64 `json:"hot_sale,omitempty" `
}

func (s *TaobaoXhotelGetFirstResult) SetHid(v int64) *TaobaoXhotelGetFirstResult {
	s.Hid = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetStatus(v int64) *TaobaoXhotelGetFirstResult {
	s.Status = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetSHotel(v TaobaoXhotelGetShotel) *TaobaoXhotelGetFirstResult {
	s.SHotel = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetErrorInfo(v string) *TaobaoXhotelGetFirstResult {
	s.ErrorInfo = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetOuterId(v string) *TaobaoXhotelGetFirstResult {
	s.OuterId = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetName(v string) *TaobaoXhotelGetFirstResult {
	s.Name = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetDomestic(v int64) *TaobaoXhotelGetFirstResult {
	s.Domestic = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetCountry(v string) *TaobaoXhotelGetFirstResult {
	s.Country = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetUsedName(v string) *TaobaoXhotelGetFirstResult {
	s.UsedName = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetProvince(v int64) *TaobaoXhotelGetFirstResult {
	s.Province = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetCity(v int64) *TaobaoXhotelGetFirstResult {
	s.City = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetDistrict(v int64) *TaobaoXhotelGetFirstResult {
	s.District = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetBusiness(v string) *TaobaoXhotelGetFirstResult {
	s.Business = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetAddress(v string) *TaobaoXhotelGetFirstResult {
	s.Address = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetLongitude(v string) *TaobaoXhotelGetFirstResult {
	s.Longitude = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetLatitude(v string) *TaobaoXhotelGetFirstResult {
	s.Latitude = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetPositionType(v string) *TaobaoXhotelGetFirstResult {
	s.PositionType = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetTel(v string) *TaobaoXhotelGetFirstResult {
	s.Tel = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetExtend(v string) *TaobaoXhotelGetFirstResult {
	s.Extend = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetMatchStatus(v int64) *TaobaoXhotelGetFirstResult {
	s.MatchStatus = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetCreatedTime(v util.LocalTime) *TaobaoXhotelGetFirstResult {
	s.CreatedTime = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetModifiedTime(v util.LocalTime) *TaobaoXhotelGetFirstResult {
	s.ModifiedTime = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetDataConfirmStr(v string) *TaobaoXhotelGetFirstResult {
	s.DataConfirmStr = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetCreditCardTypes(v string) *TaobaoXhotelGetFirstResult {
	s.CreditCardTypes = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetNameE(v string) *TaobaoXhotelGetFirstResult {
	s.NameE = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetVendor(v string) *TaobaoXhotelGetFirstResult {
	s.Vendor = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetTagJson(v string) *TaobaoXhotelGetFirstResult {
	s.TagJson = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetBillingProcessType(v int64) *TaobaoXhotelGetFirstResult {
	s.BillingProcessType = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetCurrencyCodeName(v string) *TaobaoXhotelGetFirstResult {
	s.CurrencyCodeName = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetAliNick(v string) *TaobaoXhotelGetFirstResult {
	s.AliNick = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetStandardRoomFacilities(v string) *TaobaoXhotelGetFirstResult {
	s.StandardRoomFacilities = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetStandardHotelService(v string) *TaobaoXhotelGetFirstResult {
	s.StandardHotelService = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetStandardHotelFacilities(v string) *TaobaoXhotelGetFirstResult {
	s.StandardHotelFacilities = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetStandardBookingNotice(v string) *TaobaoXhotelGetFirstResult {
	s.StandardBookingNotice = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetStandardAmuseFacilities(v string) *TaobaoXhotelGetFirstResult {
	s.StandardAmuseFacilities = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetOnSale(v int64) *TaobaoXhotelGetFirstResult {
	s.OnSale = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetHotSearch(v int64) *TaobaoXhotelGetFirstResult {
	s.HotSearch = &v
	return s
}
func (s *TaobaoXhotelGetFirstResult) SetHotSale(v int64) *TaobaoXhotelGetFirstResult {
	s.HotSale = &v
	return s
}
