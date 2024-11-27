package response

import (
)

type TaobaoXhotelRateRelationshipwithrpGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        根据条件所查询的所有结果的总数量
    */
    TotalResults  int64 `json:"total_results,omitempty" `
    /*
        所查询出的结果，是一个字符串数组
    */
    RpIds  []string `json:"rp_ids,omitempty" `
}
