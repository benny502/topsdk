package domain


type TaobaoXhotelServicetimeGetServiceTimeDataDo struct {
    /*
        supplier     */
    Supplier  *string `json:"supplier,omitempty" `

    /*
        卖家nick     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        卖家id     */
    SellerId  *int64 `json:"seller_id,omitempty" `

    /*
        业务类型：1卖家；2supplier;3酒店     */
    BusinessType  *string `json:"business_type,omitempty" `

    /*
        业务id     */
    BusinessId  *int64 `json:"business_id,omitempty" `

    /*
        timeZoneName     */
    TimeZoneName  *string `json:"time_zone_name,omitempty" `

    /*
        1: 当日订单, 2:次日及以后订单     */
    OrderConfirmType  *int64 `json:"order_confirm_type,omitempty" `

    /*
        是否在非工作时间显示商品 1:显示, 2:不显示     */
    DisplayItemInNonworkingTime  *int64 `json:"display_item_in_nonworking_time,omitempty" `

    /*
        周五服务时间（当地时间）     */
    FridayConfirmLocalTime  *string `json:"friday_confirm_local_time,omitempty" `

    /*
        周一服务时间（当地时间）     */
    MondayConfirmLocalTime  *string `json:"monday_confirm_local_time,omitempty" `

    /*
        周二服务时间（当地时间）     */
    TuesdayConfirmLocalTime  *string `json:"tuesday_confirm_local_time,omitempty" `

    /*
        id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        周三服务时间（当地时间）     */
    WednesdayConfirmLocalTime  *string `json:"wednesday_confirm_local_time,omitempty" `

    /*
        周六服务时间（当地时间）     */
    SaturdayConfirmLocalTime  *string `json:"saturday_confirm_local_time,omitempty" `

    /*
        operator     */
    Operator  *string `json:"operator,omitempty" `

    /*
        周日服务时间（当地时间）     */
    SundayConfirmLocalTime  *string `json:"sunday_confirm_local_time,omitempty" `

    /*
        周四服务时间（当地时间）     */
    ThursdayConfirmLocalTime  *string `json:"thursday_confirm_local_time,omitempty" `

    /*
        创建时间     */
    GmtCreate  *string `json:"gmt_create,omitempty" `

    /*
        最后修改时间     */
    GmtModified  *string `json:"gmt_modified,omitempty" `

}

func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetSupplier(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.Supplier = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetSellerNick(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.SellerNick = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetSellerId(v int64) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.SellerId = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetBusinessType(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.BusinessType = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetBusinessId(v int64) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.BusinessId = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetTimeZoneName(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.TimeZoneName = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetOrderConfirmType(v int64) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.OrderConfirmType = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetDisplayItemInNonworkingTime(v int64) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.DisplayItemInNonworkingTime = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetFridayConfirmLocalTime(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.FridayConfirmLocalTime = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetMondayConfirmLocalTime(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.MondayConfirmLocalTime = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetTuesdayConfirmLocalTime(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.TuesdayConfirmLocalTime = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetId(v int64) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.Id = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetWednesdayConfirmLocalTime(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.WednesdayConfirmLocalTime = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetSaturdayConfirmLocalTime(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.SaturdayConfirmLocalTime = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetOperator(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.Operator = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetSundayConfirmLocalTime(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.SundayConfirmLocalTime = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetThursdayConfirmLocalTime(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.ThursdayConfirmLocalTime = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetGmtCreate(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetServiceTimeDataDo) SetGmtModified(v string) *TaobaoXhotelServicetimeGetServiceTimeDataDo {
    s.GmtModified = &v
    return s
}
