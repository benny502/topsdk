package request


type TaobaoXhotelMultiplerateUpdateRequest struct {
    /*
        废弃，使用rate_plan_code     */
    Rpid  *int64 `json:"rpid,omitempty" required:"false" `
    /*
        价格状态。0为不可售；1为可售，默认可售     */
    Status  *int64 `json:"status,omitempty" required:"false" `
    /*
        卖家房型ID     */
    OutRid  *string `json:"out_rid,omitempty" required:"false" `
    /*
        价格开关 date：开关状态控制的那一天；rate_status：开关状态(0，关闭；1，打开); checkin_status：入住开关(0，关闭；1，打开)；checkout_status：离店开关 (0，关闭；1，打开)     */
    RateSwitchCal  *string `json:"rate_switch_cal,omitempty" required:"false" `
    /*
        连住天数(范围1~5)     */
    Lengthofstay  *int64 `json:"lengthofstay" required:"true" `
    /*
        入住人数(范围1~10)     */
    Occupancy  *int64 `json:"occupancy" required:"true" `
    /*
        币种.CNY为人民币     */
    CurrencyCode  *string `json:"currency_code,omitempty" required:"false" `
    /*
        卖家自己系统的房价code     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" required:"false" `
    /*
        价格和库存信息。 A:use_room_inventory:是否使用房型共享库存，可选值 true false ,false时：使用房价专有库存，此时要求价格和库存必填。 B:date 日期必须为 T---T+180 日内的日期（T为当天），且不能重复 C:price 价格 int类型 取值范围1-99999999 单位为分 D:quota 库存 int 类型 取值范围 0-999（数量库存） 60000(状态库存关) 61000(状态库存开) tax为税费，addBed为加床价，addPerson为加人价1,若连住大于1，price请推送总价     */
    InventoryPrice  *string `json:"inventory_price,omitempty" required:"false" `
    /*
        废弃     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        废弃，使用out_rid     */
    Gid  *int64 `json:"gid,omitempty" required:"false" `
    /*
        系统商，一般不填写，使用须申请     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        儿童人数     */
    Childnum  *int64 `json:"childnum,omitempty" required:"false" `
    /*
        婴儿人数     */
    Infantnum  *int64 `json:"infantnum,omitempty" required:"false" `
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
        -1,状态早餐，和入住人数有关系，几人价就是几份早餐；0：不含早1：含单早2：含双早N：含N早（1-99可选）；（添加RP时为必须）     */
    Breakfast  *int64 `json:"breakfast,omitempty" required:"false" `
    /*
        退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100>Y>=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100>Y>=0（不超过10条）。1.表示任意退{"cancelPolicyType":1};2.表示不能退{"cancelPolicyType":2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{"cancelPolicyType":4,"policyInfo":{"48":10,"24":20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{"cancelPolicyType":5,"policyInfo":{"timeBefore":6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{"cancelPolicyType":6,"policyInfo":{"14":1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间、多间夜扣款     */
    CancelPolicy  *string `json:"cancel_policy,omitempty" required:"false" `
    /*
        在更新rateplan时，同时新增或更新早餐日历。 date：早餐政策属于具体哪一天 breakfast_count：这一天早餐的数量。>=0,<=99 如果date为空，那么会去读取startDate和endDate（格式都为"yyyy-MM-dd"），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。     */
    BreakfastCal  *string `json:"breakfast_cal,omitempty" required:"false" `
    /*
        在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为"yyyy-MM-dd" cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。     */
    CancelPolicyCal  *string `json:"cancel_policy_cal,omitempty" required:"false" `
    /*
         是一个JSONArray 字符串 actionType  操作类型 BOUND: 绑定，UNBOUND：解绑; outXcode  元素编码 ; subTypeCode x 元素子类型， 参考：https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.9MjTPx&docType=1&articleId=121402&previewCode=787DFB0895F05C90D167579A04BD32E3; status: 状态是否生效0 失效, 1生效; shortName x元素标题; time 服务时间段(18:00-21:00); value 商品价值(100 - 999900 单位分); itemDesc 商品使用说明; dimensionType 附加产品使用维度   1:每间房维度 2:每间夜维度; picList 图片格式化信息 [{"url":"https://xxxxx/","isMain":true}]; adultCount 成人数量 (1-99); childCount 儿童数量 (0-99); itemLimit 使用限制, 文字描述,200 字内; checkInStart 入住生效开始时间; checkInEnd 入住生效结束时间; bookStartTime 预定生效开始时间; bookStartEnd 预定生效截止时间; featureDetail 详细信息json字符串 [{"detailName":"免费寄存","detailValue":[""],"type":"single","priority":1}]     */
    HotelXitemInfos  *string `json:"hotel_xitem_infos,omitempty" required:"false" `
}

func (s *TaobaoXhotelMultiplerateUpdateRequest) SetRpid(v int64) *TaobaoXhotelMultiplerateUpdateRequest {
    s.Rpid = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetStatus(v int64) *TaobaoXhotelMultiplerateUpdateRequest {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetOutRid(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetRateSwitchCal(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.RateSwitchCal = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetLengthofstay(v int64) *TaobaoXhotelMultiplerateUpdateRequest {
    s.Lengthofstay = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetOccupancy(v int64) *TaobaoXhotelMultiplerateUpdateRequest {
    s.Occupancy = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetCurrencyCode(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.CurrencyCode = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetRatePlanCode(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.RatePlanCode = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetInventoryPrice(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.InventoryPrice = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetName(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetGid(v int64) *TaobaoXhotelMultiplerateUpdateRequest {
    s.Gid = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetVendor(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetChildnum(v int64) *TaobaoXhotelMultiplerateUpdateRequest {
    s.Childnum = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetInfantnum(v int64) *TaobaoXhotelMultiplerateUpdateRequest {
    s.Infantnum = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetLockEndTime(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.LockEndTime = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetLockStartTime(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.LockStartTime = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetOnlineBookingBindingInfo(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.OnlineBookingBindingInfo = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetBreakfast(v int64) *TaobaoXhotelMultiplerateUpdateRequest {
    s.Breakfast = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetCancelPolicy(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.CancelPolicy = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetBreakfastCal(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.BreakfastCal = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetCancelPolicyCal(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.CancelPolicyCal = &v
    return s
}
func (s *TaobaoXhotelMultiplerateUpdateRequest) SetHotelXitemInfos(v string) *TaobaoXhotelMultiplerateUpdateRequest {
    s.HotelXitemInfos = &v
    return s
}

func (req *TaobaoXhotelMultiplerateUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Rpid != nil) {
        paramMap["rpid"] = *req.Rpid
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.OutRid != nil) {
        paramMap["out_rid"] = *req.OutRid
    }
    if(req.RateSwitchCal != nil) {
        paramMap["rate_switch_cal"] = *req.RateSwitchCal
    }
    if(req.Lengthofstay != nil) {
        paramMap["lengthofstay"] = *req.Lengthofstay
    }
    if(req.Occupancy != nil) {
        paramMap["occupancy"] = *req.Occupancy
    }
    if(req.CurrencyCode != nil) {
        paramMap["currency_code"] = *req.CurrencyCode
    }
    if(req.RatePlanCode != nil) {
        paramMap["rate_plan_code"] = *req.RatePlanCode
    }
    if(req.InventoryPrice != nil) {
        paramMap["inventory_price"] = *req.InventoryPrice
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.Gid != nil) {
        paramMap["gid"] = *req.Gid
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.Childnum != nil) {
        paramMap["childnum"] = *req.Childnum
    }
    if(req.Infantnum != nil) {
        paramMap["infantnum"] = *req.Infantnum
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
    if(req.Breakfast != nil) {
        paramMap["breakfast"] = *req.Breakfast
    }
    if(req.CancelPolicy != nil) {
        paramMap["cancel_policy"] = *req.CancelPolicy
    }
    if(req.BreakfastCal != nil) {
        paramMap["breakfast_cal"] = *req.BreakfastCal
    }
    if(req.CancelPolicyCal != nil) {
        paramMap["cancel_policy_cal"] = *req.CancelPolicyCal
    }
    if(req.HotelXitemInfos != nil) {
        paramMap["hotel_xitem_infos"] = *req.HotelXitemInfos
    }
    return paramMap
}

func (req *TaobaoXhotelMultiplerateUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}