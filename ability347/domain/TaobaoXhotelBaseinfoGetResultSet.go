package domain


type TaobaoXhotelBaseinfoGetResultSet struct {
    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        酒店基础信息     */
    XhotelBaseInfo  *TaobaoXhotelBaseinfoGetXHotelBaseInfo `json:"xhotel_base_info,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoXhotelBaseinfoGetResultSet) SetErrorCode(v string) *TaobaoXhotelBaseinfoGetResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetResultSet) SetErrorMsg(v string) *TaobaoXhotelBaseinfoGetResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetResultSet) SetXhotelBaseInfo(v TaobaoXhotelBaseinfoGetXHotelBaseInfo) *TaobaoXhotelBaseinfoGetResultSet {
    s.XhotelBaseInfo = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetResultSet) SetSuccess(v bool) *TaobaoXhotelBaseinfoGetResultSet {
    s.Success = &v
    return s
}
