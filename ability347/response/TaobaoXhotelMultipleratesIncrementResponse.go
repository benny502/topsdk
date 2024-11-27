package response

import (
)

type TaobaoXhotelMultipleratesIncrementResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品id-房价id-入住人数-连住天数  的集合
    */
    GidAndRpidOccupancyLengthofstay  []string `json:"gid_and_rpid_occupancy_lengthofstay,omitempty" `
    /*
        批量更新的时候，如果部分更新失败，会展示部分失败的原因
    */
    Warnmessage  string `json:"warnmessage,omitempty" `
}
