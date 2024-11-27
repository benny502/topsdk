package domain


type TaobaoXhotelBnbreviewAddBnbResult struct {
    /*
        状态，成功true，失败false     */
    Success  *bool `json:"success,omitempty" `

    /*
        响应码     */
    ResultCode  *string `json:"result_code,omitempty" `

    /*
        响应信息     */
    ResultMsg  *string `json:"result_msg,omitempty" `

}

func (s *TaobaoXhotelBnbreviewAddBnbResult) SetSuccess(v bool) *TaobaoXhotelBnbreviewAddBnbResult {
    s.Success = &v
    return s
}
func (s *TaobaoXhotelBnbreviewAddBnbResult) SetResultCode(v string) *TaobaoXhotelBnbreviewAddBnbResult {
    s.ResultCode = &v
    return s
}
func (s *TaobaoXhotelBnbreviewAddBnbResult) SetResultMsg(v string) *TaobaoXhotelBnbreviewAddBnbResult {
    s.ResultMsg = &v
    return s
}
