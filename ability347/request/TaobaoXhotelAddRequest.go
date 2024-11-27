package request


type TaobaoXhotelAddRequest struct {
    /*
        外部酒店ID, 这是卖家自己系统中的ID     */
    OuterId  *string `json:"outer_id" required:"true" `
    /*
        酒店名称,国内酒店请传中文名称     */
    Name  *string `json:"name" required:"true" `
    /*
        酒店曾用名     */
    UsedName  *string `json:"used_name,omitempty" required:"false" `
    /*
        是否国内酒店。0:国内;1:国外。默认是国内 defalutValue��0    */
    Domestic  *int64 `json:"domestic,omitempty" required:"false" `
    /*
        domestic为0时，固定China； domestic为1时，必须传定义的海外国家编码值。参见：http://hotel.alitrip.com/area.htm     */
    Country  *string `json:"country,omitempty" required:"false" `
    /*
        省份编码。选填，不填入的时候已city字段为准.参见：http://hotel.alitrip.com/area.htm，domestic为false时默认为0     */
    Province  *int64 `json:"province,omitempty" required:"false" `
    /*
        城市编码。参见：http://hotel.alitrip.com/area.htm，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（更新酒店时为可选）     */
    City  *int64 `json:"city" required:"true" `
    /*
        区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm defalutValue��0    */
    District  *int64 `json:"district,omitempty" required:"false" `
    /*
        商业区（圈）长度不超过20字     */
    Business  *string `json:"business,omitempty" required:"false" `
    /*
        酒店地址。长度不能超过255。不填入会导致不能自动匹配。     */
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
        扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析     */
    Extend  *string `json:"extend,omitempty" required:"false" `
    /*
        该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。     */
    Shid  *int64 `json:"shid,omitempty" required:"false" `
    /*
        对接系统商名称：可为空不要乱填，需要申请后使用     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5     */
    Star  *string `json:"star,omitempty" required:"false" `
    /*
        开业时间，格式为2015-01-01     */
    OpeningTime  *string `json:"opening_time,omitempty" required:"false" `
    /*
        装修时间，格式为2015-01-01装修时间     */
    DecorateTime  *string `json:"decorate_time,omitempty" required:"false" `
    /*
        楼层信息。     */
    Floors  *string `json:"floors,omitempty" required:"false" `
    /*
        房间数 0~9999之内的数字     */
    Rooms  *int64 `json:"rooms,omitempty" required:"false" `
    /*
        酒店描述     */
    Description  *string `json:"description,omitempty" required:"false" `
    /*
        酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图（主图只能有一个，如果有多个或者没有，则会报错）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 其他,餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂]，图片数量最多是能是10张。     */
    Pics  *string `json:"pics,omitempty" required:"false" `
    /*
        酒店品牌。取值为数字。枚举见链接：https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.S16vXH&docType=1&articleId=120180     */
    Brand  *string `json:"brand,omitempty" required:"false" `
    /*
        邮政编码。     */
    PostalCode  *string `json:"postal_code,omitempty" required:"false" `
    /*
        预订须知。json字段描述：hotelInMountaintop 酒店位于山顶 1在山顶、0不在；needBoat 酒店需要坐船前往 1需要、0不需要；酒店位于景区内 1在景区、0不在；extraBed 加床收费；extraCharge 额外收费；arrivalTime 到店时间；extend 其他补充项     */
    BookingNotice  *string `json:"booking_notice,omitempty" required:"false" `
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
        供应商标识，需要提前开通权限，如果需要对接请联系技术支持，请谨慎使用。注：如商家申请的应用类型为“飞猪-新业务”，此项则必填。     */
    Supplier  *string `json:"supplier,omitempty" required:"false" `
    /*
        结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用     */
    SettlementCurrency  *string `json:"settlement_currency,omitempty" required:"false" `
    /*
        标准娱乐设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891     */
    StandardAmuseFacilities  *string `json:"standard_amuse_facilities,omitempty" required:"false" `
    /*
        已废弃(房型设施信息请在房型上推送)     */
    StandardRoomFacilities  *string `json:"standard_room_facilities,omitempty" required:"false" `
    /*
        标准酒店服务,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891     */
    StandardHotelService  *string `json:"standard_hotel_service,omitempty" required:"false" `
    /*
        标准酒店设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891     */
    StandardHotelFacilities  *string `json:"standard_hotel_facilities,omitempty" required:"false" `
    /*
        标准预订须知,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891     */
    StandardBookingNotice  *string `json:"standard_booking_notice,omitempty" required:"false" `
    /*
        0:可以接待外宾；1:仅内宾     */
    ServiceType  *int64 `json:"service_type,omitempty" required:"false" `
    /*
        0:酒店；1:客栈     */
    HotelType  *int64 `json:"hotel_type,omitempty" required:"false" `
    /*
        标识坐标系类型。WGS84，表示地球坐标系 ；GCJ02，表示火星坐标系     */
    CoordinateSystem  *string `json:"coordinate_system,omitempty" required:"false" `
    /*
        废弃     */
    RoomFacilities  *string `json:"room_facilities,omitempty" required:"false" `
    /*
        废弃     */
    Service  *string `json:"service,omitempty" required:"false" `
    /*
        废弃     */
    HotelFacilities  *string `json:"hotel_facilities,omitempty" required:"false" `
    /*
        废弃     */
    HotelPolicies  *string `json:"hotel_policies,omitempty" required:"false" `
}

func (s *TaobaoXhotelAddRequest) SetOuterId(v string) *TaobaoXhotelAddRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetName(v string) *TaobaoXhotelAddRequest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetUsedName(v string) *TaobaoXhotelAddRequest {
    s.UsedName = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetDomestic(v int64) *TaobaoXhotelAddRequest {
    s.Domestic = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetCountry(v string) *TaobaoXhotelAddRequest {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetProvince(v int64) *TaobaoXhotelAddRequest {
    s.Province = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetCity(v int64) *TaobaoXhotelAddRequest {
    s.City = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetDistrict(v int64) *TaobaoXhotelAddRequest {
    s.District = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetBusiness(v string) *TaobaoXhotelAddRequest {
    s.Business = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetAddress(v string) *TaobaoXhotelAddRequest {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetLongitude(v string) *TaobaoXhotelAddRequest {
    s.Longitude = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetLatitude(v string) *TaobaoXhotelAddRequest {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetPositionType(v string) *TaobaoXhotelAddRequest {
    s.PositionType = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetTel(v string) *TaobaoXhotelAddRequest {
    s.Tel = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetExtend(v string) *TaobaoXhotelAddRequest {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetShid(v int64) *TaobaoXhotelAddRequest {
    s.Shid = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetVendor(v string) *TaobaoXhotelAddRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetStar(v string) *TaobaoXhotelAddRequest {
    s.Star = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetOpeningTime(v string) *TaobaoXhotelAddRequest {
    s.OpeningTime = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetDecorateTime(v string) *TaobaoXhotelAddRequest {
    s.DecorateTime = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetFloors(v string) *TaobaoXhotelAddRequest {
    s.Floors = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetRooms(v int64) *TaobaoXhotelAddRequest {
    s.Rooms = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetDescription(v string) *TaobaoXhotelAddRequest {
    s.Description = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetPics(v string) *TaobaoXhotelAddRequest {
    s.Pics = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetBrand(v string) *TaobaoXhotelAddRequest {
    s.Brand = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetPostalCode(v string) *TaobaoXhotelAddRequest {
    s.PostalCode = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetBookingNotice(v string) *TaobaoXhotelAddRequest {
    s.BookingNotice = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetCreditCardTypes(v string) *TaobaoXhotelAddRequest {
    s.CreditCardTypes = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetOrbitTrack(v string) *TaobaoXhotelAddRequest {
    s.OrbitTrack = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetNameE(v string) *TaobaoXhotelAddRequest {
    s.NameE = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetSupplier(v string) *TaobaoXhotelAddRequest {
    s.Supplier = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetSettlementCurrency(v string) *TaobaoXhotelAddRequest {
    s.SettlementCurrency = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetStandardAmuseFacilities(v string) *TaobaoXhotelAddRequest {
    s.StandardAmuseFacilities = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetStandardRoomFacilities(v string) *TaobaoXhotelAddRequest {
    s.StandardRoomFacilities = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetStandardHotelService(v string) *TaobaoXhotelAddRequest {
    s.StandardHotelService = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetStandardHotelFacilities(v string) *TaobaoXhotelAddRequest {
    s.StandardHotelFacilities = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetStandardBookingNotice(v string) *TaobaoXhotelAddRequest {
    s.StandardBookingNotice = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetServiceType(v int64) *TaobaoXhotelAddRequest {
    s.ServiceType = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetHotelType(v int64) *TaobaoXhotelAddRequest {
    s.HotelType = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetCoordinateSystem(v string) *TaobaoXhotelAddRequest {
    s.CoordinateSystem = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetRoomFacilities(v string) *TaobaoXhotelAddRequest {
    s.RoomFacilities = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetService(v string) *TaobaoXhotelAddRequest {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetHotelFacilities(v string) *TaobaoXhotelAddRequest {
    s.HotelFacilities = &v
    return s
}
func (s *TaobaoXhotelAddRequest) SetHotelPolicies(v string) *TaobaoXhotelAddRequest {
    s.HotelPolicies = &v
    return s
}

func (req *TaobaoXhotelAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
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
    if(req.Shid != nil) {
        paramMap["shid"] = *req.Shid
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
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
    if(req.Pics != nil) {
        paramMap["pics"] = *req.Pics
    }
    if(req.Brand != nil) {
        paramMap["brand"] = *req.Brand
    }
    if(req.PostalCode != nil) {
        paramMap["postal_code"] = *req.PostalCode
    }
    if(req.BookingNotice != nil) {
        paramMap["booking_notice"] = *req.BookingNotice
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
    if(req.Supplier != nil) {
        paramMap["supplier"] = *req.Supplier
    }
    if(req.SettlementCurrency != nil) {
        paramMap["settlement_currency"] = *req.SettlementCurrency
    }
    if(req.StandardAmuseFacilities != nil) {
        paramMap["standard_amuse_facilities"] = *req.StandardAmuseFacilities
    }
    if(req.StandardRoomFacilities != nil) {
        paramMap["standard_room_facilities"] = *req.StandardRoomFacilities
    }
    if(req.StandardHotelService != nil) {
        paramMap["standard_hotel_service"] = *req.StandardHotelService
    }
    if(req.StandardHotelFacilities != nil) {
        paramMap["standard_hotel_facilities"] = *req.StandardHotelFacilities
    }
    if(req.StandardBookingNotice != nil) {
        paramMap["standard_booking_notice"] = *req.StandardBookingNotice
    }
    if(req.ServiceType != nil) {
        paramMap["service_type"] = *req.ServiceType
    }
    if(req.HotelType != nil) {
        paramMap["hotel_type"] = *req.HotelType
    }
    if(req.CoordinateSystem != nil) {
        paramMap["coordinate_system"] = *req.CoordinateSystem
    }
    if(req.RoomFacilities != nil) {
        paramMap["room_facilities"] = *req.RoomFacilities
    }
    if(req.Service != nil) {
        paramMap["service"] = *req.Service
    }
    if(req.HotelFacilities != nil) {
        paramMap["hotel_facilities"] = *req.HotelFacilities
    }
    if(req.HotelPolicies != nil) {
        paramMap["hotel_policies"] = *req.HotelPolicies
    }
    return paramMap
}

func (req *TaobaoXhotelAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}