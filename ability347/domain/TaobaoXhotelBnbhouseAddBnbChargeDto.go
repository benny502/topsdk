package domain


type TaobaoXhotelBnbhouseAddBnbChargeDto struct {
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

func (s *TaobaoXhotelBnbhouseAddBnbChargeDto) SetNum(v int64) *TaobaoXhotelBnbhouseAddBnbChargeDto {
    s.Num = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbChargeDto) SetFee(v int64) *TaobaoXhotelBnbhouseAddBnbChargeDto {
    s.Fee = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbChargeDto) SetAddPeople(v int64) *TaobaoXhotelBnbhouseAddBnbChargeDto {
    s.AddPeople = &v
    return s
}
func (s *TaobaoXhotelBnbhouseAddBnbChargeDto) SetMinChargingAge(v int64) *TaobaoXhotelBnbhouseAddBnbChargeDto {
    s.MinChargingAge = &v
    return s
}
