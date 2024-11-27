package domain


type TaobaoXhotelPromotionResourceBindBindFailItem struct {
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

func (s *TaobaoXhotelPromotionResourceBindBindFailItem) SetRoomCode(v string) *TaobaoXhotelPromotionResourceBindBindFailItem {
    s.RoomCode = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindFailItem) SetRatePlanCode(v string) *TaobaoXhotelPromotionResourceBindBindFailItem {
    s.RatePlanCode = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindFailItem) SetErrorCode(v string) *TaobaoXhotelPromotionResourceBindBindFailItem {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoXhotelPromotionResourceBindBindFailItem) SetErrorMsg(v string) *TaobaoXhotelPromotionResourceBindBindFailItem {
    s.ErrorMsg = &v
    return s
}
