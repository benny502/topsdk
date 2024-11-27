package domain


type TaobaoXhotelPromotionQueryDirectPromotion struct {
    /*
        促销规则Id     */
    PromotionRuleId  *int64 `json:"promotion_rule_id,omitempty" `

    /*
        促销规则     */
    PromotionRules  *[]TaobaoXhotelPromotionQueryPromotionRule `json:"promotion_rules,omitempty" `

    /*
        促销类别代码     */
    PromotionCode  *int64 `json:"promotion_code,omitempty" `

    /*
        活动绑定资源     */
    Resources  *[]TaobaoXhotelPromotionQueryPromotionResource `json:"resources,omitempty" `

    /*
        促销金额计算方式, 0 Percentage | 1 Fixed; 暂时支持固定金额 Fixed:固定金额 Percentage: 折扣比例     */
    DiscountType  *int64 `json:"discount_type,omitempty" `

}

func (s *TaobaoXhotelPromotionQueryDirectPromotion) SetPromotionRuleId(v int64) *TaobaoXhotelPromotionQueryDirectPromotion {
    s.PromotionRuleId = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryDirectPromotion) SetPromotionRules(v []TaobaoXhotelPromotionQueryPromotionRule) *TaobaoXhotelPromotionQueryDirectPromotion {
    s.PromotionRules = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryDirectPromotion) SetPromotionCode(v int64) *TaobaoXhotelPromotionQueryDirectPromotion {
    s.PromotionCode = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryDirectPromotion) SetResources(v []TaobaoXhotelPromotionQueryPromotionResource) *TaobaoXhotelPromotionQueryDirectPromotion {
    s.Resources = &v
    return s
}
func (s *TaobaoXhotelPromotionQueryDirectPromotion) SetDiscountType(v int64) *TaobaoXhotelPromotionQueryDirectPromotion {
    s.DiscountType = &v
    return s
}
