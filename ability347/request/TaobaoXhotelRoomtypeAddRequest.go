package request


type TaobaoXhotelRoomtypeAddRequest struct {
    /*
        （已废弃）请使用outHid     */
    Hid  *int64 `json:"hid,omitempty" required:"false" `
    /*
        房型名称。不能超过200个字符     */
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
        （已废弃）床型请使用bed_info字段     */
    BedType  *string `json:"bed_type" required:"true" `
    /*
        床宽。按自己定义存储，比如：2.1米     */
    BedSize  *string `json:"bed_size,omitempty" required:"false" `
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
        main_bed_type母床型,sub_bed_type子床型。详情参见文档： https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.yN2mES&docType=1&articleId=118712&previewCode=1DABB73EA935608455E203BA06CF3566     */
    BedInfo  *string `json:"bed_info,omitempty" required:"false" `
    /*
        酒店房型设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891     */
    StandardRoomFacilities  *string `json:"standard_room_facilities,omitempty" required:"false" `
    /*
        窗型，窗型（windowType）： 0=无窗 1=有窗 2=部分有窗  窗型缺陷（windowTypeDefect）： 0=窗户不可打开通风 1=窗外有遮挡 2=窗外为酒店内景观  特殊窗（windowTypeSpecial）： 0=落地窗 1=飘窗 2=天窗 3=小窗。  当为有窗或部分有窗时，窗型缺陷可多选，特殊窗型需单选     */
    WindowDesc  *string `json:"window_desc,omitempty" required:"false" `
    /*
        房型加床政策。 费用(fee) 说明(desc)     */
    AddBed  *string `json:"add_bed,omitempty" required:"false" `
    /*
        儿童政策     */
    ChildrenPolicy  *string `json:"children_policy,omitempty" required:"false" `
}

func (s *TaobaoXhotelRoomtypeAddRequest) SetHid(v int64) *TaobaoXhotelRoomtypeAddRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetName(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetMaxOccupancy(v int64) *TaobaoXhotelRoomtypeAddRequest {
    s.MaxOccupancy = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetArea(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.Area = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetFloor(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.Floor = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetBedType(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.BedType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetBedSize(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.BedSize = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetInternet(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.Internet = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetService(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetExtend(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetOuterId(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetWindowType(v int64) *TaobaoXhotelRoomtypeAddRequest {
    s.WindowType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetSrid(v int64) *TaobaoXhotelRoomtypeAddRequest {
    s.Srid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetVendor(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetOutHid(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.OutHid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetPics(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.Pics = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetNameE(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.NameE = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetOperator(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.Operator = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetConnectionType(v int64) *TaobaoXhotelRoomtypeAddRequest {
    s.ConnectionType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetBedInfo(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.BedInfo = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetStandardRoomFacilities(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.StandardRoomFacilities = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetWindowDesc(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.WindowDesc = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetAddBed(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.AddBed = &v
    return s
}
func (s *TaobaoXhotelRoomtypeAddRequest) SetChildrenPolicy(v string) *TaobaoXhotelRoomtypeAddRequest {
    s.ChildrenPolicy = &v
    return s
}

func (req *TaobaoXhotelRoomtypeAddRequest) ToMap() map[string]interface{} {
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
    if(req.BedType != nil) {
        paramMap["bed_type"] = *req.BedType
    }
    if(req.BedSize != nil) {
        paramMap["bed_size"] = *req.BedSize
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
    if(req.BedInfo != nil) {
        paramMap["bed_info"] = *req.BedInfo
    }
    if(req.StandardRoomFacilities != nil) {
        paramMap["standard_room_facilities"] = *req.StandardRoomFacilities
    }
    if(req.WindowDesc != nil) {
        paramMap["window_desc"] = *req.WindowDesc
    }
    if(req.AddBed != nil) {
        paramMap["add_bed"] = *req.AddBed
    }
    if(req.ChildrenPolicy != nil) {
        paramMap["children_policy"] = *req.ChildrenPolicy
    }
    return paramMap
}

func (req *TaobaoXhotelRoomtypeAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}