package domain


type AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO struct {
    /*
        日期。yyyy-MM-dd     */
    Date  *string `json:"date,omitempty" `

    /*
        场次价库信息     */
    Sessions  *[]AlitripTravelHotelticketProductProductupdatepushProductSessionDTO `json:"sessions,omitempty" `

    /*
        是否可售卖；true：可售卖     */
    CanSell  *bool `json:"can_sell,omitempty" `

    /*
        库存     */
    Stock  *int64 `json:"stock,omitempty" `

    /*
        产品结算单价。单位：分     */
    WholesalePrice  *int64 `json:"wholesale_price,omitempty" `

    /*
        产品销售单价。单位：分     */
    RetailPrice  *int64 `json:"retail_price,omitempty" `

}

func (s *AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO) SetDate(v string) *AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO {
    s.Date = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO) SetSessions(v []AlitripTravelHotelticketProductProductupdatepushProductSessionDTO) *AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO {
    s.Sessions = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO) SetCanSell(v bool) *AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO {
    s.CanSell = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO) SetStock(v int64) *AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO {
    s.Stock = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO) SetWholesalePrice(v int64) *AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO {
    s.WholesalePrice = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO) SetRetailPrice(v int64) *AlitripTravelHotelticketProductProductupdatepushProductPriceStockDTO {
    s.RetailPrice = &v
    return s
}
