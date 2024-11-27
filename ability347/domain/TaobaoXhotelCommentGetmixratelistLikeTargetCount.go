package domain


type TaobaoXhotelCommentGetmixratelistLikeTargetCount struct {
    /*
        目标对象的赞数量或者踩数量     */
    Count  *int64 `json:"count,omitempty" `

    /*
        目标对象的ID     */
    TargetId  *int64 `json:"target_id,omitempty" `

    /*
        针对当前targetId,是否点赞过或踩过     */
    Voted  *bool `json:"voted,omitempty" `

}

func (s *TaobaoXhotelCommentGetmixratelistLikeTargetCount) SetCount(v int64) *TaobaoXhotelCommentGetmixratelistLikeTargetCount {
    s.Count = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistLikeTargetCount) SetTargetId(v int64) *TaobaoXhotelCommentGetmixratelistLikeTargetCount {
    s.TargetId = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistLikeTargetCount) SetVoted(v bool) *TaobaoXhotelCommentGetmixratelistLikeTargetCount {
    s.Voted = &v
    return s
}
