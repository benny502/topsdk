package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelBnbpromoUpdateInvalidDate struct {
    /*
        活动失效开始时间     */
    InvalidFrom  *util.LocalTime `json:"invalid_from,omitempty" `

    /*
        活动失效结束时间     */
    InvalidTo  *util.LocalTime `json:"invalid_to,omitempty" `

}

func (s *TaobaoXhotelBnbpromoUpdateInvalidDate) SetInvalidFrom(v util.LocalTime) *TaobaoXhotelBnbpromoUpdateInvalidDate {
    s.InvalidFrom = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUpdateInvalidDate) SetInvalidTo(v util.LocalTime) *TaobaoXhotelBnbpromoUpdateInvalidDate {
    s.InvalidTo = &v
    return s
}
