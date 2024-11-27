package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelUpdateSHotel struct {
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
	   0,营业中；-1，筹建中；-2，暂停营业；-3，已停业；默认为0     */
	Status *int64 `json:"status,omitempty" `

	/*
	   邮编     */
	PostalCode *string `json:"postal_code,omitempty" `

	/*
	   brand     */
	Brand *string `json:"brand,omitempty" `
}

func (s *TaobaoXhotelUpdateSHotel) SetShid(v int64) *TaobaoXhotelUpdateSHotel {
	s.Shid = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetName(v string) *TaobaoXhotelUpdateSHotel {
	s.Name = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetUsedName(v string) *TaobaoXhotelUpdateSHotel {
	s.UsedName = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetType(v string) *TaobaoXhotelUpdateSHotel {
	s.Type = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetDomestic(v int64) *TaobaoXhotelUpdateSHotel {
	s.Domestic = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetCountry(v string) *TaobaoXhotelUpdateSHotel {
	s.Country = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetCityTag(v string) *TaobaoXhotelUpdateSHotel {
	s.CityTag = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetProvince(v int64) *TaobaoXhotelUpdateSHotel {
	s.Province = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetCity(v int64) *TaobaoXhotelUpdateSHotel {
	s.City = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetDistrict(v int64) *TaobaoXhotelUpdateSHotel {
	s.District = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetBusiness(v string) *TaobaoXhotelUpdateSHotel {
	s.Business = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetAddress(v string) *TaobaoXhotelUpdateSHotel {
	s.Address = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetLevel(v string) *TaobaoXhotelUpdateSHotel {
	s.Level = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetLongitude(v string) *TaobaoXhotelUpdateSHotel {
	s.Longitude = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetLatitude(v string) *TaobaoXhotelUpdateSHotel {
	s.Latitude = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetPositionType(v int64) *TaobaoXhotelUpdateSHotel {
	s.PositionType = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetTel(v string) *TaobaoXhotelUpdateSHotel {
	s.Tel = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetFax(v string) *TaobaoXhotelUpdateSHotel {
	s.Fax = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetOpeningTime(v string) *TaobaoXhotelUpdateSHotel {
	s.OpeningTime = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetDecorateTime(v string) *TaobaoXhotelUpdateSHotel {
	s.DecorateTime = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetStoreys(v string) *TaobaoXhotelUpdateSHotel {
	s.Storeys = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetExtend(v string) *TaobaoXhotelUpdateSHotel {
	s.Extend = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetRooms(v int64) *TaobaoXhotelUpdateSHotel {
	s.Rooms = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetDesc(v string) *TaobaoXhotelUpdateSHotel {
	s.Desc = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetService(v string) *TaobaoXhotelUpdateSHotel {
	s.Service = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetHotelFacilities(v string) *TaobaoXhotelUpdateSHotel {
	s.HotelFacilities = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetRoomFacilities(v string) *TaobaoXhotelUpdateSHotel {
	s.RoomFacilities = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetPicUrl(v string) *TaobaoXhotelUpdateSHotel {
	s.PicUrl = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetCreatedTime(v util.LocalTime) *TaobaoXhotelUpdateSHotel {
	s.CreatedTime = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetModifiedTime(v util.LocalTime) *TaobaoXhotelUpdateSHotel {
	s.ModifiedTime = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetStatus(v int64) *TaobaoXhotelUpdateSHotel {
	s.Status = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetPostalCode(v string) *TaobaoXhotelUpdateSHotel {
	s.PostalCode = &v
	return s
}
func (s *TaobaoXhotelUpdateSHotel) SetBrand(v string) *TaobaoXhotelUpdateSHotel {
	s.Brand = &v
	return s
}
