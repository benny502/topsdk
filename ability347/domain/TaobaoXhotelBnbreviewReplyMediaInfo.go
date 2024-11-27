package domain


type TaobaoXhotelBnbreviewReplyMediaInfo struct {
    /*
        媒体信息，1-图片，2-视频     */
    Type  *int64 `json:"type,omitempty" `

    /*
        媒体URL地址     */
    Urls  *[]string `json:"urls,omitempty" `

}

func (s *TaobaoXhotelBnbreviewReplyMediaInfo) SetType(v int64) *TaobaoXhotelBnbreviewReplyMediaInfo {
    s.Type = &v
    return s
}
func (s *TaobaoXhotelBnbreviewReplyMediaInfo) SetUrls(v []string) *TaobaoXhotelBnbreviewReplyMediaInfo {
    s.Urls = &v
    return s
}
