package domain


type TaobaoXhotelRatesLiteIncrUpdateResultSet struct {
    /*
        多个rate的更新结果     */
    FirstResult  *string `json:"first_result,omitempty" `

    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoXhotelRatesLiteIncrUpdateResultSet) SetFirstResult(v string) *TaobaoXhotelRatesLiteIncrUpdateResultSet {
    s.FirstResult = &v
    return s
}
func (s *TaobaoXhotelRatesLiteIncrUpdateResultSet) SetErrorCode(v string) *TaobaoXhotelRatesLiteIncrUpdateResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelRatesLiteIncrUpdateResultSet) SetErrorMsg(v string) *TaobaoXhotelRatesLiteIncrUpdateResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelRatesLiteIncrUpdateResultSet) SetSuccess(v bool) *TaobaoXhotelRatesLiteIncrUpdateResultSet {
    s.Success = &v
    return s
}
