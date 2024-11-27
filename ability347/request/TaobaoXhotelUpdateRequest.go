package request


type TaobaoXhotelUpdateRequest struct {
    /*
        （已废弃）请使用outer_id来标识要修改的酒店     */
    Hid  *int64 `json:"hid,omitempty" required:"false" `
    /*
        酒店名称；（新增酒店时为必须）,国内酒店请传中文名称     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        酒店曾用名     */
    UsedName  *string `json:"used_name,omitempty" required:"false" `
    /*
        是否国内酒店。0:国内;1:国外     */
    Domestic  *int64 `json:"domestic,omitempty" required:"false" `
    /*
        domestic为true时，固定China； domestic为false时，必须传定义的海外国家编码值。参见：http://kezhan.trip.taobao.com/countrys.html     */
    Country  *string `json:"country,omitempty" required:"false" `
    /*
        省份编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时默认为0     */
    Province  *int64 `json:"province,omitempty" required:"false" `
    /*
        城市编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（新增酒店时为必须）     */
    City  *int64 `json:"city,omitempty" required:"false" `
    /*
        区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3 defalutValue��0    */
    District  *int64 `json:"district,omitempty" required:"false" `
    /*
        商业区（圈）长度不超过20字     */
    Business  *string `json:"business,omitempty" required:"false" `
    /*
        酒店地址。长度不能超过255     */
    Address  *string `json:"address" required:"true" `
    /*
        经度     */
    Longitude  *string `json:"longitude,omitempty" required:"false" `
    /*
        纬度     */
    Latitude  *string `json:"latitude,omitempty" required:"false" `
    /*
        坐标类型，现在支持：G – GoogleB – 百度A – 高德M – MapbarL – 灵图     */
    PositionType  *string `json:"position_type,omitempty" required:"false" `
    /*
        酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886     */
    Tel  *string `json:"tel" required:"true" `
    /*
        不要使用     */
    Extend  *string `json:"extend,omitempty" required:"false" `
    /*
        必传，酒店标识，商家酒店ID     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        系统商，一般情况不用，需申请使用     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。     */
    Shid  *int64 `json:"shid,omitempty" required:"false" `
    /*
        酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5     */
    Star  *string `json:"star,omitempty" required:"false" `
    /*
        开业时间，格式为2015-01-01     */
    OpeningTime  *string `json:"opening_time,omitempty" required:"false" `
    /*
        装修时间，格式为2015-10-01     */
    DecorateTime  *string `json:"decorate_time,omitempty" required:"false" `
    /*
        楼层信息     */
    Floors  *string `json:"floors,omitempty" required:"false" `
    /*
        房间数 0~9999之内的数字     */
    Rooms  *int64 `json:"rooms,omitempty" required:"false" `
    /*
        酒店描述     */
    Description  *string `json:"description,omitempty" required:"false" `
    /*
        酒店设施。json格式示例值：{"free Wi-Fi in all rooms":"true","massage":"true","meetingRoom":"true"}目前支持维护的设施枚举有：free Wi-Fi in all rooms 所有房间设有免费无线网络;meetingRoom 会议室;massage  按摩室;fitnessClub 健身房;bar 酒吧;cafe 咖啡厅;frontDeskSafe 前台贵重物品保险柜wifi 无线上网公共区域;casino 娱乐场/棋牌室;restaurant 餐厅;smoking area 吸烟区;Business Facilities 商务设施     */
    HotelFacilities  *string `json:"hotel_facilities,omitempty" required:"false" `
    /*
        酒店基础服务。json格式示例值：{"receiveForeignGuests":"true","morningCall":"true","breakfast":"true"}目前支持维护的设施枚举有：receiveForeignGuests 接待外宾;morningCall 叫醒服务; breakfast  早餐服务; airportShuttle 接机服务; luggageClaim 行李寄存; rentCar 租车; HourRoomService24 24小时客房服务; airportTransfer 酒店/机场接送; dryCleaning 干洗; expressCheckInCheckOut 快速入住/退房登记; custodyServices 保管服务     */
    Service  *string `json:"service,omitempty" required:"false" `
    /*
        房间的基础设施。json格式示例值：{"bathtub":"true","bathPub":"true"}目前支持维护的设施枚举有：bathtub 独立卫浴;bathPub 公共卫浴     */
    RoomFacilities  *string `json:"room_facilities,omitempty" required:"false" `
    /*
        酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址，main是否为主图（主图只能有一个）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 其他, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂]，图片数量最多10张。要求：无logo、水印、边框、人物，不模糊、重复、歪斜，房间图清晰，图片尺寸不小于300*225，不小于5M     */
    Pics  *string `json:"pics,omitempty" required:"false" `
    /*
        酒店品牌。取值为数字。枚举见链接：https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.S16vXH&docType=1&articleId=120180     */
    Brand  *string `json:"brand,omitempty" required:"false" `
    /*
        邮编     */
    PostalCode  *string `json:"postal_code,omitempty" required:"false" `
    /*
        酒店入住政策(针对国际酒店，儿童及加床信息)格式：{"children_age_from":"2","children_age_to":"3","children_stay_free":"True","infant_age":"1","min_guest_age":"4"}     */
    HotelPolicies  *string `json:"hotel_policies,omitempty" required:"false" `
    /*
        预订须知。json字段描述：hotelInMountaintop 酒店位于山顶 1在山顶、0不在；needBoat 酒店需要坐船前往 1需要、0不需要；酒店位于景区内 1在景区、0不在；extraBed 加床收费；extraCharge 额外收费；arrivalTime 到店时间；extend 其他补充项     */
    BookingNotice  *string `json:"booking_notice,omitempty" required:"false" `
    /*
        酒店状态 0:正常，-1:下架，-2:停售     */
    Status  *byte `json:"status,omitempty" required:"false" `
    /*
        逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡     */
    CreditCardTypes  *string `json:"credit_card_types,omitempty" required:"false" `
    /*
        扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的     */
    OrbitTrack  *string `json:"orbit_track,omitempty" required:"false" `
    /*
        卖家酒店英文名称     */
    NameE  *string `json:"name_e,omitempty" required:"false" `
    /*
        打标去标使用的 tagJson 字段     */
    TagJson  *string `json:"tag_json,omitempty" required:"false" `
    /*
        旺旺昵称     */
    AliNick  *string `json:"ali_nick,omitempty" required:"false" `
    /*
        供应商标识，如果确实需要修改原来的供应商标识才需要填写，否则不需要填写，请谨慎使用。     */
    Supplier  *string `json:"supplier,omitempty" required:"false" `
    /*
        结算流程中的结算币种，如需对接请联系飞猪技术支持，请谨慎使用     */
    SettlementCurrency  *string `json:"settlement_currency,omitempty" required:"false" `
    /*
        资源方酒店预订须知,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891     */
    StandardBookingNotice  *string `json:"standard_booking_notice,omitempty" required:"false" `
    /*
        资源方酒店设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891     */
    StandardHotelFacilities  *string `json:"standard_hotel_facilities,omitempty" required:"false" `
    /*
        资源方酒店服务,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891     */
    StandardHotelService  *string `json:"standard_hotel_service,omitempty" required:"false" `
    /*
        已废弃(房型设施信息请在房型上推送)     */
    StandardRoomFacilities  *string `json:"standard_room_facilities,omitempty" required:"false" `
    /*
        资源方娱乐设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891     */
    StandardAmuseFacilities  *string `json:"standard_amuse_facilities,omitempty" required:"false" `
    /*
        0:酒店；1:客栈     */
    HotelType  *int64 `json:"hotel_type,omitempty" required:"false" `
    /*
        0:可以接待外宾；1:仅内宾     */
    ServiceType  *int64 `json:"service_type,omitempty" required:"false" `
    /*
        标识坐标系类型。WGS84，表示地球坐标系 ；GCJ02，表示火星坐标系     */
    CoordinateSystem  *string `json:"coordinate_system,omitempty" required:"false" `
    /*
        酒店英文地址     */
    AddressEn  *string `json:"address_en,omitempty" required:"false" `
    /*
        酒店英文描述     */
    DescriptionEn  *string `json:"description_en,omitempty" required:"false" `
}

func (s *TaobaoXhotelUpdateRequest) SetHid(v int64) *TaobaoXhotelUpdateRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetName(v string) *TaobaoXhotelUpdateRequest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetUsedName(v string) *TaobaoXhotelUpdateRequest {
    s.UsedName = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetDomestic(v int64) *TaobaoXhotelUpdateRequest {
    s.Domestic = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetCountry(v string) *TaobaoXhotelUpdateRequest {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetProvince(v int64) *TaobaoXhotelUpdateRequest {
    s.Province = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetCity(v int64) *TaobaoXhotelUpdateRequest {
    s.City = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetDistrict(v int64) *TaobaoXhotelUpdateRequest {
    s.District = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetBusiness(v string) *TaobaoXhotelUpdateRequest {
    s.Business = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetAddress(v string) *TaobaoXhotelUpdateRequest {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetLongitude(v string) *TaobaoXhotelUpdateRequest {
    s.Longitude = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetLatitude(v string) *TaobaoXhotelUpdateRequest {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetPositionType(v string) *TaobaoXhotelUpdateRequest {
    s.PositionType = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetTel(v string) *TaobaoXhotelUpdateRequest {
    s.Tel = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetExtend(v string) *TaobaoXhotelUpdateRequest {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetOuterId(v string) *TaobaoXhotelUpdateRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetVendor(v string) *TaobaoXhotelUpdateRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetShid(v int64) *TaobaoXhotelUpdateRequest {
    s.Shid = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetStar(v string) *TaobaoXhotelUpdateRequest {
    s.Star = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetOpeningTime(v string) *TaobaoXhotelUpdateRequest {
    s.OpeningTime = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetDecorateTime(v string) *TaobaoXhotelUpdateRequest {
    s.DecorateTime = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetFloors(v string) *TaobaoXhotelUpdateRequest {
    s.Floors = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetRooms(v int64) *TaobaoXhotelUpdateRequest {
    s.Rooms = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetDescription(v string) *TaobaoXhotelUpdateRequest {
    s.Description = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetHotelFacilities(v string) *TaobaoXhotelUpdateRequest {
    s.HotelFacilities = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetService(v string) *TaobaoXhotelUpdateRequest {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetRoomFacilities(v string) *TaobaoXhotelUpdateRequest {
    s.RoomFacilities = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetPics(v string) *TaobaoXhotelUpdateRequest {
    s.Pics = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetBrand(v string) *TaobaoXhotelUpdateRequest {
    s.Brand = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetPostalCode(v string) *TaobaoXhotelUpdateRequest {
    s.PostalCode = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetHotelPolicies(v string) *TaobaoXhotelUpdateRequest {
    s.HotelPolicies = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetBookingNotice(v string) *TaobaoXhotelUpdateRequest {
    s.BookingNotice = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetStatus(v byte) *TaobaoXhotelUpdateRequest {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetCreditCardTypes(v string) *TaobaoXhotelUpdateRequest {
    s.CreditCardTypes = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetOrbitTrack(v string) *TaobaoXhotelUpdateRequest {
    s.OrbitTrack = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetNameE(v string) *TaobaoXhotelUpdateRequest {
    s.NameE = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetTagJson(v string) *TaobaoXhotelUpdateRequest {
    s.TagJson = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetAliNick(v string) *TaobaoXhotelUpdateRequest {
    s.AliNick = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetSupplier(v string) *TaobaoXhotelUpdateRequest {
    s.Supplier = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetSettlementCurrency(v string) *TaobaoXhotelUpdateRequest {
    s.SettlementCurrency = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetStandardBookingNotice(v string) *TaobaoXhotelUpdateRequest {
    s.StandardBookingNotice = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetStandardHotelFacilities(v string) *TaobaoXhotelUpdateRequest {
    s.StandardHotelFacilities = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetStandardHotelService(v string) *TaobaoXhotelUpdateRequest {
    s.StandardHotelService = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetStandardRoomFacilities(v string) *TaobaoXhotelUpdateRequest {
    s.StandardRoomFacilities = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetStandardAmuseFacilities(v string) *TaobaoXhotelUpdateRequest {
    s.StandardAmuseFacilities = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetHotelType(v int64) *TaobaoXhotelUpdateRequest {
    s.HotelType = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetServiceType(v int64) *TaobaoXhotelUpdateRequest {
    s.ServiceType = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetCoordinateSystem(v string) *TaobaoXhotelUpdateRequest {
    s.CoordinateSystem = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetAddressEn(v string) *TaobaoXhotelUpdateRequest {
    s.AddressEn = &v
    return s
}
func (s *TaobaoXhotelUpdateRequest) SetDescriptionEn(v string) *TaobaoXhotelUpdateRequest {
    s.DescriptionEn = &v
    return s
}

func (req *TaobaoXhotelUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Hid != nil) {
        paramMap["hid"] = *req.Hid
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.UsedName != nil) {
        paramMap["used_name"] = *req.UsedName
    }
    if(req.Domestic != nil) {
        paramMap["domestic"] = *req.Domestic
    }
    if(req.Country != nil) {
        paramMap["country"] = *req.Country
    }
    if(req.Province != nil) {
        paramMap["province"] = *req.Province
    }
    if(req.City != nil) {
        paramMap["city"] = *req.City
    }
    if(req.District != nil) {
        paramMap["district"] = *req.District
    }
    if(req.Business != nil) {
        paramMap["business"] = *req.Business
    }
    if(req.Address != nil) {
        paramMap["address"] = *req.Address
    }
    if(req.Longitude != nil) {
        paramMap["longitude"] = *req.Longitude
    }
    if(req.Latitude != nil) {
        paramMap["latitude"] = *req.Latitude
    }
    if(req.PositionType != nil) {
        paramMap["position_type"] = *req.PositionType
    }
    if(req.Tel != nil) {
        paramMap["tel"] = *req.Tel
    }
    if(req.Extend != nil) {
        paramMap["extend"] = *req.Extend
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.Shid != nil) {
        paramMap["shid"] = *req.Shid
    }
    if(req.Star != nil) {
        paramMap["star"] = *req.Star
    }
    if(req.OpeningTime != nil) {
        paramMap["opening_time"] = *req.OpeningTime
    }
    if(req.DecorateTime != nil) {
        paramMap["decorate_time"] = *req.DecorateTime
    }
    if(req.Floors != nil) {
        paramMap["floors"] = *req.Floors
    }
    if(req.Rooms != nil) {
        paramMap["rooms"] = *req.Rooms
    }
    if(req.Description != nil) {
        paramMap["description"] = *req.Description
    }
    if(req.HotelFacilities != nil) {
        paramMap["hotel_facilities"] = *req.HotelFacilities
    }
    if(req.Service != nil) {
        paramMap["service"] = *req.Service
    }
    if(req.RoomFacilities != nil) {
        paramMap["room_facilities"] = *req.RoomFacilities
    }
    if(req.Pics != nil) {
        paramMap["pics"] = *req.Pics
    }
    if(req.Brand != nil) {
        paramMap["brand"] = *req.Brand
    }
    if(req.PostalCode != nil) {
        paramMap["postal_code"] = *req.PostalCode
    }
    if(req.HotelPolicies != nil) {
        paramMap["hotel_policies"] = *req.HotelPolicies
    }
    if(req.BookingNotice != nil) {
        paramMap["booking_notice"] = *req.BookingNotice
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.CreditCardTypes != nil) {
        paramMap["credit_card_types"] = *req.CreditCardTypes
    }
    if(req.OrbitTrack != nil) {
        paramMap["orbit_track"] = *req.OrbitTrack
    }
    if(req.NameE != nil) {
        paramMap["name_e"] = *req.NameE
    }
    if(req.TagJson != nil) {
        paramMap["tag_json"] = *req.TagJson
    }
    if(req.AliNick != nil) {
        paramMap["ali_nick"] = *req.AliNick
    }
    if(req.Supplier != nil) {
        paramMap["supplier"] = *req.Supplier
    }
    if(req.SettlementCurrency != nil) {
        paramMap["settlement_currency"] = *req.SettlementCurrency
    }
    if(req.StandardBookingNotice != nil) {
        paramMap["standard_booking_notice"] = *req.StandardBookingNotice
    }
    if(req.StandardHotelFacilities != nil) {
        paramMap["standard_hotel_facilities"] = *req.StandardHotelFacilities
    }
    if(req.StandardHotelService != nil) {
        paramMap["standard_hotel_service"] = *req.StandardHotelService
    }
    if(req.StandardRoomFacilities != nil) {
        paramMap["standard_room_facilities"] = *req.StandardRoomFacilities
    }
    if(req.StandardAmuseFacilities != nil) {
        paramMap["standard_amuse_facilities"] = *req.StandardAmuseFacilities
    }
    if(req.HotelType != nil) {
        paramMap["hotel_type"] = *req.HotelType
    }
    if(req.ServiceType != nil) {
        paramMap["service_type"] = *req.ServiceType
    }
    if(req.CoordinateSystem != nil) {
        paramMap["coordinate_system"] = *req.CoordinateSystem
    }
    if(req.AddressEn != nil) {
        paramMap["address_en"] = *req.AddressEn
    }
    if(req.DescriptionEn != nil) {
        paramMap["description_en"] = *req.DescriptionEn
    }
    return paramMap
}

func (req *TaobaoXhotelUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}