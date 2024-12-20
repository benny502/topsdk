package domain


type TaobaoXhotelBaseinfoGetXHotelBaseInfo struct {
    /*
        房型基础信息     */
    RoomTypeList  *[]TaobaoXhotelBaseinfoGetRoomType `json:"room_type_list,omitempty" `

    /*
        房价基础信息(需要新增rp时绑定酒店)     */
    RatePlanList  *[]TaobaoXhotelBaseinfoGetRatePlan `json:"rate_plan_list,omitempty" `

    /*
        卖家酒店对象     */
    Hotel  *TaobaoXhotelBaseinfoGetHotel `json:"hotel,omitempty" `

    /*
        标准房型列表     */
    SRoomTypeList  *[]TaobaoXhotelBaseinfoGetSroomtypelist `json:"s_room_type_list,omitempty" `

    /*
        酒店房型可售详情查询总数     */
    TotalHotelDynamicInfo  *int64 `json:"total_hotel_dynamic_info,omitempty" `

    /*
        酒店房型可售详情查     */
    HotelDynamicInfoList  *[]TaobaoXhotelBaseinfoGetHoteldynamicinfo `json:"hotel_dynamic_info_list,omitempty" `

    /*
        酒店维度特殊标签含义, json: {"pure-direct-hotel":0,"direct-manual-order-hotel":1,"ebk-direct-hotel":0,"non-direct-hotel":1,"allow-dingding-business-travel-hotel":1,"willing-dingding-bussiness-travel-hotel":0,"calendar-room-package-hotel":1,"dijiajiajia-hotel":0,"ebk-number-of-confirm-room-hotel":1} , key含义: pure-direct-hotel 表示纯直连酒店, direct-manual-order-hotel 和 ebk-direct-hotel 和 non-direct-hotel 这三个key对应value都是0 . direct-manual-order-hotel 表示 人工承接失败订单的酒店标签。如果某个直连酒店打了该标签，那么直连下单失败以后，允许人工承接订单，由人工跟进 . ebk-direct-hotel 表示 ebk直连酒店标。如果某个酒店打了该标签，那么这个酒店下允许通过ebk发布直连rp . non-direct-hotel 表示 卖家非直连酒店标签。如果某个酒店打了该标签，那么该酒店下单不会走直连交易。 allow-dingding-business-travel-hotel 表示 允许进入阿里商旅渠道（钉钉）售卖信用住的单体酒店 willing-dingding-bussiness-travel-hotel 表示 已签协议愿意加入阿里商旅渠道售卖信用住的单体酒店 . calendar-room-package-hotel 表示 酒店可以参加日历房套餐活动包括创建，修改，删除活动信息（高星集团GMV项目） dijiajiajia-hotel 表示 底价加价酒店权限标。只有打了该标的酒店才允许维护底价加价规则和包房rp . ebk-number-of-confirm-room-hotel 表示ebk确认订单，是否要输入外部确认号     */
    TagJson  *string `json:"tag_json,omitempty" `

}

func (s *TaobaoXhotelBaseinfoGetXHotelBaseInfo) SetRoomTypeList(v []TaobaoXhotelBaseinfoGetRoomType) *TaobaoXhotelBaseinfoGetXHotelBaseInfo {
    s.RoomTypeList = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetXHotelBaseInfo) SetRatePlanList(v []TaobaoXhotelBaseinfoGetRatePlan) *TaobaoXhotelBaseinfoGetXHotelBaseInfo {
    s.RatePlanList = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetXHotelBaseInfo) SetHotel(v TaobaoXhotelBaseinfoGetHotel) *TaobaoXhotelBaseinfoGetXHotelBaseInfo {
    s.Hotel = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetXHotelBaseInfo) SetSRoomTypeList(v []TaobaoXhotelBaseinfoGetSroomtypelist) *TaobaoXhotelBaseinfoGetXHotelBaseInfo {
    s.SRoomTypeList = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetXHotelBaseInfo) SetTotalHotelDynamicInfo(v int64) *TaobaoXhotelBaseinfoGetXHotelBaseInfo {
    s.TotalHotelDynamicInfo = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetXHotelBaseInfo) SetHotelDynamicInfoList(v []TaobaoXhotelBaseinfoGetHoteldynamicinfo) *TaobaoXhotelBaseinfoGetXHotelBaseInfo {
    s.HotelDynamicInfoList = &v
    return s
}
func (s *TaobaoXhotelBaseinfoGetXHotelBaseInfo) SetTagJson(v string) *TaobaoXhotelBaseinfoGetXHotelBaseInfo {
    s.TagJson = &v
    return s
}
