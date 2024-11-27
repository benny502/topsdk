package domain


type TaobaoXhotelBnbpromoUpdateLongOrderInfo struct {
    /*
        最小连住天数     */
    MinContinuityStay  *int64 `json:"min_continuity_stay,omitempty" `

    /*
        折扣比例，填30就意味着原价的30%，也就是打3折。数字范围限定在10-95之间     */
    InvestmentNumber  *int64 `json:"investment_number,omitempty" `

}

func (s *TaobaoXhotelBnbpromoUpdateLongOrderInfo) SetMinContinuityStay(v int64) *TaobaoXhotelBnbpromoUpdateLongOrderInfo {
    s.MinContinuityStay = &v
    return s
}
func (s *TaobaoXhotelBnbpromoUpdateLongOrderInfo) SetInvestmentNumber(v int64) *TaobaoXhotelBnbpromoUpdateLongOrderInfo {
    s.InvestmentNumber = &v
    return s
}
