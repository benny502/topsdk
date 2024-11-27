package response

import (
)

type TaobaoXhotelMultiplerateUpdateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        gid-rpid-occupancy-lengthofstay
商品ID-房价ID-入住人数-连住天数
    */
    GidAndRpidOccupancyLengthofstay  string `json:"gid_and_rpid_occupancy_lengthofstay,omitempty" `
}
