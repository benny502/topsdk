package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelBnbpromoGetDailyBookingInfo struct {
	/*
	   入住开始     */
	CheckInFrom *util.LocalTime `json:"check_in_from,omitempty" `

	/*
	   固定折扣百分比     */
	InvestmentNumber *string `json:"investment_number,omitempty" `

	/*
	   一个星期内有效性约束     */
	ValidWeeks *[]int32 `json:"valid_weeks,omitempty" `

	/*
	   入住结束     */
	CheckInTo *util.LocalTime `json:"check_in_to,omitempty" `

	/*
	   失效时间     */
	InvalidDates *[]TaobaoXhotelBnbpromoGetInvalidDate `json:"invalid_dates,omitempty" `
}

func (s *TaobaoXhotelBnbpromoGetDailyBookingInfo) SetCheckInFrom(v util.LocalTime) *TaobaoXhotelBnbpromoGetDailyBookingInfo {
	s.CheckInFrom = &v
	return s
}
func (s *TaobaoXhotelBnbpromoGetDailyBookingInfo) SetInvestmentNumber(v string) *TaobaoXhotelBnbpromoGetDailyBookingInfo {
	s.InvestmentNumber = &v
	return s
}
func (s *TaobaoXhotelBnbpromoGetDailyBookingInfo) SetValidWeeks(v []int32) *TaobaoXhotelBnbpromoGetDailyBookingInfo {
	s.ValidWeeks = &v
	return s
}
func (s *TaobaoXhotelBnbpromoGetDailyBookingInfo) SetCheckInTo(v util.LocalTime) *TaobaoXhotelBnbpromoGetDailyBookingInfo {
	s.CheckInTo = &v
	return s
}
func (s *TaobaoXhotelBnbpromoGetDailyBookingInfo) SetInvalidDates(v []TaobaoXhotelBnbpromoGetInvalidDate) *TaobaoXhotelBnbpromoGetDailyBookingInfo {
	s.InvalidDates = &v
	return s
}
