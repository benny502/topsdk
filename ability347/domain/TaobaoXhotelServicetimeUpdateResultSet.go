package domain


type TaobaoXhotelServicetimeUpdateResultSet struct {
    /*
        results     */
    Results  *[]string `json:"results,omitempty" `

    /*
        totalResults     */
    TotalResults  *int64 `json:"total_results,omitempty" `

    /*
        hasNext     */
    HasNext  *bool `json:"has_next,omitempty" `

    /*
        exception     */
    Exception  *string `json:"exception,omitempty" `

    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        warnMessage     */
    WarnMessage  *string `json:"warn_message,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoXhotelServicetimeUpdateResultSet) SetResults(v []string) *TaobaoXhotelServicetimeUpdateResultSet {
    s.Results = &v
    return s
}
func (s *TaobaoXhotelServicetimeUpdateResultSet) SetTotalResults(v int64) *TaobaoXhotelServicetimeUpdateResultSet {
    s.TotalResults = &v
    return s
}
func (s *TaobaoXhotelServicetimeUpdateResultSet) SetHasNext(v bool) *TaobaoXhotelServicetimeUpdateResultSet {
    s.HasNext = &v
    return s
}
func (s *TaobaoXhotelServicetimeUpdateResultSet) SetException(v string) *TaobaoXhotelServicetimeUpdateResultSet {
    s.Exception = &v
    return s
}
func (s *TaobaoXhotelServicetimeUpdateResultSet) SetErrorCode(v string) *TaobaoXhotelServicetimeUpdateResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelServicetimeUpdateResultSet) SetWarnMessage(v string) *TaobaoXhotelServicetimeUpdateResultSet {
    s.WarnMessage = &v
    return s
}
func (s *TaobaoXhotelServicetimeUpdateResultSet) SetErrorMsg(v string) *TaobaoXhotelServicetimeUpdateResultSet {
    s.ErrorMsg = &v
    return s
}
