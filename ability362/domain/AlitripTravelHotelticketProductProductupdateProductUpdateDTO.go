package domain


type AlitripTravelHotelticketProductProductupdateProductUpdateDTO struct {
    /*
        扩展参数     */
    ExtendParams  *string `json:"extend_params,omitempty" `

    /*
        产品变更通知类型 1：价格，2：库存，3：价库     */
    NotifyType  *int64 `json:"notify_type,omitempty" `

    /*
        系统商商品码     */
    ProductId  *string `json:"product_id,omitempty" `

    /*
        场次ID信息     */
    SessionIds  *[]string `json:"session_ids,omitempty" `

    /*
        产品变更开始时间 yyyy-MM-dd     */
    StartDate  *string `json:"start_date,omitempty" `

    /*
        产品变更结束时间 yyyy-MM-dd     */
    EndDate  *string `json:"end_date,omitempty" `

}

func (s *AlitripTravelHotelticketProductProductupdateProductUpdateDTO) SetExtendParams(v string) *AlitripTravelHotelticketProductProductupdateProductUpdateDTO {
    s.ExtendParams = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdateProductUpdateDTO) SetNotifyType(v int64) *AlitripTravelHotelticketProductProductupdateProductUpdateDTO {
    s.NotifyType = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdateProductUpdateDTO) SetProductId(v string) *AlitripTravelHotelticketProductProductupdateProductUpdateDTO {
    s.ProductId = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdateProductUpdateDTO) SetSessionIds(v []string) *AlitripTravelHotelticketProductProductupdateProductUpdateDTO {
    s.SessionIds = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdateProductUpdateDTO) SetStartDate(v string) *AlitripTravelHotelticketProductProductupdateProductUpdateDTO {
    s.StartDate = &v
    return s
}
func (s *AlitripTravelHotelticketProductProductupdateProductUpdateDTO) SetEndDate(v string) *AlitripTravelHotelticketProductProductupdateProductUpdateDTO {
    s.EndDate = &v
    return s
}
