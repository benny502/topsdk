package response

import (
    "topsdk/ability359/domain"
)

type TaobaoXhotelOrderSearchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        订单信息
    */
    HotelOrders  []domain.TaobaoXhotelOrderSearchXHotelOrder `json:"hotel_orders,omitempty" `
    /*
        符合条件的结果总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
}
