package domain


type TaobaoXhotelPromotionResourceUnbindResourceRoom struct {
    /*
        供应商rate plan     */
    RatePlanCodes  *[]string `json:"rate_plan_codes,omitempty" `

    /*
        供应商房型code     */
    RoomCode  *string `json:"room_code,omitempty" `

}

func (s *TaobaoXhotelPromotionResourceUnbindResourceRoom) SetRatePlanCodes(v []string) *TaobaoXhotelPromotionResourceUnbindResourceRoom {
    s.RatePlanCodes = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindResourceRoom) SetRoomCode(v string) *TaobaoXhotelPromotionResourceUnbindResourceRoom {
    s.RoomCode = &v
    return s
}
