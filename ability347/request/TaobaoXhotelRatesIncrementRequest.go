package request


type TaobaoXhotelRatesIncrementRequest struct {
    /*
        批量修改价格和房价专有库存信息，json格式,可同时修改多套房型+价格计划的价格：A:use_room_inventory:是否使用房型共享库存，可选值 true false 1、true时：使用房型共享库存 2、false时：使用房价专有库存，此时要求房价专有库存必填。B:date  日期必须为 T---T+180 日内的日期（T为当天），不能重复。 C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 房价专有库存 int 类型 取值范围  0-999（数量库存） 60000(状态库存关) 61000(状态库存开) E:status 价格开关，0为关，1为开。lock_start_time为锁库存开始时间，lock_end_time为锁库存结束时间，如果当前时间在这个时间范围内，那么不允许修改库存。示例值：（1）[{"out_rid":"ABCDE_123","rateplan_code":"ABCDE_WHL01","vendor":"","lock_start_time":"","lock_end_time":"","data":{"use_room_inventory":false,"inventory_price":[{"date":"2013-11-18","quota":1,"price":1000,"status":1},{"date":"2013-11-19","quota":1,"price":1000,"status":0}]}},{"out_rid":"ABCDE_234","rateplan_code":"ABCDE_WHL01","vendor":"","data":{"use_room_inventory":false,"inventory_price":[{"date":"2013-11-18","quota":1,"price":1000,"status":1},{"date":"2013-11-19","quota":1,"price":1000,"status":0}]}}]     */
    RateInventoryPriceMap  *string `json:"rate_inventory_price_map" required:"true" `
}

func (s *TaobaoXhotelRatesIncrementRequest) SetRateInventoryPriceMap(v string) *TaobaoXhotelRatesIncrementRequest {
    s.RateInventoryPriceMap = &v
    return s
}

func (req *TaobaoXhotelRatesIncrementRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RateInventoryPriceMap != nil) {
        paramMap["rate_inventory_price_map"] = *req.RateInventoryPriceMap
    }
    return paramMap
}

func (req *TaobaoXhotelRatesIncrementRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}