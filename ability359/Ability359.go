package ability359

import (
	"errors"
	"log"

	"github.com/benny502/topsdk"
	"github.com/benny502/topsdk/ability359/request"
	"github.com/benny502/topsdk/ability359/response"
	"github.com/benny502/topsdk/util"
)

type Ability359 struct {
	Client *topsdk.TopClient
}

func NewAbility359(client *topsdk.TopClient) *Ability359 {
	return &Ability359{client}
}

/*
酒店产品库订单查询
*/
func (ability *Ability359) TaobaoXhotelOrderSearch(req *request.TaobaoXhotelOrderSearchRequest, session string) (*response.TaobaoXhotelOrderSearchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability359 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.order.search", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelOrderSearchResponse{}
	if err != nil {
		log.Println("taobaoXhotelOrderSearch error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
推送及更新订单确认号
*/
func (ability *Ability359) TaobaoXhotelOrderUpdateConfirmcode(req *request.TaobaoXhotelOrderUpdateConfirmcodeRequest, session string) (*response.TaobaoXhotelOrderUpdateConfirmcodeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability359 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.order.update.confirmcode", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelOrderUpdateConfirmcodeResponse{}
	if err != nil {
		log.Println("taobaoXhotelOrderUpdateConfirmcode error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
酒店订单发货接口
*/
func (ability *Ability359) TaobaoXhotelOrderUpdate(req *request.TaobaoXhotelOrderUpdateRequest, session string) (*response.TaobaoXhotelOrderUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability359 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.order.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelOrderUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelOrderUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
飞猪订单隐私号绑定
*/
func (ability *Ability359) TaobaoXhotelOrderSecretphonenumBind(req *request.TaobaoXhotelOrderSecretphonenumBindRequest, session string) (*response.TaobaoXhotelOrderSecretphonenumBindResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability359 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.order.secretphonenum.bind", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelOrderSecretphonenumBindResponse{}
	if err != nil {
		log.Println("taobaoXhotelOrderSecretphonenumBind error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
