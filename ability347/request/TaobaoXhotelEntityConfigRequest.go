package request


type TaobaoXhotelEntityConfigRequest struct {
    /*
        实体编码，例如酒店编码，房价编码     */
    EntityCode  *string `json:"entity_code,omitempty" required:"false" `
    /*
        格式为：[{"invoice_provider":"2"}] key-value形式的配置信息其中invoice_provider发票提供方1(酒店提供),2(卖家邮寄)。其余开票信息统一在卖家或酒店后台手动配置，此接口不做参数传输。     */
    ConfigData  *string `json:"config_data,omitempty" required:"false" `
    /*
        默认是taobao。和酒店，房型等实体上的vendor保持一致。     */
    Vendor  *string `json:"vendor,omitempty" required:"false" `
    /*
        实体id：飞猪平台的id，目前仅支持hid，rpid     */
    EntityId  *int64 `json:"entity_id,omitempty" required:"false" `
    /*
        目前仅支持 5：rp维度 defalutValue��0    */
    Type  *int64 `json:"type,omitempty" required:"false" `
}

func (s *TaobaoXhotelEntityConfigRequest) SetEntityCode(v string) *TaobaoXhotelEntityConfigRequest {
    s.EntityCode = &v
    return s
}
func (s *TaobaoXhotelEntityConfigRequest) SetConfigData(v string) *TaobaoXhotelEntityConfigRequest {
    s.ConfigData = &v
    return s
}
func (s *TaobaoXhotelEntityConfigRequest) SetVendor(v string) *TaobaoXhotelEntityConfigRequest {
    s.Vendor = &v
    return s
}
func (s *TaobaoXhotelEntityConfigRequest) SetEntityId(v int64) *TaobaoXhotelEntityConfigRequest {
    s.EntityId = &v
    return s
}
func (s *TaobaoXhotelEntityConfigRequest) SetType(v int64) *TaobaoXhotelEntityConfigRequest {
    s.Type = &v
    return s
}

func (req *TaobaoXhotelEntityConfigRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.EntityCode != nil) {
        paramMap["entity_code"] = *req.EntityCode
    }
    if(req.ConfigData != nil) {
        paramMap["config_data"] = *req.ConfigData
    }
    if(req.Vendor != nil) {
        paramMap["vendor"] = *req.Vendor
    }
    if(req.EntityId != nil) {
        paramMap["entity_id"] = *req.EntityId
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    return paramMap
}

func (req *TaobaoXhotelEntityConfigRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}