package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelBnbpromoUpdateDailyBookingInfo struct {
	/*
	   一个星期内有效性约束。1-7 对应周一到周日，传入的值比如[1,6]，就表示星期一和星期六营销生效     */
	ValidWeeks *[]int32 `json:"valid_weeks,omitempty" `

	/*
	   可入住的起始时间，不填默认一年，一年后自动续期     */
	CheckInFrom *util.LocalTime `json:"check_in_from,omitempty" `

	/*
	   可入住的结束时间，不填默认一年，一年后自动续期     */
	CheckInTo *util.LocalTime `json:"check_in_to,omitempty" `

	/*
	   折扣比例，填30就意味着原价的30%，也就是打3折。数字范围限定在10-95之间     */
	InvestmentNumber *int64 `json:"investment_number,omitempty" `

	/*
	   不可用日期，开始日期和结束日期: from--to  只有一天的场景，from和to传同一天； 默认：空，代表无时间限制；     */
	InvalidDates *[]TaobaoXhotelBnbpromoUpdateInvalidDate `json:"invalid_dates,omitempty" `
}

func (s *TaobaoXhotelBnbpromoUpdateDailyBookingInfo) SetValidWeeks(v []int32) *TaobaoXhotelBnbpromoUpdateDailyBookingInfo {
	s.ValidWeeks = &v
	return s
}
func (s *TaobaoXhotelBnbpromoUpdateDailyBookingInfo) SetCheckInFrom(v util.LocalTime) *TaobaoXhotelBnbpromoUpdateDailyBookingInfo {
	s.CheckInFrom = &v
	return s
}
func (s *TaobaoXhotelBnbpromoUpdateDailyBookingInfo) SetCheckInTo(v util.LocalTime) *TaobaoXhotelBnbpromoUpdateDailyBookingInfo {
	s.CheckInTo = &v
	return s
}
func (s *TaobaoXhotelBnbpromoUpdateDailyBookingInfo) SetInvestmentNumber(v int64) *TaobaoXhotelBnbpromoUpdateDailyBookingInfo {
	s.InvestmentNumber = &v
	return s
}
func (s *TaobaoXhotelBnbpromoUpdateDailyBookingInfo) SetInvalidDates(v []TaobaoXhotelBnbpromoUpdateInvalidDate) *TaobaoXhotelBnbpromoUpdateDailyBookingInfo {
	s.InvalidDates = &v
	return s
}
