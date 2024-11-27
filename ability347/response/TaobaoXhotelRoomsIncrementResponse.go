package response

import (
)

type TaobaoXhotelRoomsIncrementResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        成功的gids LIST
    */
    Gids  []string `json:"gids,omitempty" `
}
