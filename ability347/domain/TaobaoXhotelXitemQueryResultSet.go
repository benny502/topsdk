package domain


type TaobaoXhotelXitemQueryResultSet struct {
    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        查询到的 x 元素     */
    XItems  *[]TaobaoXhotelXitemQueryHotelXitemDO `json:"x_items,omitempty" `

    /*
        记录总数     */
    TotalCount  *int64 `json:"total_count,omitempty" `

}

func (s *TaobaoXhotelXitemQueryResultSet) SetErrorMsg(v string) *TaobaoXhotelXitemQueryResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelXitemQueryResultSet) SetErrorCode(v string) *TaobaoXhotelXitemQueryResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelXitemQueryResultSet) SetXItems(v []TaobaoXhotelXitemQueryHotelXitemDO) *TaobaoXhotelXitemQueryResultSet {
    s.XItems = &v
    return s
}
func (s *TaobaoXhotelXitemQueryResultSet) SetTotalCount(v int64) *TaobaoXhotelXitemQueryResultSet {
    s.TotalCount = &v
    return s
}
