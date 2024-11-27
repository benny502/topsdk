package domain


type TaobaoXhotelRateDeleteResultSet struct {
    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        rateid-房型id-房价id     */
    RateidGidRpid  *string `json:"rateid_gid_rpid,omitempty" `

}

func (s *TaobaoXhotelRateDeleteResultSet) SetErrorCode(v string) *TaobaoXhotelRateDeleteResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelRateDeleteResultSet) SetErrorMsg(v string) *TaobaoXhotelRateDeleteResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelRateDeleteResultSet) SetRateidGidRpid(v string) *TaobaoXhotelRateDeleteResultSet {
    s.RateidGidRpid = &v
    return s
}
