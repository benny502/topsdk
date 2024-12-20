package request


type TaobaoXhotelRatesLiteIncrUpdateRequest struct {
    /*
        json格式: [{"rp_code":"031612_181220000074_F1_1","out_rid":"031612_HD","vendor":"taobao","useRoomInv":0,"invAndSwitch":[{"date":"2020-01-09","price":1200,"quota":23,"rateStatus":1},{"date":"2020-01-10","price":1200,"quota":23,"rateStatus":0}]},{"rp_code":"1234","out_rid":"123456","vendor":"taobao","useRoomInv":0,"invAndSwitch":[{"date":"yyyy-MM-dd","price":1200,"quota":23,"rateStatus":1},{"date":"yyyy-MM-dd","price":1200,"quota":23,"rateStatus":1}]}] rp_code:房价code,out_rid:房型code,必传字段 useRoomInv: 表示是否使用房型库存(1表示是,0表示否),可选字段 invAndSwitch: 表示要更新的日历化价格库存以及开关,增量更新,更新哪天就传那天的日历化单元, 可选字段 invAndSwitch.date: 表示日历化单元里的日期,格式:yyyy-MM-dd, 在 invAndSwitch 里的每个日历化单元里是必填字段 invAndSwitch.price: 表示要更新的价格,单位 分, 可选字段 invAndSwitch.quota: 表示要更新的普通库存,可选字段  invAndSwitch.rateStatus: 表示要更新的价格开关,1表示开,0表示关,可选字段。请注意，该接口为轻量级批量增量更新接口，只能更新未来4天内价格库存开关信息，传入的参数日期超过4天了会报错。该接口核心是解决近几天价格库存实时性。     */
    RateListInvInfo  *string `json:"rate_list_inv_info" required:"true" `
}

func (s *TaobaoXhotelRatesLiteIncrUpdateRequest) SetRateListInvInfo(v string) *TaobaoXhotelRatesLiteIncrUpdateRequest {
    s.RateListInvInfo = &v
    return s
}

func (req *TaobaoXhotelRatesLiteIncrUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RateListInvInfo != nil) {
        paramMap["rate_list_inv_info"] = *req.RateListInvInfo
    }
    return paramMap
}

func (req *TaobaoXhotelRatesLiteIncrUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}