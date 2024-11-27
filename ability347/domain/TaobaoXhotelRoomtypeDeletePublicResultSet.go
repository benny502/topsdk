package domain


type TaobaoXhotelRoomtypeDeletePublicResultSet struct {
    /*
        success     */
    Success  *bool `json:"success,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        deleteResult     */
    DeleteResult  *string `json:"delete_result,omitempty" `

}

func (s *TaobaoXhotelRoomtypeDeletePublicResultSet) SetSuccess(v bool) *TaobaoXhotelRoomtypeDeletePublicResultSet {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelRoomtypeDeletePublicResultSet) SetErrorMsg(v string) *TaobaoXhotelRoomtypeDeletePublicResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelRoomtypeDeletePublicResultSet) SetErrorCode(v string) *TaobaoXhotelRoomtypeDeletePublicResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelRoomtypeDeletePublicResultSet) SetDeleteResult(v string) *TaobaoXhotelRoomtypeDeletePublicResultSet {
    s.DeleteResult = &v
    return s
}
