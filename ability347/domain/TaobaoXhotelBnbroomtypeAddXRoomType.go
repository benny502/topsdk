package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelBnbroomtypeAddXRoomType struct {
    /*
        飞猪门店id     */
    Hid  *int64 `json:"hid,omitempty" `

    /*
        飞猪房源id     */
    Rid  *int64 `json:"rid,omitempty" `

    /*
        房源状态     */
    Status  *int64 `json:"status,omitempty" `

    /*
        创建时间     */
    GmtCreate  *util.LocalTime `json:"gmt_create,omitempty" `

    /*
        修改时间     */
    GmtModified  *util.LocalTime `json:"gmt_modified,omitempty" `

    /*
        匹配状态: 0：待系统匹配 1：已系统匹配，匹配成功，待卖家确认 2：已系统匹配，匹配失败，待人工匹配 3：已人工匹配，匹配成功，待卖家确认 4：已人工匹配，匹配失败 5：卖家已确认，确认&ldquo;YES&rdquo; 6：卖家已确认，确认&ldquo;NO&rdquo; 7:已系统匹配，但是匹配重复，待人工确认     */
    MatchStatus  *int64 `json:"match_status,omitempty" `

    /*
        卖家系统id     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        房源名称     */
    Name  *string `json:"name,omitempty" `

    /*
        最大入住人数     */
    MaxOccupancy  *int64 `json:"max_occupancy,omitempty" `

    /*
        面积     */
    Area  *string `json:"area,omitempty" `

    /*
        暂时不对外     */
    Extend  *string `json:"extend,omitempty" `

    /*
        窗型,0：无窗/1：有窗     */
    WindowType  *int64 `json:"window_type,omitempty" `

    /*
        设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {&quot;bar&quot;:false,&quot;catv&quot;:false,&quot;ddd&quot;:false,&quot;idd&quot;:false,&quot;pubtoilet&quot;:false,&quot;toilet&quot;:false}     */
    Service  *string `json:"service,omitempty" `

    /*
        宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带     */
    Internet  *string `json:"internet,omitempty" `

    /*
        客房在建筑的第几层，隔层为1-2层，4-5层，7-8层     */
    Floor  *string `json:"floor,omitempty" `

}

func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetHid(v int64) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetRid(v int64) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.Rid = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetStatus(v int64) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetGmtCreate(v util.LocalTime) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetGmtModified(v util.LocalTime) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.GmtModified = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetMatchStatus(v int64) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.MatchStatus = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetOuterId(v string) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetName(v string) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetMaxOccupancy(v int64) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.MaxOccupancy = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetArea(v string) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.Area = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetExtend(v string) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetWindowType(v int64) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.WindowType = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetService(v string) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetInternet(v string) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.Internet = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddXRoomType) SetFloor(v string) *TaobaoXhotelBnbroomtypeAddXRoomType {
    s.Floor = &v
    return s
}
