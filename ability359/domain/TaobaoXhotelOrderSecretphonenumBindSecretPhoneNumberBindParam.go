package domain


type TaobaoXhotelOrderSecretphonenumBindSecretPhoneNumberBindParam struct {
    /*
        酒店订单tid     */
    Tid  *int64 `json:"tid,omitempty" `

}

func (s *TaobaoXhotelOrderSecretphonenumBindSecretPhoneNumberBindParam) SetTid(v int64) *TaobaoXhotelOrderSecretphonenumBindSecretPhoneNumberBindParam {
    s.Tid = &v
    return s
}
