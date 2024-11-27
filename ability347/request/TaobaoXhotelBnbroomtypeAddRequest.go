package request

import (
        "topsdk/ability347/domain"
        "topsdk/util"
    )

type TaobaoXhotelBnbroomtypeAddRequest struct {
    /*
        单间面积，单位平方米     */
    RentSize  *int64 `json:"rent_size,omitempty" required:"false" `
    /*
        是否支持IM聊天 0不支持 1支持     */
    SupportIm  *int64 `json:"support_im,omitempty" required:"false" `
    /*
        房源外部标签 标签信息，逗号(,)分隔，最多1000字符     */
    OuterTags  *string `json:"outer_tags,omitempty" required:"false" `
    /*
        清洁费是否收取 0：否 1：是     */
    CleaningCharge  *int64 `json:"cleaning_charge,omitempty" required:"false" `
    /*
        发票，0：卖家提供发票，1：房东提供发票     */
    Invoice  *int64 `json:"invoice" required:"true" `
    /*
        装修等级 1 精装；2普通；3简装     */
    DecorateLevel  *int64 `json:"decorate_level,omitempty" required:"false" `
    /*
        民宿入住要求&附加信息     */
    BnbBookingTime  *domain.TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto `json:"bnb_booking_time,omitempty" required:"false" `
    /*
        外部门店id     */
    OutHid  *string `json:"out_hid" required:"true" `
    /*
        酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886     */
    Tel  *string `json:"tel" required:"true" `
    /*
        是否可接待外宾 0：否 1：是；默认值: 0 defalutValue��0    */
    ReceiveForeigners  *int64 `json:"receive_foreigners,omitempty" required:"false" `
    /*
        位置描述     */
    LocalInfo  *string `json:"local_info,omitempty" required:"false" `
    /*
        品牌名称，最多100字符     */
    Brand  *string `json:"brand,omitempty" required:"false" `
    /*
        房源图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。     */
    Pics  *[]domain.TaobaoXhotelBnbroomtypeAddBnbPictureDTO `json:"pics" required:"true" `
    /*
        清洁费类型 0.线下；1.线上     */
    CleaningType  *int64 `json:"cleaning_type,omitempty" required:"false" `
    /*
        押金金额     */
    DepositAmount  *int64 `json:"deposit_amount,omitempty" required:"false" `
    /*
        房源英文名     */
    NameE  *string `json:"name_e,omitempty" required:"false" `
    /*
        装修时间，格式为2015-01-01装修时间     */
    DecorateTime  *string `json:"decorate_time,omitempty" required:"false" `
    /*
        0-n；若不可加床，值为0     */
    ExtraBedsNum  *int64 `json:"extra_beds_num,omitempty" required:"false" `
    /*
        可提供发票类型，1.专票 2.纸质普票 3.电子普票     */
    InvoiceType  *int64 `json:"invoice_type" required:"true" `
    /*
        是否有前台 0没有 1有     */
    HasFrontDesk  *int64 `json:"has_front_desk,omitempty" required:"false" `
    /*
        是否接待儿童、老人；成年人必接待，详见“可接待客人”https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1     */
    GuestAge  *int64 `json:"guest_age,omitempty" required:"false" `
    /*
        结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用     */
    SettlementCurrency  *string `json:"settlement_currency,omitempty" required:"false" `
    /*
        房源名     */
    Name  *string `json:"name" required:"true" `
    /*
        周边描述     */
    SurroundInfo  *string `json:"surround_info,omitempty" required:"false" `
    /*
        最大入住人数 1-99     */
    MaxOccupancy  *int64 `json:"max_occupancy,omitempty" required:"false" `
    /*
        是否使用实拍图片 0不使用 1使用     */
    IsUseShootImage  *int64 `json:"is_use_shoot_image,omitempty" required:"false" `
    /*
        状态 0：在线 -1：不在线 -2:停售     */
    Status  *byte `json:"status" required:"true" `
    /*
        0：不限制，1：只限男性，2：只限女性'     */
    GuestGender  *int64 `json:"guest_gender,omitempty" required:"false" `
    /*
        详见“允许活动”：https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1  ，最多500字符     */
    ActivitiesAllowed  *string `json:"activities_allowed,omitempty" required:"false" `
    /*
        清洁费金额；整数[1，9999999]     */
    ExtraCleaningCharge  *int64 `json:"extra_cleaning_charge,omitempty" required:"false" `
    /*
        开业时间，格式为2015-01-01     */
    OpeningTime  *string `json:"opening_time,omitempty" required:"false" `
    /*
        出租类型，1整租；2分租。3床位 默认整租，该字段不能更新     */
    RentType  *int64 `json:"rent_type" required:"true" `
    /*
        如果要变更商品房源编码请使用该字段。     */
    NewOuterId  *string `json:"new_outer_id,omitempty" required:"false" `
    /*
        房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房,bedroom和livingroom不能为空     */
    HouseModel  *string `json:"house_model" required:"true" `
    /*
        0-无窗  1-有窗  2-部分有窗 3-暗窗 4-部分暗窗  5-落地窗     */
    WindowType  *int64 `json:"window_type,omitempty" required:"false" `
    /*
        有无资质执照 0 没有 1有     */
    HasLicense  *int64 `json:"has_license" required:"true" `
    /*
        视频地址，最多1000字符     */
    VideoUrl  *string `json:"video_url,omitempty" required:"false" `
    /*
        销售渠道,默认taobao     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        亮点描述，最多1000字符     */
    Brightspot  *string `json:"brightspot,omitempty" required:"false" `
    /*
        是否开启闪订 0不开启 1开启     */
    QuickOrder  *int64 `json:"quick_order,omitempty" required:"false" `
    /*
        客房在建筑的第几层，隔层为1-2层，4-5层，7-8层     */
    Floor  *string `json:"floor,omitempty" required:"false" `
    /*
        单间面积，单位平方米     */
    HouseSize  *int64 `json:"house_size" required:"true" `
    /*
        房源类型,见https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1     */
    ProductType  *int64 `json:"product_type,omitempty" required:"false" `
    /*
        房源介绍,最多2000字符     */
    Introduction  *string `json:"introduction,omitempty" required:"false" `
    /*
        是否与房东同住 0 不同住 1同住     */
    HasLandlord  *int64 `json:"has_landlord,omitempty" required:"false" `
    /*
        加人收费信息     */
    Charge  *domain.TaobaoXhotelBnbroomtypeAddBnbChargeDto `json:"charge,omitempty" required:"false" `
    /*
        装修风格https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1     */
    DecorateStyle  *int64 `json:"decorate_style,omitempty" required:"false" `
    /*
        是否信用免押金0：否 1：是     */
    Supportcredit  *int64 `json:"supportcredit,omitempty" required:"false" `
    /*
        入住须知，最多2000字符     */
    CheckInNotes  *string `json:"check_in_notes,omitempty" required:"false" `
    /*
        真实联系方式     */
    RealTel  *string `json:"real_tel" required:"true" `
    /*
        设施服务。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类     */
    Service  *string `json:"service,omitempty" required:"false" `
    /*
        “打扫类型1(1客1扫/换),2(1天1扫/换),https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1     */
    CleaningFrequency  *int64 `json:"cleaning_frequency,omitempty" required:"false" `
    /*
        房源id, 这是卖家自己系统中的ID     */
    OuterId  *string `json:"outer_id" required:"true" `
    /*
        民宿房源位置信息     */
    Location  *domain.TaobaoXhotelBnbroomtypeAddBnbLocationDto `json:"location" required:"true" `
    /*
        风景类型(枚举)https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1     */
    ScenicFeature  *int64 `json:"scenic_feature,omitempty" required:"false" `
    /*
        押金类型0.线下；1.线上     */
    DepositType  *int64 `json:"deposit_type,omitempty" required:"false" `
    /*
        床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347     */
    BedInfo  *string `json:"bed_info" required:"true" `
    /*
        额外收费     */
    ExtraFee  *string `json:"extra_fee,omitempty" required:"false" `
    /*
        加床费,分为单位     */
    ExtraBedsFee  *int64 `json:"extra_beds_fee,omitempty" required:"false" `
    /*
        民宿名称，默认取bnbName     */
    BnbName  *string `json:"bnb_name,omitempty" required:"false" `
    /*
        添加标准房源匹配     */
    Srid  *int64 `json:"srid,omitempty" required:"false" `
    /*
        标准酒店服务,参考文档https://fliggy.open.taobao.com/doc.htm?docId=120362&docType=1     */
    StandardRoomFacilities  *string `json:"standard_room_facilities" required:"true" `
    /*
        民宿扩展信息     */
    BnbExtend  *string `json:"bnb_extend,omitempty" required:"false" `
}

func (s *TaobaoXhotelBnbroomtypeAddRequest) SetRentSize(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.RentSize = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetSupportIm(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.SupportIm = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetOuterTags(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.OuterTags = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetCleaningCharge(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.CleaningCharge = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetInvoice(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Invoice = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetDecorateLevel(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.DecorateLevel = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetBnbBookingTime(v domain.TaobaoXhotelBnbroomtypeAddBnbBookingTimeDto) *TaobaoXhotelBnbroomtypeAddRequest {
    s.BnbBookingTime = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetOutHid(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.OutHid = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetTel(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Tel = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetReceiveForeigners(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.ReceiveForeigners = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetLocalInfo(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.LocalInfo = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetBrand(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Brand = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetPics(v []domain.TaobaoXhotelBnbroomtypeAddBnbPictureDTO) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Pics = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetCleaningType(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.CleaningType = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetDepositAmount(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.DepositAmount = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetNameE(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.NameE = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetDecorateTime(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.DecorateTime = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetExtraBedsNum(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.ExtraBedsNum = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetInvoiceType(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.InvoiceType = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetHasFrontDesk(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.HasFrontDesk = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetGuestAge(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.GuestAge = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetSettlementCurrency(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.SettlementCurrency = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetName(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetSurroundInfo(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.SurroundInfo = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetMaxOccupancy(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.MaxOccupancy = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetIsUseShootImage(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.IsUseShootImage = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetStatus(v byte) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetGuestGender(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.GuestGender = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetActivitiesAllowed(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.ActivitiesAllowed = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetExtraCleaningCharge(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.ExtraCleaningCharge = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetOpeningTime(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.OpeningTime = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetRentType(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.RentType = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetNewOuterId(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.NewOuterId = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetHouseModel(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.HouseModel = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetWindowType(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.WindowType = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetHasLicense(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.HasLicense = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetVideoUrl(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.VideoUrl = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetVendor(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetBrightspot(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Brightspot = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetQuickOrder(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.QuickOrder = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetFloor(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Floor = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetHouseSize(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.HouseSize = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetProductType(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.ProductType = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetIntroduction(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Introduction = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetHasLandlord(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.HasLandlord = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetCharge(v domain.TaobaoXhotelBnbroomtypeAddBnbChargeDto) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Charge = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetDecorateStyle(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.DecorateStyle = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetSupportcredit(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Supportcredit = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetCheckInNotes(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.CheckInNotes = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetRealTel(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.RealTel = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetService(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetCleaningFrequency(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.CleaningFrequency = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetOuterId(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetLocation(v domain.TaobaoXhotelBnbroomtypeAddBnbLocationDto) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Location = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetScenicFeature(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.ScenicFeature = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetDepositType(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.DepositType = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetBedInfo(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.BedInfo = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetExtraFee(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.ExtraFee = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetExtraBedsFee(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.ExtraBedsFee = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetBnbName(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.BnbName = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetSrid(v int64) *TaobaoXhotelBnbroomtypeAddRequest {
    s.Srid = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetStandardRoomFacilities(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.StandardRoomFacilities = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddRequest) SetBnbExtend(v string) *TaobaoXhotelBnbroomtypeAddRequest {
    s.BnbExtend = &v
    return s
}

func (req *TaobaoXhotelBnbroomtypeAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RentSize != nil) {
        paramMap["rent_size"] = *req.RentSize
    }
    if(req.SupportIm != nil) {
        paramMap["support_im"] = *req.SupportIm
    }
    if(req.OuterTags != nil) {
        paramMap["outer_tags"] = *req.OuterTags
    }
    if(req.CleaningCharge != nil) {
        paramMap["cleaning_charge"] = *req.CleaningCharge
    }
    if(req.Invoice != nil) {
        paramMap["invoice"] = *req.Invoice
    }
    if(req.DecorateLevel != nil) {
        paramMap["decorate_level"] = *req.DecorateLevel
    }
    if(req.BnbBookingTime != nil) {
        paramMap["bnb_booking_time"] = util.ConvertStruct(*req.BnbBookingTime)
    }
    if(req.OutHid != nil) {
        paramMap["out_hid"] = *req.OutHid
    }
    if(req.Tel != nil) {
        paramMap["tel"] = *req.Tel
    }
    if(req.ReceiveForeigners != nil) {
        paramMap["receive_foreigners"] = *req.ReceiveForeigners
    }
    if(req.LocalInfo != nil) {
        paramMap["local_info"] = *req.LocalInfo
    }
    if(req.Brand != nil) {
        paramMap["brand"] = *req.Brand
    }
    if(req.Pics != nil) {
        paramMap["pics"] = util.ConvertStructList(*req.Pics)
    }
    if(req.CleaningType != nil) {
        paramMap["cleaning_type"] = *req.CleaningType
    }
    if(req.DepositAmount != nil) {
        paramMap["deposit_amount"] = *req.DepositAmount
    }
    if(req.NameE != nil) {
        paramMap["name_e"] = *req.NameE
    }
    if(req.DecorateTime != nil) {
        paramMap["decorate_time"] = *req.DecorateTime
    }
    if(req.ExtraBedsNum != nil) {
        paramMap["extra_beds_num"] = *req.ExtraBedsNum
    }
    if(req.InvoiceType != nil) {
        paramMap["invoice_type"] = *req.InvoiceType
    }
    if(req.HasFrontDesk != nil) {
        paramMap["has_front_desk"] = *req.HasFrontDesk
    }
    if(req.GuestAge != nil) {
        paramMap["guest_age"] = *req.GuestAge
    }
    if(req.SettlementCurrency != nil) {
        paramMap["settlement_currency"] = *req.SettlementCurrency
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.SurroundInfo != nil) {
        paramMap["surround_info"] = *req.SurroundInfo
    }
    if(req.MaxOccupancy != nil) {
        paramMap["max_occupancy"] = *req.MaxOccupancy
    }
    if(req.IsUseShootImage != nil) {
        paramMap["is_use_shoot_image"] = *req.IsUseShootImage
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.GuestGender != nil) {
        paramMap["guest_gender"] = *req.GuestGender
    }
    if(req.ActivitiesAllowed != nil) {
        paramMap["activities_allowed"] = *req.ActivitiesAllowed
    }
    if(req.ExtraCleaningCharge != nil) {
        paramMap["extra_cleaning_charge"] = *req.ExtraCleaningCharge
    }
    if(req.OpeningTime != nil) {
        paramMap["opening_time"] = *req.OpeningTime
    }
    if(req.RentType != nil) {
        paramMap["rent_type"] = *req.RentType
    }
    if(req.NewOuterId != nil) {
        paramMap["new_outer_id"] = *req.NewOuterId
    }
    if(req.HouseModel != nil) {
        paramMap["house_model"] = *req.HouseModel
    }
    if(req.WindowType != nil) {
        paramMap["window_type"] = *req.WindowType
    }
    if(req.HasLicense != nil) {
        paramMap["has_license"] = *req.HasLicense
    }
    if(req.VideoUrl != nil) {
        paramMap["video_url"] = *req.VideoUrl
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.Brightspot != nil) {
        paramMap["brightspot"] = *req.Brightspot
    }
    if(req.QuickOrder != nil) {
        paramMap["quick_order"] = *req.QuickOrder
    }
    if(req.Floor != nil) {
        paramMap["floor"] = *req.Floor
    }
    if(req.HouseSize != nil) {
        paramMap["house_size"] = *req.HouseSize
    }
    if(req.ProductType != nil) {
        paramMap["product_type"] = *req.ProductType
    }
    if(req.Introduction != nil) {
        paramMap["introduction"] = *req.Introduction
    }
    if(req.HasLandlord != nil) {
        paramMap["has_landlord"] = *req.HasLandlord
    }
    if(req.Charge != nil) {
        paramMap["charge"] = util.ConvertStruct(*req.Charge)
    }
    if(req.DecorateStyle != nil) {
        paramMap["decorate_style"] = *req.DecorateStyle
    }
    if(req.Supportcredit != nil) {
        paramMap["supportcredit"] = *req.Supportcredit
    }
    if(req.CheckInNotes != nil) {
        paramMap["check_in_notes"] = *req.CheckInNotes
    }
    if(req.RealTel != nil) {
        paramMap["real_tel"] = *req.RealTel
    }
    if(req.Service != nil) {
        paramMap["service"] = *req.Service
    }
    if(req.CleaningFrequency != nil) {
        paramMap["cleaning_frequency"] = *req.CleaningFrequency
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Location != nil) {
        paramMap["location"] = util.ConvertStruct(*req.Location)
    }
    if(req.ScenicFeature != nil) {
        paramMap["scenic_feature"] = *req.ScenicFeature
    }
    if(req.DepositType != nil) {
        paramMap["deposit_type"] = *req.DepositType
    }
    if(req.BedInfo != nil) {
        paramMap["bed_info"] = *req.BedInfo
    }
    if(req.ExtraFee != nil) {
        paramMap["extra_fee"] = *req.ExtraFee
    }
    if(req.ExtraBedsFee != nil) {
        paramMap["extra_beds_fee"] = *req.ExtraBedsFee
    }
    if(req.BnbName != nil) {
        paramMap["bnb_name"] = *req.BnbName
    }
    if(req.Srid != nil) {
        paramMap["srid"] = *req.Srid
    }
    if(req.StandardRoomFacilities != nil) {
        paramMap["standard_room_facilities"] = *req.StandardRoomFacilities
    }
    if(req.BnbExtend != nil) {
        paramMap["bnb_extend"] = *req.BnbExtend
    }
    return paramMap
}

func (req *TaobaoXhotelBnbroomtypeAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}