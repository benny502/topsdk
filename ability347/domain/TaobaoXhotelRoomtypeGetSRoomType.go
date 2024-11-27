package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelRoomtypeGetSRoomType struct {
	/*
	   匹配的标准房型     */
	Srid *int64 `json:"srid,omitempty" `

	/*
	   房型名     */
	Name *string `json:"name,omitempty" `

	/*
	   楼层     */
	Floor *string `json:"floor,omitempty" `

	/*
		        宽带服务
		"0","有线上网(免费),
		"1","有线上网(无)",
		"2","有线上网(收费)",
		"3","有线上网(部分有且免费)",
		"4","有线上网(部分有且收费)"     */
	Internet *string `json:"internet,omitempty" `

	/*
	   shid     */
	Shid *int64 `json:"shid,omitempty" `

	/*
	   pic_url     */
	PicUrl *string `json:"pic_url,omitempty" `

	/*
	   facility     */
	Facility *string `json:"facility,omitempty" `

	/*
	   最大入住人数     */
	MaxOccupancy *int64 `json:"max_occupancy,omitempty" `

	/*
	   面积     */
	Area *string `json:"area,omitempty" `

	/*
	   扩展字段     */
	Extend *string `json:"extend,omitempty" `

	/*
	   创建时间     */
	CreatedTime *util.LocalTime `json:"created_time,omitempty" `

	/*
	   修改时间     */
	ModifiedTime *util.LocalTime `json:"modified_time,omitempty" `

	/*
		        窗型，枚举类型
		0, 无窗,
		1, 有窗;     */
	WindowType *string `json:"window_type,omitempty" `

	/*
	   床型。json格式：[{"bedType":"大床","bedSize":"1.5m"},{"bedType":"双床","bedSize":"1.2m"}]     */
	Bed *string `json:"bed,omitempty" `

	/*
	   状态。0:正常;-1:删除     */
	Status *int64 `json:"status,omitempty" `
}

func (s *TaobaoXhotelRoomtypeGetSRoomType) SetSrid(v int64) *TaobaoXhotelRoomtypeGetSRoomType {
	s.Srid = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetName(v string) *TaobaoXhotelRoomtypeGetSRoomType {
	s.Name = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetFloor(v string) *TaobaoXhotelRoomtypeGetSRoomType {
	s.Floor = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetInternet(v string) *TaobaoXhotelRoomtypeGetSRoomType {
	s.Internet = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetShid(v int64) *TaobaoXhotelRoomtypeGetSRoomType {
	s.Shid = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetPicUrl(v string) *TaobaoXhotelRoomtypeGetSRoomType {
	s.PicUrl = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetFacility(v string) *TaobaoXhotelRoomtypeGetSRoomType {
	s.Facility = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetMaxOccupancy(v int64) *TaobaoXhotelRoomtypeGetSRoomType {
	s.MaxOccupancy = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetArea(v string) *TaobaoXhotelRoomtypeGetSRoomType {
	s.Area = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetExtend(v string) *TaobaoXhotelRoomtypeGetSRoomType {
	s.Extend = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetCreatedTime(v util.LocalTime) *TaobaoXhotelRoomtypeGetSRoomType {
	s.CreatedTime = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetModifiedTime(v util.LocalTime) *TaobaoXhotelRoomtypeGetSRoomType {
	s.ModifiedTime = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetWindowType(v string) *TaobaoXhotelRoomtypeGetSRoomType {
	s.WindowType = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetBed(v string) *TaobaoXhotelRoomtypeGetSRoomType {
	s.Bed = &v
	return s
}
func (s *TaobaoXhotelRoomtypeGetSRoomType) SetStatus(v int64) *TaobaoXhotelRoomtypeGetSRoomType {
	s.Status = &v
	return s
}
