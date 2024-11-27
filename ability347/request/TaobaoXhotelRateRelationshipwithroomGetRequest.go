package request


type TaobaoXhotelRateRelationshipwithroomGetRequest struct {
    /*
        rpId     */
    RpId  *int64 `json:"rp_id" required:"true" `
    /*
        页数 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
}

func (s *TaobaoXhotelRateRelationshipwithroomGetRequest) SetRpId(v int64) *TaobaoXhotelRateRelationshipwithroomGetRequest {
    s.RpId = &v
    return s
}
func (s *TaobaoXhotelRateRelationshipwithroomGetRequest) SetPageNo(v int64) *TaobaoXhotelRateRelationshipwithroomGetRequest {
    s.PageNo = &v
    return s
}

func (req *TaobaoXhotelRateRelationshipwithroomGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RpId != nil) {
        paramMap["rp_id"] = *req.RpId
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    return paramMap
}

func (req *TaobaoXhotelRateRelationshipwithroomGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}