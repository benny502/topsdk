package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelRoomtypeUpdateXRoomType struct {
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
        窗型,0：无窗/1：有窗     */
    WindowType  *int64 `json:"window_type,omitempty" `

    /*
        设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {&quot;bar&quot;:false,&quot;catv&quot;:false,&quot;ddd&quot;:false,&quot;idd&quot;:false,&quot;pubtoilet&quot;:false,&quot;toilet&quot;:false}     */
    Service  *string `json:"service,omitempty" `

    /*
        扩展信息的JSON。 注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析     */
    Extend  *string `json:"extend,omitempty" `

    /*
        标准房型信息     */
    SRoomtype  *TaobaoXhotelRoomtypeUpdateSRoomType `json:"s_roomtype,omitempty" `

    /*
        卖家房型英文名称     */
    NameE  *string `json:"name_e,omitempty" `

}

func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetRid(v int64) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.Rid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetHid(v int64) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetCreatedTime(v util.LocalTime) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.CreatedTime = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetModifiedTime(v util.LocalTime) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetMatchStatus(v int64) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.MatchStatus = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetStatus(v int64) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetErrorInfo(v string) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.ErrorInfo = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetOuterId(v string) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetName(v string) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetMaxOccupancy(v int64) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.MaxOccupancy = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetArea(v string) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.Area = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetFloor(v string) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.Floor = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetBedType(v string) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.BedType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetBedSize(v string) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.BedSize = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetInternet(v string) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.Internet = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetWindowType(v int64) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.WindowType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetService(v string) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetExtend(v string) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetSRoomtype(v TaobaoXhotelRoomtypeUpdateSRoomType) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.SRoomtype = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateXRoomType) SetNameE(v string) *TaobaoXhotelRoomtypeUpdateXRoomType {
    s.NameE = &v
    return s
}
