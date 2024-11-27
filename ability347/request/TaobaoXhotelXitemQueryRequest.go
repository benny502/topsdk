package request

import (
        "topsdk/util"
    )

type TaobaoXhotelXitemQueryRequest struct {
    /*
        系统商，一般不填写，使用须申请     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        商家酒店ID，指明该 x 元素属于哪家酒店     */
    OutHid  *string `json:"out_hid" required:"true" `
    /*
        商家房型ID，指明该 x 元素关联哪个房型     */
    OutRid  *string `json:"out_rid,omitempty" required:"false" `
    /*
        商家 RP ID，指明该 x 元素关联了哪个 RP     */
    RatePlanCode  *string `json:"rate_plan_code,omitempty" required:"false" `
    /*
        需要查询的 x_code 编码     */
    OutXCodes  *[]string `json:"out_x_codes,omitempty" required:"false" `
}

func (s *TaobaoXhotelXitemQueryRequest) SetVendor(v string) *TaobaoXhotelXitemQueryRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelXitemQueryRequest) SetOutHid(v string) *TaobaoXhotelXitemQueryRequest {
    s.OutHid = &v
    return s
}
func (s *TaobaoXhotelXitemQueryRequest) SetOutRid(v string) *TaobaoXhotelXitemQueryRequest {
    s.OutRid = &v
    return s
}
func (s *TaobaoXhotelXitemQueryRequest) SetRatePlanCode(v string) *TaobaoXhotelXitemQueryRequest {
    s.RatePlanCode = &v
    return s
}
func (s *TaobaoXhotelXitemQueryRequest) SetOutXCodes(v []string) *TaobaoXhotelXitemQueryRequest {
    s.OutXCodes = &v
    return s
}

func (req *TaobaoXhotelXitemQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.OutHid != nil) {
        paramMap["out_hid"] = *req.OutHid
    }
    if(req.OutRid != nil) {
        paramMap["out_rid"] = *req.OutRid
    }
    if(req.RatePlanCode != nil) {
        paramMap["rate_plan_code"] = *req.RatePlanCode
    }
    if(req.OutXCodes != nil) {
        paramMap["out_x_codes"] = util.ConvertBasicList(*req.OutXCodes)
    }
    return paramMap
}

func (req *TaobaoXhotelXitemQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}