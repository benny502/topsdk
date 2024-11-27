package domain


type TaobaoXhotelBnbhouseAddXsHotel struct {
    /*
        酒店地址     */
    Address  *string `json:"address,omitempty" `

    /*
        商圈     */
    Business  *string `json:"business,omitempty" `

    /*
        区域编码     */
    District  *int64 `json:"district,omitempty" `

    /*
        城市编码     */
    City  *int64 `json:"city,omitempty" `

    /*
        省份编码     */
    Province  *int64 `json:"province,omitempty" `

    /*
        国家编码     */
    Country  *string `json:"country,omitempty" `

    /*
        0:国内;1:国外     */
    Domestic  *int64 `json:"domestic,omitempty" `

    /*
        门店名称     */
    Name  *string `json:"name,omitempty" `

    /*
        系统自动生成     */
    Shid  *int64 `json:"shid,omitempty" `

    /*
        英文名称     */
    NameE  *string `json:"name_e,omitempty" `

    /*
        门店状态：0: 正常;-2:停售；-1：删除     */
    Status  *int64 `json:"status,omitempty" `

    /*
        品牌     */
    Brand  *string `json:"brand,omitempty" `

    /*
        照片,json格式     */
    Pics  *string `json:"pics,omitempty" `

    /*
        交通距离与设施服务。JSON格式。     */
    Service  *string `json:"service,omitempty" `

    /*
        酒店设施。json格式     */
    HotelFacilities  *string `json:"hotel_facilities,omitempty" `

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
        开业年份     */
    OpeningTime  *string `json:"opening_time,omitempty" `

    /*
        门店电话     */
    Tel  *string `json:"tel,omitempty" `

    /*
        纬度     */
    Latitude  *string `json:"latitude,omitempty" `

    /*
        position_type     */
    PositionType  *int64 `json:"position_type,omitempty" `

    /*
        精度     */
    Longitude  *string `json:"longitude,omitempty" `

}

func (s *TaobaoXhotelBnbhouseAddXsHotel) SetAddress(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetBusiness(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Business = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetDistrict(v int64) *TaobaoXhotelBnbhouseAddXsHotel {
    s.District = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetCity(v int64) *TaobaoXhotelBnbhouseAddXsHotel {
    s.City = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetProvince(v int64) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Province = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetCountry(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetDomestic(v int64) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Domestic = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetName(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetShid(v int64) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Shid = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetNameE(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.NameE = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetStatus(v int64) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetBrand(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Brand = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetPics(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Pics = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetService(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetHotelFacilities(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.HotelFacilities = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetDescription(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Description = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetRooms(v int64) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Rooms = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetFloors(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Floors = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetDecorateTime(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.DecorateTime = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetOpeningTime(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.OpeningTime = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetTel(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Tel = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetLatitude(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetPositionType(v int64) *TaobaoXhotelBnbhouseAddXsHotel {
    s.PositionType = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddXsHotel) SetLongitude(v string) *TaobaoXhotelBnbhouseAddXsHotel {
    s.Longitude = &v
    return s
}
