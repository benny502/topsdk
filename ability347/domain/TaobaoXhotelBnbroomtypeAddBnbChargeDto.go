package domain


type TaobaoXhotelBnbroomtypeAddBnbChargeDto struct {
    /*
        允许加人数     */
    Num  *int64 `json:"num,omitempty" `

    /*
        加人费用金额     */
    Fee  *int64 `json:"fee,omitempty" `

    /*
        是否允许加人 0不允许 1允许     */
    AddPeople  *int64 `json:"add_people,omitempty" `

    /*
        最小收费年龄     */
    MinChargingAge  *int64 `json:"min_charging_age,omitempty" `

}

func (s *TaobaoXhotelBnbroomtypeAddBnbChargeDto) SetNum(v int64) *TaobaoXhotelBnbroomtypeAddBnbChargeDto {
    s.Num = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbChargeDto) SetFee(v int64) *TaobaoXhotelBnbroomtypeAddBnbChargeDto {
    s.Fee = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbChargeDto) SetAddPeople(v int64) *TaobaoXhotelBnbroomtypeAddBnbChargeDto {
    s.AddPeople = &v
    return s
}
func (s *TaobaoXhotelBnbroomtypeAddBnbChargeDto) SetMinChargingAge(v int64) *TaobaoXhotelBnbroomtypeAddBnbChargeDto {
    s.MinChargingAge = &v
    return s
}
