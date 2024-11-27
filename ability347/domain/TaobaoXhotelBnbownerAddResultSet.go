package domain


type TaobaoXhotelBnbownerAddResultSet struct {
    /*
        firstResult     */
    FirstResult  *TaobaoXhotelBnbownerAddAddOwnerParam `json:"first_result,omitempty" `

}

func (s *TaobaoXhotelBnbownerAddResultSet) SetFirstResult(v TaobaoXhotelBnbownerAddAddOwnerParam) *TaobaoXhotelBnbownerAddResultSet {
    s.FirstResult = &v
    return s
}
