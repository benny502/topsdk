package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelUpdateXHotel struct {
    /*
        酒店ID     */
    Hid  *int64 `json:"hid,omitempty" `

    /*
        酒店状态：0: 正常;-2:停售；-1：删除     */
    Status  *int64 `json:"status,omitempty" `

    /*
        淘宝标准酒店信息     */
    SHotel  *TaobaoXhotelUpdateSHotel `json:"s_hotel,omitempty" `

    /*
        未通过时的拒绝原因等。     */
    ErrorInfo  *string `json:"error_info,omitempty" `

    /*
        此字段已废弃     */
    MatchStatus  *int64 `json:"match_status,omitempty" `

    /*
        创建时间     */
    CreatedTime  *util.LocalTime `json:"created_time,omitempty" `

    /*
        修改时间     */
    ModifiedTime  *util.LocalTime `json:"modified_time,omitempty" `

    /*
        卖家自己系统的id     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        酒店名称     */
    Name  *string `json:"name,omitempty" `

    /*
        0:国内;1:国外     */
    Domestic  *int64 `json:"domestic,omitempty" `

    /*
        国家编码     */
    Country  *string `json:"country,omitempty" `

    /*
        曾用名     */
    UsedName  *string `json:"used_name,omitempty" `

    /*
        省份编码     */
    Province  *int64 `json:"province,omitempty" `

    /*
        城市编码     */
    City  *int64 `json:"city,omitempty" `

    /*
        地区编码     */
    District  *int64 `json:"district,omitempty" `

    /*
        商圈信息     */
    Business  *string `json:"business,omitempty" `

    /*
        酒店地址     */
    Address  *string `json:"address,omitempty" `

    /*
        经度     */
    Longitude  *string `json:"longitude,omitempty" `

    /*
        纬度     */
    Latitude  *string `json:"latitude,omitempty" `

    /*
        坐标类型     */
    PositionType  *string `json:"position_type,omitempty" `

    /*
        酒店电话     */
    Tel  *string `json:"tel,omitempty" `

    /*
        扩展信息     */
    Extend  *string `json:"extend,omitempty" `

    /*
        逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡     */
    CreditCardTypes  *string `json:"credit_card_types,omitempty" `

    /*
        卖家酒店英文名称     */
    NameE  *string `json:"name_e,omitempty" `

    /*
        0:酒店；1:客栈     */
    HotelType  *int64 `json:"hotel_type,omitempty" `

    /*
        0:可以接待外宾；1:仅内宾     */
    ServiceType  *int64 `json:"service_type,omitempty" `

}

func (s *TaobaoXhotelUpdateXHotel) SetHid(v int64) *TaobaoXhotelUpdateXHotel {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetStatus(v int64) *TaobaoXhotelUpdateXHotel {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetSHotel(v TaobaoXhotelUpdateSHotel) *TaobaoXhotelUpdateXHotel {
    s.SHotel = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetErrorInfo(v string) *TaobaoXhotelUpdateXHotel {
    s.ErrorInfo = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetMatchStatus(v int64) *TaobaoXhotelUpdateXHotel {
    s.MatchStatus = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetCreatedTime(v util.LocalTime) *TaobaoXhotelUpdateXHotel {
    s.CreatedTime = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetModifiedTime(v util.LocalTime) *TaobaoXhotelUpdateXHotel {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetOuterId(v string) *TaobaoXhotelUpdateXHotel {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetName(v string) *TaobaoXhotelUpdateXHotel {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetDomestic(v int64) *TaobaoXhotelUpdateXHotel {
    s.Domestic = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetCountry(v string) *TaobaoXhotelUpdateXHotel {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetUsedName(v string) *TaobaoXhotelUpdateXHotel {
    s.UsedName = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetProvince(v int64) *TaobaoXhotelUpdateXHotel {
    s.Province = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetCity(v int64) *TaobaoXhotelUpdateXHotel {
    s.City = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetDistrict(v int64) *TaobaoXhotelUpdateXHotel {
    s.District = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetBusiness(v string) *TaobaoXhotelUpdateXHotel {
    s.Business = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetAddress(v string) *TaobaoXhotelUpdateXHotel {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetLongitude(v string) *TaobaoXhotelUpdateXHotel {
    s.Longitude = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetLatitude(v string) *TaobaoXhotelUpdateXHotel {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetPositionType(v string) *TaobaoXhotelUpdateXHotel {
    s.PositionType = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetTel(v string) *TaobaoXhotelUpdateXHotel {
    s.Tel = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetExtend(v string) *TaobaoXhotelUpdateXHotel {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetCreditCardTypes(v string) *TaobaoXhotelUpdateXHotel {
    s.CreditCardTypes = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetNameE(v string) *TaobaoXhotelUpdateXHotel {
    s.NameE = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetHotelType(v int64) *TaobaoXhotelUpdateXHotel {
    s.HotelType = &v
    return s
}
func (s *TaobaoXhotelUpdateXHotel) SetServiceType(v int64) *TaobaoXhotelUpdateXHotel {
    s.ServiceType = &v
    return s
}
