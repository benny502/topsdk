package domain


type TaobaoXhotelCommentGetmixratelistAddRateEntrance struct {
    /*
        是否可以发表评价,0:不可以发表评价,1:可以发表评价.     */
    Available  *int64 `json:"available,omitempty" `

    /*
        写点评的跳转链接     */
    JumpInfo  *TaobaoXhotelCommentGetmixratelistJumpInfo `json:"jump_info,omitempty" `

    /*
        名称     */
    Name  *string `json:"name,omitempty" `

    /*
        是否展示入口 0:不展示,1:展示     */
    ShowStatus  *int64 `json:"show_status,omitempty" `

}

func (s *TaobaoXhotelCommentGetmixratelistAddRateEntrance) SetAvailable(v int64) *TaobaoXhotelCommentGetmixratelistAddRateEntrance {
    s.Available = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistAddRateEntrance) SetJumpInfo(v TaobaoXhotelCommentGetmixratelistJumpInfo) *TaobaoXhotelCommentGetmixratelistAddRateEntrance {
    s.JumpInfo = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistAddRateEntrance) SetName(v string) *TaobaoXhotelCommentGetmixratelistAddRateEntrance {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistAddRateEntrance) SetShowStatus(v int64) *TaobaoXhotelCommentGetmixratelistAddRateEntrance {
    s.ShowStatus = &v
    return s
}
