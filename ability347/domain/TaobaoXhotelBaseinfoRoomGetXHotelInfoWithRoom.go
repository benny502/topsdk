package domain


type TaobaoXhotelBaseinfoRoomGetXHotelInfoWithRoom struct {
    /*
        房型基础信息     */
    RoomTypeList  *[]TaobaoXhotelBaseinfoRoomGetRoomType `json:"room_type_list,omitempty" `

}

func (s *TaobaoXhotelBaseinfoRoomGetXHotelInfoWithRoom) SetRoomTypeList(v []TaobaoXhotelBaseinfoRoomGetRoomType) *TaobaoXhotelBaseinfoRoomGetXHotelInfoWithRoom {
    s.RoomTypeList = &v
    return s
}
