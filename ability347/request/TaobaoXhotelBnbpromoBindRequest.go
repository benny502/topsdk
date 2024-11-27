package request

import (
        "topsdk/ability347/domain"
        "topsdk/util"
    )

type TaobaoXhotelBnbpromoBindRequest struct {
    /*
        营销活动code     */
    ActivityCode  *string `json:"activity_code,omitempty" required:"false" `
    /*
        外部价格信息     */
    RateInfos  *[]domain.TaobaoXhotelBnbpromoBindPromoRateInfo `json:"rate_infos,omitempty" required:"false" `
    /*
        活动入住时间，民宿通用营销必填     */
    CheckInFrom  *util.LocalTime `json:"check_in_from,omitempty" required:"false" `
    /*
        活动离店时间，民宿通用营销必填     */
    CheckOutTo  *util.LocalTime `json:"check_out_to,omitempty" required:"false" `
}

func (s *TaobaoXhotelBnbpromoBindRequest) SetActivityCode(v string) *TaobaoXhotelBnbpromoBindRequest {
    s.ActivityCode = &v
    return s
}
func (s *TaobaoXhotelBnbpromoBindRequest) SetRateInfos(v []domain.TaobaoXhotelBnbpromoBindPromoRateInfo) *TaobaoXhotelBnbpromoBindRequest {
    s.RateInfos = &v
    return s
}
func (s *TaobaoXhotelBnbpromoBindRequest) SetCheckInFrom(v util.LocalTime) *TaobaoXhotelBnbpromoBindRequest {
    s.CheckInFrom = &v
    return s
}
func (s *TaobaoXhotelBnbpromoBindRequest) SetCheckOutTo(v util.LocalTime) *TaobaoXhotelBnbpromoBindRequest {
    s.CheckOutTo = &v
    return s
}

func (req *TaobaoXhotelBnbpromoBindRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ActivityCode != nil) {
        paramMap["activity_code"] = *req.ActivityCode
    }
    if(req.RateInfos != nil) {
        paramMap["rate_infos"] = util.ConvertStructList(*req.RateInfos)
    }
    if(req.CheckInFrom != nil) {
        paramMap["check_in_from"] = *req.CheckInFrom
    }
    if(req.CheckOutTo != nil) {
        paramMap["check_out_to"] = *req.CheckOutTo
    }
    return paramMap
}

func (req *TaobaoXhotelBnbpromoBindRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}