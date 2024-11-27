package domain


type TaobaoXhotelCommentGetmixratelistScoreInfo struct {
    /*
        数量     */
    Count  *int64 `json:"count,omitempty" `

    /*
        描述     */
    Desc  *string `json:"desc,omitempty" `

    /*
        标签     */
    Label  *string `json:"label,omitempty" `

    /*
        分数     */
    Score  *string `json:"score,omitempty" `

}

func (s *TaobaoXhotelCommentGetmixratelistScoreInfo) SetCount(v int64) *TaobaoXhotelCommentGetmixratelistScoreInfo {
    s.Count = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistScoreInfo) SetDesc(v string) *TaobaoXhotelCommentGetmixratelistScoreInfo {
    s.Desc = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistScoreInfo) SetLabel(v string) *TaobaoXhotelCommentGetmixratelistScoreInfo {
    s.Label = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistScoreInfo) SetScore(v string) *TaobaoXhotelCommentGetmixratelistScoreInfo {
    s.Score = &v
    return s
}
