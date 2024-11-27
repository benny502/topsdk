package domain


type TaobaoXhotelServicetimeGetResultSet struct {
    /*
        firstResult     */
    FirstResults  *[]TaobaoXhotelServicetimeGetServiceTimeDataDo `json:"first_results,omitempty" `

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

    /*
        results     */
    Results  *[]string `json:"results,omitempty" `

}

func (s *TaobaoXhotelServicetimeGetResultSet) SetFirstResults(v []TaobaoXhotelServicetimeGetServiceTimeDataDo) *TaobaoXhotelServicetimeGetResultSet {
    s.FirstResults = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetResultSet) SetTotalResults(v int64) *TaobaoXhotelServicetimeGetResultSet {
    s.TotalResults = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetResultSet) SetHasNext(v bool) *TaobaoXhotelServicetimeGetResultSet {
    s.HasNext = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetResultSet) SetException(v string) *TaobaoXhotelServicetimeGetResultSet {
    s.Exception = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetResultSet) SetErrorCode(v string) *TaobaoXhotelServicetimeGetResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetResultSet) SetWarnMessage(v string) *TaobaoXhotelServicetimeGetResultSet {
    s.WarnMessage = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetResultSet) SetErrorMsg(v string) *TaobaoXhotelServicetimeGetResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelServicetimeGetResultSet) SetResults(v []string) *TaobaoXhotelServicetimeGetResultSet {
    s.Results = &v
    return s
}
