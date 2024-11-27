package domain


type TaobaoXhotelCommentGetmixratelistItemStatisticVo struct {
    /*
        最佳得分项     */
    BestItem  *string `json:"best_item,omitempty" `

    /*
        五分制标记, 1为五分制， 0为十分制     */
    IsFiveGrade  *int64 `json:"is_five_grade,omitempty" `

    /*
        评论总数     */
    RateCnt  *int64 `json:"rate_cnt,omitempty" `

    /*
        有图的评论总数     */
    RatePicCnt  *int64 `json:"rate_pic_cnt,omitempty" `

    /*
        推荐率     */
    RecommendStr  *string `json:"recommend_str,omitempty" `

    /*
        回复总数     */
    ReplyCnt  *int64 `json:"reply_cnt,omitempty" `

    /*
        tab信息     */
    RoomTabInfos  *[]TaobaoXhotelCommentGetmixratelistTabInfo `json:"room_tab_infos,omitempty" `

    /*
        评分描述： 非常好     */
    ScoreDesc  *string `json:"score_desc,omitempty" `

    /*
        评分详情，json格式     */
    ScoreDetail  *string `json:"score_detail,omitempty" `

    /*
        不同分数的数量     */
    ScoreInfos  *[]TaobaoXhotelCommentGetmixratelistScoreInfo `json:"score_infos,omitempty" `

    /*
        评分星级     */
    ScoreLevel  *int64 `json:"score_level,omitempty" `

    /*
        source来源 0自采 1共享 21agoda 22艺龙 23tripAdvisor     */
    Source  *int64 `json:"source,omitempty" `

    /*
        tab信息     */
    SubItemInfos  *[]TaobaoXhotelCommentGetmixratelistTabInfo `json:"sub_item_infos,omitempty" `

    /*
        tab信息     */
    TabInfoS  *[]TaobaoXhotelCommentGetmixratelistTabInfo `json:"tab_info_s,omitempty" `

    /*
        热词显示的行数     */
    TabShowLines  *int64 `json:"tab_show_lines,omitempty" `

    /*
        总评分     */
    TotalScore  *string `json:"total_score,omitempty" `

    /*
        酒店类目是标准酒店ID     */
    TravelItemId  *int64 `json:"travel_item_id,omitempty" `

    /*
        旅游商品信息     */
    TravelItemInfo  *string `json:"travel_item_info,omitempty" `

    /*
        tripAdv评论数     */
    TripAdvateCnt  *int64 `json:"trip_advate_cnt,omitempty" `

}

func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetBestItem(v string) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.BestItem = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetIsFiveGrade(v int64) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.IsFiveGrade = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetRateCnt(v int64) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.RateCnt = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetRatePicCnt(v int64) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.RatePicCnt = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetRecommendStr(v string) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.RecommendStr = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetReplyCnt(v int64) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.ReplyCnt = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetRoomTabInfos(v []TaobaoXhotelCommentGetmixratelistTabInfo) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.RoomTabInfos = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetScoreDesc(v string) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.ScoreDesc = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetScoreDetail(v string) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.ScoreDetail = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetScoreInfos(v []TaobaoXhotelCommentGetmixratelistScoreInfo) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.ScoreInfos = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetScoreLevel(v int64) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.ScoreLevel = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetSource(v int64) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.Source = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetSubItemInfos(v []TaobaoXhotelCommentGetmixratelistTabInfo) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.SubItemInfos = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetTabInfoS(v []TaobaoXhotelCommentGetmixratelistTabInfo) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.TabInfoS = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetTabShowLines(v int64) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.TabShowLines = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetTotalScore(v string) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.TotalScore = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetTravelItemId(v int64) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.TravelItemId = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetTravelItemInfo(v string) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.TravelItemInfo = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistItemStatisticVo) SetTripAdvateCnt(v int64) *TaobaoXhotelCommentGetmixratelistItemStatisticVo {
    s.TripAdvateCnt = &v
    return s
}
