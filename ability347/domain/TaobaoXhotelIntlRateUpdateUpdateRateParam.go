package domain

import (
        "topsdk/util"
    )

type TaobaoXhotelIntlRateUpdateUpdateRateParam struct {
    /*
        rate更新列表     */
    UpdateRateDOList  *[]TaobaoXhotelIntlRateUpdateUpdateRateDo `json:"update_rate_d_o_list,omitempty" `

    /*
        过期时间     */
    ExpireTime  *util.LocalTime `json:"expire_time,omitempty" `

    /*
        供应商     */
    Supplier  *string `json:"supplier,omitempty" `

}

func (s *TaobaoXhotelIntlRateUpdateUpdateRateParam) SetUpdateRateDOList(v []TaobaoXhotelIntlRateUpdateUpdateRateDo) *TaobaoXhotelIntlRateUpdateUpdateRateParam {
    s.UpdateRateDOList = &v
    return s
}
func (s *TaobaoXhotelIntlRateUpdateUpdateRateParam) SetExpireTime(v util.LocalTime) *TaobaoXhotelIntlRateUpdateUpdateRateParam {
    s.ExpireTime = &v
    return s
}
func (s *TaobaoXhotelIntlRateUpdateUpdateRateParam) SetSupplier(v string) *TaobaoXhotelIntlRateUpdateUpdateRateParam {
    s.Supplier = &v
    return s
}
