package domain


type TaobaoXhotelBnbhouseAddBnbPictureDTO struct {
    /*
        图片属性 取值范围只能是：[普通图, 平面图, 全景图]     */
    Attribute  *string `json:"attribute,omitempty" `

    /*
        是否主图  主图只能有一个，如果有多个或者没有，则会报错     */
    Ismain  *bool `json:"ismain,omitempty" `

    /*
        type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]     */
    Type  *string `json:"type,omitempty" `

    /*
        图片地址     */
    Url  *string `json:"url,omitempty" `

    /*
        图片描述     */
    Des  *string `json:"des,omitempty" `

}

func (s *TaobaoXhotelBnbhouseAddBnbPictureDTO) SetAttribute(v string) *TaobaoXhotelBnbhouseAddBnbPictureDTO {
    s.Attribute = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbPictureDTO) SetIsmain(v bool) *TaobaoXhotelBnbhouseAddBnbPictureDTO {
    s.Ismain = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbPictureDTO) SetType(v string) *TaobaoXhotelBnbhouseAddBnbPictureDTO {
    s.Type = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbPictureDTO) SetUrl(v string) *TaobaoXhotelBnbhouseAddBnbPictureDTO {
    s.Url = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbPictureDTO) SetDes(v string) *TaobaoXhotelBnbhouseAddBnbPictureDTO {
    s.Des = &v
    return s
}
