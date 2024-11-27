package domain


type TaobaoXhotelXitemDeleteResultSet struct {
    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

}

func (s *TaobaoXhotelXitemDeleteResultSet) SetErrorMsg(v string) *TaobaoXhotelXitemDeleteResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelXitemDeleteResultSet) SetErrorCode(v string) *TaobaoXhotelXitemDeleteResultSet {
    s.ErrorCode = &v
    return s
}
