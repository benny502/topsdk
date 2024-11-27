package domain


type TaobaoXhotelIntlRateUpdateResultSet struct {
    /*
        total_results     */
    TotalResults  *int64 `json:"total_results,omitempty" `

    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        sub_error_msg     */
    SubErrorMsg  *string `json:"sub_error_msg,omitempty" `

    /*
        sub_error_code     */
    SubErrorCode  *string `json:"sub_error_code,omitempty" `

    /*
        detail_msg_info     */
    DetailMsgInfo  *string `json:"detail_msg_info,omitempty" `

    /*
        总记录条数     */
    TotalCount  *int64 `json:"total_count,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoXhotelIntlRateUpdateResultSet) SetTotalResults(v int64) *TaobaoXhotelIntlRateUpdateResultSet {
    s.TotalResults = &v
    return s
}
func (s *TaobaoXhotelIntlRateUpdateResultSet) SetErrorCode(v string) *TaobaoXhotelIntlRateUpdateResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelIntlRateUpdateResultSet) SetSubErrorMsg(v string) *TaobaoXhotelIntlRateUpdateResultSet {
    s.SubErrorMsg = &v
    return s
}
func (s *TaobaoXhotelIntlRateUpdateResultSet) SetSubErrorCode(v string) *TaobaoXhotelIntlRateUpdateResultSet {
    s.SubErrorCode = &v
    return s
}
func (s *TaobaoXhotelIntlRateUpdateResultSet) SetDetailMsgInfo(v string) *TaobaoXhotelIntlRateUpdateResultSet {
    s.DetailMsgInfo = &v
    return s
}
func (s *TaobaoXhotelIntlRateUpdateResultSet) SetTotalCount(v int64) *TaobaoXhotelIntlRateUpdateResultSet {
    s.TotalCount = &v
    return s
}
func (s *TaobaoXhotelIntlRateUpdateResultSet) SetErrorMsg(v string) *TaobaoXhotelIntlRateUpdateResultSet {
    s.ErrorMsg = &v
    return s
}
