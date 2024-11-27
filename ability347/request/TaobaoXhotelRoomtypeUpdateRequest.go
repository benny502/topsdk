package request


type TaobaoXhotelRoomtypeUpdateRequest struct {
    /*
        （已废弃）     */
    Rid  *int64 `json:"rid,omitempty" required:"false" `
    /*
        房型名称。不能超过200字；添加房型时为必须     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        最大入住人数，默认2（1-99） defalutValue��2    */
    MaxOccupancy  *int64 `json:"max_occupancy,omitempty" required:"false" `
    /*
        具体面积大小，请按照正确格式填写。两种格式，例如：40 或者 10-20     */
    Area  *string `json:"area,omitempty" required:"false" `
    /*
        客房在建筑的第几层，隔层为1-2层，4-5层，7-8层     */
    Floor  *string `json:"floor,omitempty" required:"false" `
    /*
        （已废弃）床型请使用bed_info字段     */
    BedType  *string `json:"bed_type,omitempty" required:"false" `
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
        0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗     */
    WindowType  *int64 `json:"window_type,omitempty" required:"false" `
    /*
        扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析     */
    Extend  *string `json:"extend,omitempty" required:"false" `
    /*
        （必传）商家房型ID     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。     */
    Srid  *int64 `json:"srid,omitempty" required:"false" `
    /*
        系统商，不要使用，只有申请才可用     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        （已废弃）     */
    Hid  *int64 `json:"hid,omitempty" required:"false" `
    /*
        商家酒店ID(如果更新房型的时候房型不存在，会拿该code去新增房型)     */
    HotelCode  *string `json:"hotel_code,omitempty" required:"false" `
    /*
        房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。要求：无logo、水印、边框、人物，不模糊，不重复，不歪斜，房间图清晰，图片尺寸不小于300*225，不小于5M     */
    Pics  *string `json:"pics,omitempty" required:"false" `
    /*
        房型状态。0:正常，-1:删除，-2:停售     */
    Status  *byte `json:"status,omitempty" required:"false" `
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
        新的房型编码，请确实需要修改某个房型的编码的时候才使用，如需使用，请联系飞猪技术支持开通权限，否则会更新失败     */
    NewOuterId  *string `json:"new_outer_id,omitempty" required:"false" `
    /*
        房间设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891     */
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

func (s *TaobaoXhotelRoomtypeUpdateRequest) SetRid(v int64) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Rid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetName(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Name = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetMaxOccupancy(v int64) *TaobaoXhotelRoomtypeUpdateRequest {
    s.MaxOccupancy = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetArea(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Area = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetFloor(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Floor = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetBedType(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.BedType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetBedSize(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.BedSize = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetInternet(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Internet = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetService(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Service = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetWindowType(v int64) *TaobaoXhotelRoomtypeUpdateRequest {
    s.WindowType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetExtend(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Extend = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetOuterId(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetSrid(v int64) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Srid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetVendor(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetHid(v int64) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetHotelCode(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.HotelCode = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetPics(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Pics = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetStatus(v byte) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetNameE(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.NameE = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetOperator(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.Operator = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetConnectionType(v int64) *TaobaoXhotelRoomtypeUpdateRequest {
    s.ConnectionType = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetBedInfo(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.BedInfo = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetNewOuterId(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.NewOuterId = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetStandardRoomFacilities(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.StandardRoomFacilities = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetWindowDesc(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.WindowDesc = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetAddBed(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.AddBed = &v
    return s
}
func (s *TaobaoXhotelRoomtypeUpdateRequest) SetChildrenPolicy(v string) *TaobaoXhotelRoomtypeUpdateRequest {
    s.ChildrenPolicy = &v
    return s
}

func (req *TaobaoXhotelRoomtypeUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Rid != nil) {
        paramMap["rid"] = *req.Rid
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
    if(req.WindowType != nil) {
        paramMap["window_type"] = *req.WindowType
    }
    if(req.Extend != nil) {
        paramMap["extend"] = *req.Extend
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Srid != nil) {
        paramMap["srid"] = *req.Srid
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.Hid != nil) {
        paramMap["hid"] = *req.Hid
    }
    if(req.HotelCode != nil) {
        paramMap["hotel_code"] = *req.HotelCode
    }
    if(req.Pics != nil) {
        paramMap["pics"] = *req.Pics
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
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
    if(req.NewOuterId != nil) {
        paramMap["new_outer_id"] = *req.NewOuterId
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

func (req *TaobaoXhotelRoomtypeUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}