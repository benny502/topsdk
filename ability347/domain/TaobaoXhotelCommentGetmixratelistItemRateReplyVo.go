package domain


type TaobaoXhotelCommentGetmixratelistItemRateReplyVo struct {
    /*
        业务类型     */
    BizType  *int64 `json:"biz_type,omitempty" `

    /*
        内容     */
    Content  *string `json:"content,omitempty" `

    /*
        创建日期     */
    GmtCreate  *string `json:"gmt_create,omitempty" `

    /*
        与主评间隔天数     */
    IntervalDay  *int64 `json:"interval_day,omitempty" `

    /*
        图片信息     */
    MediaInfo  *string `json:"media_info,omitempty" `

    /*
        当前用户的评价 0:不是, 1:是     */
    Owner  *int64 `json:"owner,omitempty" `

    /*
        回复的是那一条，如果是回复主评为0，否则为追评id，组成树形结构     */
    ParentId  *int64 `json:"parent_id,omitempty" `

    /*
        被回复人的冗余信息     */
    ParentInfo  *TaobaoXhotelCommentGetmixratelistParentInfo `json:"parent_info,omitempty" `

    /*
        该条回复的id     */
    ReplyId  *int64 `json:"reply_id,omitempty" `

    /*
        回复类型 0买家追评 1卖家回复     */
    ReplyType  *int64 `json:"reply_type,omitempty" `

    /*
        状态     */
    Status  *int64 `json:"status,omitempty" `

    /*
        酒店类目是标准酒店ID ，用来分库     */
    TravelItemId  *int64 `json:"travel_item_id,omitempty" `

    /*
        用户名称,已脱敏     */
    UserNick  *string `json:"user_nick,omitempty" `

}

func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetBizType(v int64) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.BizType = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetContent(v string) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.Content = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetGmtCreate(v string) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetIntervalDay(v int64) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.IntervalDay = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetMediaInfo(v string) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.MediaInfo = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetOwner(v int64) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.Owner = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetParentId(v int64) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.ParentId = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetParentInfo(v TaobaoXhotelCommentGetmixratelistParentInfo) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.ParentInfo = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetReplyId(v int64) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.ReplyId = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetReplyType(v int64) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.ReplyType = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetStatus(v int64) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetTravelItemId(v int64) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.TravelItemId = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemRateReplyVo) SetUserNick(v string) *TaobaoXhotelCommentGetmixratelistItemRateReplyVo {
    s.UserNick = &v
    return s
}
