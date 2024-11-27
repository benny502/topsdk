package domain


type TaobaoXhotelCommentGetmixratelistGetMixRateListResult struct {
    /*
        页面布局信息     */
    ConfigInfo  *string `json:"config_info,omitempty" `

    /*
        错误代码     */
    ErrCode  *string `json:"err_code,omitempty" `

    /*
        错误详细描述     */
    ErrMsg  *string `json:"err_msg,omitempty" `

    /*
        是否可以向下翻页，0不可以，1可以     */
    HasNextPage  *int64 `json:"has_next_page,omitempty" `

    /*
        统计信息     */
    ItemStatistic  *TaobaoXhotelCommentGetmixratelistItemStatisticVo `json:"item_statistic,omitempty" `

    /*
        商品评论列表     */
    MixRates  *[]TaobaoXhotelCommentGetmixratelistMixRateVo `json:"mix_rates,omitempty" `

    /*
        写点评页跳转信息     */
    RateEntrance  *TaobaoXhotelCommentGetmixratelistAddRateEntrance `json:"rate_entrance,omitempty" `

    /*
        成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        总数     */
    TotalNum  *int64 `json:"total_num,omitempty" `

}

func (s *TaobaoXhotelCommentGetmixratelistGetMixRateListResult) SetConfigInfo(v string) *TaobaoXhotelCommentGetmixratelistGetMixRateListResult {
    s.ConfigInfo = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistGetMixRateListResult) SetErrCode(v string) *TaobaoXhotelCommentGetmixratelistGetMixRateListResult {
    s.ErrCode = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistGetMixRateListResult) SetErrMsg(v string) *TaobaoXhotelCommentGetmixratelistGetMixRateListResult {
    s.ErrMsg = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistGetMixRateListResult) SetHasNextPage(v int64) *TaobaoXhotelCommentGetmixratelistGetMixRateListResult {
    s.HasNextPage = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistGetMixRateListResult) SetItemStatistic(v TaobaoXhotelCommentGetmixratelistItemStatisticVo) *TaobaoXhotelCommentGetmixratelistGetMixRateListResult {
    s.ItemStatistic = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistGetMixRateListResult) SetMixRates(v []TaobaoXhotelCommentGetmixratelistMixRateVo) *TaobaoXhotelCommentGetmixratelistGetMixRateListResult {
    s.MixRates = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistGetMixRateListResult) SetRateEntrance(v TaobaoXhotelCommentGetmixratelistAddRateEntrance) *TaobaoXhotelCommentGetmixratelistGetMixRateListResult {
    s.RateEntrance = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistGetMixRateListResult) SetSuccess(v bool) *TaobaoXhotelCommentGetmixratelistGetMixRateListResult {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistGetMixRateListResult) SetTotalNum(v int64) *TaobaoXhotelCommentGetmixratelistGetMixRateListResult {
    s.TotalNum = &v
    return s
}
