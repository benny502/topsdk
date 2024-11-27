package domain


type TaobaoXhotelPromotionResourceUnbindUnBindSuccItem struct {
    /*
        供应商房型code     */
    RoomCode  *string `json:"room_code,omitempty" `

    /*
        供应商rate plan     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" `

}

func (s *TaobaoXhotelPromotionResourceUnbindUnBindSuccItem) SetRoomCode(v string) *TaobaoXhotelPromotionResourceUnbindUnBindSuccItem {
    s.RoomCode = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindUnBindSuccItem) SetRatePlanCode(v string) *TaobaoXhotelPromotionResourceUnbindUnBindSuccItem {
    s.RatePlanCode = &v
    return s
}
