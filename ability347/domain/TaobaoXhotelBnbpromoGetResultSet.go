package domain


type TaobaoXhotelBnbpromoGetResultSet struct {
    /*
        民宿活动信息     */
    BnbPromo  *TaobaoXhotelBnbpromoGetBnbPromoDTO `json:"bnb_promo,omitempty" `

    /*
        错误code     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误码     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoXhotelBnbpromoGetResultSet) SetBnbPromo(v TaobaoXhotelBnbpromoGetBnbPromoDTO) *TaobaoXhotelBnbpromoGetResultSet {
    s.BnbPromo = &v
    return s
}
func (s *TaobaoXhotelBnbpromoGetResultSet) SetErrorCode(v string) *TaobaoXhotelBnbpromoGetResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelBnbpromoGetResultSet) SetErrorMsg(v string) *TaobaoXhotelBnbpromoGetResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelBnbpromoGetResultSet) SetSuccess(v bool) *TaobaoXhotelBnbpromoGetResultSet {
    s.Success = &v
    return s
}
