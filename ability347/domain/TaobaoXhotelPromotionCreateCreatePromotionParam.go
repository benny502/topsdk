package domain


type TaobaoXhotelPromotionCreateCreatePromotionParam struct {
    /*
        促销规则Id,传值代表修改，不传代表新建     */
    PromotionRuleId  *int64 `json:"promotion_rule_id,omitempty" `

    /*
        当前一个活动ID只可传1条规则     */
    PromotionRules  *[]TaobaoXhotelPromotionCreatePromotionRule `json:"promotion_rules,omitempty" `

    /*
        折扣类型，1 代表折扣，2 代表立减，创建后不支持修改     */
    DiscountType  *int64 `json:"discount_type,omitempty" `

    /*
        促销类别代码: [天天特惠 4151]， [早订优惠 4154]， [连住优惠 4153]， [限时特惠 4152]， [今夜甩卖 4367]， [门店新客 4155]， [出行特惠 4322] 活动创建成功后，不可修改活动类型。     */
    PromotionCode  *string `json:"promotion_code,omitempty" `

}

func (s *TaobaoXhotelPromotionCreateCreatePromotionParam) SetPromotionRuleId(v int64) *TaobaoXhotelPromotionCreateCreatePromotionParam {
    s.PromotionRuleId = &v
    return s
}
func (s *TaobaoXhotelPromotionCreateCreatePromotionParam) SetPromotionRules(v []TaobaoXhotelPromotionCreatePromotionRule) *TaobaoXhotelPromotionCreateCreatePromotionParam {
    s.PromotionRules = &v
    return s
}
func (s *TaobaoXhotelPromotionCreateCreatePromotionParam) SetDiscountType(v int64) *TaobaoXhotelPromotionCreateCreatePromotionParam {
    s.DiscountType = &v
    return s
}
func (s *TaobaoXhotelPromotionCreateCreatePromotionParam) SetPromotionCode(v string) *TaobaoXhotelPromotionCreateCreatePromotionParam {
    s.PromotionCode = &v
    return s
}
