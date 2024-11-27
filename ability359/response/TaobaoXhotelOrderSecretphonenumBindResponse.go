package response

import (
)

type TaobaoXhotelOrderSecretphonenumBindResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        隐私号
    */
    SecretPhoneNumber  string `json:"secret_phone_number,omitempty" `
}
