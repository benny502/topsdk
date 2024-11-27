package domain


type TaobaoXhotelOrderSearchXOrderGuest struct {
    /*
        入住人姓名     */
    Name  *string `json:"name,omitempty" `

    /*
        房间序号     */
    RoomPos  *int64 `json:"room_pos,omitempty" `

    /*
        入住人序号     */
    PersonPos  *int64 `json:"person_pos,omitempty" `

}

func (s *TaobaoXhotelOrderSearchXOrderGuest) SetName(v string) *TaobaoXhotelOrderSearchXOrderGuest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelOrderSearchXOrderGuest) SetRoomPos(v int64) *TaobaoXhotelOrderSearchXOrderGuest {
    s.RoomPos = &v
    return s
}
func (s *TaobaoXhotelOrderSearchXOrderGuest) SetPersonPos(v int64) *TaobaoXhotelOrderSearchXOrderGuest {
    s.PersonPos = &v
    return s
}
