package domain


type TaobaoXhotelBnbhouseAddBnbBookingTimeDto struct {
    /*
        开始接待时间 hh:mm,24小时时间格式     */
    StartReceptionTime  *string `json:"start_reception_time,omitempty" `

    /*
        结束接待时间 hh:mm,24小时时间格式     */
    EndReceptionTime  *string `json:"end_reception_time,omitempty" `

    /*
        最早入住时间 hh:mm,24小时时间格式住时间,默认值: 14:00 defalutValue:14：00    */
    EarliestCheckInTime  *string `json:"earliest_check_in_time,omitempty" `

    /*
        最晚预定时间 hh:mm,24小时时间格式     */
    LatestBookingTime  *string `json:"latest_booking_time,omitempty" `

    /*
        最晚入住时间 hh:mm,24小时时间格式     */
    LatestCheckInTime  *string `json:"latest_check_in_time,omitempty" `

    /*
        最晚离店时间 hh:mm,24小时时间格式,默认值: 12:00 defalutValue:12：00    */
    LatestCheckOutTime  *string `json:"latest_check_out_time,omitempty" `

}

func (s *TaobaoXhotelBnbhouseAddBnbBookingTimeDto) SetStartReceptionTime(v string) *TaobaoXhotelBnbhouseAddBnbBookingTimeDto {
    s.StartReceptionTime = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbBookingTimeDto) SetEndReceptionTime(v string) *TaobaoXhotelBnbhouseAddBnbBookingTimeDto {
    s.EndReceptionTime = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbBookingTimeDto) SetEarliestCheckInTime(v string) *TaobaoXhotelBnbhouseAddBnbBookingTimeDto {
    s.EarliestCheckInTime = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbBookingTimeDto) SetLatestBookingTime(v string) *TaobaoXhotelBnbhouseAddBnbBookingTimeDto {
    s.LatestBookingTime = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbBookingTimeDto) SetLatestCheckInTime(v string) *TaobaoXhotelBnbhouseAddBnbBookingTimeDto {
    s.LatestCheckInTime = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbBookingTimeDto) SetLatestCheckOutTime(v string) *TaobaoXhotelBnbhouseAddBnbBookingTimeDto {
    s.LatestCheckOutTime = &v
    return s
}
