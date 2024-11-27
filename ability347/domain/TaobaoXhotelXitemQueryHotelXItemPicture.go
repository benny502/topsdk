package domain


type TaobaoXhotelXitemQueryHotelXItemPicture struct {
    /*
        图片地址     */
    Url  *string `json:"url,omitempty" `

    /*
        是否主图     */
    IsMain  *bool `json:"is_main,omitempty" `

}

func (s *TaobaoXhotelXitemQueryHotelXItemPicture) SetUrl(v string) *TaobaoXhotelXitemQueryHotelXItemPicture {
    s.Url = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXItemPicture) SetIsMain(v bool) *TaobaoXhotelXitemQueryHotelXItemPicture {
    s.IsMain = &v
    return s
}
