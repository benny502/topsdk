package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelXitemQueryHotelXitemDO struct {
    /*
        创建时间     */
    GmtCreate  *util.LocalTime `json:"gmt_create,omitempty" `

    /*
        修改时间     */
    GmtModified  *util.LocalTime `json:"gmt_modified,omitempty" `

    /*
        外部code     */
    OutXCode  *string `json:"out_x_code,omitempty" `

    /*
        子类型code     */
    SubTypeCode  *string `json:"sub_type_code,omitempty" `

    /*
        外部酒店code     */
    OutHid  *string `json:"out_hid,omitempty" `

    /*
        元素类型简写     */
    ShortName  *string `json:"short_name,omitempty" `

    /*
        服务时间段     */
    Time  *string `json:"time,omitempty" `

    /*
        商品价值     */
    Value  *int64 `json:"value,omitempty" `

    /*
        商品使用说明     */
    ItemDesc  *string `json:"item_desc,omitempty" `

    /*
        状态是否生效0 失效, 1生效     */
    Status  *int64 `json:"status,omitempty" `

    /*
        附加产品使用维度   1:每间房维度 2:每间夜维度     */
    DimensionType  *int64 `json:"dimension_type,omitempty" `

    /*
         审核状态-1：拒绝，0：审核中，1：审核通过     */
    AuditStatus  *int64 `json:"audit_status,omitempty" `

    /*
        审核拒绝原因     */
    AuditRejectReason  *string `json:"audit_reject_reason,omitempty" `

    /*
        详细信息json字符串     */
    FeatureDetail  *string `json:"feature_detail,omitempty" `

    /*
        酒+X 图片格式化信息     */
    Pictures  *[]TaobaoXhotelXitemQueryHotelXItemPicture `json:"pictures,omitempty" `

}

func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetGmtCreate(v util.LocalTime) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetGmtModified(v util.LocalTime) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.GmtModified = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetOutXCode(v string) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.OutXCode = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetSubTypeCode(v string) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.SubTypeCode = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetOutHid(v string) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.OutHid = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetShortName(v string) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.ShortName = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetTime(v string) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.Time = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetValue(v int64) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.Value = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetItemDesc(v string) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.ItemDesc = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetStatus(v int64) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.Status = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetDimensionType(v int64) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.DimensionType = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetAuditStatus(v int64) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.AuditStatus = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetAuditRejectReason(v string) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.AuditRejectReason = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetFeatureDetail(v string) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.FeatureDetail = &v
    return s
}
func (s *TaobaoXhotelXitemQueryHotelXitemDO) SetPictures(v []TaobaoXhotelXitemQueryHotelXItemPicture) *TaobaoXhotelXitemQueryHotelXitemDO {
    s.Pictures = &v
    return s
}
