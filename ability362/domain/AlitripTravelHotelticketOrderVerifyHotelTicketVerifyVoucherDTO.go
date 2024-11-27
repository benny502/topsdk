package domain


type AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO struct {
    /*
        凭证码     */
    Code  *string `json:"code,omitempty" `

    /*
        使用时间：yyyy-MM-dd HH:mm:ss     */
    UseDate  *string `json:"use_date,omitempty" `

    /*
        证件号     */
    CertificateId  *string `json:"certificate_id,omitempty" `

    /*
        凭证类型 1：票码， 2：券码     */
    Type  *int64 `json:"type,omitempty" `

    /*
        票或券 核销使用数量     */
    UsageNums  *int64 `json:"usage_nums,omitempty" `

    /*
        业务类型：1：门票， 2：酒店     */
    BizType  *int64 `json:"biz_type,omitempty" `

}

func (s *AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO) SetCode(v string) *AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO {
    s.Code = &v
    return s
}
func (s *AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO) SetUseDate(v string) *AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO {
    s.UseDate = &v
    return s
}
func (s *AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO) SetCertificateId(v string) *AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO {
    s.CertificateId = &v
    return s
}
func (s *AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO) SetType(v int64) *AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO {
    s.Type = &v
    return s
}
func (s *AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO) SetUsageNums(v int64) *AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO {
    s.UsageNums = &v
    return s
}
func (s *AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO) SetBizType(v int64) *AlitripTravelHotelticketOrderVerifyHotelTicketVerifyVoucherDTO {
    s.BizType = &v
    return s
}
