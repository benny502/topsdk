package domain


type TaobaoXhotelHouseAddXHotel struct {
    /*
        酒店ID     */
    Hid  *int64 `json:"hid,omitempty" `

    /*
        酒店状态：0: 正常;-2:停售；-1：删除     */
    Status  *int64 `json:"status,omitempty" `

    /*
        淘宝标准酒店信息     */
    SHotel  *TaobaoXhotelHouseAddSHotel `json:"s_hotel,omitempty" `

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
        hotel匹配状态: 0：待系统匹配 1：已系统匹配，匹配成功，待卖家确认 2：已系统匹配，匹配失败，待人工匹配 3：已人工匹配，匹配成功，待卖家确认 4：已人工匹配，匹配失败 5：卖家已确认，确认&ldquo;YES&rdquo; 6：卖家已确认，确认&ldquo;NO&rdquo; 7:已系统匹配，但是匹配重复，待人工确认     */
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

func (s *TaobaoXhotelHouseAddXHotel) SetHid(v int64) *TaobaoXhotelHouseAddXHotel {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetStatus(v int64) *TaobaoXhotelHouseAddXHotel {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetSHotel(v TaobaoXhotelHouseAddSHotel) *TaobaoXhotelHouseAddXHotel {
    s.SHotel = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetOuterId(v string) *TaobaoXhotelHouseAddXHotel {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetName(v string) *TaobaoXhotelHouseAddXHotel {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetDomestic(v int64) *TaobaoXhotelHouseAddXHotel {
    s.Domestic = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetCountry(v string) *TaobaoXhotelHouseAddXHotel {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetUsedName(v string) *TaobaoXhotelHouseAddXHotel {
    s.UsedName = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetProvince(v int64) *TaobaoXhotelHouseAddXHotel {
    s.Province = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetCity(v int64) *TaobaoXhotelHouseAddXHotel {
    s.City = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetDistrict(v int64) *TaobaoXhotelHouseAddXHotel {
    s.District = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetBusiness(v string) *TaobaoXhotelHouseAddXHotel {
    s.Business = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetAddress(v string) *TaobaoXhotelHouseAddXHotel {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetLongitude(v string) *TaobaoXhotelHouseAddXHotel {
    s.Longitude = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetLatitude(v string) *TaobaoXhotelHouseAddXHotel {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetPositionType(v string) *TaobaoXhotelHouseAddXHotel {
    s.PositionType = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetTel(v string) *TaobaoXhotelHouseAddXHotel {
    s.Tel = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetExtend(v string) *TaobaoXhotelHouseAddXHotel {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetMatchStatus(v int64) *TaobaoXhotelHouseAddXHotel {
    s.MatchStatus = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetErrorInfo(v string) *TaobaoXhotelHouseAddXHotel {
    s.ErrorInfo = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetCreditCardTypes(v string) *TaobaoXhotelHouseAddXHotel {
    s.CreditCardTypes = &v
    return s
}
func (s *TaobaoXhotelHouseAddXHotel) SetNameE(v string) *TaobaoXhotelHouseAddXHotel {
    s.NameE = &v
    return s
}
