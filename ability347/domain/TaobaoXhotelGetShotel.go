package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelGetShotel struct {
	/*
	   酒店ID     */
	Shid *int64 `json:"shid,omitempty" `

	/*
	   name     */
	Name *string `json:"name,omitempty" `

	/*
	   used_name     */
	UsedName *string `json:"used_name,omitempty" `

	/*
	   酒店类型     */
	Type *string `json:"type,omitempty" `

	/*
	   0:国内;1:国外     */
	Domestic *int64 `json:"domestic,omitempty" `

	/*
	   国家编码     */
	Country *string `json:"country,omitempty" `

	/*
	   地区标签     */
	CityTag *string `json:"city_tag,omitempty" `

	/*
	   省份编码     */
	Province *int64 `json:"province,omitempty" `

	/*
	   城市编码     */
	City *int64 `json:"city,omitempty" `

	/*
	   区域编码     */
	District *int64 `json:"district,omitempty" `

	/*
	   business     */
	Business *string `json:"business,omitempty" `

	/*
	   酒店地址     */
	Address *string `json:"address,omitempty" `

	/*
	   酒店级别     */
	Level *string `json:"level,omitempty" `

	/*
	   longitude     */
	Longitude *string `json:"longitude,omitempty" `

	/*
	   latitude     */
	Latitude *string `json:"latitude,omitempty" `

	/*
	   position_type     */
	PositionType *int64 `json:"position_type,omitempty" `

	/*
	   酒店电话     */
	Tel *string `json:"tel,omitempty" `

	/*
	   传真     */
	Fax *string `json:"fax,omitempty" `

	/*
	   开业年份     */
	OpeningTime *string `json:"opening_time,omitempty" `

	/*
	   装修年份     */
	DecorateTime *string `json:"decorate_time,omitempty" `

	/*
	   楼层数     */
	Storeys *string `json:"storeys,omitempty" `

	/*
	   扩展信息的JSON     */
	Extend *string `json:"extend,omitempty" `

	/*
	   房间数     */
	Rooms *int64 `json:"rooms,omitempty" `

	/*
	   酒店介绍     */
	Desc *string `json:"desc,omitempty" `

	/*
	   交通距离与设施服务。JSON格式。     */
	Service *string `json:"service,omitempty" `

	/*
	   酒店设施     */
	HotelFacilities *string `json:"hotel_facilities,omitempty" `

	/*
	   房间设施     */
	RoomFacilities *string `json:"room_facilities,omitempty" `

	/*
	   酒店图片url     */
	PicUrl *string `json:"pic_url,omitempty" `

	/*
	   创建时间     */
	CreatedTime *util.LocalTime `json:"created_time,omitempty" `

	/*
	   修改时间     */
	ModifiedTime *util.LocalTime `json:"modified_time,omitempty" `

	/*
	   0：营业中；-1：筹建中；- 4：失效     */
	Status *int64 `json:"status,omitempty" `

	/*
	   邮编     */
	PostalCode *string `json:"postal_code,omitempty" `

	/*
	   brand     */
	Brand *string `json:"brand,omitempty" `

	/*
	   杂费     */
	Fee *string `json:"fee,omitempty" `

	/*
	   标准酒店英文名     */
	NameE *string `json:"name_e,omitempty" `

	/*
	   0：酒店 1：客栈     */
	IsKezhan *int64 `json:"is_kezhan,omitempty" `
}

func (s *TaobaoXhotelGetShotel) SetShid(v int64) *TaobaoXhotelGetShotel {
	s.Shid = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetName(v string) *TaobaoXhotelGetShotel {
	s.Name = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetUsedName(v string) *TaobaoXhotelGetShotel {
	s.UsedName = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetType(v string) *TaobaoXhotelGetShotel {
	s.Type = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetDomestic(v int64) *TaobaoXhotelGetShotel {
	s.Domestic = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetCountry(v string) *TaobaoXhotelGetShotel {
	s.Country = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetCityTag(v string) *TaobaoXhotelGetShotel {
	s.CityTag = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetProvince(v int64) *TaobaoXhotelGetShotel {
	s.Province = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetCity(v int64) *TaobaoXhotelGetShotel {
	s.City = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetDistrict(v int64) *TaobaoXhotelGetShotel {
	s.District = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetBusiness(v string) *TaobaoXhotelGetShotel {
	s.Business = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetAddress(v string) *TaobaoXhotelGetShotel {
	s.Address = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetLevel(v string) *TaobaoXhotelGetShotel {
	s.Level = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetLongitude(v string) *TaobaoXhotelGetShotel {
	s.Longitude = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetLatitude(v string) *TaobaoXhotelGetShotel {
	s.Latitude = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetPositionType(v int64) *TaobaoXhotelGetShotel {
	s.PositionType = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetTel(v string) *TaobaoXhotelGetShotel {
	s.Tel = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetFax(v string) *TaobaoXhotelGetShotel {
	s.Fax = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetOpeningTime(v string) *TaobaoXhotelGetShotel {
	s.OpeningTime = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetDecorateTime(v string) *TaobaoXhotelGetShotel {
	s.DecorateTime = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetStoreys(v string) *TaobaoXhotelGetShotel {
	s.Storeys = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetExtend(v string) *TaobaoXhotelGetShotel {
	s.Extend = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetRooms(v int64) *TaobaoXhotelGetShotel {
	s.Rooms = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetDesc(v string) *TaobaoXhotelGetShotel {
	s.Desc = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetService(v string) *TaobaoXhotelGetShotel {
	s.Service = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetHotelFacilities(v string) *TaobaoXhotelGetShotel {
	s.HotelFacilities = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetRoomFacilities(v string) *TaobaoXhotelGetShotel {
	s.RoomFacilities = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetPicUrl(v string) *TaobaoXhotelGetShotel {
	s.PicUrl = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetCreatedTime(v util.LocalTime) *TaobaoXhotelGetShotel {
	s.CreatedTime = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetModifiedTime(v util.LocalTime) *TaobaoXhotelGetShotel {
	s.ModifiedTime = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetStatus(v int64) *TaobaoXhotelGetShotel {
	s.Status = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetPostalCode(v string) *TaobaoXhotelGetShotel {
	s.PostalCode = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetBrand(v string) *TaobaoXhotelGetShotel {
	s.Brand = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetFee(v string) *TaobaoXhotelGetShotel {
	s.Fee = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetNameE(v string) *TaobaoXhotelGetShotel {
	s.NameE = &v
	return s
}
func (s *TaobaoXhotelGetShotel) SetIsKezhan(v int64) *TaobaoXhotelGetShotel {
	s.IsKezhan = &v
	return s
}
