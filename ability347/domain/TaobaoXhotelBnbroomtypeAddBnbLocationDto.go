package domain


type TaobaoXhotelBnbroomtypeAddBnbLocationDto struct {
    /*
        domestic为0时，固定China； domestic为1时，必须传定义的海外国家编码值。参见：http://hotel.alitrip.com/area.htm     */
    Country  *string `json:"country,omitempty" `

    /*
        地址，最多200字符     */
    Address  *string `json:"address,omitempty" `

    /*
        商圈，最多64字符     */
    Business  *string `json:"business,omitempty" `

    /*
        城市编码。参见：http://hotel.alitrip.com/area.htm，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（更新酒店时为可选）     */
    City  *int64 `json:"city,omitempty" `

    /*
        纬度     */
    Latitude  *string `json:"latitude,omitempty" `

    /*
        坐标类型，现在支持：G – Google; B – 百度; A – 高德; M – Mapbar; L – 灵图 开芯     */
    PositionType  *string `json:"position_type,omitempty" `

    /*
        是否国内酒店。0:国内;1:国外。默认是国内     */
    Domestic  *int64 `json:"domestic,omitempty" `

    /*
        省份编码http://hotel.alitrip.com/area.htm     */
    Province  *int64 `json:"province,omitempty" `

    /*
        英文地址，最多500字符     */
    EnAddress  *string `json:"en_address,omitempty" `

    /*
        接待地址，最多200字符     */
    ReceptionAddress  *string `json:"reception_address,omitempty" `

    /*
        区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm     */
    District  *int64 `json:"district,omitempty" `

    /*
        时区0到+11或者0到-11     */
    Timezone  *string `json:"timezone,omitempty" `

    /*
        门牌号，最多200字符     */
    Doorplate  *string `json:"doorplate,omitempty" `

    /*
        经度     */
    Longitude  *string `json:"longitude,omitempty" `

}

func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetCountry(v string) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetAddress(v string) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetBusiness(v string) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.Business = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetCity(v int64) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.City = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetLatitude(v string) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetPositionType(v string) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.PositionType = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetDomestic(v int64) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.Domestic = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetProvince(v int64) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.Province = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetEnAddress(v string) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.EnAddress = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetReceptionAddress(v string) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.ReceptionAddress = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetDistrict(v int64) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.District = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetTimezone(v string) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.Timezone = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetDoorplate(v string) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.Doorplate = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbLocationDto) SetLongitude(v string) *TaobaoXhotelBnbroomtypeAddBnbLocationDto {
    s.Longitude = &v
    return s
}
