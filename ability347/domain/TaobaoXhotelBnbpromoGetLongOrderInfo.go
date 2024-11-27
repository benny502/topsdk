package domain


type TaobaoXhotelBnbpromoGetLongOrderInfo struct {
    /*
        连住天数可选择范围为（2，3，4，5，7，15，30）     */
    MinContinuityStay  *int64 `json:"min_continuity_stay,omitempty" `

    /*
        固定折扣百分比     */
    InvestmentNumber  *int64 `json:"investment_number,omitempty" `

}

func (s *TaobaoXhotelBnbpromoGetLongOrderInfo) SetMinContinuityStay(v int64) *TaobaoXhotelBnbpromoGetLongOrderInfo {
    s.MinContinuityStay = &v
    return s
}
func (s *TaobaoXhotelBnbpromoGetLongOrderInfo) SetInvestmentNumber(v int64) *TaobaoXhotelBnbpromoGetLongOrderInfo {
    s.InvestmentNumber = &v
    return s
}
