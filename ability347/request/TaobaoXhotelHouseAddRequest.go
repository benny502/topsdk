package request


type TaobaoXhotelHouseAddRequest struct {
    /*
        外部酒店ID, 这是卖家自己系统中的ID     */
    OuterId  *string `json:"outer_id" required:"true" `
    /*
        酒店名称     */
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
        酒店地址。长度不能超过120。不填入会导致不能自动匹配,此地址为下单前展示给用户地址     */
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
        酒店入住政策，{"10003":"仅2岁以上儿童可随同入住"}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=10、code=10003,value为文字描述     */
    HotelPolicies  *string `json:"hotel_policies,omitempty" required:"false" `
    /*
        酒店设施。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=20的分类     */
    HotelFacilities  *string `json:"hotel_facilities,omitempty" required:"false" `
    /*
        酒店服务。json格式示例值：{"24058":"可以接待外宾","24198":"叫醒服务","24200":"洗衣服务"}，key-24101为属性编码，value-为"true"时表示有该属性，为文字时表示具体描述，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类     */
    Service  *string `json:"service,omitempty" required:"false" `
    /*
        房间设施。json格式示例值：{"24101": true,"24091": true,"24095": true}，key-24101为属性编码，value-为"true"时表示有该属性，为文字时表示具体描述，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类     */
    RoomFacilities  *string `json:"room_facilities,omitempty" required:"false" `
    /*
        酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图（主图只能有一个，如果有多个或者没有，则会报错）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多是能是10张。     */
    Pics  *string `json:"pics,omitempty" required:"false" `
    /*
        酒店品牌。取值为数字。枚举如下（只给出top30，如果不满足，请联系去啊接口人）：    ruJia("1", "rujiakuaijie", "如家快捷", 1),    qiTian("2", "7 days", "7天连锁", 1),    hanTing("3", "Hanting Inns & Hotels", "汉庭酒店", 1),    geLinHaoTai("4", "Green Tree Inn", "格林豪泰", 1),    jinJiang("5", "Jinjiang Inn", "锦江之星", 1),    su8("6", "Super 8", "速8", 1),    moTai("7", "Motel", "莫泰", 1),    zhouji("8", "InterContinental", "洲际", 4),    budint("9", "Pod Inn", "布丁", 1),    jiuJiu("10", "jiujiuliansuo", "99连锁", 1),    piaoHome("11", "Piao Home Inn", "飘HOME", 1),    juzi("12", "Orange Hotels", "桔子酒店", 1),    yibai("13", "yibai", "易佰", 1),    weiyena("14","weiyena","维也纳",2),    huangguanjiari("15", "huangguanjiari", "皇冠假日", 4),    xidawu("16", "xidawu", "喜达屋", 3),    chengshiBJ("17", "chengshibianjie", "城市便捷", 1),    shagnKeYou("18", "shagnkeyou", "尚客优", 1),    jinjiang("19", "jinjiang", "锦江酒店", 3),    wendemu("20", "Hawthorn Suites", "温德姆", 4),    yibisi("21", "Ibis Hotels", "宜必思", 1),    wanhao("22", "JM Hoteles", "万豪", 4),    yijia365("23", "yijia365", "驿家365", 1),    shoulv("24", "shoulvjituan", "首旅建国", 3),    kaiyuan("25", "New Century Hotel", "开元大酒店", 4),    yagao("26", "yagao", "雅高", 3),    daisi("27", "daisi", "戴斯", 3),    jinling("28", "jinlingliansuo", "金陵", 4),    xianggelila("29", "Shangri-La City Hotels", "香格里拉", 4),    xierdun("30", "Hilton", "希尔顿", 4),     */
    Brand  *string `json:"brand,omitempty" required:"false" `
    /*
        邮政编码。     */
    PostalCode  *string `json:"postal_code,omitempty" required:"false" `
    /*
        预订须知。json格式，示例:{"10001":"14:00","10002":"12:00","10005":"清洁福50元","10006":"请准备好您的身份证件，我需要登记 不允许吸烟"},预订须知，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=10的分类     */
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
        供应商标识，需要提前开通权限，如果需要对接请联系技术支持，请谨慎使用     */
    Supplier  *string `json:"supplier,omitempty" required:"false" `
    /*
        结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用     */
    SettlementCurrency  *string `json:"settlement_currency,omitempty" required:"false" `
    /*
        房东信息,{"outerId: 外部房东ID,": "xxxx", "nickName": "张三", "avatarUrl": "http://test.com/1.jpg", "telephone": "0571-1234567", "mobilePhone": "12334567678", "email":"test@test.com", "gender": "F", "avgConfirmTime": 30, "responseRate": 100, "description": "房东太懒,什么也没有填", "birthday":"2018-01-01", "qualifacation": 1, "bloodType": 1, "profession":"交互设计师", "country":"CN", "province":"420000", "city":"421200", "real_name_status": true, "validate":"1,2,4,8","confirmRate": 98} JSON字段描述: outerId: 商家房东ID, nickName: 房东昵称, avatarUrl: 房东头像地址, telephone: 固定电话, mobilePhone: 移动电话, email: 邮箱地址, gender: 性别 M男性， F女性， avgConfirmTime: 平均确认时间, 单位分钟, responseRate: 房东回复率, description: 房东介绍, birthday:生日，格式yyyy-MM-dd, qualifacation:学历,1:小学,2:初中,3:高中,4:本科,5:硕士,6:博士,7:博士后,0:其他, profession: 职业 country: 国家code province: 省code city: 城市code realNameStatus: 实名认证状态, true已认证 validate: 认证情况:1:身份验证,2:头像验证,4:手机验证,8:邮箱验证,二进制各位代表含义, bloodType: 血型: 0未知,1:A型,2:B型,3:AB型,4:O型;confirmRate: 订单接单率，0-100     */
    OwnerInfo  *string `json:"owner_info,omitempty" required:"false" `
    /*
        描述信息，inside： 内部描述,traffic:交通情况,arround:周边情况     */
    ArroundDesc  *string `json:"arround_desc,omitempty" required:"false" `
    /*
        用户下单之后展示的详细地址     */
    RealAddress  *string `json:"real_address,omitempty" required:"false" `
    /*
        数据状态 0:正常，-1:删除，-2:停售     */
    Status  *int64 `json:"status,omitempty" required:"false" `
}

func (s *TaobaoXhotelHouseAddRequest) SetOuterId(v string) *TaobaoXhotelHouseAddRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetName(v string) *TaobaoXhotelHouseAddRequest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetUsedName(v string) *TaobaoXhotelHouseAddRequest {
    s.UsedName = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetDomestic(v int64) *TaobaoXhotelHouseAddRequest {
    s.Domestic = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetCountry(v string) *TaobaoXhotelHouseAddRequest {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetProvince(v int64) *TaobaoXhotelHouseAddRequest {
    s.Province = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetCity(v int64) *TaobaoXhotelHouseAddRequest {
    s.City = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetDistrict(v int64) *TaobaoXhotelHouseAddRequest {
    s.District = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetBusiness(v string) *TaobaoXhotelHouseAddRequest {
    s.Business = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetAddress(v string) *TaobaoXhotelHouseAddRequest {
    s.Address = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetLongitude(v string) *TaobaoXhotelHouseAddRequest {
    s.Longitude = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetLatitude(v string) *TaobaoXhotelHouseAddRequest {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetPositionType(v string) *TaobaoXhotelHouseAddRequest {
    s.PositionType = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetTel(v string) *TaobaoXhotelHouseAddRequest {
    s.Tel = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetExtend(v string) *TaobaoXhotelHouseAddRequest {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetShid(v int64) *TaobaoXhotelHouseAddRequest {
    s.Shid = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetVendor(v string) *TaobaoXhotelHouseAddRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetStar(v string) *TaobaoXhotelHouseAddRequest {
    s.Star = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetOpeningTime(v string) *TaobaoXhotelHouseAddRequest {
    s.OpeningTime = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetDecorateTime(v string) *TaobaoXhotelHouseAddRequest {
    s.DecorateTime = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetFloors(v string) *TaobaoXhotelHouseAddRequest {
    s.Floors = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetRooms(v int64) *TaobaoXhotelHouseAddRequest {
    s.Rooms = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetDescription(v string) *TaobaoXhotelHouseAddRequest {
    s.Description = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetHotelPolicies(v string) *TaobaoXhotelHouseAddRequest {
    s.HotelPolicies = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetHotelFacilities(v string) *TaobaoXhotelHouseAddRequest {
    s.HotelFacilities = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetService(v string) *TaobaoXhotelHouseAddRequest {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetRoomFacilities(v string) *TaobaoXhotelHouseAddRequest {
    s.RoomFacilities = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetPics(v string) *TaobaoXhotelHouseAddRequest {
    s.Pics = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetBrand(v string) *TaobaoXhotelHouseAddRequest {
    s.Brand = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetPostalCode(v string) *TaobaoXhotelHouseAddRequest {
    s.PostalCode = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetBookingNotice(v string) *TaobaoXhotelHouseAddRequest {
    s.BookingNotice = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetCreditCardTypes(v string) *TaobaoXhotelHouseAddRequest {
    s.CreditCardTypes = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetOrbitTrack(v string) *TaobaoXhotelHouseAddRequest {
    s.OrbitTrack = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetNameE(v string) *TaobaoXhotelHouseAddRequest {
    s.NameE = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetSupplier(v string) *TaobaoXhotelHouseAddRequest {
    s.Supplier = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetSettlementCurrency(v string) *TaobaoXhotelHouseAddRequest {
    s.SettlementCurrency = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetOwnerInfo(v string) *TaobaoXhotelHouseAddRequest {
    s.OwnerInfo = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetArroundDesc(v string) *TaobaoXhotelHouseAddRequest {
    s.ArroundDesc = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetRealAddress(v string) *TaobaoXhotelHouseAddRequest {
    s.RealAddress = &v
    return s
}
func (s *TaobaoXhotelHouseAddRequest) SetStatus(v int64) *TaobaoXhotelHouseAddRequest {
    s.Status = &v
    return s
}

func (req *TaobaoXhotelHouseAddRequest) ToMap() map[string]interface{} {
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
    if(req.HotelPolicies != nil) {
        paramMap["hotel_policies"] = *req.HotelPolicies
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
    if(req.OwnerInfo != nil) {
        paramMap["owner_info"] = *req.OwnerInfo
    }
    if(req.ArroundDesc != nil) {
        paramMap["arround_desc"] = *req.ArroundDesc
    }
    if(req.RealAddress != nil) {
        paramMap["real_address"] = *req.RealAddress
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    return paramMap
}

func (req *TaobaoXhotelHouseAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}