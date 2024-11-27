package domain


type TaobaoXhotelBnbpromoUpdateResultSet struct {
    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoXhotelBnbpromoUpdateResultSet) SetErrorCode(v string) *TaobaoXhotelBnbpromoUpdateResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUpdateResultSet) SetErrorMsg(v string) *TaobaoXhotelBnbpromoUpdateResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUpdateResultSet) SetSuccess(v bool) *TaobaoXhotelBnbpromoUpdateResultSet {
    s.Success = &v
    return s
}
