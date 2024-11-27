package request


type TaobaoXhotelRateAddRequest struct {
    /*
        gid酒店商品id     */
    Gid  *int64 `json:"gid,omitempty" required:"false" `
    /*
        酒店RPID     */
    Rpid  *int64 `json:"rpid,omitempty" required:"false" `
    /*
        用于标示该宝贝的售卖渠道信息，允许同一个卖家酒店房型在淘宝系统发布多个售卖渠道的宝贝的价格。     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        卖家自己系统的Code，简称RateCode     */
    RateplanCode  *string `json:"rateplan_code,omitempty" required:"false" `
    /*
        卖家房型ID, 这是卖家自己系统中的房型ID，注意：需按照规则组合     */
    OutRid  *string `json:"out_rid,omitempty" required:"false" `
    /*
        在添加新rate时，同时添加rate开关日历。可以只设定想设定的某些天，可以不连续。date：开关状态控制的是那一天rate_status：开关状态。0，关闭；1，打开     */
    RateSwitchCal  *string `json:"rate_switch_cal,omitempty" required:"false" `
    /*
        名称     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        价格和库存信息。A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。B:date  日期必须为 T---T+90 日内的日期（T为当天），且不能重复C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)     */
    InventoryPrice  *string `json:"inventory_price" required:"true" `
    /*
        额外服务-是否可以加床，1：不可以，2：可以     */
    AddBed  *int64 `json:"add_bed,omitempty" required:"false" `
    /*
        额外服务-加床价格     */
    AddBedPrice  *int64 `json:"add_bed_price,omitempty" required:"false" `
    /*
        币种（仅支持CNY）     */
    CurrencyCode  *int64 `json:"currency_code,omitempty" required:"false" `
    /*
        实价有房标签（RP支付类型为全额支付）     */
    ShijiaTag  *int64 `json:"shijia_tag,omitempty" required:"false" `
    /*
        “即时确认”标识，此类商品预订后直接发货。     */
    JishiquerenTag  *int64 `json:"jishiqueren_tag,omitempty" required:"false" `
    /*
        锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）     */
    LockEndTime  *string `json:"lock_end_time,omitempty" required:"false" `
    /*
        锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）     */
    LockStartTime  *string `json:"lock_start_time,omitempty" required:"false" `
    /*
        币种信息,默认是CNY,  @see com.taobao.trip.hotel.model.enums.CurrencyEnum     */
    CurrencyCodeName  *string `json:"currency_code_name,omitempty" required:"false" `
    /*
        操作人信息     */
    Operator  *string `json:"operator,omitempty" required:"false" `
    /*
        默认是2 ,     */
    Source  *int64 `json:"source,omitempty" required:"false" `
    /*
        1是开,0是关, 不填默认是开, rate状态     */
    Status  *int64 `json:"status,omitempty" required:"false" `
    /*
        在线预约关联关系推送，priceRuleNumber：加价规则序号     */
    OnlineBookingBindingInfo  *string `json:"online_booking_binding_info,omitempty" required:"false" `
    /*
         是一个JSONArray 字符串 actionType  操作类型 BOUND: 绑定，UNBOUND：解绑; outXcode  元素编码 ; subTypeCode x 元素子类型， 参考：https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.9MjTPx&docType=1&articleId=121402&previewCode=787DFB0895F05C90D167579A04BD32E3; status: 状态是否生效0 失效, 1生效; shortName x元素标题; time 服务时间段(18:00-21:00); value 商品价值(100 - 999900 单位分); itemDesc 商品使用说明; dimensionType 附加产品使用维度   1:每间房维度 2:每间夜维度; picList 图片格式化信息 [{"url":"https://xxxxx/","isMain":true}]; adultCount 成人数量 (1-99); childCount 儿童数量 (0-99); itemLimit 使用限制, 文字描述,200 字内; checkInStart 入住生效开始时间; checkInEnd 入住生效结束时间; bookStartTime 预定生效开始时间; bookStartEnd 预定生效截止时间; featureDetail 详细信息json字符串 [{"detailName":"免费寄存","detailValue":[""],"type":"single","priority":1}]     */
    HotelXitemInfos  *string `json:"hotel_xitem_infos,omitempty" required:"false" `
}

func (s *TaobaoXhotelRateAddRequest) SetGid(v int64) *TaobaoXhotelRateAddRequest {
    s.Gid = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetRpid(v int64) *TaobaoXhotelRateAddRequest {
    s.Rpid = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetVendor(v string) *TaobaoXhotelRateAddRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetRateplanCode(v string) *TaobaoXhotelRateAddRequest {
    s.RateplanCode = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetOutRid(v string) *TaobaoXhotelRateAddRequest {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetRateSwitchCal(v string) *TaobaoXhotelRateAddRequest {
    s.RateSwitchCal = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetName(v string) *TaobaoXhotelRateAddRequest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetInventoryPrice(v string) *TaobaoXhotelRateAddRequest {
    s.InventoryPrice = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetAddBed(v int64) *TaobaoXhotelRateAddRequest {
    s.AddBed = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetAddBedPrice(v int64) *TaobaoXhotelRateAddRequest {
    s.AddBedPrice = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetCurrencyCode(v int64) *TaobaoXhotelRateAddRequest {
    s.CurrencyCode = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetShijiaTag(v int64) *TaobaoXhotelRateAddRequest {
    s.ShijiaTag = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetJishiquerenTag(v int64) *TaobaoXhotelRateAddRequest {
    s.JishiquerenTag = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetLockEndTime(v string) *TaobaoXhotelRateAddRequest {
    s.LockEndTime = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetLockStartTime(v string) *TaobaoXhotelRateAddRequest {
    s.LockStartTime = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetCurrencyCodeName(v string) *TaobaoXhotelRateAddRequest {
    s.CurrencyCodeName = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetOperator(v string) *TaobaoXhotelRateAddRequest {
    s.Operator = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetSource(v int64) *TaobaoXhotelRateAddRequest {
    s.Source = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetStatus(v int64) *TaobaoXhotelRateAddRequest {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetOnlineBookingBindingInfo(v string) *TaobaoXhotelRateAddRequest {
    s.OnlineBookingBindingInfo = &v
    return s
}
func (s *TaobaoXhotelRateAddRequest) SetHotelXitemInfos(v string) *TaobaoXhotelRateAddRequest {
    s.HotelXitemInfos = &v
    return s
}

func (req *TaobaoXhotelRateAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Gid != nil) {
        paramMap["gid"] = *req.Gid
    }
    if(req.Rpid != nil) {
        paramMap["rpid"] = *req.Rpid
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
    if(req.LockEndTime != nil) {
        paramMap["lock_end_time"] = *req.LockEndTime
    }
    if(req.LockStartTime != nil) {
        paramMap["lock_start_time"] = *req.LockStartTime
    }
    if(req.CurrencyCodeName != nil) {
        paramMap["currency_code_name"] = *req.CurrencyCodeName
    }
    if(req.Operator != nil) {
        paramMap["operator"] = *req.Operator
    }
    if(req.Source != nil) {
        paramMap["source"] = *req.Source
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.OnlineBookingBindingInfo != nil) {
        paramMap["online_booking_binding_info"] = *req.OnlineBookingBindingInfo
    }
    if(req.HotelXitemInfos != nil) {
        paramMap["hotel_xitem_infos"] = *req.HotelXitemInfos
    }
    return paramMap
}

func (req *TaobaoXhotelRateAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}