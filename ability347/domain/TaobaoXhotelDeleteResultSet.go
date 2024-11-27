package domain


type TaobaoXhotelDeleteResultSet struct {
    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        是否出错     */
    Error  *bool `json:"error,omitempty" `

}

func (s *TaobaoXhotelDeleteResultSet) SetErrorCode(v string) *TaobaoXhotelDeleteResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelDeleteResultSet) SetErrorMsg(v string) *TaobaoXhotelDeleteResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelDeleteResultSet) SetError(v bool) *TaobaoXhotelDeleteResultSet {
    s.Error = &v
    return s
}
