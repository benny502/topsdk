package domain


type TaobaoXhotelBnbownerDeleteResultSet struct {
    /*
        系统自动生成     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        是否出错     */
    Error  *bool `json:"error,omitempty" `

    /*
        系统自动生成     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoXhotelBnbownerDeleteResultSet) SetErrorCode(v string) *TaobaoXhotelBnbownerDeleteResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelBnbownerDeleteResultSet) SetError(v bool) *TaobaoXhotelBnbownerDeleteResultSet {
    s.Error = &v
    return s
}
func (s *TaobaoXhotelBnbownerDeleteResultSet) SetErrorMsg(v string) *TaobaoXhotelBnbownerDeleteResultSet {
    s.ErrorMsg = &v
    return s
}
