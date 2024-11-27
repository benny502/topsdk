package domain


type TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam struct {
    /*
        业务类型     */
    BizType  *int64 `json:"biz_type,omitempty" `

    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        加载点评     */
    LoadReply  *int64 `json:"load_reply,omitempty" `

    /*
        页码     */
    PageNo  *int64 `json:"page_no,omitempty" `

    /*
        分页大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        排序     */
    Sort  *int64 `json:"sort,omitempty" `

    /*
        tab过滤     */
    TabFilter  *string `json:"tab_filter,omitempty" `

}

func (s *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam) SetBizType(v int64) *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam {
    s.BizType = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam) SetItemId(v int64) *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam {
    s.ItemId = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam) SetLoadReply(v int64) *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam {
    s.LoadReply = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam) SetPageNo(v int64) *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam {
    s.PageNo = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam) SetPageSize(v int64) *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam {
    s.PageSize = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam) SetSort(v int64) *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam {
    s.Sort = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam) SetTabFilter(v string) *TaobaoXhotelCommentGetmixratelistTopGetMixRateListParam {
    s.TabFilter = &v
    return s
}
