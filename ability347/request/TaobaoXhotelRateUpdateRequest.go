package request


type TaobaoXhotelRateUpdateRequest struct {
    /*
        不推荐使用，请使用out_rid     */
    Gid  *int64 `json:"gid,omitempty" required:"false" `
    /*
        不推荐使用，请使用rateplan_code     */
    Rpid  *int64 `json:"rpid,omitempty" required:"false" `
    /*
        废弃     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        每日价格和房价专有库存信息。A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。B:date  日期必须为 T---T+180 日内的日期（T为当天），且不能重复C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)     */
    InventoryPrice  *string `json:"inventory_price,omitempty" required:"false" `
    /*
        不推荐使用     */
    AddBed  *int64 `json:"add_bed,omitempty" required:"false" `
    /*
        不推荐使用     */
    AddBedPrice  *int64 `json:"add_bed_price,omitempty" required:"false" `
    /*
        不推荐使用（仅支持CNY）     */
    CurrencyCode  *int64 `json:"currency_code,omitempty" required:"false" `
    /*
        不推荐使用     */
    ShijiaTag  *int64 `json:"shijia_tag,omitempty" required:"false" `
    /*
        不推荐使用     */
    JishiquerenTag  *int64 `json:"jishiqueren_tag,omitempty" required:"false" `
    /*
        系统商，一般不用填写，使用需要申请     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        商家价格计划编码     */
    RateplanCode  *string `json:"rateplan_code,omitempty" required:"false" `
    /*
        商家房型ID     */
    OutRid  *string `json:"out_rid,omitempty" required:"false" `
    /*
        日历价格开关， date：开关状态控制的是那一天 rate_status：开关状态。0，关闭；1，打开     */
    RateSwitchCal  *string `json:"rate_switch_cal,omitempty" required:"false" `
    /*
        锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）     */
    LockEndTime  *string `json:"lock_end_time,omitempty" required:"false" `
    /*
        锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）     */
    LockStartTime  *string `json:"lock_start_time,omitempty" required:"false" `
    /*
        在线预约关联关系推送，priceRuleNumber：加价规则序号     */
    OnlineBookingBindingInfo  *string `json:"online_booking_binding_info,omitempty" required:"false" `
    /*
         是一个JSONArray 字符串 actionType  操作类型 BOUND: 绑定，UNBOUND：解绑; outXcode  元素编码 ; subTypeCode x 元素子类型， 参考：https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.9MjTPx&docType=1&articleId=121402&previewCode=787DFB0895F05C90D167579A04BD32E3; status: 状态是否生效0 失效, 1生效; shortName x元素标题; time 服务时间段(18:00-21:00); value 商品价值(100 - 999900 单位分); itemDesc 商品使用说明; dimensionType 附加产品使用维度   1:每间房维度 2:每间夜维度; picList 图片格式化信息 [{"url":"https://xxxxx/","isMain":true}]; adultCount 成人数量 (1-99); childCount 儿童数量 (0-99); itemLimit 使用限制, 文字描述,200 字内; checkInStart 入住生效开始时间; checkInEnd 入住生效结束时间; bookStartTime 预定生效开始时间; bookStartEnd 预定生效截止时间; featureDetail 详细信息json字符串 [{"detailName":"免费寄存","detailValue":[""],"type":"single","priority":1}]     */
    HotelXitemInfos  *string `json:"hotel_xitem_infos,omitempty" required:"false" `
}

func (s *TaobaoXhotelRateUpdateRequest) SetGid(v int64) *TaobaoXhotelRateUpdateRequest {
    s.Gid = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetRpid(v int64) *TaobaoXhotelRateUpdateRequest {
    s.Rpid = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetName(v string) *TaobaoXhotelRateUpdateRequest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetInventoryPrice(v string) *TaobaoXhotelRateUpdateRequest {
    s.InventoryPrice = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetAddBed(v int64) *TaobaoXhotelRateUpdateRequest {
    s.AddBed = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetAddBedPrice(v int64) *TaobaoXhotelRateUpdateRequest {
    s.AddBedPrice = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetCurrencyCode(v int64) *TaobaoXhotelRateUpdateRequest {
    s.CurrencyCode = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetShijiaTag(v int64) *TaobaoXhotelRateUpdateRequest {
    s.ShijiaTag = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetJishiquerenTag(v int64) *TaobaoXhotelRateUpdateRequest {
    s.JishiquerenTag = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetVendor(v string) *TaobaoXhotelRateUpdateRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetRateplanCode(v string) *TaobaoXhotelRateUpdateRequest {
    s.RateplanCode = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetOutRid(v string) *TaobaoXhotelRateUpdateRequest {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetRateSwitchCal(v string) *TaobaoXhotelRateUpdateRequest {
    s.RateSwitchCal = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetLockEndTime(v string) *TaobaoXhotelRateUpdateRequest {
    s.LockEndTime = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetLockStartTime(v string) *TaobaoXhotelRateUpdateRequest {
    s.LockStartTime = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetOnlineBookingBindingInfo(v string) *TaobaoXhotelRateUpdateRequest {
    s.OnlineBookingBindingInfo = &v
    return s
}
func (s *TaobaoXhotelRateUpdateRequest) SetHotelXitemInfos(v string) *TaobaoXhotelRateUpdateRequest {
    s.HotelXitemInfos = &v
    return s
}

func (req *TaobaoXhotelRateUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Gid != nil) {
        paramMap["gid"] = *req.Gid
    }
    if(req.Rpid != nil) {
        paramMap["rpid"] = *req.Rpid
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.InventoryPrice != nil) {
        paramMap["inventory_price"] = *req.InventoryPrice
    }
    if(req.AddBed != nil) {
        paramMap["add_bed"] = *req.AddBed
    }
    if(req.AddBedPrice != nil) {
        paramMap["add_bed_price"] = *req.AddBedPrice
    }
    if(req.CurrencyCode != nil) {
        paramMap["currency_code"] = *req.CurrencyCode
    }
    if(req.ShijiaTag != nil) {
        paramMap["shijia_tag"] = *req.ShijiaTag
    }
    if(req.JishiquerenTag != nil) {
        paramMap["jishiqueren_tag"] = *req.JishiquerenTag
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.RateplanCode != nil) {
        paramMap["rateplan_code"] = *req.RateplanCode
    }
    if(req.OutRid != nil) {
        paramMap["out_rid"] = *req.OutRid
    }
    if(req.RateSwitchCal != nil) {
        paramMap["rate_switch_cal"] = *req.RateSwitchCal
    }
    if(req.LockEndTime != nil) {
        paramMap["lock_end_time"] = *req.LockEndTime
    }
    if(req.LockStartTime != nil) {
        paramMap["lock_start_time"] = *req.LockStartTime
    }
    if(req.OnlineBookingBindingInfo != nil) {
        paramMap["online_booking_binding_info"] = *req.OnlineBookingBindingInfo
    }
    if(req.HotelXitemInfos != nil) {
        paramMap["hotel_xitem_infos"] = *req.HotelXitemInfos
    }
    return paramMap
}

func (req *TaobaoXhotelRateUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}