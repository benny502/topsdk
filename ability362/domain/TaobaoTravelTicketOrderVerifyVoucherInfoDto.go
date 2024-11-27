package domain


type TaobaoTravelTicketOrderVerifyVoucherInfoDto struct {
    /*
        用户短信会收到的确认号     */
    ConfirmCode  *string `json:"confirm_code,omitempty" `

    /*
        凭证使用次数     */
    UsedQuantity  *int64 `json:"used_quantity,omitempty" `

    /*
        凭证使用时间，格式:yyyy-MM-dd HH:mm:ss     */
    UsedDate  *string `json:"used_date,omitempty" `

}

func (s *TaobaoTravelTicketOrderVerifyVoucherInfoDto) SetConfirmCode(v string) *TaobaoTravelTicketOrderVerifyVoucherInfoDto {
    s.ConfirmCode = &v
    return s
}
func (s *TaobaoTravelTicketOrderVerifyVoucherInfoDto) SetUsedQuantity(v int64) *TaobaoTravelTicketOrderVerifyVoucherInfoDto {
    s.UsedQuantity = &v
    return s
}
func (s *TaobaoTravelTicketOrderVerifyVoucherInfoDto) SetUsedDate(v string) *TaobaoTravelTicketOrderVerifyVoucherInfoDto {
    s.UsedDate = &v
    return s
}
