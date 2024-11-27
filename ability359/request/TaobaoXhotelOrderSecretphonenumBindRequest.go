package request

import (
	"github.com/benny502/topsdk/ability359/domain"
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelOrderSecretphonenumBindRequest struct {
	/*
	   隐私号绑定参数     */
	SecretPhoneNumberBindParam *domain.TaobaoXhotelOrderSecretphonenumBindSecretPhoneNumberBindParam `json:"secret_phone_number_bind_param" required:"true" `
}

func (s *TaobaoXhotelOrderSecretphonenumBindRequest) SetSecretPhoneNumberBindParam(v domain.TaobaoXhotelOrderSecretphonenumBindSecretPhoneNumberBindParam) *TaobaoXhotelOrderSecretphonenumBindRequest {
	s.SecretPhoneNumberBindParam = &v
	return s
}

func (req *TaobaoXhotelOrderSecretphonenumBindRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SecretPhoneNumberBindParam != nil {
		paramMap["secret_phone_number_bind_param"] = util.ConvertStruct(*req.SecretPhoneNumberBindParam)
	}
	return paramMap
}

func (req *TaobaoXhotelOrderSecretphonenumBindRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
