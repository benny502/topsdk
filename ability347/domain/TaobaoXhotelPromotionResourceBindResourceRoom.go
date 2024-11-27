package domain


type TaobaoXhotelPromotionResourceBindResourceRoom struct {
    /*
        供应商rpcodes     */
    RatePlanCodes  *[]string `json:"rate_plan_codes,omitempty" `

    /*
        供应商房型code     */
    RoomCode  *string `json:"room_code,omitempty" `

}

func (s *TaobaoXhotelPromotionResourceBindResourceRoom) SetRatePlanCodes(v []string) *TaobaoXhotelPromotionResourceBindResourceRoom {
    s.RatePlanCodes = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindResourceRoom) SetRoomCode(v string) *TaobaoXhotelPromotionResourceBindResourceRoom {
    s.RoomCode = &v
    return s
}
