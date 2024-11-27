package domain


type TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO struct {
    /*
        床型数据检查结果，false代表有冲突     */
    Bedtypecheckresult  *string `json:"bedtypecheckresult,omitempty" `

    /*
        结构化床型     */
    BedType  *string `json:"bed_type,omitempty" `

    /*
        原始床型     */
    OriginBedType  *string `json:"origin_bed_type,omitempty" `

    /*
        房型英文名称     */
    Roomtypeenname  *string `json:"roomtypeenname,omitempty" `

    /*
        房型名称     */
    Roomtypename  *string `json:"roomtypename,omitempty" `

    /*
        房型编码     */
    OuterRoomTypeId  *string `json:"outer_room_type_id,omitempty" `

    /*
        酒店名称     */
    Hotelname  *string `json:"hotelname,omitempty" `

    /*
        酒店编码     */
    OuterHotelId  *string `json:"outer_hotel_id,omitempty" `

    /*
        冲突数据创建时间（供商家参考，如果时间是过去三天以外的，那么可以不用关注）     */
    GmtCreate  *string `json:"gmt_create,omitempty" `

    /*
        床型英文描叙     */
    Ennamebedtypedesc  *string `json:"ennamebedtypedesc,omitempty" `

    /*
        床型描叙     */
    Namebedtypedesc  *string `json:"namebedtypedesc,omitempty" `

    /*
        床型描叙     */
    Bedtypedesc  *string `json:"bedtypedesc,omitempty" `

    /*
        床型检查信息     */
    Bedtypecheckmsg  *string `json:"bedtypecheckmsg,omitempty" `

}

func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetBedtypecheckresult(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.Bedtypecheckresult = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetBedType(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.BedType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetOriginBedType(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.OriginBedType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetRoomtypeenname(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.Roomtypeenname = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetRoomtypename(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.Roomtypename = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetOuterRoomTypeId(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.OuterRoomTypeId = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetHotelname(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.Hotelname = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetOuterHotelId(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.OuterHotelId = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetGmtCreate(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetEnnamebedtypedesc(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.Ennamebedtypedesc = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetNamebedtypedesc(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.Namebedtypedesc = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetBedtypedesc(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.Bedtypedesc = &v
    return s
}
func (s *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO) SetBedtypecheckmsg(v string) *TaobaoXhotelRoomtypeConflictDataRoomTypeCheckResultDO {
    s.Bedtypecheckmsg = &v
    return s
}
