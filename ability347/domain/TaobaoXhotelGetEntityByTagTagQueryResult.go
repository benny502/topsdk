package domain


type TaobaoXhotelGetEntityByTagTagQueryResult struct {
    /*
        总数     */
    TotalAmount  *int64 `json:"total_amount,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        列表     */
    TagEntityDoList  *[]TaobaoXhotelGetEntityByTagTagEntityDoList `json:"tag_entity_do_list,omitempty" `

    /*
        token     */
    TokenStr  *string `json:"token_str,omitempty" `

    /*
        耗时     */
    SpentTime  *int64 `json:"spent_time,omitempty" `

    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoXhotelGetEntityByTagTagQueryResult) SetTotalAmount(v int64) *TaobaoXhotelGetEntityByTagTagQueryResult {
    s.TotalAmount = &v
    return s
}
func (s *TaobaoXhotelGetEntityByTagTagQueryResult) SetSuccess(v bool) *TaobaoXhotelGetEntityByTagTagQueryResult {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelGetEntityByTagTagQueryResult) SetTagEntityDoList(v []TaobaoXhotelGetEntityByTagTagEntityDoList) *TaobaoXhotelGetEntityByTagTagQueryResult {
    s.TagEntityDoList = &v
    return s
}
func (s *TaobaoXhotelGetEntityByTagTagQueryResult) SetTokenStr(v string) *TaobaoXhotelGetEntityByTagTagQueryResult {
    s.TokenStr = &v
    return s
}
func (s *TaobaoXhotelGetEntityByTagTagQueryResult) SetSpentTime(v int64) *TaobaoXhotelGetEntityByTagTagQueryResult {
    s.SpentTime = &v
    return s
}
func (s *TaobaoXhotelGetEntityByTagTagQueryResult) SetErrorMsg(v string) *TaobaoXhotelGetEntityByTagTagQueryResult {
    s.ErrorMsg = &v
    return s
}
