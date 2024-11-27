package request

import (
        "topsdk/util"
    )

type TaobaoXhotelRateplanAddRequest struct {
    /*
        卖家自己系统的Code，简称RateCode     */
    RateplanCode  *string `json:"rateplan_code" required:"true" `
    /*
        在淘宝搜索页面展示的房价名称。请注意名称里不要维护早餐信息，如果想设置早餐信息，请设置breakfast_count字段即可     */
    Name  *string `json:"name" required:"true" `
    /*
        RP的英文名称     */
    EnglishName  *string `json:"english_name,omitempty" required:"false" `
    /*
        支付类型，只支持：1：预付5：现付6: 信用住7:预付在线预约8:信用住在线预约。其中5,6,7,8四种类型需要申请权限     */
    PaymentType  *int64 `json:"payment_type" required:"true" `
    /*
        -1：状态早餐,有具体几人价有关系，几人价是几份早餐;0：不含早1：含单早2：含双早N：含N早（-1-99可选）     */
    BreakfastCount  *int64 `json:"breakfast_count" required:"true" `
    /*
        废弃     */
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
        最小入住天数（1-90）。默认1,小时房RP请设置为1 defalutValue��1    */
    MinDays  *int64 `json:"min_days,omitempty" required:"false" `
    /*
        最大入住天数（1-90）。默认90 信用住不超过9天,小时房RP请设置为1,特殊商家支持90天 defalutValue��90    */
    MaxDays  *int64 `json:"max_days,omitempty" required:"false" `
    /*
        首日入住房间数（1-99）。默认1。【废弃】 defalutValue��1    */
    MinAmount  *int64 `json:"min_amount,omitempty" required:"false" `
    /*
        最小提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表必须至少提前两天预定，那么如果想预定24号入住，,必须在23号零点前下单。 defalutValue��0    */
    MinAdvHours  *int64 `json:"min_adv_hours,omitempty" required:"false" `
    /*
        最大提前预定小时数，从入住当天的24点往前计算。例如如果这个字段设置了48，代表最多提前两天预定，那么如果想预定24号入住，,必须在23号零点以后下单。 defalutValue��2160    */
    MaxAdvHours  *int64 `json:"max_adv_hours,omitempty" required:"false" `
    /*
        产品每日开始销售时间，start_time一定为当天时间 defalutValue��00:00    */
    StartTime  *string `json:"start_time,omitempty" required:"false" `
    /*
        产品每日结束销售时间,当end_time<start_time时，表示end_time为第二天，此时附加限制end_time<=06:00:00并且start_time>=12:00:00,表明可售时间从当天12点到次日的凌晨6点（扩展此信息主要为了描述尾房的rp）注意start_time一定是当天的时间。尾房18：00起可售 defalutValue��00:00    */
    EndTime  *string `json:"end_time,omitempty" required:"false" `
    /*
        退订政策字段，是个json串，参考示例值设置改字段的值。允许变更/取消：在XX年XX月XX日XX时前取消收取Y%的手续费，100>Y>=0允许变更/取消：在入住前X小时前取消收取Y%的手续费，100>Y>=0（不超过10条）。1.表示任意退{"cancelPolicyType":1};2.表示不能退{"cancelPolicyType":2};4.从入住当天24点往前推X小时前取消收取Y%手续费，否则不可取消{"cancelPolicyType":4,"policyInfo":{"48":10,"24":20}}表示，从入住日24点往前推提前至少48小时取消，收取10%的手续费，从入住日24点往前推提前至少24小时取消，收取20%的手续费;5.从24点往前推多少小时可退{"cancelPolicyType":5,"policyInfo":{"timeBefore":6}}表示从入住日24点往前推至少6个小时即入住日18点前可免费取消;6.从入住日24点往前推，至少提前小时数扣取首晚房费{"cancelPolicyType":6,"policyInfo":{"14":1}}表示入住日24点往前推14小时，即入住日10点前取消收取首晚房费。 注意：支付类型为预付，那么可以使用所有的退订类型,但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。支持多段时间,多间夜扣款。阶梯退手续费限制请查看https://hot.bbs.taobao.com/detail.html?postId=8892814     */
    CancelPolicy  *string `json:"cancel_policy" required:"true" `
    /*
        个性化定制扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析     */
    Extend  *string `json:"extend,omitempty" required:"false" `
    /*
        1：开启（默认）2：关闭。如果没传值那么默认默认值为1     */
    Status  *int64 `json:"status" required:"true" `
    /*
        担保类型，只支持： 0  无担保  1  峰时首晚担保 2峰时全额担保 3全天首晚担保 4全天全额担保     */
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
        不推送则默认2人，如有低于2人的RP限制请推送该字段。     */
    Occupancy  *int64 `json:"occupancy,omitempty" required:"false" `
    /*
        系统商，一般不填写，使用须申请     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        需申请会员权限。是否是新用户首住优惠rp。1-代表是。0或者不填写代表否     */
    FirstStay  *int64 `json:"first_stay,omitempty" required:"false" `
    /*
        废弃。价格类型字段：0.非协议价；1.集采协议价；如果不是协议价，请不要填写该字段。该字段有权限控制，如需使用，请联系阿里旅行运营。 如果不填写或者填写为0，默认是阿里旅行价     */
    Agreement  *int64 `json:"agreement,omitempty" required:"false" `
    /*
        不推荐使用，使用改规则     */
    CancelBeforeDay  *int64 `json:"cancel_before_day,omitempty" required:"false" `
    /*
        不推荐使用，使用改规则     */
    CancelBeforeHour  *string `json:"cancel_before_hour,omitempty" required:"false" `
    /*
        在添加rateplan时，同时新增早餐日历。date：说明这条记录的早餐政策breakfast_count：这一天早餐的数量。>=-1,<=99。如果date为空，那么会去读取startDate和endDate（格式都为"yyyy-MM-dd"），即早餐正常属于一个时间段。-1为状态早餐，和最终绑定的几人价有关，如果是一人价那么就是我一份早餐，二人价就是两份早餐。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。     */
    BreakfastCal  *string `json:"breakfast_cal,omitempty" required:"false" `
    /*
        在新增rateplan的同时新增取消政策日历。 json格式。 date：日历的某一天，格式为"yyyy-MM-dd" cancel_policy：日历某一天的价格政策。格式和限制同cancel_policy。 如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即取消政策属于某一个时间段。 注意：支付类型为预付，那么可以使用所有的退订类型，但是必须是非担保；支付类型为面付或者信任住并且是无担保，那么只能使用1类型的退订；支付类型为面付或者信任住并且为担保，那么只能使用2,5类型的退订；支付类型为在线预约，那么只能使用1,2,5类型的退改。如果支付类型是面付或者信任住并且为担保，那么如果传了4或者6的退订，那么会强制转成类型5。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。阶梯退手续费限制请查看https://hot.bbs.taobao.com/detail.html?postId=8892814     */
    CancelPolicyCal  *string `json:"cancel_policy_cal,omitempty" required:"false" `
    /*
        在新增rateplan的同时，新增担保日历。date：担保日历的某一天。guarantee:担保政策。其中有两个属性：guaranteeType,guaranteeStartTime。 guaranteeType的可选值同guaranteeType字段，详见guaranteeType字段。guaranteeStartTime格式为"HH:mm"。如果date为空，那么会读取startDate和endDate（格式都为"yyyy-MM-dd"），即担保政策属于某一个时间段。（如果设置了峰时担保类型，那么峰时担保时间不能为空，并且必须大于等于8点）。请注意，该字段仅能维护从当前时间开始，10年以内的数据，如果超过10年，会报错。     */
    GuaranteeCal  *string `json:"guarantee_cal,omitempty" required:"false" `
    /*
        生效开始时间，用来控制此rateplan生效的开始时间，配合字段deadline_time一起限定rp的有效期     */
    EffectiveTime  *util.LocalTime `json:"effective_time,omitempty" required:"false" `
    /*
        生效截止时间，用来控制此rateplan生效的截止时间，配合字段effective_time一起限定rp的有效期     */
    DeadlineTime  *util.LocalTime `json:"deadline_time,omitempty" required:"false" `
    /*
        rp类型，1为小时房。目前只支持小时房。如果不是小时房rateplan,则不要填写。如果想要清空该字段可以传入none     */
    RpType  *string `json:"rp_type,omitempty" required:"false" `
    /*
        小时房入住时间跨度。小时房特有字段。比如4小时的小时房，那么值为4     */
    Hourage  *string `json:"hourage,omitempty" required:"false" `
    /*
        最早可选入住时间，小时房特有字段。格式为HH:mm     */
    CanCheckinEnd  *string `json:"can_checkin_end,omitempty" required:"false" `
    /*
        最晚可选入住时间，小时房特有字段。格式为HH:mm     */
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
        是否学生价，0：否；1：是。     */
    IsStudent  *int64 `json:"is_student,omitempty" required:"false" `
    /*
        酒店id     */
    Hid  *int64 `json:"hid,omitempty" required:"false" `
    /*
        房型id     */
    Rid  *int64 `json:"rid,omitempty" required:"false" `
    /*
        外部酒店id     */
    OutHid  *string `json:"out_hid,omitempty" required:"false" `
    /*
        外部房型id     */
    OutRid  *string `json:"out_rid,omitempty" required:"false" `
    /*
        super rp标记，1是；0否     */
    SuperRpFlag  *int64 `json:"super_rp_flag,omitempty" required:"false" `
    /*
        base rp标记，1是；0否     */
    BaseRpFlag  *int64 `json:"base_rp_flag,omitempty" required:"false" `
    /*
        2 VCC担保 1 PCI担保 0 支付宝担保(信用住产品担保方式只能为支付宝担保)     */
    GuaranteeMode  *int64 `json:"guarantee_mode,omitempty" required:"false" `
    /*
        父rpid,使用场景：当一个rp发布变价rate的时候，记录父rp信息，用于下单时候传递约定的rpcode给外部     */
    ParentRpCode  *string `json:"parent_rp_code,omitempty" required:"false" `
    /*
        父rpcode，使用场景：当一个rp发布变价rate的时候，用于下单时候传递约定的rpcode给外部     */
    ParentRpid  *int64 `json:"parent_rpid,omitempty" required:"false" `
    /*
        操作rateplan的来源     */
    Operator  *string `json:"operator,omitempty" required:"false" `
    /*
        新增RP时的 打标和去标 需求,     */
    TagJson  *string `json:"tag_json,omitempty" required:"false" `
    /*
        来源     */
    Source  *int64 `json:"source,omitempty" required:"false" `
    /*
        保留房提前x小时自动确认时间，比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认     */
    AllotmentReleaseTime  *string `json:"allotment_release_time,omitempty" required:"false" `
    /*
        普通保留房提前x小时自动确认时间，比如设置为6 那么从入住当日24点往前推6小时即18:00以前可以自动确认有房，否则是待确认     */
    CommonAllotReleaseTime  *string `json:"common_allot_release_time,omitempty" required:"false" `
    /*
        商品来源渠道。1：直采（直连酒店PMS）, 1-1：直采（非直连） 2：携程系, 3：美团, 4：国内旅行社分销, 5：海外供应商。非酒店资源方卖家必须提供商品来源渠道，携程系包括携程、去哪儿、艺龙。     */
    ResourceType  *string `json:"resource_type,omitempty" required:"false" `
    /*
        是否底价加价，1是底价加价,0 非底价加价rp     */
    BottomPriceFlag  *int64 `json:"bottom_price_flag,omitempty" required:"false" `
    /*
        最晚可选离店时间，小时房特有字段。格式为HH:mm     */
    CanCheckoutEnd  *string `json:"can_checkout_end,omitempty" required:"false" `
    /*
        会员价支持标识,1表示支持会员价规则     */
    MemDiscFlag  *int64 `json:"mem_disc_flag,omitempty" required:"false" `
    /*
        会员价加价规则。c:表示折扣百分比,例子8,意为会员价优惠8%,s:标识起始日期,e:表示截止日期，t:表示加价类型，0:代表折扣。会员价=售价*(1-c%)     */
    MemberDiscountCal  *string `json:"member_discount_cal,omitempty" required:"false" `
    /*
        RP入住人限制信息。JSON格式     */
    GuestLimit  *string `json:"guest_limit,omitempty" required:"false" `
    /*
        RP参与的活动，3为尾房,4超级房券,8直连免房     */
    ActivityType  *string `json:"activity_type,omitempty" required:"false" `
    /*
        在线预约关联关系推送，priceRuleNumber：加价规则序号     */
    OnlineBookingBindingInfo  *string `json:"online_booking_binding_info,omitempty" required:"false" `
    /*
        rp的权益信息, 调用该字段请优先联系对接业务同学。type枚举: eeo,meo, value取值:1. 额外积分 2. 优惠价格 3. 套餐权益 4.行政礼遇。     */
    Rights  *string `json:"rights,omitempty" required:"false" `
    /*
        商业化充值类型 seller充值到卖家 hotel充值到门店     */
    FreeRoomChargeDstRole  *string `json:"free_room_charge_dst_role,omitempty" required:"false" `
    /*
        儿童价格政策 年龄区间必须连续且有一个从0开始 max：年龄区间上限 min：年龄区间下限 t：加价类型，1-百分比金额加价，2-固定金额加价 v：加价因子，固定加价单位为分，百分比加价单位为百分比     */
    ChildrenPricePolicy  *string `json:"children_price_policy,omitempty" required:"false" `
}

func (s *TaobaoXhotelRateplanAddRequest) SetRateplanCode(v string) *TaobaoXhotelRateplanAddRequest {
    s.RateplanCode = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetName(v string) *TaobaoXhotelRateplanAddRequest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetEnglishName(v string) *TaobaoXhotelRateplanAddRequest {
    s.EnglishName = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetPaymentType(v int64) *TaobaoXhotelRateplanAddRequest {
    s.PaymentType = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetBreakfastCount(v int64) *TaobaoXhotelRateplanAddRequest {
    s.BreakfastCount = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetFeeBreakfastCount(v int64) *TaobaoXhotelRateplanAddRequest {
    s.FeeBreakfastCount = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetFeeBreakfastAmount(v int64) *TaobaoXhotelRateplanAddRequest {
    s.FeeBreakfastAmount = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetFeeGovTaxAmount(v int64) *TaobaoXhotelRateplanAddRequest {
    s.FeeGovTaxAmount = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetFeeGovTaxPercent(v int64) *TaobaoXhotelRateplanAddRequest {
    s.FeeGovTaxPercent = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetFeeServiceAmount(v int64) *TaobaoXhotelRateplanAddRequest {
    s.FeeServiceAmount = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetFeeServicePercent(v int64) *TaobaoXhotelRateplanAddRequest {
    s.FeeServicePercent = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetExtendFee(v string) *TaobaoXhotelRateplanAddRequest {
    s.ExtendFee = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetMinDays(v int64) *TaobaoXhotelRateplanAddRequest {
    s.MinDays = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetMaxDays(v int64) *TaobaoXhotelRateplanAddRequest {
    s.MaxDays = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetMinAmount(v int64) *TaobaoXhotelRateplanAddRequest {
    s.MinAmount = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetMinAdvHours(v int64) *TaobaoXhotelRateplanAddRequest {
    s.MinAdvHours = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetMaxAdvHours(v int64) *TaobaoXhotelRateplanAddRequest {
    s.MaxAdvHours = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetStartTime(v string) *TaobaoXhotelRateplanAddRequest {
    s.StartTime = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetEndTime(v string) *TaobaoXhotelRateplanAddRequest {
    s.EndTime = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetCancelPolicy(v string) *TaobaoXhotelRateplanAddRequest {
    s.CancelPolicy = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetExtend(v string) *TaobaoXhotelRateplanAddRequest {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetStatus(v int64) *TaobaoXhotelRateplanAddRequest {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetGuaranteeType(v int64) *TaobaoXhotelRateplanAddRequest {
    s.GuaranteeType = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetGuaranteeStartTime(v string) *TaobaoXhotelRateplanAddRequest {
    s.GuaranteeStartTime = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetMemberLevel(v string) *TaobaoXhotelRateplanAddRequest {
    s.MemberLevel = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetChannel(v string) *TaobaoXhotelRateplanAddRequest {
    s.Channel = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetOccupancy(v int64) *TaobaoXhotelRateplanAddRequest {
    s.Occupancy = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetVendor(v string) *TaobaoXhotelRateplanAddRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetFirstStay(v int64) *TaobaoXhotelRateplanAddRequest {
    s.FirstStay = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetAgreement(v int64) *TaobaoXhotelRateplanAddRequest {
    s.Agreement = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetCancelBeforeDay(v int64) *TaobaoXhotelRateplanAddRequest {
    s.CancelBeforeDay = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetCancelBeforeHour(v string) *TaobaoXhotelRateplanAddRequest {
    s.CancelBeforeHour = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetBreakfastCal(v string) *TaobaoXhotelRateplanAddRequest {
    s.BreakfastCal = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetCancelPolicyCal(v string) *TaobaoXhotelRateplanAddRequest {
    s.CancelPolicyCal = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetGuaranteeCal(v string) *TaobaoXhotelRateplanAddRequest {
    s.GuaranteeCal = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetEffectiveTime(v util.LocalTime) *TaobaoXhotelRateplanAddRequest {
    s.EffectiveTime = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetDeadlineTime(v util.LocalTime) *TaobaoXhotelRateplanAddRequest {
    s.DeadlineTime = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetRpType(v string) *TaobaoXhotelRateplanAddRequest {
    s.RpType = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetHourage(v string) *TaobaoXhotelRateplanAddRequest {
    s.Hourage = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetCanCheckinEnd(v string) *TaobaoXhotelRateplanAddRequest {
    s.CanCheckinEnd = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetCanCheckinStart(v string) *TaobaoXhotelRateplanAddRequest {
    s.CanCheckinStart = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetMaxChildAge(v int64) *TaobaoXhotelRateplanAddRequest {
    s.MaxChildAge = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetMinChildAge(v int64) *TaobaoXhotelRateplanAddRequest {
    s.MinChildAge = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetMaxInfantAge(v int64) *TaobaoXhotelRateplanAddRequest {
    s.MaxInfantAge = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetMinInfantAge(v int64) *TaobaoXhotelRateplanAddRequest {
    s.MinInfantAge = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetDinningDesc(v string) *TaobaoXhotelRateplanAddRequest {
    s.DinningDesc = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetIsStudent(v int64) *TaobaoXhotelRateplanAddRequest {
    s.IsStudent = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetHid(v int64) *TaobaoXhotelRateplanAddRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetRid(v int64) *TaobaoXhotelRateplanAddRequest {
    s.Rid = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetOutHid(v string) *TaobaoXhotelRateplanAddRequest {
    s.OutHid = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetOutRid(v string) *TaobaoXhotelRateplanAddRequest {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetSuperRpFlag(v int64) *TaobaoXhotelRateplanAddRequest {
    s.SuperRpFlag = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetBaseRpFlag(v int64) *TaobaoXhotelRateplanAddRequest {
    s.BaseRpFlag = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetGuaranteeMode(v int64) *TaobaoXhotelRateplanAddRequest {
    s.GuaranteeMode = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetParentRpCode(v string) *TaobaoXhotelRateplanAddRequest {
    s.ParentRpCode = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetParentRpid(v int64) *TaobaoXhotelRateplanAddRequest {
    s.ParentRpid = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetOperator(v string) *TaobaoXhotelRateplanAddRequest {
    s.Operator = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetTagJson(v string) *TaobaoXhotelRateplanAddRequest {
    s.TagJson = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetSource(v int64) *TaobaoXhotelRateplanAddRequest {
    s.Source = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetAllotmentReleaseTime(v string) *TaobaoXhotelRateplanAddRequest {
    s.AllotmentReleaseTime = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetCommonAllotReleaseTime(v string) *TaobaoXhotelRateplanAddRequest {
    s.CommonAllotReleaseTime = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetResourceType(v string) *TaobaoXhotelRateplanAddRequest {
    s.ResourceType = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetBottomPriceFlag(v int64) *TaobaoXhotelRateplanAddRequest {
    s.BottomPriceFlag = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetCanCheckoutEnd(v string) *TaobaoXhotelRateplanAddRequest {
    s.CanCheckoutEnd = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetMemDiscFlag(v int64) *TaobaoXhotelRateplanAddRequest {
    s.MemDiscFlag = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetMemberDiscountCal(v string) *TaobaoXhotelRateplanAddRequest {
    s.MemberDiscountCal = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetGuestLimit(v string) *TaobaoXhotelRateplanAddRequest {
    s.GuestLimit = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetActivityType(v string) *TaobaoXhotelRateplanAddRequest {
    s.ActivityType = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetOnlineBookingBindingInfo(v string) *TaobaoXhotelRateplanAddRequest {
    s.OnlineBookingBindingInfo = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetRights(v string) *TaobaoXhotelRateplanAddRequest {
    s.Rights = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetFreeRoomChargeDstRole(v string) *TaobaoXhotelRateplanAddRequest {
    s.FreeRoomChargeDstRole = &v
    return s
}
func (s *TaobaoXhotelRateplanAddRequest) SetChildrenPricePolicy(v string) *TaobaoXhotelRateplanAddRequest {
    s.ChildrenPricePolicy = &v
    return s
}

func (req *TaobaoXhotelRateplanAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RateplanCode != nil) {
        paramMap["rateplan_code"] = *req.RateplanCode
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.EnglishName != nil) {
        paramMap["english_name"] = *req.EnglishName
    }
    if(req.PaymentType != nil) {
        paramMap["payment_type"] = *req.PaymentType
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
    if(req.OutHid != nil) {
        paramMap["out_hid"] = *req.OutHid
    }
    if(req.OutRid != nil) {
        paramMap["out_rid"] = *req.OutRid
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
    if(req.ParentRpCode != nil) {
        paramMap["parent_rp_code"] = *req.ParentRpCode
    }
    if(req.ParentRpid != nil) {
        paramMap["parent_rpid"] = *req.ParentRpid
    }
    if(req.Operator != nil) {
        paramMap["operator"] = *req.Operator
    }
    if(req.TagJson != nil) {
        paramMap["tag_json"] = *req.TagJson
    }
    if(req.Source != nil) {
        paramMap["source"] = *req.Source
    }
    if(req.AllotmentReleaseTime != nil) {
        paramMap["allotment_release_time"] = *req.AllotmentReleaseTime
    }
    if(req.CommonAllotReleaseTime != nil) {
        paramMap["common_allot_release_time"] = *req.CommonAllotReleaseTime
    }
    if(req.ResourceType != nil) {
        paramMap["resource_type"] = *req.ResourceType
    }
    if(req.BottomPriceFlag != nil) {
        paramMap["bottom_price_flag"] = *req.BottomPriceFlag
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
    if(req.GuestLimit != nil) {
        paramMap["guest_limit"] = *req.GuestLimit
    }
    if(req.ActivityType != nil) {
        paramMap["activity_type"] = *req.ActivityType
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

func (req *TaobaoXhotelRateplanAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}