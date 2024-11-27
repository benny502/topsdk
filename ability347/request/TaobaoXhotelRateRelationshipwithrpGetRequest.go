package request


type TaobaoXhotelRateRelationshipwithrpGetRequest struct {
    /*
        宝贝的gid     */
    Gid  *int64 `json:"gid" required:"true" `
    /*
        页数，可根据此值展示某页的数据。不填默认为1 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
}

func (s *TaobaoXhotelRateRelationshipwithrpGetRequest) SetGid(v int64) *TaobaoXhotelRateRelationshipwithrpGetRequest {
    s.Gid = &v
    return s
}
func (s *TaobaoXhotelRateRelationshipwithrpGetRequest) SetPageNo(v int64) *TaobaoXhotelRateRelationshipwithrpGetRequest {
    s.PageNo = &v
    return s
}

func (req *TaobaoXhotelRateRelationshipwithrpGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Gid != nil) {
        paramMap["gid"] = *req.Gid
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    return paramMap
}

func (req *TaobaoXhotelRateRelationshipwithrpGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}