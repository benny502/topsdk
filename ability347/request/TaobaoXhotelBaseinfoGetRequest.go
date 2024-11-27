package request


type TaobaoXhotelBaseinfoGetRequest struct {
    /*
        淘宝酒店ID     */
    Hid  *int64 `json:"hid,omitempty" required:"false" `
    /*
        推荐使用卖家系统中的酒店ID。     */
    OutHid  *string `json:"out_hid,omitempty" required:"false" `
    /*
        用于标示该酒店发布的渠道信息     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        是否需要房价基本信息（false为不需要），默认为需要 defalutValue��true    */
    IsNeedRatePlan  *bool `json:"is_need_rate_plan,omitempty" required:"false" `
    /*
        是否需要房型基本信息（false为不需要），默认为需要 defalutValue��true    */
    IsNeedRoomType  *bool `json:"is_need_room_type,omitempty" required:"false" `
    /*
        是否需要 根据 hid 查询 标准房型列表 defalutValue��false    */
    NeedSRoomTypeList  *bool `json:"need_s_room_type_list,omitempty" required:"false" `
    /*
        是否需要酒店房型可售详情 defalutValue��false    */
    NeedHotelDynamicInfo  *bool `json:"need_hotel_dynamic_info,omitempty" required:"false" `
    /*
        在查询酒店房型可售详情 时的入参JSON , {@link com.taobao.trip.hpc.client.query.HotelSellerInvQuery}     */
    JsonHotelSellerInvQuery  *string `json:"json_hotel_seller_inv_query,omitempty" required:"false" `
}

func (s *TaobaoXhotelBaseinfoGetRequest) SetHid(v int64) *TaobaoXhotelBaseinfoGetRequest {
    s.Hid = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRequest) SetOutHid(v string) *TaobaoXhotelBaseinfoGetRequest {
    s.OutHid = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRequest) SetVendor(v string) *TaobaoXhotelBaseinfoGetRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRequest) SetIsNeedRatePlan(v bool) *TaobaoXhotelBaseinfoGetRequest {
    s.IsNeedRatePlan = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRequest) SetIsNeedRoomType(v bool) *TaobaoXhotelBaseinfoGetRequest {
    s.IsNeedRoomType = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRequest) SetNeedSRoomTypeList(v bool) *TaobaoXhotelBaseinfoGetRequest {
    s.NeedSRoomTypeList = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRequest) SetNeedHotelDynamicInfo(v bool) *TaobaoXhotelBaseinfoGetRequest {
    s.NeedHotelDynamicInfo = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetRequest) SetJsonHotelSellerInvQuery(v string) *TaobaoXhotelBaseinfoGetRequest {
    s.JsonHotelSellerInvQuery = &v
    return s
}

func (req *TaobaoXhotelBaseinfoGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Hid != nil) {
        paramMap["hid"] = *req.Hid
    }
    if(req.OutHid != nil) {
        paramMap["out_hid"] = *req.OutHid
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.IsNeedRatePlan != nil) {
        paramMap["is_need_rate_plan"] = *req.IsNeedRatePlan
    }
    if(req.IsNeedRoomType != nil) {
        paramMap["is_need_room_type"] = *req.IsNeedRoomType
    }
    if(req.NeedSRoomTypeList != nil) {
        paramMap["need_s_room_type_list"] = *req.NeedSRoomTypeList
    }
    if(req.NeedHotelDynamicInfo != nil) {
        paramMap["need_hotel_dynamic_info"] = *req.NeedHotelDynamicInfo
    }
    if(req.JsonHotelSellerInvQuery != nil) {
        paramMap["json_hotel_seller_inv_query"] = *req.JsonHotelSellerInvQuery
    }
    return paramMap
}

func (req *TaobaoXhotelBaseinfoGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}