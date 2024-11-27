package domain


type TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto struct {
    /*
        开始接待时间 hh:mm,24小时时间格式     */
    StartReceptionTime  *string `json:"start_reception_time,omitempty" `

    /*
        结束接待时间 hh:mm,24小时时间格式     */
    EndReceptionTime  *string `json:"end_reception_time,omitempty" `

    /*
        最早入住时间 hh:mm,24小时时间格式;默认值: 14:00 defalutValue:14:00    */
    EarliestCheckInTime  *string `json:"earliest_check_in_time,omitempty" `

    /*
        最晚预订时间 hh:mm,24小时时间格式     */
    LatestBookingTime  *string `json:"latest_booking_time,omitempty" `

    /*
        最晚入住时间 hh:mm,24小时时间格式     */
    LatestCheckInTime  *string `json:"latest_check_in_time,omitempty" `

    /*
        最晚离店时间 hh:mm,24小时时间格式;默认值: 12:00 defalutValue:12:00    */
    LatestCheckOutTime  *string `json:"latest_check_out_time,omitempty" `

}

func (s *TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto) SetStartReceptionTime(v string) *TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto {
    s.StartReceptionTime = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto) SetEndReceptionTime(v string) *TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto {
    s.EndReceptionTime = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto) SetEarliestCheckInTime(v string) *TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto {
    s.EarliestCheckInTime = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto) SetLatestBookingTime(v string) *TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto {
    s.LatestBookingTime = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto) SetLatestCheckInTime(v string) *TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto {
    s.LatestCheckInTime = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto) SetLatestCheckOutTime(v string) *TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto {
    s.LatestCheckOutTime = &v
    return s
}
