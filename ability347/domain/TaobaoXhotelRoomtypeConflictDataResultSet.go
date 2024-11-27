package domain


type TaobaoXhotelRoomtypeConflictDataResultSet struct {
    /*
        警告信息     */
    WarnMessage  *string `json:"warn_message,omitempty" `

    /*
        总数     */
    TotalResults  *int64 `json:"total_results,omitempty" `

    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        是否还有下一页     */
    HasNext  *bool `json:"has_next,omitempty" `

    /*
        结果集     */
    Results  *[]TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO `json:"results,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoXhotelRoomtypeConflictDataResultSet) SetWarnMessage(v string) *TaobaoXhotelRoomtypeConflictDataResultSet {
    s.WarnMessage = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataResultSet) SetTotalResults(v int64) *TaobaoXhotelRoomtypeConflictDataResultSet {
    s.TotalResults = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataResultSet) SetErrorCode(v string) *TaobaoXhotelRoomtypeConflictDataResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataResultSet) SetHasNext(v bool) *TaobaoXhotelRoomtypeConflictDataResultSet {
    s.HasNext = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataResultSet) SetResults(v []TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) *TaobaoXhotelRoomtypeConflictDataResultSet {
    s.Results = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataResultSet) SetErrorMsg(v string) *TaobaoXhotelRoomtypeConflictDataResultSet {
    s.ErrorMsg = &v
    return s
}
