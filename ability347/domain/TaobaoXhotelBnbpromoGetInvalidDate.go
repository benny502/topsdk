package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelBnbpromoGetInvalidDate struct {
	/*
	   活动失效开始时间     */
	InvalidFrom *util.LocalTime `json:"invalid_from,omitempty" `

	/*
	   活动失效结束时间     */
	InvalidTo *util.LocalTime `json:"invalid_to,omitempty" `
}

func (s *TaobaoXhotelBnbpromoGetInvalidDate) SetInvalidFrom(v util.LocalTime) *TaobaoXhotelBnbpromoGetInvalidDate {
	s.InvalidFrom = &v
	return s
}
func (s *TaobaoXhotelBnbpromoGetInvalidDate) SetInvalidTo(v util.LocalTime) *TaobaoXhotelBnbpromoGetInvalidDate {
	s.InvalidTo = &v
	return s
}
