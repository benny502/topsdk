package response

import (
)

type TaobaoXhotelRateRelationshipwithroomGetResponse struct {

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
        返回值
    */
    Gids  []string `json:"gids,omitempty" `
}
