package domain


type TaobaoXhotelAddXHotel struct {
    /*
        酒店ID     */
    Hid  *int64 `json:"hid,omitempty" `

    /*
        酒店状态：0: 正常;-2:停售；-1：删除     */
    Status  *int64 `json:"status,omitempty" `

    /*
        淘宝标准酒店信息     */
    SHotel  *TaobaoXhotelAddSHotel `json:"s_hotel,omitempty" `

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
        此字段已废弃     */
    MatchStatus  *int64 `json:"match_status,omitempty" `

    /*
        未通过时的拒绝原因等。     */
    ErrorInfo  *string `json:"error_info,omitempty" `

    /*
        逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡     */
    CreditCardTypes  *string `json:"credit_card_types,omitempty" `

    /*
        卖家酒店英文名称     */
    NameE  *string `json:"name_e,omitempty" `

}

func (s *TaobaoXhotelAddXHotel) SetHid(v int64) *TaobaoXhotelAddXHotel {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetStatus(v int64) *TaobaoXhotelAddXHotel {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetSHotel(v TaobaoXhotelAddSHotel) *TaobaoXhotelAddXHotel {
    s.SHotel = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetOuterId(v string) *TaobaoXhotelAddXHotel {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetName(v string) *TaobaoXhotelAddXHotel {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetDomestic(v int64) *TaobaoXhotelAddXHotel {
    s.Domestic = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetCountry(v string) *TaobaoXhotelAddXHotel {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetUsedName(v string) *TaobaoXhotelAddXHotel {
    s.UsedName = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetProvince(v int64) *TaobaoXhotelAddXHotel {
    s.Province = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetCity(v int64) *TaobaoXhotelAddXHotel {
    s.City = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetDistrict(v int64) *TaobaoXhotelAddXHotel {
    s.District = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetBusiness(v string) *TaobaoXhotelAddXHotel {
    s.Business = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetAddress(v string) *TaobaoXhotelAddXHotel {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetLongitude(v string) *TaobaoXhotelAddXHotel {
    s.Longitude = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetLatitude(v string) *TaobaoXhotelAddXHotel {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetPositionType(v string) *TaobaoXhotelAddXHotel {
    s.PositionType = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetTel(v string) *TaobaoXhotelAddXHotel {
    s.Tel = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetExtend(v string) *TaobaoXhotelAddXHotel {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetMatchStatus(v int64) *TaobaoXhotelAddXHotel {
    s.MatchStatus = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetErrorInfo(v string) *TaobaoXhotelAddXHotel {
    s.ErrorInfo = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetCreditCardTypes(v string) *TaobaoXhotelAddXHotel {
    s.CreditCardTypes = &v
    return s
}
func (s *TaobaoXhotelAddXHotel) SetNameE(v string) *TaobaoXhotelAddXHotel {
    s.NameE = &v
    return s
}
