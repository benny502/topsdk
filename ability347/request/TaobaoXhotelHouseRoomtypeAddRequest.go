package request


type TaobaoXhotelHouseRoomtypeAddRequest struct {
    /*
        （已废弃）请使用outHid     */
    Hid  *int64 `json:"hid,omitempty" required:"false" `
    /*
        房型名称。不能超过30字     */
    Name  *string `json:"name" required:"true" `
    /*
        最大入住人数，默认2（1-99） defalutValue��2    */
    MaxOccupancy  *int64 `json:"max_occupancy,omitempty" required:"false" `
    /*
        具体面积大小，请按照正确格式填写。两种格式，例如：40或者 10-20     */
    Area  *string `json:"area,omitempty" required:"false" `
    /*
        客房在建筑的第几层，隔层为1-2层，4-5层，7-8层     */
    Floor  *string `json:"floor,omitempty" required:"false" `
    /*
        宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带     */
    Internet  *string `json:"internet,omitempty" required:"false" `
    /*
        设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}     */
    Service  *string `json:"service,omitempty" required:"false" `
    /*
        不要使用     */
    Extend  *string `json:"extend,omitempty" required:"false" `
    /*
        卖家房型ID，不能重复建议格式是:酒店code_房型code     */
    OuterId  *string `json:"outer_id" required:"true" `
    /*
        0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗     */
    WindowType  *int64 `json:"window_type,omitempty" required:"false" `
    /*
        该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。     */
    Srid  *int64 `json:"srid,omitempty" required:"false" `
    /*
        系统商，无申请不可使用     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        （必传）商家酒店ID，指明该房型属于哪家酒店     */
    OutHid  *string `json:"out_hid,omitempty" required:"false" `
    /*
        房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。     */
    Pics  *string `json:"pics,omitempty" required:"false" `
    /*
        卖家房型英文名称     */
    NameE  *string `json:"name_e,omitempty" required:"false" `
    /*
        操作人信息     */
    Operator  *string `json:"operator,omitempty" required:"false" `
    /*
        属性值为1: 含义是非直连房型     */
    ConnectionType  *int64 `json:"connection_type,omitempty" required:"false" `
    /*
        房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房     */
    HouseModel  *string `json:"house_model,omitempty" required:"false" `
    /*
        房屋面积     */
    HouseSize  *int64 `json:"house_size,omitempty" required:"false" `
    /*
        出租类型:1.整租;2.单间;3.床位     */
    RentType  *int64 `json:"rent_type,omitempty" required:"false" `
    /*
        出租面积,单位平方米     */
    RentSize  *int64 `json:"rent_size,omitempty" required:"false" `
    /*
        是否和房东合住:0.不和房东合住;1.和房东合住;     */
    HasLandlord  *int64 `json:"has_landlord,omitempty" required:"false" `
    /*
        床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347     */
    BedInfo  *string `json:"bed_info,omitempty" required:"false" `
    /*
        数据状态 0:正常，-1:删除，-2:停售     */
    Status  *int64 `json:"status,omitempty" required:"false" `
}

func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetHid(v int64) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetName(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetMaxOccupancy(v int64) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.MaxOccupancy = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetArea(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.Area = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetFloor(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.Floor = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetInternet(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.Internet = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetService(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetExtend(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetOuterId(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetWindowType(v int64) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.WindowType = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetSrid(v int64) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.Srid = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetVendor(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetOutHid(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.OutHid = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetPics(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.Pics = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetNameE(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.NameE = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetOperator(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.Operator = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetConnectionType(v int64) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.ConnectionType = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetHouseModel(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.HouseModel = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetHouseSize(v int64) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.HouseSize = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetRentType(v int64) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.RentType = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetRentSize(v int64) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.RentSize = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetHasLandlord(v int64) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.HasLandlord = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetBedInfo(v string) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.BedInfo = &v
    return s
}
func (s *TaobaoXhotelHouseRoomtypeAddRequest) SetStatus(v int64) *TaobaoXhotelHouseRoomtypeAddRequest {
    s.Status = &v
    return s
}

func (req *TaobaoXhotelHouseRoomtypeAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Hid != nil) {
        paramMap["hid"] = *req.Hid
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.MaxOccupancy != nil) {
        paramMap["max_occupancy"] = *req.MaxOccupancy
    }
    if(req.Area != nil) {
        paramMap["area"] = *req.Area
    }
    if(req.Floor != nil) {
        paramMap["floor"] = *req.Floor
    }
    if(req.Internet != nil) {
        paramMap["internet"] = *req.Internet
    }
    if(req.Service != nil) {
        paramMap["service"] = *req.Service
    }
    if(req.Extend != nil) {
        paramMap["extend"] = *req.Extend
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.WindowType != nil) {
        paramMap["window_type"] = *req.WindowType
    }
    if(req.Srid != nil) {
        paramMap["srid"] = *req.Srid
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.OutHid != nil) {
        paramMap["out_hid"] = *req.OutHid
    }
    if(req.Pics != nil) {
        paramMap["pics"] = *req.Pics
    }
    if(req.NameE != nil) {
        paramMap["name_e"] = *req.NameE
    }
    if(req.Operator != nil) {
        paramMap["operator"] = *req.Operator
    }
    if(req.ConnectionType != nil) {
        paramMap["connection_type"] = *req.ConnectionType
    }
    if(req.HouseModel != nil) {
        paramMap["house_model"] = *req.HouseModel
    }
    if(req.HouseSize != nil) {
        paramMap["house_size"] = *req.HouseSize
    }
    if(req.RentType != nil) {
        paramMap["rent_type"] = *req.RentType
    }
    if(req.RentSize != nil) {
        paramMap["rent_size"] = *req.RentSize
    }
    if(req.HasLandlord != nil) {
        paramMap["has_landlord"] = *req.HasLandlord
    }
    if(req.BedInfo != nil) {
        paramMap["bed_info"] = *req.BedInfo
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    return paramMap
}

func (req *TaobaoXhotelHouseRoomtypeAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}