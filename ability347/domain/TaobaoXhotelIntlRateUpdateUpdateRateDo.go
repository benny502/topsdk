package domain


type TaobaoXhotelIntlRateUpdateUpdateRateDo struct {
    /*
        入离日期     */
    CheckDateDOList  *[]TaobaoXhotelIntlRateUpdateCheckDateDo `json:"check_date_d_o_list,omitempty" `

    /*
        成人数 defalutValue:0    */
    Adults  *int64 `json:"adults,omitempty" `

    /*
        酒店id     */
    OutHid  *string `json:"out_hid,omitempty" `

    /*
        儿童数 defalutValue:0    */
    Children  *int64 `json:"children,omitempty" `

    /*
        儿童年龄，多个儿童年龄用竖线分割：4|5，默认儿童年龄为5岁     */
    Ages  *string `json:"ages,omitempty" `

}

func (s *TaobaoXhotelIntlRateUpdateUpdateRateDo) SetCheckDateDOList(v []TaobaoXhotelIntlRateUpdateCheckDateDo) *TaobaoXhotelIntlRateUpdateUpdateRateDo {
    s.CheckDateDOList = &v
    return s
}
func (s *TaobaoXhotelIntlRateUpdateUpdateRateDo) SetAdults(v int64) *TaobaoXhotelIntlRateUpdateUpdateRateDo {
    s.Adults = &v
    return s
}
func (s *TaobaoXhotelIntlRateUpdateUpdateRateDo) SetOutHid(v string) *TaobaoXhotelIntlRateUpdateUpdateRateDo {
    s.OutHid = &v
    return s
}
func (s *TaobaoXhotelIntlRateUpdateUpdateRateDo) SetChildren(v int64) *TaobaoXhotelIntlRateUpdateUpdateRateDo {
    s.Children = &v
    return s
}
func (s *TaobaoXhotelIntlRateUpdateUpdateRateDo) SetAges(v string) *TaobaoXhotelIntlRateUpdateUpdateRateDo {
    s.Ages = &v
    return s
}
