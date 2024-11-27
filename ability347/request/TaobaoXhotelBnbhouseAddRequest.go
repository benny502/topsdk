package request

import (
	"github.com/benny502/topsdk/ability347/domain"
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelBnbhouseAddRequest struct {
	/*
	   外部房东id     */
	OutOwnerId *string `json:"out_owner_id" required:"true" `
	/*
	   入住要求&附加信息     */
	BnbBookingTime *domain.TaobaoXhotelBnbhouseAddBnbBookingTimeDto `json:"bnb_booking_time,omitempty" required:"false" `
	/*
	   装修等级 1 精装 2普通 3简装     */
	DecorateLevel *int64 `json:"decorate_level,omitempty" required:"false" `
	/*
	   联系方式。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886     */
	Tel *string `json:"tel" required:"true" `
	/*
	   是否可接待外宾 0不接待 1接待,默认值: 0 defalutValue��0    */
	ReceiveForeigners *int64 `json:"receive_foreigners,omitempty" required:"false" `
	/*
	   门店英文名称     */
	NameE *string `json:"name_e,omitempty" required:"false" `
	/*
	   装修时间，格式为2015-01-01     */
	DecorateTime *string `json:"decorate_time,omitempty" required:"false" `
	/*
	   可加床数     */
	ExtraBedsNum *int64 `json:"extra_beds_num,omitempty" required:"false" `
	/*
	   门店标签 标签信息，逗号(,)分隔     */
	Tags *string `json:"tags,omitempty" required:"false" `
	/*
	   是否有前台 0没有 1有     */
	HasFrontDesk *int64 `json:"has_front_desk,omitempty" required:"false" `
	/*
	   可接待客人年龄情况：是否接待儿童、老人；成年人必接待，详见“可接待客人”list     */
	GuestAge *int64 `json:"guest_age,omitempty" required:"false" `
	/*
	   结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用     */
	SettlementCurrency *string `json:"settlement_currency,omitempty" required:"false" `
	/*
	   门店名称     */
	Name *string `json:"name" required:"true" `
	/*
	   是否使用实拍图片 0不使用 1使用     */
	IsUseShootImage *int64 `json:"is_use_shoot_image,omitempty" required:"false" `
	/*
	   房型状态。0:正常，-1:删除，-2:停售     */
	Status *byte `json:"status" required:"true" `
	/*
	   可接待客人性别 0：不限制，1：只限男性，2：只限女性'     */
	GuestGender *int64 `json:"guest_gender,omitempty" required:"false" `
	/*
	   详见“允许活动”list 12,32,33,     */
	ActivitiesAllowed *string `json:"activities_allowed,omitempty" required:"false" `
	/*
	   开业时间，格式为2015-01-01     */
	OpeningTime *string `json:"opening_time,omitempty" required:"false" `
	/*
	   门店简介     */
	Description *string `json:"description,omitempty" required:"false" `
	/*
	   民宿门店添加     */
	Pictures *[]domain.TaobaoXhotelBnbhouseAddBnbPictureDTO `json:"pictures" required:"true" `
	/*
	   有无资质执照 0 无资质 1有资质     */
	HasLicense *int64 `json:"has_license" required:"true" `
	/*
	   楼层     */
	Floors *string `json:"floors,omitempty" required:"false" `
	/*
	   视频地址     */
	VideoUrl *string `json:"video_url,omitempty" required:"false" `
	/*
	   对接系统商名称：可为空不要乱填，需要申请后使用     */
	Vendor *string `json:"vendor,omitempty" required:"false" `
	/*
	   门店类型，详见“房源类型list     */
	ProductType *int64 `json:"product_type" required:"true" `
	/*
	   加人收费信息     */
	Charge *domain.TaobaoXhotelBnbhouseAddBnbChargeDto `json:"charge,omitempty" required:"false" `
	/*
	   装修风格，详见装修风格枚举表     */
	DecorateStyle *int64 `json:"decorate_style,omitempty" required:"false" `
	/*
	   入住须知     */
	CheckInNotes *string `json:"check_in_notes,omitempty" required:"false" `
	/*
	   真实联系方式     */
	RealTel *string `json:"real_tel" required:"true" `
	/*
	   位置信息     */
	Location *domain.TaobaoXhotelBnbhouseAddBnbLocationDto `json:"location" required:"true" `
	/*
	   门店属性 1 个人房源 2 城市公寓 3 乡村民宿等     */
	Attributes *int64 `json:"attributes" required:"true" `
	/*
	   供应商渠道门店ID     */
	OuterId *string `json:"outer_id" required:"true" `
	/*
	   风景类型，详见风景类型枚举表     */
	ScenicFeature *int64 `json:"scenic_feature,omitempty" required:"false" `
	/*
	   品牌名称     */
	Brand *string `json:"brand,omitempty" required:"false" `
	/*
	   酒店设施。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=20的分类     */
	Facilities *string `json:"facilities,omitempty" required:"false" `
	/*
	   面积大小     */
	HouseSize *int64 `json:"house_size,omitempty" required:"false" `
	/*
	   匹配的标准门店     */
	Shid *int64 `json:"shid,omitempty" required:"false" `
	/*
	   标准酒店服务,参考文档https://fliggy.open.taobao.com/doc.htm?docId=120362&docType=1     */
	StandardHotelFacilities *string `json:"standard_hotel_facilities" required:"true" `
	/*
	   传入是或者否，是表明为菲住合作模式，hid打标；“否”表示为正常合作模式，hid去标；不传保持原有的合作模式不变     */
	IsFeizhuHotel *bool `json:"is_feizhu_hotel,omitempty" required:"false" `
	/*
	   作为菲住门店签约的佣金比率,范围为10-50     */
	CommissionRate *int64 `json:"commission_rate,omitempty" required:"false" `
}

func (s *TaobaoXhotelBnbhouseAddRequest) SetOutOwnerId(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.OutOwnerId = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetBnbBookingTime(v domain.TaobaoXhotelBnbhouseAddBnbBookingTimeDto) *TaobaoXhotelBnbhouseAddRequest {
	s.BnbBookingTime = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetDecorateLevel(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.DecorateLevel = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetTel(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.Tel = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetReceiveForeigners(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.ReceiveForeigners = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetNameE(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.NameE = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetDecorateTime(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.DecorateTime = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetExtraBedsNum(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.ExtraBedsNum = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetTags(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.Tags = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetHasFrontDesk(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.HasFrontDesk = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetGuestAge(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.GuestAge = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetSettlementCurrency(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.SettlementCurrency = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetName(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.Name = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetIsUseShootImage(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.IsUseShootImage = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetStatus(v byte) *TaobaoXhotelBnbhouseAddRequest {
	s.Status = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetGuestGender(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.GuestGender = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetActivitiesAllowed(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.ActivitiesAllowed = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetOpeningTime(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.OpeningTime = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetDescription(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.Description = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetPictures(v []domain.TaobaoXhotelBnbhouseAddBnbPictureDTO) *TaobaoXhotelBnbhouseAddRequest {
	s.Pictures = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetHasLicense(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.HasLicense = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetFloors(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.Floors = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetVideoUrl(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.VideoUrl = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetVendor(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.Vendor = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetProductType(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.ProductType = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetCharge(v domain.TaobaoXhotelBnbhouseAddBnbChargeDto) *TaobaoXhotelBnbhouseAddRequest {
	s.Charge = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetDecorateStyle(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.DecorateStyle = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetCheckInNotes(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.CheckInNotes = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetRealTel(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.RealTel = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetLocation(v domain.TaobaoXhotelBnbhouseAddBnbLocationDto) *TaobaoXhotelBnbhouseAddRequest {
	s.Location = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetAttributes(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.Attributes = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetOuterId(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.OuterId = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetScenicFeature(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.ScenicFeature = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetBrand(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.Brand = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetFacilities(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.Facilities = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetHouseSize(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.HouseSize = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetShid(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.Shid = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetStandardHotelFacilities(v string) *TaobaoXhotelBnbhouseAddRequest {
	s.StandardHotelFacilities = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetIsFeizhuHotel(v bool) *TaobaoXhotelBnbhouseAddRequest {
	s.IsFeizhuHotel = &v
	return s
}
func (s *TaobaoXhotelBnbhouseAddRequest) SetCommissionRate(v int64) *TaobaoXhotelBnbhouseAddRequest {
	s.CommissionRate = &v
	return s
}

func (req *TaobaoXhotelBnbhouseAddRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.OutOwnerId != nil {
		paramMap["out_owner_id"] = *req.OutOwnerId
	}
	if req.BnbBookingTime != nil {
		paramMap["bnb_booking_time"] = util.ConvertStruct(*req.BnbBookingTime)
	}
	if req.DecorateLevel != nil {
		paramMap["decorate_level"] = *req.DecorateLevel
	}
	if req.Tel != nil {
		paramMap["tel"] = *req.Tel
	}
	if req.ReceiveForeigners != nil {
		paramMap["receive_foreigners"] = *req.ReceiveForeigners
	}
	if req.NameE != nil {
		paramMap["name_e"] = *req.NameE
	}
	if req.DecorateTime != nil {
		paramMap["decorate_time"] = *req.DecorateTime
	}
	if req.ExtraBedsNum != nil {
		paramMap["extra_beds_num"] = *req.ExtraBedsNum
	}
	if req.Tags != nil {
		paramMap["tags"] = *req.Tags
	}
	if req.HasFrontDesk != nil {
		paramMap["has_front_desk"] = *req.HasFrontDesk
	}
	if req.GuestAge != nil {
		paramMap["guest_age"] = *req.GuestAge
	}
	if req.SettlementCurrency != nil {
		paramMap["settlement_currency"] = *req.SettlementCurrency
	}
	if req.Name != nil {
		paramMap["name"] = *req.Name
	}
	if req.IsUseShootImage != nil {
		paramMap["is_use_shoot_image"] = *req.IsUseShootImage
	}
	if req.Status != nil {
		paramMap["status"] = *req.Status
	}
	if req.GuestGender != nil {
		paramMap["guest_gender"] = *req.GuestGender
	}
	if req.ActivitiesAllowed != nil {
		paramMap["activities_allowed"] = *req.ActivitiesAllowed
	}
	if req.OpeningTime != nil {
		paramMap["opening_time"] = *req.OpeningTime
	}
	if req.Description != nil {
		paramMap["description"] = *req.Description
	}
	if req.Pictures != nil {
		paramMap["pictures"] = util.ConvertStructList(*req.Pictures)
	}
	if req.HasLicense != nil {
		paramMap["has_license"] = *req.HasLicense
	}
	if req.Floors != nil {
		paramMap["floors"] = *req.Floors
	}
	if req.VideoUrl != nil {
		paramMap["video_url"] = *req.VideoUrl
	}
	if req.Vendor != nil {
		paramMap["vendor"] = *req.Vendor
	}
	if req.ProductType != nil {
		paramMap["product_type"] = *req.ProductType
	}
	if req.Charge != nil {
		paramMap["charge"] = util.ConvertStruct(*req.Charge)
	}
	if req.DecorateStyle != nil {
		paramMap["decorate_style"] = *req.DecorateStyle
	}
	if req.CheckInNotes != nil {
		paramMap["check_in_notes"] = *req.CheckInNotes
	}
	if req.RealTel != nil {
		paramMap["real_tel"] = *req.RealTel
	}
	if req.Location != nil {
		paramMap["location"] = util.ConvertStruct(*req.Location)
	}
	if req.Attributes != nil {
		paramMap["attributes"] = *req.Attributes
	}
	if req.OuterId != nil {
		paramMap["outer_id"] = *req.OuterId
	}
	if req.ScenicFeature != nil {
		paramMap["scenic_feature"] = *req.ScenicFeature
	}
	if req.Brand != nil {
		paramMap["brand"] = *req.Brand
	}
	if req.Facilities != nil {
		paramMap["facilities"] = *req.Facilities
	}
	if req.HouseSize != nil {
		paramMap["house_size"] = *req.HouseSize
	}
	if req.Shid != nil {
		paramMap["shid"] = *req.Shid
	}
	if req.StandardHotelFacilities != nil {
		paramMap["standard_hotel_facilities"] = *req.StandardHotelFacilities
	}
	if req.IsFeizhuHotel != nil {
		paramMap["is_feizhu_hotel"] = *req.IsFeizhuHotel
	}
	if req.CommissionRate != nil {
		paramMap["commission_rate"] = *req.CommissionRate
	}
	return paramMap
}

func (req *TaobaoXhotelBnbhouseAddRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
