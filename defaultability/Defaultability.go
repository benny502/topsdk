package defaultability

import (
	"errors"
	"log"

	"github.com/benny502/topsdk"
	"github.com/benny502/topsdk/defaultability/request"
	"github.com/benny502/topsdk/defaultability/response"
	"github.com/benny502/topsdk/util"
)

type Defaultability struct {
	Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability {
	return &Defaultability{client}
}

/*
获取商家所在分组及其已授权(广播)消息topics
*/
func (ability *Defaultability) TaobaoTmcUserGet(req *request.TaobaoTmcUserGetRequest) (*response.TaobaoTmcUserGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.user.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcUserGetResponse{}
	if err != nil {
		log.Println("taobaoTmcUserGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
关键词过滤匹配
*/
func (ability *Defaultability) TaobaoKfcKeywordSearch(req *request.TaobaoKfcKeywordSearchRequest, session string) (*response.TaobaoKfcKeywordSearchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.kfc.keyword.search", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoKfcKeywordSearchResponse{}
	if err != nil {
		log.Println("taobaoKfcKeywordSearch error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
发布单条消息
*/
func (ability *Defaultability) TaobaoTmcMessageProduce(req *request.TaobaoTmcMessageProduceRequest) (*response.TaobaoTmcMessageProduceResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.message.produce", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcMessageProduceResponse{}
	if err != nil {
		log.Println("taobaoTmcMessageProduce error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
取消用户的消息服务
*/
func (ability *Defaultability) TaobaoTmcUserCancel(req *request.TaobaoTmcUserCancelRequest) (*response.TaobaoTmcUserCancelResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.user.cancel", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcUserCancelResponse{}
	if err != nil {
		log.Println("taobaoTmcUserCancel error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
为已授权的用户开通消息服务
*/
func (ability *Defaultability) TaobaoTmcUserPermit(req *request.TaobaoTmcUserPermitRequest, session string) (*response.TaobaoTmcUserPermitResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tmc.user.permit", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTmcUserPermitResponse{}
	if err != nil {
		log.Println("taobaoTmcUserPermit error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
