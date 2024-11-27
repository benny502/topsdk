package domain


type TaobaoXhotelCommentGetmixratelistParentInfo struct {
    /*
        用户昵称,已脱敏     */
    UserNick  *string `json:"user_nick,omitempty" `

}

func (s *TaobaoXhotelCommentGetmixratelistParentInfo) SetUserNick(v string) *TaobaoXhotelCommentGetmixratelistParentInfo {
    s.UserNick = &v
    return s
}
