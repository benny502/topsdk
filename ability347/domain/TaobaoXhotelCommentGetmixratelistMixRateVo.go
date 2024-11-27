package domain


type TaobaoXhotelCommentGetmixratelistMixRateVo struct {
    /*
        顶数量     */
    AgreeCount  *int64 `json:"agree_count,omitempty" `

    /*
        业务类型     */
    BizType  *int64 `json:"biz_type,omitempty" `

    /*
        内容     */
    Content  *string `json:"content,omitempty" `

    /*
        默认icon     */
    DefaultIcon  *int64 `json:"default_icon,omitempty" `

    /*
        差评数     */
    DisagreeCount  *int64 `json:"disagree_count,omitempty" `

    /*
        创建时间     */
    GmtCreate  *string `json:"gmt_create,omitempty" `

    /*
        id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        商品信息     */
    ItemInfo  *string `json:"item_info,omitempty" `

    /*
        商品评论ID     */
    ItemRateId  *int64 `json:"item_rate_id,omitempty" `

    /*
        所有商品评论回复     */
    ItemReplies  *[]TaobaoXhotelCommentGetmixratelistItemRateReplyVo `json:"item_replies,omitempty" `

    /*
        点赞数据     */
    Like  *TaobaoXhotelCommentGetmixratelistLikeTargetCount `json:"like,omitempty" `

    /*
        图片信息,图片URL的list     */
    MediaInfo  *string `json:"media_info,omitempty" `

    /*
        订单ID     */
    OrderId  *int64 `json:"order_id,omitempty" `

    /*
        订单ID     */
    OrderIdStr  *string `json:"order_id_str,omitempty" `

    /*
        交易信息     */
    OrderInfo  *string `json:"order_info,omitempty" `

    /*
        当前用户的评价 0:不是, 1:是     */
    Owner  *int64 `json:"owner,omitempty" `

    /*
        图片URL     */
    PictureUrls  *[]string `json:"picture_urls,omitempty" `

    /*
        POI固定文本     */
    PoiStr  *string `json:"poi_str,omitempty" `

    /*
        跳转URL     */
    RedirectUrl  *string `json:"redirect_url,omitempty" `

    /*
        回复数量     */
    ReplyCount  *int64 `json:"reply_count,omitempty" `

    /*
        评分详情     */
    ScoreDetail  *string `json:"score_detail,omitempty" `

    /*
        评分详情     */
    SellerScoreDetail  *string `json:"seller_score_detail,omitempty" `

    /*
        SKU     */
    Sku  *string `json:"sku,omitempty" `

    /*
        点评来源     */
    Source  *int64 `json:"source,omitempty" `

    /*
        点评来源名称     */
    SourceTypeName  *string `json:"source_type_name,omitempty" `

    /*
        分割线内容     */
    SplitLineContent  *string `json:"split_line_content,omitempty" `

    /*
        置顶状态 0 普通，1置顶     */
    Status  *int64 `json:"status,omitempty" `

    /*
        1     */
    TabInfos  *[]TaobaoXhotelCommentGetmixratelistTabInfo `json:"tab_infos,omitempty" `

    /*
        标签信息     */
    TagInfo  *string `json:"tag_info,omitempty" `

    /*
        评论title,留作备用     */
    Title  *string `json:"title,omitempty" `

    /*
        置顶状态 0 普通，1置顶     */
    TopStatus  *int64 `json:"top_status,omitempty" `

    /*
        置顶图标     */
    TopStatusIcon  *string `json:"top_status_icon,omitempty" `

    /*
        置顶图标     */
    TotalScore  *int64 `json:"total_score,omitempty" `

    /*
        总评分     */
    TravelName  *string `json:"travel_name,omitempty" `

    /*
        航旅标准商品子类型id     */
    TravelSubItemId  *int64 `json:"travel_sub_item_id,omitempty" `

    /*
        航旅标准商品子类型信息     */
    TravelSubItemInfo  *string `json:"travel_sub_item_info,omitempty" `

    /*
        行程ID     */
    TripGuidId  *int64 `json:"trip_guid_id,omitempty" `

    /*
        冗余     */
    Ttid  *string `json:"ttid,omitempty" `

    /*
        用户昵称,已脱敏     */
    UserNick  *string `json:"user_nick,omitempty" `

}

func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetAgreeCount(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.AgreeCount = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetBizType(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.BizType = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetContent(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.Content = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetDefaultIcon(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.DefaultIcon = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetDisagreeCount(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.DisagreeCount = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetGmtCreate(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetId(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.Id = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetItemId(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.ItemId = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetItemInfo(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.ItemInfo = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetItemRateId(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.ItemRateId = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetItemReplies(v []TaobaoXhotelCommentGetmixratelistItemRateReplyVo) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.ItemReplies = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetLike(v TaobaoXhotelCommentGetmixratelistLikeTargetCount) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.Like = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetMediaInfo(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.MediaInfo = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetOrderId(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.OrderId = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetOrderIdStr(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.OrderIdStr = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetOrderInfo(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.OrderInfo = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetOwner(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.Owner = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetPictureUrls(v []string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.PictureUrls = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetPoiStr(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.PoiStr = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetRedirectUrl(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.RedirectUrl = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetReplyCount(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.ReplyCount = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetScoreDetail(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.ScoreDetail = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetSellerScoreDetail(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.SellerScoreDetail = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetSku(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.Sku = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetSource(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.Source = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetSourceTypeName(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.SourceTypeName = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetSplitLineContent(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.SplitLineContent = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetStatus(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetTabInfos(v []TaobaoXhotelCommentGetmixratelistTabInfo) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.TabInfos = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetTagInfo(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.TagInfo = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetTitle(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.Title = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetTopStatus(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.TopStatus = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetTopStatusIcon(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.TopStatusIcon = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetTotalScore(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.TotalScore = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetTravelName(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.TravelName = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetTravelSubItemId(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.TravelSubItemId = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetTravelSubItemInfo(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.TravelSubItemInfo = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetTripGuidId(v int64) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.TripGuidId = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetTtid(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.Ttid = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistMixRateVo) SetUserNick(v string) *TaobaoXhotelCommentGetmixratelistMixRateVo {
    s.UserNick = &v
    return s
}
