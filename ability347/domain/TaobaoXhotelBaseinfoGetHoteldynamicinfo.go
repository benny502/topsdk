package domain


type TaobaoXhotelBaseinfoGetHoteldynamicinfo struct {
    /*
        不可售原因     */
    UnsaleReason  *string `json:"unsale_reason,omitempty" `

    /*
        酒店hid     */
    Hid  *int64 `json:"hid,omitempty" `

    /*
        酒店状态     */
    Status  *int64 `json:"status,omitempty" `

    /*
        酒店的销售渠道     */
    Vendor  *string `json:"vendor,omitempty" `

    /*
        可售健康房型数     */
    KsHeathyRoomNum  *int64 `json:"ks_heathy_room_num,omitempty" `

    /*
        电话     */
    Tel  *string `json:"tel,omitempty" `

    /*
        标准酒店ID     */
    Shid  *int64 `json:"shid,omitempty" `

    /*
        城市名称     */
    CityStr  *string `json:"city_str,omitempty" `

    /*
        城市编码     */
    City  *int64 `json:"city,omitempty" `

    /*
        id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        房型数     */
    RoomNun  *int64 `json:"room_nun,omitempty" `

    /*
        卖家昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        地址     */
    Address  *string `json:"address,omitempty" `

    /*
        可售房型数     */
    KsRoomNum  *int64 `json:"ks_room_num,omitempty" `

    /*
        卖家ID     */
    SellerId  *int64 `json:"seller_id,omitempty" `

    /*
        不可售原因     */
    UnsaleType  *int64 `json:"unsale_type,omitempty" `

    /*
        酒店名字     */
    Name  *string `json:"name,omitempty" `

    /*
        分析日期     */
    CalculateDate  *string `json:"calculate_date,omitempty" `

    /*
        酒店匹配     */
    DataConfirm  *int64 `json:"data_confirm,omitempty" `

    /*
        酒店外部ID     */
    OuterId  *string `json:"outer_id,omitempty" `

}

func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetUnsaleReason(v string) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.UnsaleReason = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetHid(v int64) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetStatus(v int64) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetVendor(v string) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetKsHeathyRoomNum(v int64) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.KsHeathyRoomNum = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetTel(v string) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.Tel = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetShid(v int64) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.Shid = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetCityStr(v string) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.CityStr = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetCity(v int64) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.City = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetId(v int64) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.Id = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetRoomNun(v int64) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.RoomNun = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetSellerNick(v string) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.SellerNick = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetAddress(v string) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetKsRoomNum(v int64) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.KsRoomNum = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetSellerId(v int64) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.SellerId = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetUnsaleType(v int64) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.UnsaleType = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetName(v string) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetCalculateDate(v string) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.CalculateDate = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetDataConfirm(v int64) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.DataConfirm = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetHoteldynamicinfo) SetOuterId(v string) *TaobaoXhotelBaseinfoGetHoteldynamicinfo {
    s.OuterId = &v
    return s
}
