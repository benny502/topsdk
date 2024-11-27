package domain


type TaobaoXhotelBnbhouseAddBnbLocationDto struct {
    /*
        domestic为0时，固定China； domestic为1时，必须传定义的海外国家编码值。参见：http://hotel.alitrip.com/area.htm     */
    Country  *string `json:"country,omitempty" `

    /*
        门店地址     */
    Address  *string `json:"address,omitempty" `

    /*
        商圈     */
    Business  *string `json:"business,omitempty" `

    /*
        城市编码。参见：http://hotel.alitrip.com/area.htm，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（更新时为可选）     */
    City  *int64 `json:"city,omitempty" `

    /*
        纬度     */
    Latitude  *string `json:"latitude,omitempty" `

    /*
        坐标类型，现在支持：G – Google; B – 百度; A – 高德; M – Mapbar; L – 灵图     */
    PositionType  *string `json:"position_type,omitempty" `

    /*
        国别 0:国内;1:国外。默认是国内     */
    Domestic  *int64 `json:"domestic,omitempty" `

    /*
        省份编码http://hotel.alitrip.com/area.htm     */
    Province  *int64 `json:"province,omitempty" `

    /*
        门店英文地址     */
    EnAddress  *string `json:"en_address,omitempty" `

    /*
        接待地址     */
    ReceptionAddress  *string `json:"reception_address,omitempty" `

    /*
        区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm     */
    District  *int64 `json:"district,omitempty" `

    /*
        时区0到+11或者0到-11     */
    Timezone  *string `json:"timezone,omitempty" `

    /*
        门牌号     */
    Doorplate  *string `json:"doorplate,omitempty" `

    /*
        经度     */
    Longitude  *string `json:"longitude,omitempty" `

    /*
        城市名称，优先取city字段，city字段如果为空会校验cityName     */
    CityName  *string `json:"city_name,omitempty" `

}

func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetCountry(v string) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetAddress(v string) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetBusiness(v string) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.Business = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetCity(v int64) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.City = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetLatitude(v string) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetPositionType(v string) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.PositionType = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetDomestic(v int64) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.Domestic = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetProvince(v int64) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.Province = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetEnAddress(v string) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.EnAddress = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetReceptionAddress(v string) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.ReceptionAddress = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetDistrict(v int64) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.District = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetTimezone(v string) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.Timezone = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetDoorplate(v string) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.Doorplate = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetLongitude(v string) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.Longitude = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbLocationDto) SetCityName(v string) *TaobaoXhotelBnbhouseAddBnbLocationDto {
    s.CityName = &v
    return s
}
