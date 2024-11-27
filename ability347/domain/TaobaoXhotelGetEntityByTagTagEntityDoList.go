package domain


type TaobaoXhotelGetEntityByTagTagEntityDoList struct {
    /*
        实体id     */
    EntityId  *int64 `json:"entity_id,omitempty" `

}

func (s *TaobaoXhotelGetEntityByTagTagEntityDoList) SetEntityId(v int64) *TaobaoXhotelGetEntityByTagTagEntityDoList {
    s.EntityId = &v
    return s
}
