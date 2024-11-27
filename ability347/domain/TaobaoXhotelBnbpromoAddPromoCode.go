package domain


type TaobaoXhotelBnbpromoAddPromoCode struct {
    /*
        营销活动code     */
    ActivityCode  *string `json:"activity_code,omitempty" `

}

func (s *TaobaoXhotelBnbpromoAddPromoCode) SetActivityCode(v string) *TaobaoXhotelBnbpromoAddPromoCode {
    s.ActivityCode = &v
    return s
}
