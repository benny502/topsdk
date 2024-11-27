package domain


type TaobaoXhotelCommentGetmixratelistJumpInfo struct {
    /*
        h5地址     */
    JumpH5Url  *string `json:"jump_h5_url,omitempty" `

    /*
        是否跳转native     */
    JumpNative  *bool `json:"jump_native,omitempty" `

    /*
        native地址     */
    JumpNativeUrl  *string `json:"jump_native_url,omitempty" `

    /*
        要往native跳转的客户端版本,当该字段存在时，只有大于该版本才会走native跳转     */
    NativeVersion  *string `json:"native_version,omitempty" `

    /*
        native页面信息     */
    PageName  *string `json:"page_name,omitempty" `

}

func (s *TaobaoXhotelCommentGetmixratelistJumpInfo) SetJumpH5Url(v string) *TaobaoXhotelCommentGetmixratelistJumpInfo {
    s.JumpH5Url = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistJumpInfo) SetJumpNative(v bool) *TaobaoXhotelCommentGetmixratelistJumpInfo {
    s.JumpNative = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistJumpInfo) SetJumpNativeUrl(v string) *TaobaoXhotelCommentGetmixratelistJumpInfo {
    s.JumpNativeUrl = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistJumpInfo) SetNativeVersion(v string) *TaobaoXhotelCommentGetmixratelistJumpInfo {
    s.NativeVersion = &v
    return s
}
func (s *TaobaoXhotelCommentGetmixratelistJumpInfo) SetPageName(v string) *TaobaoXhotelCommentGetmixratelistJumpInfo {
    s.PageName = &v
    return s
}
