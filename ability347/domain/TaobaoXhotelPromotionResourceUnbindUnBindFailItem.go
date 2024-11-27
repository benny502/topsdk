package domain


type TaobaoXhotelPromotionResourceUnbindUnBindFailItem struct {
    /*
        供应商房型code     */
    RoomCode  *string `json:"room_code,omitempty" `

    /*
        供应商rate plan     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误描述     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoXhotelPromotionResourceUnbindUnBindFailItem) SetRoomCode(v string) *TaobaoXhotelPromotionResourceUnbindUnBindFailItem {
    s.RoomCode = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindUnBindFailItem) SetRatePlanCode(v string) *TaobaoXhotelPromotionResourceUnbindUnBindFailItem {
    s.RatePlanCode = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindUnBindFailItem) SetErrorCode(v string) *TaobaoXhotelPromotionResourceUnbindUnBindFailItem {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceUnbindUnBindFailItem) SetErrorMsg(v string) *TaobaoXhotelPromotionResourceUnbindUnBindFailItem {
    s.ErrorMsg = &v
    return s
}
