package domain


type TaobaoXhotelBaseinfoGetSroomtypelist struct {
    /*
        窗型     */
    WindowType  *string `json:"window_type,omitempty" `

    /*
        酒店图片结构化信息     */
    PicsInfo  *string `json:"pics_info,omitempty" `

    /*
        宽带服务     */
    NetworkService  *string `json:"network_service,omitempty" `

    /*
        最后变更人     */
    LastModify  *string `json:"last_modify,omitempty" `

    /*
        includeTypes     */
    IncludeTypes  *string `json:"include_types,omitempty" `

    /*
        创建人     */
    Auditor  *string `json:"auditor,omitempty" `

    /*
        标准酒店ID     */
    Shid  *int64 `json:"shid,omitempty" `

    /*
        房型原始图片     */
    OriginalPics  *string `json:"original_pics,omitempty" `

    /*
        版本号     */
    Version  *int64 `json:"version,omitempty" `

    /*
        标准房型ID     */
    Srid  *int64 `json:"srid,omitempty" `

    /*
        房型图片     */
    Pics  *string `json:"pics,omitempty" `

    /*
        图片扩展字段     */
    PicsExt  *string `json:"pics_ext,omitempty" `

    /*
        面积     */
    Area  *string `json:"area,omitempty" `

    /*
        设施     */
    Facility  *string `json:"facility,omitempty" `

    /*
        标准房型名称     */
    Name  *string `json:"name,omitempty" `

    /*
        propertiesDOs     */
    PropertiesDOs  *string `json:"properties_d_os,omitempty" `

    /*
        是否可加床     */
    AddBed  *int64 `json:"add_bed,omitempty" `

    /*
        标准房型的英文名     */
    NameE  *string `json:"name_e,omitempty" `

    /*
        状态     */
    Status  *int64 `json:"status,omitempty" `

    /*
        扩展信息     */
    Extend  *string `json:"extend,omitempty" `

    /*
        床型     */
    Bed  *string `json:"bed,omitempty" `

    /*
        transferPics     */
    TransferPics  *string `json:"transfer_pics,omitempty" `

    /*
        楼层     */
    Floor  *string `json:"floor,omitempty" `

    /*
        来源     */
    Source  *int64 `json:"source,omitempty" `

    /*
        bedList     */
    BedList  *string `json:"bed_list,omitempty" `

    /*
        最大入住人数     */
    MaxOccupancy  *int64 `json:"max_occupancy,omitempty" `

    /*
        外部id     */
    OuterId  *string `json:"outer_id,omitempty" `

}

func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetWindowType(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.WindowType = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetPicsInfo(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.PicsInfo = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetNetworkService(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.NetworkService = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetLastModify(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.LastModify = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetIncludeTypes(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.IncludeTypes = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetAuditor(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Auditor = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetShid(v int64) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Shid = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetOriginalPics(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.OriginalPics = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetVersion(v int64) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Version = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetSrid(v int64) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Srid = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetPics(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Pics = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetPicsExt(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.PicsExt = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetArea(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Area = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetFacility(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Facility = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetName(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetPropertiesDOs(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.PropertiesDOs = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetAddBed(v int64) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.AddBed = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetNameE(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.NameE = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetStatus(v int64) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetExtend(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetBed(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Bed = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetTransferPics(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.TransferPics = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetFloor(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Floor = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetSource(v int64) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.Source = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetBedList(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.BedList = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetMaxOccupancy(v int64) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.MaxOccupancy = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetSroomtypelist) SetOuterId(v string) *TaobaoXhotelBaseinfoGetSroomtypelist {
    s.OuterId = &v
    return s
}
