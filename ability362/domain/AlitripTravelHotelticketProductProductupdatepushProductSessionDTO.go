package domain


type AlitripTravelHotelticketProductProductupdatepushProductSessionDTO struct {
    /*
        开始时间。HH:mm     */
    StartTime  *string `json:"start_time,omitempty" `

    /*
        场次ID     */
    SessionId  *string `json:"session_id,omitempty" `

    /*
        结束时间。HH:mm     */
    EndTime  *string `json:"end_time,omitempty" `

    /*
        场次库存     */
    Stock  *int64 `json:"stock,omitempty" `

    /*
        产品场次结算单价。单位：分     */
    WholesalePrice  *int64 `json:"wholesale_price,omitempty" `

    /*
        产品场次销售单价。单位：分     */
    RetailPrice  *int64 `json:"retail_price,omitempty" `

}

func (s *AlitripTravelHotelticketProductProductupdatepushProductSessionDTO) SetStartTime(v string) *AlitripTravelHotelticketProductProductupdatepushProductSessionDTO {
    s.StartTime = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductSessionDTO) SetSessionId(v string) *AlitripTravelHotelticketProductProductupdatepushProductSessionDTO {
    s.SessionId = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductSessionDTO) SetEndTime(v string) *AlitripTravelHotelticketProductProductupdatepushProductSessionDTO {
    s.EndTime = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductSessionDTO) SetStock(v int64) *AlitripTravelHotelticketProductProductupdatepushProductSessionDTO {
    s.Stock = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductSessionDTO) SetWholesalePrice(v int64) *AlitripTravelHotelticketProductProductupdatepushProductSessionDTO {
    s.WholesalePrice = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdatepushProductSessionDTO) SetRetailPrice(v int64) *AlitripTravelHotelticketProductProductupdatepushProductSessionDTO {
    s.RetailPrice = &v
    return s
}
