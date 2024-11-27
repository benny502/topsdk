package request

import (
        "topsdk/util"
    )

type TaobaoXhotelQuotaUpdateRequest struct {
    /*
        库存类型,0: 普通库存, 1: 普通保留房库存, 2:协议保留房库存     */
    QuotaType  *int64 `json:"quota_type" required:"true" `
    /*
        是否使用room库存,true使用，false不使用 defalutValue��false    */
    UseRoomInventory  *bool `json:"use_room_inventory,omitempty" required:"false" `
    /*
        room的gid defalutValue��0    */
    Gid  *int64 `json:"gid,omitempty" required:"false" `
    /*
        增减的值     */
    Quota  *int64 `json:"quota" required:"true" `
    /*
        数量类型, 2:增加房量,3:减少房量     */
    QuotaNumType  *int64 `json:"quota_num_type" required:"true" `
    /*
        修改日期列表     */
    Dates  *[]util.LocalTime `json:"dates" required:"true" `
    /*
        rate的id, rate库存时必填 defalutValue��0    */
    RateId  *int64 `json:"rate_id,omitempty" required:"false" `
}

func (s *TaobaoXhotelQuotaUpdateRequest) SetQuotaType(v int64) *TaobaoXhotelQuotaUpdateRequest {
    s.QuotaType = &v
    return s
}
func (s *TaobaoXhotelQuotaUpdateRequest) SetUseRoomInventory(v bool) *TaobaoXhotelQuotaUpdateRequest {
    s.UseRoomInventory = &v
    return s
}
func (s *TaobaoXhotelQuotaUpdateRequest) SetGid(v int64) *TaobaoXhotelQuotaUpdateRequest {
    s.Gid = &v
    return s
}
func (s *TaobaoXhotelQuotaUpdateRequest) SetQuota(v int64) *TaobaoXhotelQuotaUpdateRequest {
    s.Quota = &v
    return s
}
func (s *TaobaoXhotelQuotaUpdateRequest) SetQuotaNumType(v int64) *TaobaoXhotelQuotaUpdateRequest {
    s.QuotaNumType = &v
    return s
}
func (s *TaobaoXhotelQuotaUpdateRequest) SetDates(v []util.LocalTime) *TaobaoXhotelQuotaUpdateRequest {
    s.Dates = &v
    return s
}
func (s *TaobaoXhotelQuotaUpdateRequest) SetRateId(v int64) *TaobaoXhotelQuotaUpdateRequest {
    s.RateId = &v
    return s
}

func (req *TaobaoXhotelQuotaUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.QuotaType != nil) {
        paramMap["quota_type"] = *req.QuotaType
    }
    if(req.UseRoomInventory != nil) {
        paramMap["use_room_inventory"] = *req.UseRoomInventory
    }
    if(req.Gid != nil) {
        paramMap["gid"] = *req.Gid
    }
    if(req.Quota != nil) {
        paramMap["quota"] = *req.Quota
    }
    if(req.QuotaNumType != nil) {
        paramMap["quota_num_type"] = *req.QuotaNumType
    }
    if(req.Dates != nil) {
        paramMap["dates"] = util.ConvertBasicList(*req.Dates)
    }
    if(req.RateId != nil) {
        paramMap["rate_id"] = *req.RateId
    }
    return paramMap
}

func (req *TaobaoXhotelQuotaUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}