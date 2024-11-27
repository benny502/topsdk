package response

import (
)

type TaobaoXhotelRatesUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        gid_and_rateplan_ids
    */
    GidAndRpids  []string `json:"gid_and_rpids,omitempty" `
}
