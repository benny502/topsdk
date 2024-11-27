package domain


type TaobaoXhotelBnbroomtypeAddBnbPictureDTO struct {
    /*
        图片属性 取值范围只能是：[普通图, 平面图, 全景图]     */
    Attribute  *string `json:"attribute,omitempty" `

    /*
        是否主图 false不是， true是,是否主图 主图只能有一个，如果有多个或者没有，则会报错     */
    Ismain  *bool `json:"ismain,omitempty" `

    /*
        图片描述     */
    Des  *string `json:"des,omitempty" `

    /*
        type表示图片类型，取值范围只能是：【厨房、卫生间、客厅、卧室、其他】     */
    Type  *string `json:"type,omitempty" `

    /*
        图片地址     */
    Url  *string `json:"url,omitempty" `

}

func (s *TaobaoXhotelBnbroomtypeAddBnbPictureDTO) SetAttribute(v string) *TaobaoXhotelBnbroomtypeAddBnbPictureDTO {
    s.Attribute = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbPictureDTO) SetIsmain(v bool) *TaobaoXhotelBnbroomtypeAddBnbPictureDTO {
    s.Ismain = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbPictureDTO) SetDes(v string) *TaobaoXhotelBnbroomtypeAddBnbPictureDTO {
    s.Des = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbPictureDTO) SetType(v string) *TaobaoXhotelBnbroomtypeAddBnbPictureDTO {
    s.Type = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbPictureDTO) SetUrl(v string) *TaobaoXhotelBnbroomtypeAddBnbPictureDTO {
    s.Url = &v
    return s
}
