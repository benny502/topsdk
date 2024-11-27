package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelHouseAddSHotel struct {
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
        状态:0：待系统匹配1：已系统匹配，匹配成功，待卖家确认2：已系统匹配，匹配失败，待人工匹配3：已人工匹配，匹配成功，待卖家确认4：已人工匹配，匹配失败5：卖家已确认，确认“YES”6：卖家已确认，确认“NO”7：停售     */
    Status  *int64 `json:"status,omitempty" `

    /*
        邮编     */
    PostalCode  *string `json:"postal_code,omitempty" `

    /*
        brand     */
    Brand  *string `json:"brand,omitempty" `

}

func (s *TaobaoXhotelHouseAddSHotel) SetShid(v int64) *TaobaoXhotelHouseAddSHotel {
    s.Shid = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetName(v string) *TaobaoXhotelHouseAddSHotel {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetUsedName(v string) *TaobaoXhotelHouseAddSHotel {
    s.UsedName = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetType(v string) *TaobaoXhotelHouseAddSHotel {
    s.Type = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetDomestic(v int64) *TaobaoXhotelHouseAddSHotel {
    s.Domestic = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetCountry(v string) *TaobaoXhotelHouseAddSHotel {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetCityTag(v string) *TaobaoXhotelHouseAddSHotel {
    s.CityTag = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetProvince(v int64) *TaobaoXhotelHouseAddSHotel {
    s.Province = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetCity(v int64) *TaobaoXhotelHouseAddSHotel {
    s.City = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetDistrict(v int64) *TaobaoXhotelHouseAddSHotel {
    s.District = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetBusiness(v string) *TaobaoXhotelHouseAddSHotel {
    s.Business = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetAddress(v string) *TaobaoXhotelHouseAddSHotel {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetLevel(v string) *TaobaoXhotelHouseAddSHotel {
    s.Level = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetLongitude(v string) *TaobaoXhotelHouseAddSHotel {
    s.Longitude = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetLatitude(v string) *TaobaoXhotelHouseAddSHotel {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetPositionType(v int64) *TaobaoXhotelHouseAddSHotel {
    s.PositionType = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetTel(v string) *TaobaoXhotelHouseAddSHotel {
    s.Tel = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetFax(v string) *TaobaoXhotelHouseAddSHotel {
    s.Fax = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetOpeningTime(v string) *TaobaoXhotelHouseAddSHotel {
    s.OpeningTime = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetDecorateTime(v string) *TaobaoXhotelHouseAddSHotel {
    s.DecorateTime = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetStoreys(v string) *TaobaoXhotelHouseAddSHotel {
    s.Storeys = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetExtend(v string) *TaobaoXhotelHouseAddSHotel {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetRooms(v int64) *TaobaoXhotelHouseAddSHotel {
    s.Rooms = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetDesc(v string) *TaobaoXhotelHouseAddSHotel {
    s.Desc = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetService(v string) *TaobaoXhotelHouseAddSHotel {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetHotelFacilities(v string) *TaobaoXhotelHouseAddSHotel {
    s.HotelFacilities = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetRoomFacilities(v string) *TaobaoXhotelHouseAddSHotel {
    s.RoomFacilities = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetPicUrl(v string) *TaobaoXhotelHouseAddSHotel {
    s.PicUrl = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetCreatedTime(v util.LocalTime) *TaobaoXhotelHouseAddSHotel {
    s.CreatedTime = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetModifiedTime(v util.LocalTime) *TaobaoXhotelHouseAddSHotel {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetStatus(v int64) *TaobaoXhotelHouseAddSHotel {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetPostalCode(v string) *TaobaoXhotelHouseAddSHotel {
    s.PostalCode = &v
    return s
}
func (s *TaobaoXhotelHouseAddSHotel) SetBrand(v string) *TaobaoXhotelHouseAddSHotel {
    s.Brand = &v
    return s
}
