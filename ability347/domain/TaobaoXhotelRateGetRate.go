package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelRateGetRate struct {
	/*
	   酒店商品id     */
	Gid *int64 `json:"gid,omitempty" `

	/*
	   酒店RPID     */
	Rpid *int64 `json:"rpid,omitempty" `

	/*
	   名称     */
	Name *string `json:"name,omitempty" `

	/*
		        价格和库存信息。
		A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。
		B:date  日期必须为 T---T+90 日内的日期（T为当天），且不能重复
		C:price 价格 int类型 取值范围1-99999999 单位为分
		D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)     */
	InventoryPrice *string `json:"inventory_price,omitempty" `

	/*
	   额外服务-是否可以加床，1：不可以，2：可以     */
	AddBed *int64 `json:"add_bed,omitempty" `

	/*
	   额外服务-加床价格     */
	AddBedPrice *int64 `json:"add_bed_price,omitempty" `

	/*
	   币种（仅支持CNY）     */
	CurrencyCode *int64 `json:"currency_code,omitempty" `

	/*
	   实价有房标签（RP支付类型为全额支付）     */
	ShijiaTag *int64 `json:"shijia_tag,omitempty" `

	/*
	   即时确认状态，表示此rate预订后是否可以直接发货。可取范围：0,1。可以为空     */
	JishiquerenTag *int64 `json:"jishiqueren_tag,omitempty" `

	/*
	   创建时间     */
	CreatedTime *util.LocalTime `json:"created_time,omitempty" `

	/*
	   修改时间     */
	ModifiedTime *util.LocalTime `json:"modified_time,omitempty" `

	/*
	   是否使用RoomInventory库存   仅当Rate上使用时有意义     */
	UseRoomInventory *bool `json:"use_room_inventory,omitempty" `

	/*
	   结构化的库存和开关,   	 date 日期 	 price 价格 int 类型, 取值范围1-99999999 单位为分  	 quota 普通库存 int 类型 取值范围 0-999（数量库存） 60000(状态库存关) 61000(状态库存开)  	 alQuota 协议保留房库存  int 类型 取值范围 0-999（数量库存） 60000(状态库存关) 61000(状态库存开) 	 genAlQuota 普通保留房库存, int 类型 取值范围 0-999（数量库存） 60000(状态库存关) 61000(状态库存开) 	 rateSwitch  date日期的价格开关, 值为true时,表示当天价格开, false表示价格关     */
	InvPriceWithSwitch *string `json:"inv_price_with_switch,omitempty" `

	/*
	   rate 维度下特殊标签含义 json: {"ebk-tail-room-Rate":1}, key:ebk-tail-room-Rate 表示rate维度ebk尾房标     */
	TagJson *string `json:"tag_json,omitempty" `
}

func (s *TaobaoXhotelRateGetRate) SetGid(v int64) *TaobaoXhotelRateGetRate {
	s.Gid = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetRpid(v int64) *TaobaoXhotelRateGetRate {
	s.Rpid = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetName(v string) *TaobaoXhotelRateGetRate {
	s.Name = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetInventoryPrice(v string) *TaobaoXhotelRateGetRate {
	s.InventoryPrice = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetAddBed(v int64) *TaobaoXhotelRateGetRate {
	s.AddBed = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetAddBedPrice(v int64) *TaobaoXhotelRateGetRate {
	s.AddBedPrice = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetCurrencyCode(v int64) *TaobaoXhotelRateGetRate {
	s.CurrencyCode = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetShijiaTag(v int64) *TaobaoXhotelRateGetRate {
	s.ShijiaTag = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetJishiquerenTag(v int64) *TaobaoXhotelRateGetRate {
	s.JishiquerenTag = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetCreatedTime(v util.LocalTime) *TaobaoXhotelRateGetRate {
	s.CreatedTime = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetModifiedTime(v util.LocalTime) *TaobaoXhotelRateGetRate {
	s.ModifiedTime = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetUseRoomInventory(v bool) *TaobaoXhotelRateGetRate {
	s.UseRoomInventory = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetInvPriceWithSwitch(v string) *TaobaoXhotelRateGetRate {
	s.InvPriceWithSwitch = &v
	return s
}
func (s *TaobaoXhotelRateGetRate) SetTagJson(v string) *TaobaoXhotelRateGetRate {
	s.TagJson = &v
	return s
}
