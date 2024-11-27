package ability362

import (
	"errors"
	"log"

	"github.com/benny502/topsdk"
	"github.com/benny502/topsdk/ability362/request"
	"github.com/benny502/topsdk/ability362/response"
	"github.com/benny502/topsdk/util"
)

type Ability362 struct {
	Client *topsdk.TopClient
}

func NewAbility362(client *topsdk.TopClient) *Ability362 {
	return &Ability362{client}
}

/*
创单(支付订单)通知
*/
func (ability *Ability362) AlitripTravelHotelticketOrderCreate(req *request.AlitripTravelHotelticketOrderCreateRequest) (*response.AlitripTravelHotelticketOrderCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability362 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alitrip.travel.hotelticket.order.create", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlitripTravelHotelticketOrderCreateResponse{}
	if err != nil {
		log.Println("alitripTravelHotelticketOrderCreate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
退款结结果通知
*/
func (ability *Ability362) AlitripTravelHotelticketOrderRefund(req *request.AlitripTravelHotelticketOrderRefundRequest) (*response.AlitripTravelHotelticketOrderRefundResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability362 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alitrip.travel.hotelticket.order.refund", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlitripTravelHotelticketOrderRefundResponse{}
	if err != nil {
		log.Println("alitripTravelHotelticketOrderRefund error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
订单核销通知
*/
func (ability *Ability362) AlitripTravelHotelticketOrderVerify(req *request.AlitripTravelHotelticketOrderVerifyRequest) (*response.AlitripTravelHotelticketOrderVerifyResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability362 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alitrip.travel.hotelticket.order.verify", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlitripTravelHotelticketOrderVerifyResponse{}
	if err != nil {
		log.Println("alitripTravelHotelticketOrderVerify error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
产品批量变更通知
*/
func (ability *Ability362) AlitripTravelHotelticketProductProductupdate(req *request.AlitripTravelHotelticketProductProductupdateRequest) (*response.AlitripTravelHotelticketProductProductupdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability362 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alitrip.travel.hotelticket.product.productupdate", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlitripTravelHotelticketProductProductupdateResponse{}
	if err != nil {
		log.Println("alitripTravelHotelticketProductProductupdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
产品批量变更推送通知
*/
func (ability *Ability362) AlitripTravelHotelticketProductProductupdatepush(req *request.AlitripTravelHotelticketProductProductupdatepushRequest) (*response.AlitripTravelHotelticketProductProductupdatepushResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability362 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alitrip.travel.hotelticket.product.productupdatepush", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlitripTravelHotelticketProductProductupdatepushResponse{}
	if err != nil {
		log.Println("alitripTravelHotelticketProductProductupdatepush error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
飞猪门票核销通知
*/
func (ability *Ability362) TaobaoTravelTicketOrderVerify(req *request.TaobaoTravelTicketOrderVerifyRequest) (*response.TaobaoTravelTicketOrderVerifyResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability362 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.travel.ticket.order.verify", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTravelTicketOrderVerifyResponse{}
	if err != nil {
		log.Println("taobaoTravelTicketOrderVerify error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
飞猪门票退票结果通知
*/
func (ability *Ability362) TaobaoTravelTicketOrderRefund(req *request.TaobaoTravelTicketOrderRefundRequest) (*response.TaobaoTravelTicketOrderRefundResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability362 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.travel.ticket.order.refund", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTravelTicketOrderRefundResponse{}
	if err != nil {
		log.Println("taobaoTravelTicketOrderRefund error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
