package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelIntlRateUpdateCheckDateDo struct {
	/*
	   入住时间     */
	CheckOut *util.LocalTime `json:"check_out,omitempty" `

	/*
	   离店时间     */
	CheckIn *util.LocalTime `json:"check_in,omitempty" `
}

func (s *TaobaoXhotelIntlRateUpdateCheckDateDo) SetCheckOut(v util.LocalTime) *TaobaoXhotelIntlRateUpdateCheckDateDo {
	s.CheckOut = &v
	return s
}
func (s *TaobaoXhotelIntlRateUpdateCheckDateDo) SetCheckIn(v util.LocalTime) *TaobaoXhotelIntlRateUpdateCheckDateDo {
	s.CheckIn = &v
	return s
}
