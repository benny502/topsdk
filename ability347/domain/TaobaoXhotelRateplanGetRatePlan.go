package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelRateplanGetRatePlan struct {
    /*
        rateplan_id     */
    Rpid  *int64 `json:"rpid,omitempty" `

    /*
        卖家自己系统的Code，简称RateCode     */
    RateplanCode  *string `json:"rateplan_code,omitempty" `

    /*
        RP名称     */
    Name  *string `json:"name,omitempty" `

    /*
        英文名称     */
    EnglishName  *string `json:"english_name,omitempty" `

    /*
        支付类型 可选值 1：预付 5：前台面付     */
    PaymentType  *int64 `json:"payment_type,omitempty" `

    /*
        早餐数量     */
    BreakfastCount  *int64 `json:"breakfast_count,omitempty" `

    /*
        另加早餐数量     */
    FeeBreakfastCount  *int64 `json:"fee_breakfast_count,omitempty" `

    /*
        另加早餐金额     */
    FeeBreakfastAmount  *int64 `json:"fee_breakfast_amount,omitempty" `

    /*
        额外服务-政府税-金额（1-9999）     */
    FeeGovTaxAmount  *int64 `json:"fee_gov_tax_amount,omitempty" `

    /*
        额外服务-政府税-百分比（0%-99%）     */
    FeeGovTaxPercent  *int64 `json:"fee_gov_tax_percent,omitempty" `

    /*
        额外服务-服务费-金额（0-9999）     */
    FeeServiceAmount  *int64 `json:"fee_service_amount,omitempty" `

    /*
        额外服务-服务费-百分比（0%-99%）     */
    FeeServicePercent  *int64 `json:"fee_service_percent,omitempty" `

    /*
        额外服务的扩展，是一段JSON     */
    ExtendFee  *string `json:"extend_fee,omitempty" `

    /*
        最小入住天数（1-365）     */
    MinDays  *int64 `json:"min_days,omitempty" `

    /*
        最大入住天数（1-365）     */
    MaxDays  *int64 `json:"max_days,omitempty" `

    /*
        首日入住房间数（1-99）【废弃】     */
    MinAmount  *int64 `json:"min_amount,omitempty" `

    /*
        最小提前预订小时按入住时间的23:59:59(一般认为24点)来计算     */
    MinAdvHours  *int64 `json:"min_adv_hours,omitempty" `

    /*
        最大提前预订小时按入住时间的23:59:59(一般认为24点)来计算     */
    MaxAdvHours  *int64 `json:"max_adv_hours,omitempty" `

    /*
        每日生效时间默认00:00:00。生效时间＜结束时间     */
    StartTime  *string `json:"start_time,omitempty" `

    /*
        每日结束时间默认24:00:00。生效时间＜结束时间     */
    EndTime  *string `json:"end_time,omitempty" `

    /*
        退订政策     */
    CancelPolicy  *string `json:"cancel_policy,omitempty" `

    /*
        extend     */
    Extend  *string `json:"extend,omitempty" `

    /*
        状态     */
    Status  *int64 `json:"status,omitempty" `

    /*
        创建时间     */
    CreatedTime  *util.LocalTime `json:"created_time,omitempty" `

    /*
        修改时间     */
    ModifiedTime  *util.LocalTime `json:"modified_time,omitempty" `

    /*
        扩展字段1     */
    ExtendInfo1  *string `json:"extend_info1,omitempty" `

    /*
        扩展字段2     */
    ExtendInfo2  *string `json:"extend_info2,omitempty" `

    /*
        扩展字段3     */
    ExtendInfo3  *string `json:"extend_info3,omitempty" `

    /*
        担保类型，只支持： 0 无担保 1 峰时首晚担保 2峰时全额担保 3全天首晚担保 4全天全额担保     */
    GuaranteeType  *int64 `json:"guarantee_type,omitempty" `

    /*
        每日开始担保时间     */
    GuaranteeStartTime  *string `json:"guarantee_start_time,omitempty" `

    /*
        会员等级。支持多个等级","分隔     */
    MemberLevel  *string `json:"member_level,omitempty" `

    /*
        销售渠道。如需开通，需要申请权限。目前支持的渠道有 H:飞猪全渠道（选择H，可实现飞猪、高德、支付宝、手淘均可售卖） O:钉钉商旅 。如果有多个用","分开，比如H,O。如果需要投放其他渠道，请联系飞猪运营或者技术支持。     */
    Channel  *string `json:"channel,omitempty" `

    /*
        入住人数     */
    Occupancy  *int64 `json:"occupancy,omitempty" `

    /*
        是否是首住优惠rp。1代表是     */
    FirstStay  *int64 `json:"first_stay,omitempty" `

    /*
        是否是协议价。1代表是     */
    Agreement  *int64 `json:"agreement,omitempty" `

    /*
        rateplan生效开始时间     */
    EffectiveTime  *util.LocalTime `json:"effective_time,omitempty" `

    /*
        rateplan生效截止时间     */
    DeadlineTime  *util.LocalTime `json:"deadline_time,omitempty" `

    /*
        0支付宝担保 1PCI担保     */
    GuaranteeMode  *int64 `json:"guarantee_mode,omitempty" `

    /*
        协议保留房提前确认时间     */
    AllotmentReleaseTime  *string `json:"allotment_release_time,omitempty" `

    /*
        是否包房RP 1包房RP,0 非包房rp     */
    PackRoomFlag  *int64 `json:"pack_room_flag,omitempty" `

    /*
        是否底价加价，1是底价加价,0 非底价加价rp     */
    BottomPriceFlag  *int64 `json:"bottom_price_flag,omitempty" `

    /*
        是否为学生价     */
    IsStudent  *int64 `json:"is_student,omitempty" `

    /*
        rp维度的发票信息,type:1.酒店提供发票；2.卖家提供发票,desc:发票描叙，比如：卖家包邮提供发票,格式为：{"type":1;"desc":""}     */
    InvoiceContent  *string `json:"invoice_content,omitempty" `

    /*
        来源     */
    Source  *int64 `json:"source,omitempty" `

    /*
        key的含义:    non-direct-RP 表示非直连RP,      super-could-price-change-RP 表示rp的super标，打上这个tag，表明这个rateplan下单的时候支持变价，商家系统直接放开价格校验。      base-could-derived-RP 表示base rateplan标签，打上了这个tag，表明这是一个base的rateplan，基于该rateplan可以衍生出子rateplan  .        ebk-tail-room-RP 表示 ebk尾房rate plan级别标     */
    TagJson  *string `json:"tag_json,omitempty" `

    /*
        sell 端特殊RP 对应的 gid     */
    SellGid  *int64 `json:"sell_gid,omitempty" `

    /*
        可入住的最晚时间（小时房相关字段）     */
    CanCheckinEnd  *string `json:"can_checkin_end,omitempty" `

    /*
        可入住的最早时间（小时房相关字段）     */
    CanCheckinStart  *string `json:"can_checkin_start,omitempty" `

    /*
        每日生效结束时间（仅时分秒有效）     */
    EndTimeDaily  *util.LocalTime `json:"end_time_daily,omitempty" `

    /*
        rateplan类型 1为小时房     */
    RpType  *string `json:"rp_type,omitempty" `

    /*
        入住的开始跨度（小时房相关字段）     */
    Hourage  *string `json:"hourage,omitempty" `

    /*
        父rpid     */
    ParentRpid  *int64 `json:"parent_rpid,omitempty" `

    /*
        每日生效开始时间（仅时分秒有效）     */
    StartTimeDaily  *string `json:"start_time_daily,omitempty" `

    /*
        普通保留房提前确认时间     */
    CommonAllotReleaseTime  *string `json:"common_allot_release_time,omitempty" `

    /*
        companyAssist     */
    CompanyAssist  *int64 `json:"company_assist,omitempty" `

    /*
        hotelCompanyMappingDOS     */
    HotelCompanyMappingDOS  *string `json:"hotel_company_mapping_d_o_s,omitempty" `

    /*
        calBreakfastStr     */
    CalBreakfastStr  *string `json:"cal_breakfast_str,omitempty" `

    /*
        calGuaranteeStr     */
    CalGuaranteeStr  *string `json:"cal_guarantee_str,omitempty" `

    /*
        calChangeRuleStr     */
    CalChangeRuleStr  *string `json:"cal_change_rule_str,omitempty" `

    /*
        可离店的最晚时间（小时房相关字段）     */
    CanCheckoutEnd  *string `json:"can_checkout_end,omitempty" `

    /*
        会员价加价规则。c:表示折扣百分比,例子8,意为会员价优惠8%,s:标识起始日期,e:表示截止日期，t:表示加价类型，0:代表折扣。会员价=售价*(1-c%)     */
    MemberDiscountCal  *string `json:"member_discount_cal,omitempty" `

    /*
        会员价支持标识,1表示支持会员价规则     */
    MemDiscFlag  *int64 `json:"mem_disc_flag,omitempty" `

    /*
        酒+X特色     */
    Benefits  *string `json:"benefits,omitempty" `

    /*
        活动类型: 1通兑,2秒杀,3尾房,4超级房券     */
    ActivityType  *string `json:"activity_type,omitempty" `

    /*
        儿童价格政策 年龄区间必须连续且有一个从0开始 max：年龄区间上限 min：年龄区间下限 t：加价类型，1-百分比金额加价，2-固定金额加价 v：加价因子，固定加价单位为分，百分比加价单位为百分比     */
    ChildrenPricePolicy  *string `json:"children_price_policy,omitempty" `

}

func (s *TaobaoXhotelRateplanGetRatePlan) SetRpid(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.Rpid = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetRateplanCode(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.RateplanCode = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetName(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetEnglishName(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.EnglishName = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetPaymentType(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.PaymentType = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetBreakfastCount(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.BreakfastCount = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetFeeBreakfastCount(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.FeeBreakfastCount = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetFeeBreakfastAmount(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.FeeBreakfastAmount = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetFeeGovTaxAmount(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.FeeGovTaxAmount = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetFeeGovTaxPercent(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.FeeGovTaxPercent = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetFeeServiceAmount(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.FeeServiceAmount = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetFeeServicePercent(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.FeeServicePercent = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetExtendFee(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.ExtendFee = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetMinDays(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.MinDays = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetMaxDays(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.MaxDays = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetMinAmount(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.MinAmount = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetMinAdvHours(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.MinAdvHours = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetMaxAdvHours(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.MaxAdvHours = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetStartTime(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.StartTime = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetEndTime(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.EndTime = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetCancelPolicy(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.CancelPolicy = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetExtend(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetStatus(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetCreatedTime(v util.LocalTime) *TaobaoXhotelRateplanGetRatePlan {
    s.CreatedTime = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetModifiedTime(v util.LocalTime) *TaobaoXhotelRateplanGetRatePlan {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetExtendInfo1(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.ExtendInfo1 = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetExtendInfo2(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.ExtendInfo2 = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetExtendInfo3(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.ExtendInfo3 = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetGuaranteeType(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.GuaranteeType = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetGuaranteeStartTime(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.GuaranteeStartTime = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetMemberLevel(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.MemberLevel = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetChannel(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.Channel = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetOccupancy(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.Occupancy = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetFirstStay(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.FirstStay = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetAgreement(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.Agreement = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetEffectiveTime(v util.LocalTime) *TaobaoXhotelRateplanGetRatePlan {
    s.EffectiveTime = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetDeadlineTime(v util.LocalTime) *TaobaoXhotelRateplanGetRatePlan {
    s.DeadlineTime = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetGuaranteeMode(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.GuaranteeMode = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetAllotmentReleaseTime(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.AllotmentReleaseTime = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetPackRoomFlag(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.PackRoomFlag = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetBottomPriceFlag(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.BottomPriceFlag = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetIsStudent(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.IsStudent = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetInvoiceContent(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.InvoiceContent = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetSource(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.Source = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetTagJson(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.TagJson = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetSellGid(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.SellGid = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetCanCheckinEnd(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.CanCheckinEnd = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetCanCheckinStart(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.CanCheckinStart = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetEndTimeDaily(v util.LocalTime) *TaobaoXhotelRateplanGetRatePlan {
    s.EndTimeDaily = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetRpType(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.RpType = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetHourage(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.Hourage = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetParentRpid(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.ParentRpid = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetStartTimeDaily(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.StartTimeDaily = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetCommonAllotReleaseTime(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.CommonAllotReleaseTime = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetCompanyAssist(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.CompanyAssist = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetHotelCompanyMappingDOS(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.HotelCompanyMappingDOS = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetCalBreakfastStr(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.CalBreakfastStr = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetCalGuaranteeStr(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.CalGuaranteeStr = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetCalChangeRuleStr(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.CalChangeRuleStr = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetCanCheckoutEnd(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.CanCheckoutEnd = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetMemberDiscountCal(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.MemberDiscountCal = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetMemDiscFlag(v int64) *TaobaoXhotelRateplanGetRatePlan {
    s.MemDiscFlag = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetBenefits(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.Benefits = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetActivityType(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.ActivityType = &v
    return s
}
func (s *TaobaoXhotelRateplanGetRatePlan) SetChildrenPricePolicy(v string) *TaobaoXhotelRateplanGetRatePlan {
    s.ChildrenPricePolicy = &v
    return s
}
