package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelRoomtypeGetXRoomType struct {
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
    SRoomtype  *TaobaoXhotelRoomtypeGetSRoomType `json:"s_roomtype,omitempty" `

    /*
        酒店数据状态：匹配成功；待匹配；待审核；审核失败；疑似错误；请注意：只有状态为&ldquo;匹配成功&rdquo;才是正常状态。其他状态都不会上架商品。     */
    DataConfirmStr  *string `json:"data_confirm_str,omitempty" `

    /*
        卖家房型英文名称     */
    NameE  *string `json:"name_e,omitempty" `

    /*
        房型维度下特殊标签含义 json: {"non-direct-roomType":1} , key:non-direct-roomType 表示非直连房型     */
    TagJson  *string `json:"tag_json,omitempty" `

}

func (s *TaobaoXhotelRoomtypeGetXRoomType) SetRid(v int64) *TaobaoXhotelRoomtypeGetXRoomType {
    s.Rid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetHid(v int64) *TaobaoXhotelRoomtypeGetXRoomType {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetCreatedTime(v util.LocalTime) *TaobaoXhotelRoomtypeGetXRoomType {
    s.CreatedTime = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetModifiedTime(v util.LocalTime) *TaobaoXhotelRoomtypeGetXRoomType {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetMatchStatus(v int64) *TaobaoXhotelRoomtypeGetXRoomType {
    s.MatchStatus = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetStatus(v int64) *TaobaoXhotelRoomtypeGetXRoomType {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetErrorInfo(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.ErrorInfo = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetOuterId(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetName(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetMaxOccupancy(v int64) *TaobaoXhotelRoomtypeGetXRoomType {
    s.MaxOccupancy = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetArea(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.Area = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetFloor(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.Floor = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetBedType(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.BedType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetBedSize(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.BedSize = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetInternet(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.Internet = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetService(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetWindowType(v int64) *TaobaoXhotelRoomtypeGetXRoomType {
    s.WindowType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetExtend(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetSRoomtype(v TaobaoXhotelRoomtypeGetSRoomType) *TaobaoXhotelRoomtypeGetXRoomType {
    s.SRoomtype = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetDataConfirmStr(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.DataConfirmStr = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetNameE(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.NameE = &v
    return s
}
func (s *TaobaoXhotelRoomtypeGetXRoomType) SetTagJson(v string) *TaobaoXhotelRoomtypeGetXRoomType {
    s.TagJson = &v
    return s
}
