package domain


type AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO struct {
    /*
        日历价格库存信息  日历价格库存信息     */
    PriceStocks  *[]AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO `json:"price_stocks,omitempty" `

    /*
        产品变更通知类型 1：价格，2：库存，3：价库     */
    NotifyType  *int64 `json:"notify_type,omitempty" `

    /*
        系统商商品码     */
    ProductId  *string `json:"product_id,omitempty" `

    /*
        床型ID     */
    BedId  *string `json:"bed_id,omitempty" `

    /*
        酒店ID     */
    HotelId  *string `json:"hotel_id,omitempty" `

    /*
        房型ID     */
    RoomId  *string `json:"room_id,omitempty" `

    /*
        扩展参数     */
    ExtendParams  *string `json:"extend_params,omitempty" `

    /*
        模式 默认值1；1:普通日历/预约商品（非通兑和非任选） defalutValue:1    */
    Mode  *int64 `json:"mode,omitempty" `

}

func (s *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO) SetPriceStocks(v []AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO) *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO {
    s.PriceStocks = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO) SetNotifyType(v int64) *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO {
    s.NotifyType = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO) SetProductId(v string) *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO {
    s.ProductId = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO) SetBedId(v string) *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO {
    s.BedId = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO) SetHotelId(v string) *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO {
    s.HotelId = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO) SetRoomId(v string) *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO {
    s.RoomId = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO) SetExtendParams(v string) *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO {
    s.ExtendParams = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO) SetMode(v int64) *AlitripTravelHotelticketProductProductupdatepushProductUpdatePushDTO {
    s.Mode = &v
    return s
}
