package domain


type TaobaoXhotelBnbpromoAddLongOrderInfo struct {
    /*
        互动折扣     */
    InvestmentNumber  *int64 `json:"investment_number,omitempty" `

    /*
        最小连住天数     */
    MinContinuityStay  *int64 `json:"min_continuity_stay,omitempty" `

}

func (s *TaobaoXhotelBnbpromoAddLongOrderInfo) SetInvestmentNumber(v int64) *TaobaoXhotelBnbpromoAddLongOrderInfo {
    s.InvestmentNumber = &v
    return s
}
func (s *TaobaoXhotelBnbpromoAddLongOrderInfo) SetMinContinuityStay(v int64) *TaobaoXhotelBnbpromoAddLongOrderInfo {
    s.MinContinuityStay = &v
    return s
}
