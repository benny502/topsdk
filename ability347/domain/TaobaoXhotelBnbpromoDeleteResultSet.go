package domain


type TaobaoXhotelBnbpromoDeleteResultSet struct {
    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误code     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误码     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoXhotelBnbpromoDeleteResultSet) SetSuccess(v bool) *TaobaoXhotelBnbpromoDeleteResultSet {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelBnbpromoDeleteResultSet) SetErrorCode(v string) *TaobaoXhotelBnbpromoDeleteResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelBnbpromoDeleteResultSet) SetErrorMsg(v string) *TaobaoXhotelBnbpromoDeleteResultSet {
    s.ErrorMsg = &v
    return s
}
