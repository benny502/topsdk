package domain


type TaobaoXhotelPromotionResourceBindBindSuccItem struct {
    /*
        供应商房型code     */
    RoomCode  *string `json:"room_code,omitempty" `

    /*
        供应商rate plan     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" `

}

func (s *TaobaoXhotelPromotionResourceBindBindSuccItem) SetRoomCode(v string) *TaobaoXhotelPromotionResourceBindBindSuccItem {
    s.RoomCode = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindSuccItem) SetRatePlanCode(v string) *TaobaoXhotelPromotionResourceBindBindSuccItem {
    s.RatePlanCode = &v
    return s
}
