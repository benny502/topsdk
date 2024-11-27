package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelAddSHotel struct {
    /*
        酒店ID     */
    Shid  *int64 `json:"shid,omitempty" `

    /*
        name     */
    Name  *string `json:"name,omitempty" `

    /*
        used_name     */
    UsedName  *string `json:"used_name,omitempty" `

    /*
        酒店类型     */
    Type  *string `json:"type,omitempty" `

    /*
        0:国内;1:国外     */
    Domestic  *int64 `json:"domestic,omitempty" `

    /*
        国家编码     */
    Country  *string `json:"country,omitempty" `

    /*
        地区标签     */
    CityTag  *string `json:"city_tag,omitempty" `

    /*
        省份编码     */
    Province  *int64 `json:"province,omitempty" `

    /*
        城市编码     */
    City  *int64 `json:"city,omitempty" `

    /*
        区域编码     */
    District  *int64 `json:"district,omitempty" `

    /*
        business     */
    Business  *string `json:"business,omitempty" `

    /*
        酒店地址     */
    Address  *string `json:"address,omitempty" `

    /*
        酒店级别     */
    Level  *string `json:"level,omitempty" `

    /*
        longitude     */
    Longitude  *string `json:"longitude,omitempty" `

    /*
        latitude     */
    Latitude  *string `json:"latitude,omitempty" `

    /*
        position_type     */
    PositionType  *int64 `json:"position_type,omitempty" `

    /*
        酒店电话     */
    Tel  *string `json:"tel,omitempty" `

    /*
        传真     */
    Fax  *string `json:"fax,omitempty" `

    /*
        开业年份     */
    OpeningTime  *string `json:"opening_time,omitempty" `

    /*
        装修年份     */
    DecorateTime  *string `json:"decorate_time,omitempty" `

    /*
        楼层数     */
    Storeys  *string `json:"storeys,omitempty" `

    /*
        扩展信息的JSON     */
    Extend  *string `json:"extend,omitempty" `

    /*
        房间数     */
    Rooms  *int64 `json:"rooms,omitempty" `

    /*
        酒店介绍     */
    Desc  *string `json:"desc,omitempty" `

    /*
        交通距离与设施服务。JSON格式。     */
    Service  *string `json:"service,omitempty" `

    /*
        酒店设施     */
    HotelFacilities  *string `json:"hotel_facilities,omitempty" `

    /*
        房间设施     */
    RoomFacilities  *string `json:"room_facilities,omitempty" `

    /*
        酒店图片url     */
    PicUrl  *string `json:"pic_url,omitempty" `

    /*
        创建时间     */
    CreatedTime  *util.LocalTime `json:"created_time,omitempty" `

    /*
        修改时间     */
    ModifiedTime  *util.LocalTime `json:"modified_time,omitempty" `

    /*
        0,营业中；-1，筹建中；-2，暂停营业；-3，已停业；默认为0     */
    Status  *int64 `json:"status,omitempty" `

    /*
        邮编     */
    PostalCode  *string `json:"postal_code,omitempty" `

    /*
        brand     */
    Brand  *string `json:"brand,omitempty" `

}

func (s *TaobaoXhotelAddSHotel) SetShid(v int64) *TaobaoXhotelAddSHotel {
    s.Shid = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetName(v string) *TaobaoXhotelAddSHotel {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetUsedName(v string) *TaobaoXhotelAddSHotel {
    s.UsedName = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetType(v string) *TaobaoXhotelAddSHotel {
    s.Type = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetDomestic(v int64) *TaobaoXhotelAddSHotel {
    s.Domestic = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetCountry(v string) *TaobaoXhotelAddSHotel {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetCityTag(v string) *TaobaoXhotelAddSHotel {
    s.CityTag = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetProvince(v int64) *TaobaoXhotelAddSHotel {
    s.Province = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetCity(v int64) *TaobaoXhotelAddSHotel {
    s.City = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetDistrict(v int64) *TaobaoXhotelAddSHotel {
    s.District = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetBusiness(v string) *TaobaoXhotelAddSHotel {
    s.Business = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetAddress(v string) *TaobaoXhotelAddSHotel {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetLevel(v string) *TaobaoXhotelAddSHotel {
    s.Level = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetLongitude(v string) *TaobaoXhotelAddSHotel {
    s.Longitude = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetLatitude(v string) *TaobaoXhotelAddSHotel {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetPositionType(v int64) *TaobaoXhotelAddSHotel {
    s.PositionType = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetTel(v string) *TaobaoXhotelAddSHotel {
    s.Tel = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetFax(v string) *TaobaoXhotelAddSHotel {
    s.Fax = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetOpeningTime(v string) *TaobaoXhotelAddSHotel {
    s.OpeningTime = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetDecorateTime(v string) *TaobaoXhotelAddSHotel {
    s.DecorateTime = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetStoreys(v string) *TaobaoXhotelAddSHotel {
    s.Storeys = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetExtend(v string) *TaobaoXhotelAddSHotel {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetRooms(v int64) *TaobaoXhotelAddSHotel {
    s.Rooms = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetDesc(v string) *TaobaoXhotelAddSHotel {
    s.Desc = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetService(v string) *TaobaoXhotelAddSHotel {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetHotelFacilities(v string) *TaobaoXhotelAddSHotel {
    s.HotelFacilities = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetRoomFacilities(v string) *TaobaoXhotelAddSHotel {
    s.RoomFacilities = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetPicUrl(v string) *TaobaoXhotelAddSHotel {
    s.PicUrl = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetCreatedTime(v util.LocalTime) *TaobaoXhotelAddSHotel {
    s.CreatedTime = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetModifiedTime(v util.LocalTime) *TaobaoXhotelAddSHotel {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetStatus(v int64) *TaobaoXhotelAddSHotel {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetPostalCode(v string) *TaobaoXhotelAddSHotel {
    s.PostalCode = &v
    return s
}
func (s *TaobaoXhotelAddSHotel) SetBrand(v string) *TaobaoXhotelAddSHotel {
    s.Brand = &v
    return s
}
