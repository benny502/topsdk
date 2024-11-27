package request

import (
        "topsdk/util"
    )

type TaobaoXhotelRateplanUpdateRequest struct {
    /*
        不推荐使用，使用ratePlanCode来标识要修改的RP     */
    Rpid  *int64 `json:"rpid,omitempty" required:"false" `
    /*
        在淘宝搜索页面展示的房价名称；（添加RP时为必须）。注意该名称不要包含早餐相关信息，如果想维护早餐信息，请设置breakfast_count字段即可。     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        RP的英文名称     */
    EnglishName  *string `json:"english_name,omitempty" required:"false" `
    /*
        -1,状态早餐，和入住人数有关系，几人价就是几份早餐；0：不含早1：含单早2：含双早N：含N早（1-99可选）；（添加RP时为必须）     */
    BreakfastCount  *int64 `json:"breakfast_count,omitempty" required:"false" `
    /*
        不推荐使用     */
    FeeBreakfastCount  *int64 `json:"fee_breakfast_count,omitempty" required:"false" `
    /*
        不推荐使用     */
    FeeBreakfastAmount  *int64 `json:"fee_breakfast_amount,omitempty" required:"false" `
    /*
        不推荐使用     */
    FeeGovTaxAmount  *int64 `json:"fee_gov_tax_amount,omitempty" required:"false" `
    /*
        不推荐使用     */
    FeeGovTaxPercent  *int64 `json:"fee_gov_tax_percent,omitempty" required:"false" `
    /*
        不推荐使用     */
    FeeServiceAmount  *int64 `json:"fee_service_amount,omitempty" required:"false" `
    /*
        不推荐使用     */
    FeeServicePercent  *int64 `json:"fee_service_percent,omitempty" required:"false" `
    /*
        不推荐使用     */
    ExtendFee  *string `json:"extend_fee,omitempty" required:"false" `
    /*
        最小入住天数（1-90）。默认1,小时房RP请设置为1     */
    MinDays  *int64 `json:"min_days,omitempty" required:"false" `
    /*
        最大入住天数（1-90）。默认90,信用住最大入住天数不超过9天,小时房RP请设置为1,个别卖家支持180     */
    MaxDays  *int64 `json:"max_days,omitempty" required:"false" `
    /*
        首日入住房间数（1-99）。默认1。【废弃】     */
    MinAmount  *int64 `json:"min_amount,omitempty" required:"false" `
    /*
        最小提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表必须至少提前两天预定，那么如果想预定24号入住，,必须在23号零点前下单。     */
    MinAdvHours  *int64 `json:"min_adv_hours,omitempty" required:"false" `
    /*
        最大提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表最多提前两天预定，那么如果想预定24号入住，,必须在23号零点以后下单。     */
    MaxAdvHours  *int64 `json:"max_adv_hours,omitempty" required:"false" `
    /*
        产品每日开始销售时间，start_time一定为当天时间     */
    StartTime  *string `json:"start_time,omitempty" required:"false" `
    /*
        产品每日开始销售时间，start_time一定为当天时间     */
    EndTime  *string `json:"end_time,omitempty" required:"false" `
    /*
        退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100>Y>=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100>Y>=0（不超过10条）。1.表示任意退{"cancelPolicyType":1};2.表示不能退{"cancelPolicyType":2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{"cancelPolicyType":4,"policyInfo":{"48":10,"24":20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{"cancelPolicyType":5,"policyInfo":{"timeBefore":6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{"cancelPolicyType":6,"policyInfo":{"14":1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间、多间夜扣款     */
    CancelPolicy  *string `json:"cancel_policy,omitempty" required:"false" `
    /*
        不推荐使用     */
    Extend  *string `json:"extend,omitempty" required:"false" `
    /*
        1：开启（默认）2：关闭。如果没传值那么默认默认值为1；（添加RP时为必须）     */
    Status  *int64 `json:"status,omitempty" required:"false" `
    /*
        担保类型，只支持： 0 无担保 1 峰时首晚担保 2峰时全额担保 3全天首晚担保 4全天全额担保     */
    GuaranteeType  *int64 `json:"guarantee_type,omitempty" required:"false" `
    /*
        分时担保每日开始担保时间。 （如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）     */
    GuaranteeStartTime  *string `json:"guarantee_start_time,omitempty" required:"false" `
    /*
        双方映射后的会员等级。如需开通，需要申请权限，取值范围为：1,2,3,4,5,6,none。比如飞猪F3对应商家V4,则传4.（如果有疑问请联系对接技术支持）     */
    MemberLevel  *string `json:"member_level,omitempty" required:"false" `
    /*
        销售渠道。如需开通，需要申请权限。目前支持的渠道有 H:飞猪全渠道（选择H，可实现飞猪、高德、支付宝、手淘均可售卖） O:钉钉商旅 。如果有多个用","分开，比如H,O。如果需要投放其他渠道，请联系飞猪运营或者技术支持。     */
    Channel  *string `json:"channel,omitempty" required:"false" `
    /*
        不推荐使用     */
    Occupancy  *int64 `json:"occupancy,omitempty" required:"false" `
    /*
        系统商，一般不用填写，使用须申请     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        商家价格政策编码     */
    RateplanCode  *string `json:"rateplan_code,omitempty" required:"false" `
    /*
        需申请会员权限。是否是新用户首住优惠rp。1-代表是。0-代表否。不填写代表不更新该字段。     */
    FirstStay  *int64 `json:"first_stay,omitempty" required:"false" `
    /*
        价格类型字段：0.非协议价；1.集采协议价；如果不是协议价，请不要填写该字段。该字段有权限控制，如需使用，请联系阿里旅行运营。 如果不填写或者填写为0，默认是阿里旅行价     */
    Agreement  *int64 `json:"agreement,omitempty" required:"false" `
    /*
        不推荐使用，见退改规则     */
    CancelBeforeDay  *int64 `json:"cancel_before_day,omitempty" required:"false" `
    /*
        不推荐使用，见退改规则     */
    CancelBeforeHour  *string `json:"cancel_before_hour,omitempty" required:"false" `
    /*
        在更新rateplan时，同时新增或更新早餐日历。 date：早餐政策属于具体哪一天 breakfast_count：这一天早餐的数量。>=0,<=99 如果date为空，那么会去读取startDate和endDate（格式都为"yyyy-MM-dd"），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。     */
    BreakfastCal  *string `json:"breakfast_cal,omitempty" required:"false" `
    /*
        在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为"yyyy-MM-dd" cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。     */
    CancelPolicyCal  *string `json:"cancel_policy_cal,omitempty" required:"false" `
    /*
        在更新rateplan的同时，新增或更新担保日历。 date：担保日历的某一天。 guarantee:担保政策。 其中有两个属性：guaranteeType,guaranteeStartTime。 guaranteeType的可选值同guaranteeType字段，详见guaranteeType字段。guaranteeStartTime格式为"HH:mm" 。如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即担保政策属于某一个时间段。（如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点） 请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。     */
    GuaranteeCal  *string `json:"guarantee_cal,omitempty" required:"false" `
    /*
        生效开始时间，用来控制此rateplan生效的开始时间，配合字段deadline_time一起限定rp的有效期     */
    EffectiveTime  *util.LocalTime `json:"effective_time,omitempty" required:"false" `
    /*
        生效截止时间，用来控制此rateplan生效的截止时间，配合字段effective_time一起限定rp的有效期     */
    DeadlineTime  *util.LocalTime `json:"deadline_time,omitempty" required:"false" `
    /*
        支付类型，只支持：1：预付5：现付6: 信用住7:在线预约8:在线预约信用住。其中5,6,7,8三种类型需要申请权限     */
    PaymentType  *int64 `json:"payment_type,omitempty" required:"false" `
    /*
        rp类型，1为小时房。目前只支持小时房。如果不是小时房rateplan,则不要填写。如果想要清空该字段可以传入none     */
    RpType  *string `json:"rp_type,omitempty" required:"false" `
    /*
        小时房入住时间跨度。小时房特有字段。比如4小时的小时房，那么值为4     */
    Hourage  *string `json:"hourage,omitempty" required:"false" `
    /*
        最晚可选入住时间，小时房特有字段。格式为HH:mm     */
    CanCheckinEnd  *string `json:"can_checkin_end,omitempty" required:"false" `
    /*
        最早可选入住时间，小时房特有字段。格式为HH:mm     */
    CanCheckinStart  *string `json:"can_checkin_start,omitempty" required:"false" `
    /*
        儿童最大年龄(0-18)     */
    MaxChildAge  *int64 `json:"max_child_age,omitempty" required:"false" `
    /*
        儿童最小年龄(0-18)     */
    MinChildAge  *int64 `json:"min_child_age,omitempty" required:"false" `
    /*
        婴儿最大年龄(0-18)     */
    MaxInfantAge  *int64 `json:"max_infant_age,omitempty" required:"false" `
    /*
        婴儿最小年龄(0-18)     */
    MinInfantAge  *int64 `json:"min_infant_age,omitempty" required:"false" `
    /*
        餐食描述     */
    DinningDesc  *string `json:"dinning_desc,omitempty" required:"false" `
    /*
        学生价，1：是；0：否     */
    IsStudent  *int64 `json:"is_student,omitempty" required:"false" `
    /*
        酒店id     */
    Hid  *int64 `json:"hid,omitempty" required:"false" `
    /*
        房型id     */
    Rid  *int64 `json:"rid,omitempty" required:"false" `
    /*
        外部房型id     */
    OutRid  *string `json:"out_rid,omitempty" required:"false" `
    /*
        外部酒店id     */
    OutHid  *string `json:"out_hid,omitempty" required:"false" `
    /*
        super rp标记，1是；0否     */
    SuperRpFlag  *int64 `json:"super_rp_flag,omitempty" required:"false" `
    /*
        base rp标记，1是；0否     */
    BaseRpFlag  *int64 `json:"base_rp_flag,omitempty" required:"false" `
    /*
        -9999 清空担保, 2 VCC担保, 1 PCI担保，0 支付宝担保(信用住产品担保方式只能为支付宝担保)     */
    GuaranteeMode  *int64 `json:"guarantee_mode,omitempty" required:"false" `
    /*
        operator     */
    Operator  *string `json:"operator,omitempty" required:"false" `
    /*
        父rpcode，使用场景：当一个rp发布变价rate的时候，用于下单时候传递约定的rpcode给外部     */
    ParentRpCode  *string `json:"parent_rp_code,omitempty" required:"false" `
    /*
        父rpid，使用场景：当一个rp发布变价rate的时候，记录父rp信息，用于下单时候传递约定的rpcode给外部     */
    ParentRpid  *int64 `json:"parent_rpid,omitempty" required:"false" `
    /*
        更新RP时的 打标和去标 需求, 0 就是 去标, 1 就是打标,    key的含义:    non-direct-RP 表示非直连RP,    super-could-price-change-RP 表示rp的super标，打上这个tag，表明这个rateplan下单的时候支持变价，商家系统直接放开价格校验。   base-could-derived-RP 表示base rateplan标签，打上了这个tag，表明这是一个base的rateplan，基于该rateplan可以衍生出子rateplan,       ebk-tail-room-RP 表示 ebk尾房rate plan级别标, free-room 表示免房商品     */
    TagJson  *string `json:"tag_json,omitempty" required:"false" `
    /*
        协议保留房提前x小时自动确认时间 比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认     */
    AllotmentReleaseTime  *string `json:"allotment_release_time,omitempty" required:"false" `
    /*
        是否包房RP 1包房RP,0 非包房rp     */
    PackRoomFlag  *string `json:"pack_room_flag,omitempty" required:"false" `
    /*
        是否底价加价，1是底价加价,0 非底价加价rp     */
    BottomPriceFlag  *string `json:"bottom_price_flag,omitempty" required:"false" `
    /*
        价格计划名称name通过加工处理以后的值     */
    DisplayName  *string `json:"display_name,omitempty" required:"false" `
    /*
        来源     */
    Source  *int64 `json:"source,omitempty" required:"false" `
    /*
        普通保留房提前x小时自动确认时间 比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认     */
    CommonAllotReleaseTime  *string `json:"common_allot_release_time,omitempty" required:"false" `
    /*
        是否企业托管RP 0-普通rp,1-企业托管rp     */
    CompanyAssist  *int64 `json:"company_assist,omitempty" required:"false" `
    /*
        酒店-企业-rp映射实体集合     */
    HotelCompanyMappingDOS  *string `json:"hotel_company_mapping_d_o_s,omitempty" required:"false" `
    /*
        商品来源渠道。1：直采（直连酒店PMS）, 1-1：直采（非直连） 2：携程系, 3：美团, 4：国内旅行社分销, 5：海外供应商。非酒店资源方卖家必须提供商品来源渠道，携程系包括携程、去哪儿、艺龙。     */
    ResourceType  *string `json:"resource_type,omitempty" required:"false" `
    /*
        最晚可选离店时间，小时房特有字段。格式为HH:mm     */
    CanCheckoutEnd  *string `json:"can_checkout_end,omitempty" required:"false" `
    /*
        会员价支持标识,1表示支持会员价规则     */
    MemDiscFlag  *int64 `json:"mem_disc_flag,omitempty" required:"false" `
    /*
        会员价加价规则。c:表示折扣百分比,例子8,意为会员价优惠8%,s:标识起始日期,e:表示截止日期，t:表示加价类型，0:代表折扣。会员价=售价*(1-c%)。该字段使用需要联系小二     */
    MemberDiscountCal  *string `json:"member_discount_cal,omitempty" required:"false" `
    /*
        卖点     */
    Benefits  *string `json:"benefits,omitempty" required:"false" `
    /*
        活动类型。1通兑 4超级房券 8直连免房     */
    ActivityType  *string `json:"activity_type,omitempty" required:"false" `
    /*
        入住人限制     */
    GuestLimit  *string `json:"guest_limit,omitempty" required:"false" `
    /*
        在线预约关联关系推送，priceRuleNumber：加价规则序号     */
    OnlineBookingBindingInfo  *string `json:"online_booking_binding_info,omitempty" required:"false" `
    /*
        rp的权益信息。1. 额外积分 2. 优惠价格 3. 套餐权益 4.行政礼遇     */
    Rights  *string `json:"rights,omitempty" required:"false" `
    /*
        商业化充值类型 seller充值到卖家 hotel充值到门店     */
    FreeRoomChargeDstRole  *string `json:"free_room_charge_dst_role,omitempty" required:"false" `
    /*
        儿童价格政策 年龄区间必须连续且有一个从0开始 max：年龄区间上限 min：年龄区间下限 t：加价类型，1-百分比金额加价，2-固定金额加价 加价因子，固定加价单位为分，百分比加价单位为百分比     */
    ChildrenPricePolicy  *string `json:"children_price_policy,omitempty" required:"false" `
}

func (s *TaobaoXhotelRateplanUpdateRequest) SetRpid(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.Rpid = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetName(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetEnglishName(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.EnglishName = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetBreakfastCount(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.BreakfastCount = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetFeeBreakfastCount(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.FeeBreakfastCount = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetFeeBreakfastAmount(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.FeeBreakfastAmount = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetFeeGovTaxAmount(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.FeeGovTaxAmount = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetFeeGovTaxPercent(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.FeeGovTaxPercent = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetFeeServiceAmount(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.FeeServiceAmount = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetFeeServicePercent(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.FeeServicePercent = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetExtendFee(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.ExtendFee = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetMinDays(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.MinDays = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetMaxDays(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.MaxDays = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetMinAmount(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.MinAmount = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetMinAdvHours(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.MinAdvHours = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetMaxAdvHours(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.MaxAdvHours = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetStartTime(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.StartTime = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetEndTime(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.EndTime = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetCancelPolicy(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.CancelPolicy = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetExtend(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetStatus(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetGuaranteeType(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.GuaranteeType = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetGuaranteeStartTime(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.GuaranteeStartTime = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetMemberLevel(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.MemberLevel = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetChannel(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.Channel = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetOccupancy(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.Occupancy = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetVendor(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetRateplanCode(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.RateplanCode = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetFirstStay(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.FirstStay = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetAgreement(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.Agreement = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetCancelBeforeDay(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.CancelBeforeDay = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetCancelBeforeHour(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.CancelBeforeHour = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetBreakfastCal(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.BreakfastCal = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetCancelPolicyCal(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.CancelPolicyCal = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetGuaranteeCal(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.GuaranteeCal = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetEffectiveTime(v util.LocalTime) *TaobaoXhotelRateplanUpdateRequest {
    s.EffectiveTime = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetDeadlineTime(v util.LocalTime) *TaobaoXhotelRateplanUpdateRequest {
    s.DeadlineTime = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetPaymentType(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.PaymentType = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetRpType(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.RpType = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetHourage(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.Hourage = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetCanCheckinEnd(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.CanCheckinEnd = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetCanCheckinStart(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.CanCheckinStart = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetMaxChildAge(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.MaxChildAge = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetMinChildAge(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.MinChildAge = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetMaxInfantAge(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.MaxInfantAge = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetMinInfantAge(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.MinInfantAge = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetDinningDesc(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.DinningDesc = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetIsStudent(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.IsStudent = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetHid(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetRid(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.Rid = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetOutRid(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetOutHid(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.OutHid = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetSuperRpFlag(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.SuperRpFlag = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetBaseRpFlag(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.BaseRpFlag = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetGuaranteeMode(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.GuaranteeMode = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetOperator(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.Operator = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetParentRpCode(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.ParentRpCode = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetParentRpid(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.ParentRpid = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetTagJson(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.TagJson = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetAllotmentReleaseTime(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.AllotmentReleaseTime = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetPackRoomFlag(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.PackRoomFlag = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetBottomPriceFlag(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.BottomPriceFlag = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetDisplayName(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.DisplayName = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetSource(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.Source = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetCommonAllotReleaseTime(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.CommonAllotReleaseTime = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetCompanyAssist(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.CompanyAssist = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetHotelCompanyMappingDOS(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.HotelCompanyMappingDOS = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetResourceType(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.ResourceType = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetCanCheckoutEnd(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.CanCheckoutEnd = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetMemDiscFlag(v int64) *TaobaoXhotelRateplanUpdateRequest {
    s.MemDiscFlag = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetMemberDiscountCal(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.MemberDiscountCal = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetBenefits(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.Benefits = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetActivityType(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.ActivityType = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetGuestLimit(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.GuestLimit = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetOnlineBookingBindingInfo(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.OnlineBookingBindingInfo = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetRights(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.Rights = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetFreeRoomChargeDstRole(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.FreeRoomChargeDstRole = &v
    return s
}
func (s *TaobaoXhotelRateplanUpdateRequest) SetChildrenPricePolicy(v string) *TaobaoXhotelRateplanUpdateRequest {
    s.ChildrenPricePolicy = &v
    return s
}

func (req *TaobaoXhotelRateplanUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Rpid != nil) {
        paramMap["rpid"] = *req.Rpid
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.EnglishName != nil) {
        paramMap["english_name"] = *req.EnglishName
    }
    if(req.BreakfastCount != nil) {
        paramMap["breakfast_count"] = *req.BreakfastCount
    }
    if(req.FeeBreakfastCount != nil) {
        paramMap["fee_breakfast_count"] = *req.FeeBreakfastCount
    }
    if(req.FeeBreakfastAmount != nil) {
        paramMap["fee_breakfast_amount"] = *req.FeeBreakfastAmount
    }
    if(req.FeeGovTaxAmount != nil) {
        paramMap["fee_gov_tax_amount"] = *req.FeeGovTaxAmount
    }
    if(req.FeeGovTaxPercent != nil) {
        paramMap["fee_gov_tax_percent"] = *req.FeeGovTaxPercent
    }
    if(req.FeeServiceAmount != nil) {
        paramMap["fee_service_amount"] = *req.FeeServiceAmount
    }
    if(req.FeeServicePercent != nil) {
        paramMap["fee_service_percent"] = *req.FeeServicePercent
    }
    if(req.ExtendFee != nil) {
        paramMap["extend_fee"] = *req.ExtendFee
    }
    if(req.MinDays != nil) {
        paramMap["min_days"] = *req.MinDays
    }
    if(req.MaxDays != nil) {
        paramMap["max_days"] = *req.MaxDays
    }
    if(req.MinAmount != nil) {
        paramMap["min_amount"] = *req.MinAmount
    }
    if(req.MinAdvHours != nil) {
        paramMap["min_adv_hours"] = *req.MinAdvHours
    }
    if(req.MaxAdvHours != nil) {
        paramMap["max_adv_hours"] = *req.MaxAdvHours
    }
    if(req.StartTime != nil) {
        paramMap["start_time"] = *req.StartTime
    }
    if(req.EndTime != nil) {
        paramMap["end_time"] = *req.EndTime
    }
    if(req.CancelPolicy != nil) {
        paramMap["cancel_policy"] = *req.CancelPolicy
    }
    if(req.Extend != nil) {
        paramMap["extend"] = *req.Extend
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.GuaranteeType != nil) {
        paramMap["guarantee_type"] = *req.GuaranteeType
    }
    if(req.GuaranteeStartTime != nil) {
        paramMap["guarantee_start_time"] = *req.GuaranteeStartTime
    }
    if(req.MemberLevel != nil) {
        paramMap["member_level"] = *req.MemberLevel
    }
    if(req.Channel != nil) {
        paramMap["channel"] = *req.Channel
    }
    if(req.Occupancy != nil) {
        paramMap["occupancy"] = *req.Occupancy
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.RateplanCode != nil) {
        paramMap["rateplan_code"] = *req.RateplanCode
    }
    if(req.FirstStay != nil) {
        paramMap["first_stay"] = *req.FirstStay
    }
    if(req.Agreement != nil) {
        paramMap["agreement"] = *req.Agreement
    }
    if(req.CancelBeforeDay != nil) {
        paramMap["cancel_before_day"] = *req.CancelBeforeDay
    }
    if(req.CancelBeforeHour != nil) {
        paramMap["cancel_before_hour"] = *req.CancelBeforeHour
    }
    if(req.BreakfastCal != nil) {
        paramMap["breakfast_cal"] = *req.BreakfastCal
    }
    if(req.CancelPolicyCal != nil) {
        paramMap["cancel_policy_cal"] = *req.CancelPolicyCal
    }
    if(req.GuaranteeCal != nil) {
        paramMap["guarantee_cal"] = *req.GuaranteeCal
    }
    if(req.EffectiveTime != nil) {
        paramMap["effective_time"] = *req.EffectiveTime
    }
    if(req.DeadlineTime != nil) {
        paramMap["deadline_time"] = *req.DeadlineTime
    }
    if(req.PaymentType != nil) {
        paramMap["payment_type"] = *req.PaymentType
    }
    if(req.RpType != nil) {
        paramMap["rp_type"] = *req.RpType
    }
    if(req.Hourage != nil) {
        paramMap["hourage"] = *req.Hourage
    }
    if(req.CanCheckinEnd != nil) {
        paramMap["can_checkin_end"] = *req.CanCheckinEnd
    }
    if(req.CanCheckinStart != nil) {
        paramMap["can_checkin_start"] = *req.CanCheckinStart
    }
    if(req.MaxChildAge != nil) {
        paramMap["max_child_age"] = *req.MaxChildAge
    }
    if(req.MinChildAge != nil) {
        paramMap["min_child_age"] = *req.MinChildAge
    }
    if(req.MaxInfantAge != nil) {
        paramMap["max_infant_age"] = *req.MaxInfantAge
    }
    if(req.MinInfantAge != nil) {
        paramMap["min_infant_age"] = *req.MinInfantAge
    }
    if(req.DinningDesc != nil) {
        paramMap["dinning_desc"] = *req.DinningDesc
    }
    if(req.IsStudent != nil) {
        paramMap["is_student"] = *req.IsStudent
    }
    if(req.Hid != nil) {
        paramMap["hid"] = *req.Hid
    }
    if(req.Rid != nil) {
        paramMap["rid"] = *req.Rid
    }
    if(req.OutRid != nil) {
        paramMap["out_rid"] = *req.OutRid
    }
    if(req.OutHid != nil) {
        paramMap["out_hid"] = *req.OutHid
    }
    if(req.SuperRpFlag != nil) {
        paramMap["super_rp_flag"] = *req.SuperRpFlag
    }
    if(req.BaseRpFlag != nil) {
        paramMap["base_rp_flag"] = *req.BaseRpFlag
    }
    if(req.GuaranteeMode != nil) {
        paramMap["guarantee_mode"] = *req.GuaranteeMode
    }
    if(req.Operator != nil) {
        paramMap["operator"] = *req.Operator
    }
    if(req.ParentRpCode != nil) {
        paramMap["parent_rp_code"] = *req.ParentRpCode
    }
    if(req.ParentRpid != nil) {
        paramMap["parent_rpid"] = *req.ParentRpid
    }
    if(req.TagJson != nil) {
        paramMap["tag_json"] = *req.TagJson
    }
    if(req.AllotmentReleaseTime != nil) {
        paramMap["allotment_release_time"] = *req.AllotmentReleaseTime
    }
    if(req.PackRoomFlag != nil) {
        paramMap["pack_room_flag"] = *req.PackRoomFlag
    }
    if(req.BottomPriceFlag != nil) {
        paramMap["bottom_price_flag"] = *req.BottomPriceFlag
    }
    if(req.DisplayName != nil) {
        paramMap["display_name"] = *req.DisplayName
    }
    if(req.Source != nil) {
        paramMap["source"] = *req.Source
    }
    if(req.CommonAllotReleaseTime != nil) {
        paramMap["common_allot_release_time"] = *req.CommonAllotReleaseTime
    }
    if(req.CompanyAssist != nil) {
        paramMap["company_assist"] = *req.CompanyAssist
    }
    if(req.HotelCompanyMappingDOS != nil) {
        paramMap["hotel_company_mapping_d_o_s"] = *req.HotelCompanyMappingDOS
    }
    if(req.ResourceType != nil) {
        paramMap["resource_type"] = *req.ResourceType
    }
    if(req.CanCheckoutEnd != nil) {
        paramMap["can_checkout_end"] = *req.CanCheckoutEnd
    }
    if(req.MemDiscFlag != nil) {
        paramMap["mem_disc_flag"] = *req.MemDiscFlag
    }
    if(req.MemberDiscountCal != nil) {
        paramMap["member_discount_cal"] = *req.MemberDiscountCal
    }
    if(req.Benefits != nil) {
        paramMap["benefits"] = *req.Benefits
    }
    if(req.ActivityType != nil) {
        paramMap["activity_type"] = *req.ActivityType
    }
    if(req.GuestLimit != nil) {
        paramMap["guest_limit"] = *req.GuestLimit
    }
    if(req.OnlineBookingBindingInfo != nil) {
        paramMap["online_booking_binding_info"] = *req.OnlineBookingBindingInfo
    }
    if(req.Rights != nil) {
        paramMap["rights"] = *req.Rights
    }
    if(req.FreeRoomChargeDstRole != nil) {
        paramMap["free_room_charge_dst_role"] = *req.FreeRoomChargeDstRole
    }
    if(req.ChildrenPricePolicy != nil) {
        paramMap["children_price_policy"] = *req.ChildrenPricePolicy
    }
    return paramMap
}

func (req *TaobaoXhotelRateplanUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}