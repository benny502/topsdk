package domain


type TaobaoXhotelPromotionQueryPromotionResource struct {
    /*
        房型信息     */
    Rooms  *[]TaobaoXhotelPromotionQueryResourceRoom `json:"rooms,omitempty" `

    /*
        供应商酒店code     */
    HotelCode  *string `json:"hotel_code,omitempty" `

}

func (s *TaobaoXhotelPromotionQueryPromotionResource) SetRooms(v []TaobaoXhotelPromotionQueryResourceRoom) *TaobaoXhotelPromotionQueryPromotionResource {
    s.Rooms = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryPromotionResource) SetHotelCode(v string) *TaobaoXhotelPromotionQueryPromotionResource {
    s.HotelCode = &v
    return s
}
