package domain


type TaobaoXhotelBnbreviewAddReviewPicInfo struct {
    /*
        图片地址     */
    Url  *string `json:"url,omitempty" `

}

func (s *TaobaoXhotelBnbreviewAddReviewPicInfo) SetUrl(v string) *TaobaoXhotelBnbreviewAddReviewPicInfo {
    s.Url = &v
    return s
}
