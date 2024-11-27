package domain


type TaobaoXhotelBaseinfoRoomGetResultSet struct {
    /*
        success     */
    Success  *bool `json:"success,omitempty" `

    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        酒店基础信息     */
    XhotelBaseInfo  *TaobaoXhotelBaseinfoRoomGetXHotelInfoWithRoom `json:"xhotel_base_info,omitempty" `

}

func (s *TaobaoXhotelBaseinfoRoomGetResultSet) SetSuccess(v bool) *TaobaoXhotelBaseinfoRoomGetResultSet {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelBaseinfoRoomGetResultSet) SetErrorCode(v string) *TaobaoXhotelBaseinfoRoomGetResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelBaseinfoRoomGetResultSet) SetErrorMsg(v string) *TaobaoXhotelBaseinfoRoomGetResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelBaseinfoRoomGetResultSet) SetXhotelBaseInfo(v TaobaoXhotelBaseinfoRoomGetXHotelInfoWithRoom) *TaobaoXhotelBaseinfoRoomGetResultSet {
    s.XhotelBaseInfo = &v
    return s
}
