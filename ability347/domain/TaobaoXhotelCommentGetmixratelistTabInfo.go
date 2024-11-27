package domain


type TaobaoXhotelCommentGetmixratelistTabInfo struct {
    /*
        标签的态度 1好 0中性 -1 差     */
    Attitude  *int64 `json:"attitude,omitempty" `

    /*
        tab是否点击     */
    IsClick  *bool `json:"is_click,omitempty" `

    /*
        tab编码     */
    TabCode  *string `json:"tab_code,omitempty" `

    /*
        tab下的统计数据     */
    TabDetail  *string `json:"tab_detail,omitempty" `

    /*
        id     */
    TabId  *int64 `json:"tab_id,omitempty" `

    /*
        名称     */
    TabName  *string `json:"tab_name,omitempty" `

    /*
        标签的类型 0正常 1热词 2房型 3度假     */
    Type  *int64 `json:"type,omitempty" `

}

func (s *TaobaoXhotelCommentGetmixratelistTabInfo) SetAttitude(v int64) *TaobaoXhotelCommentGetmixratelistTabInfo {
    s.Attitude = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistTabInfo) SetIsClick(v bool) *TaobaoXhotelCommentGetmixratelistTabInfo {
    s.IsClick = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistTabInfo) SetTabCode(v string) *TaobaoXhotelCommentGetmixratelistTabInfo {
    s.TabCode = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistTabInfo) SetTabDetail(v string) *TaobaoXhotelCommentGetmixratelistTabInfo {
    s.TabDetail = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistTabInfo) SetTabId(v int64) *TaobaoXhotelCommentGetmixratelistTabInfo {
    s.TabId = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistTabInfo) SetTabName(v string) *TaobaoXhotelCommentGetmixratelistTabInfo {
    s.TabName = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistTabInfo) SetType(v int64) *TaobaoXhotelCommentGetmixratelistTabInfo {
    s.Type = &v
    return s
}
