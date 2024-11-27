package domain


type TaobaoXhotelPromotionQueryResourceRoom struct {
    /*
        供应商rpcodes     */
    RatePlanCodes  *[]string `json:"rate_plan_codes,omitempty" `

    /*
        供应商房型code     */
    RoomCode  *string `json:"room_code,omitempty" `

}

func (s *TaobaoXhotelPromotionQueryResourceRoom) SetRatePlanCodes(v []string) *TaobaoXhotelPromotionQueryResourceRoom {
    s.RatePlanCodes = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryResourceRoom) SetRoomCode(v string) *TaobaoXhotelPromotionQueryResourceRoom {
    s.RoomCode = &v
    return s
}
