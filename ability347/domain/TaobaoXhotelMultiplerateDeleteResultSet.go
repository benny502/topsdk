package domain


type TaobaoXhotelMultiplerateDeleteResultSet struct {
    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        房型编码_房价编码_入住人数_连住天数     */
    OutRidRateplanCodeOccupancyLengthofstay  *string `json:"out_rid_rateplan_code_occupancy_lengthofstay,omitempty" `

}

func (s *TaobaoXhotelMultiplerateDeleteResultSet) SetErrorCode(v string) *TaobaoXhotelMultiplerateDeleteResultSet {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelMultiplerateDeleteResultSet) SetErrorMsg(v string) *TaobaoXhotelMultiplerateDeleteResultSet {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoXhotelMultiplerateDeleteResultSet) SetOutRidRateplanCodeOccupancyLengthofstay(v string) *TaobaoXhotelMultiplerateDeleteResultSet {
    s.OutRidRateplanCodeOccupancyLengthofstay = &v
    return s
}
