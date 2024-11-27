package request

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelXitemDeleteRequest struct {
	/*
	   系统商，一般不填写，使用须申请     */
	Vendor *string `json:"vendor,omitempty" required:"false" `
	/*
	   商家酒店ID，指明该 x 元素属于哪家酒店     */
	OutHid *string `json:"out_hid" required:"true" `
	/*
	   需要删除的 x_code 编码     */
	OutXCodes *[]string `json:"out_x_codes" required:"true" `
}

func (s *TaobaoXhotelXitemDeleteRequest) SetVendor(v string) *TaobaoXhotelXitemDeleteRequest {
	s.Vendor = &v
	return s
}
func (s *TaobaoXhotelXitemDeleteRequest) SetOutHid(v string) *TaobaoXhotelXitemDeleteRequest {
	s.OutHid = &v
	return s
}
func (s *TaobaoXhotelXitemDeleteRequest) SetOutXCodes(v []string) *TaobaoXhotelXitemDeleteRequest {
	s.OutXCodes = &v
	return s
}

func (req *TaobaoXhotelXitemDeleteRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Vendor != nil {
		paramMap["vendor"] = *req.Vendor
	}
	if req.OutHid != nil {
		paramMap["out_hid"] = *req.OutHid
	}
	if req.OutXCodes != nil {
		paramMap["out_x_codes"] = util.ConvertBasicList(*req.OutXCodes)
	}
	return paramMap
}

func (req *TaobaoXhotelXitemDeleteRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
