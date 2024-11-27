package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelRoomtypeAddXRoomType struct {
    /*
        rid     */
    Rid  *int64 `json:"rid,omitempty" `

    /*
        hid     */
    Hid  *int64 `json:"hid,omitempty" `

    /*
        创建时间     */
    CreatedTime  *util.LocalTime `json:"created_time,omitempty" `

    /*
        修改时间     */
    ModifiedTime  *util.LocalTime `json:"modified_time,omitempty" `

    /*
        此字段已废弃     */
    MatchStatus  *int64 `json:"match_status,omitempty" `

    /*
        房型状态。0:正常，-1:删除，-2:停售     */
    Status  *int64 `json:"status,omitempty" `

    /*
        出错原因,没有匹配上标准房型时，小二拒绝的理由     */
    ErrorInfo  *string `json:"error_info,omitempty" `

    /*
        卖家系统id     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        房型名称     */
    Name  *string `json:"name,omitempty" `

    /*
        最大入住人数     */
    MaxOccupancy  *int64 `json:"max_occupancy,omitempty" `

    /*
        可选值：A,B,C,D。分别代表： A：15平米以下，B：16－30平米，C：31－50平米，D：50平米以上 2）也可以自己定义，比如：40平方米     */
    Area  *string `json:"area,omitempty" `

    /*
        客房在建筑的第几层，隔层为1-2层，4-5层，7-8层     */
    Floor  *string `json:"floor,omitempty" `

    /*
        床型。按自己定义存储，比如：高低床、上下床     */
    BedType  *string `json:"bed_type,omitempty" `

    /*
        床宽。     */
    BedSize  *string `json:"bed_size,omitempty" `

    /*
        宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带     */
    Internet  *string `json:"internet,omitempty" `

    /*
        设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {&quot;bar&quot;:false,&quot;catv&quot;:false,&quot;ddd&quot;:false,&quot;idd&quot;:false,&quot;pubtoilet&quot;:false,&quot;toilet&quot;:false}     */
    Service  *string `json:"service,omitempty" `

    /*
        窗型,0：无窗/1：有窗     */
    WindowType  *int64 `json:"window_type,omitempty" `

    /*
        扩展信息的JSON。 注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析     */
    Extend  *string `json:"extend,omitempty" `

    /*
        标准房型信息     */
    SRoomtype  *TaobaoXhotelRoomtypeAddSRoomType `json:"s_roomtype,omitempty" `

    /*
        卖家房型英文名称     */
    NameE  *string `json:"name_e,omitempty" `

}

func (s *TaobaoXhotelRoomtypeAddXRoomType) SetRid(v int64) *TaobaoXhotelRoomtypeAddXRoomType {
    s.Rid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetHid(v int64) *TaobaoXhotelRoomtypeAddXRoomType {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetCreatedTime(v util.LocalTime) *TaobaoXhotelRoomtypeAddXRoomType {
    s.CreatedTime = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetModifiedTime(v util.LocalTime) *TaobaoXhotelRoomtypeAddXRoomType {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetMatchStatus(v int64) *TaobaoXhotelRoomtypeAddXRoomType {
    s.MatchStatus = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetStatus(v int64) *TaobaoXhotelRoomtypeAddXRoomType {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetErrorInfo(v string) *TaobaoXhotelRoomtypeAddXRoomType {
    s.ErrorInfo = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetOuterId(v string) *TaobaoXhotelRoomtypeAddXRoomType {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetName(v string) *TaobaoXhotelRoomtypeAddXRoomType {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetMaxOccupancy(v int64) *TaobaoXhotelRoomtypeAddXRoomType {
    s.MaxOccupancy = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetArea(v string) *TaobaoXhotelRoomtypeAddXRoomType {
    s.Area = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetFloor(v string) *TaobaoXhotelRoomtypeAddXRoomType {
    s.Floor = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetBedType(v string) *TaobaoXhotelRoomtypeAddXRoomType {
    s.BedType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetBedSize(v string) *TaobaoXhotelRoomtypeAddXRoomType {
    s.BedSize = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetInternet(v string) *TaobaoXhotelRoomtypeAddXRoomType {
    s.Internet = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetService(v string) *TaobaoXhotelRoomtypeAddXRoomType {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetWindowType(v int64) *TaobaoXhotelRoomtypeAddXRoomType {
    s.WindowType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetExtend(v string) *TaobaoXhotelRoomtypeAddXRoomType {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetSRoomtype(v TaobaoXhotelRoomtypeAddSRoomType) *TaobaoXhotelRoomtypeAddXRoomType {
    s.SRoomtype = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddXRoomType) SetNameE(v string) *TaobaoXhotelRoomtypeAddXRoomType {
    s.NameE = &v
    return s
}
