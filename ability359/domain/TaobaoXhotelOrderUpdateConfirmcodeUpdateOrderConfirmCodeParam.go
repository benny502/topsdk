package domain


type TaobaoXhotelOrderUpdateConfirmcodeUpdateOrderConfirmCodeParam struct {
    /*
        PMS确认号     */
    PmsResId  *string `json:"pms_res_id,omitempty" `

    /*
        飞猪订单号     */
    Tid  *int64 `json:"tid,omitempty" `

    /*
        商家系统订单号     */
    OutOrderId  *string `json:"out_order_id,omitempty" `

}

func (s *TaobaoXhotelOrderUpdateConfirmcodeUpdateOrderConfirmCodeParam) SetPmsResId(v string) *TaobaoXhotelOrderUpdateConfirmcodeUpdateOrderConfirmCodeParam {
    s.PmsResId = &v
    return s
}
func (s *TaobaoXhotelOrderUpdateConfirmcodeUpdateOrderConfirmCodeParam) SetTid(v int64) *TaobaoXhotelOrderUpdateConfirmcodeUpdateOrderConfirmCodeParam {
    s.Tid = &v
    return s
}
func (s *TaobaoXhotelOrderUpdateConfirmcodeUpdateOrderConfirmCodeParam) SetOutOrderId(v string) *TaobaoXhotelOrderUpdateConfirmcodeUpdateOrderConfirmCodeParam {
    s.OutOrderId = &v
    return s
}
