package domain


type AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO struct {
    /*
        凭证码     */
    Code  *string `json:"code,omitempty" `

    /*
        每张票或券可使用次数（如针对一码多刷，往返索道3张票1个码，每张票可使用次数为2，则该码可以刷6次     */
    AvailableNums  *int64 `json:"available_nums,omitempty" `

    /*
        证件号     */
    CertificateId  *string `json:"certificate_id,omitempty" `

    /*
        凭证类型 1：票码， 2：券码      */
    Type  *int64 `json:"type,omitempty" `

    /*
        二维码图片链接     */
    Url  *string `json:"url,omitempty" `

    /*
        凭证 可用/不可用     */
    CanUse  *bool `json:"can_use,omitempty" `

    /*
        已使用次数     */
    UsageNums  *int64 `json:"usage_nums,omitempty" `

    /*
        业务类型：1：门票， 2：酒店     */
    BizType  *int64 `json:"biz_type,omitempty" `

}

func (s *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO) SetCode(v string) *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO {
    s.Code = &v
    return s
}
func (s *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO) SetAvailableNums(v int64) *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO {
    s.AvailableNums = &v
    return s
}
func (s *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO) SetCertificateId(v string) *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO {
    s.CertificateId = &v
    return s
}
func (s *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO) SetType(v int64) *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO {
    s.Type = &v
    return s
}
func (s *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO) SetUrl(v string) *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO {
    s.Url = &v
    return s
}
func (s *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO) SetCanUse(v bool) *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO {
    s.CanUse = &v
    return s
}
func (s *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO) SetUsageNums(v int64) *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO {
    s.UsageNums = &v
    return s
}
func (s *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO) SetBizType(v int64) *AlitripTravelHotelticketOrderCreateHotelTicketVoucherDTO {
    s.BizType = &v
    return s
}
