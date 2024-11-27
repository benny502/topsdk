package domain


type TaobaoXhotelBaseinfoGetRoomType struct {
    /*
        房型名称     */
    Name  *string `json:"name,omitempty" `

    /*
        阿里房型id     */
    Rid  *int64 `json:"rid,omitempty" `

    /*
        房型状态。0:正常，-1:删除，-2:停售     */
    Status  *byte `json:"status,omitempty" `

    /*
        系统商，一般不填写，使用须申请     */
    Vendor  *string `json:"vendor,omitempty" `

    /*
        商家房型ID     */
    OuterId  *string `json:"outer_id,omitempty" `

}

func (s *TaobaoXhotelBaseinfoGetRoomType) SetName(v string) *TaobaoXhotelBaseinfoGetRoomType {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRoomType) SetRid(v int64) *TaobaoXhotelBaseinfoGetRoomType {
    s.Rid = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRoomType) SetStatus(v byte) *TaobaoXhotelBaseinfoGetRoomType {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRoomType) SetVendor(v string) *TaobaoXhotelBaseinfoGetRoomType {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRoomType) SetOuterId(v string) *TaobaoXhotelBaseinfoGetRoomType {
    s.OuterId = &v
    return s
}
