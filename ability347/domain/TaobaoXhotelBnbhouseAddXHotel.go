package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelBnbhouseAddXHotel struct {
    /*
        国家编码     */
    Country  *string `json:"country,omitempty" `

    /*
        0:国内;1:国外     */
    Domestic  *int64 `json:"domestic,omitempty" `

    /*
        对接系统商名称     */
    Vendor  *string `json:"vendor,omitempty" `

    /*
        门店名称     */
    Name  *string `json:"name,omitempty" `

    /*
        外部id     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        修改时间     */
    ModifiedTime  *util.LocalTime `json:"modified_time,omitempty" `

    /*
        创建时间     */
    CreatedTime  *util.LocalTime `json:"created_time,omitempty" `

    /*
        未通过时的拒绝原因等。     */
    ErrorInfo  *string `json:"error_info,omitempty" `

    /*
        淘宝标准门店信息     */
    Shotel  *TaobaoXhotelBnbhouseAddXsHotel `json:"shotel,omitempty" `

    /*
        卖家门店ID     */
    Hid  *int64 `json:"hid,omitempty" `

    /*
        英文名称     */
    NameE  *string `json:"name_e,omitempty" `

    /*
        门店电话     */
    Tel  *string `json:"tel,omitempty" `

    /*
        position_type     */
    PositionType  *string `json:"position_type,omitempty" `

    /*
        纬度     */
    Latitude  *string `json:"latitude,omitempty" `

    /*
        精度     */
    Longitude  *string `json:"longitude,omitempty" `

    /*
        地址     */
    Address  *string `json:"address,omitempty" `

    /*
        商圈     */
    Business  *string `json:"business,omitempty" `

    /*
        地区     */
    District  *int64 `json:"district,omitempty" `

    /*
        城市     */
    City  *int64 `json:"city,omitempty" `

    /*
        省     */
    Province  *int64 `json:"province,omitempty" `

    /*
        品牌     */
    Brand  *string `json:"brand,omitempty" `

    /*
        照片格式，json     */
    Pics  *string `json:"pics,omitempty" `

    /*
        设施     */
    HotelFacilities  *string `json:"hotel_facilities,omitempty" `

    /*
        入住政策     */
    HotelPolicies  *string `json:"hotel_policies,omitempty" `

    /*
        描述     */
    Description  *string `json:"description,omitempty" `

    /*
        房间数     */
    Rooms  *int64 `json:"rooms,omitempty" `

    /*
        楼层     */
    Floors  *string `json:"floors,omitempty" `

    /*
        装修时间     */
    DecorateTime  *string `json:"decorate_time,omitempty" `

    /*
        开业时间     */
    OpeningTime  *string `json:"opening_time,omitempty" `

    /*
        预订须知     */
    BookingNotice  *string `json:"booking_notice,omitempty" `

}

func (s *TaobaoXhotelBnbhouseAddXHotel) SetCountry(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetDomestic(v int64) *TaobaoXhotelBnbhouseAddXHotel {
    s.Domestic = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetVendor(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetName(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetOuterId(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetModifiedTime(v util.LocalTime) *TaobaoXhotelBnbhouseAddXHotel {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetCreatedTime(v util.LocalTime) *TaobaoXhotelBnbhouseAddXHotel {
    s.CreatedTime = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetErrorInfo(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.ErrorInfo = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetShotel(v TaobaoXhotelBnbhouseAddXsHotel) *TaobaoXhotelBnbhouseAddXHotel {
    s.Shotel = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetHid(v int64) *TaobaoXhotelBnbhouseAddXHotel {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetNameE(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.NameE = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetTel(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.Tel = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetPositionType(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.PositionType = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetLatitude(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetLongitude(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.Longitude = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetAddress(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetBusiness(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.Business = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetDistrict(v int64) *TaobaoXhotelBnbhouseAddXHotel {
    s.District = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetCity(v int64) *TaobaoXhotelBnbhouseAddXHotel {
    s.City = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetProvince(v int64) *TaobaoXhotelBnbhouseAddXHotel {
    s.Province = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetBrand(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.Brand = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetPics(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.Pics = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetHotelFacilities(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.HotelFacilities = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetHotelPolicies(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.HotelPolicies = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetDescription(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.Description = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetRooms(v int64) *TaobaoXhotelBnbhouseAddXHotel {
    s.Rooms = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetFloors(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.Floors = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetDecorateTime(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.DecorateTime = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetOpeningTime(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.OpeningTime = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXHotel) SetBookingNotice(v string) *TaobaoXhotelBnbhouseAddXHotel {
    s.BookingNotice = &v
    return s
}
