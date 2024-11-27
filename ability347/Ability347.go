package ability347

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability347/request"
	"topsdk/ability347/response"
	"topsdk/util"
)

type Ability347 struct {
	Client *topsdk.TopClient
}

func NewAbility347(client *topsdk.TopClient) *Ability347 {
	return &Ability347{client}
}

/*
自促活动申请接口
*/
func (ability *Ability347) TaobaoXhotelBnbpromoAdd(req *request.TaobaoXhotelBnbpromoAddRequest, session string) (*response.TaobaoXhotelBnbpromoAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbpromo.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbpromoAddResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbpromoAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
飞猪酒店多维度服务时间维护接口
*/
func (ability *Ability347) TaobaoXhotelServicetimeUpdate(req *request.TaobaoXhotelServicetimeUpdateRequest, session string) (*response.TaobaoXhotelServicetimeUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.servicetime.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelServicetimeUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelServicetimeUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询实体对应的服务时间数据
*/
func (ability *Ability347) TaobaoXhotelServicetimeGet(req *request.TaobaoXhotelServicetimeGetRequest, session string) (*response.TaobaoXhotelServicetimeGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.servicetime.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelServicetimeGetResponse{}
	if err != nil {
		log.Println("taobaoXhotelServicetimeGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
复杂价格推送接口（批量全量）
*/
func (ability *Ability347) TaobaoXhotelMultipleratesUpdate(req *request.TaobaoXhotelMultipleratesUpdateRequest, session string) (*response.TaobaoXhotelMultipleratesUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.multiplerates.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelMultipleratesUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelMultipleratesUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
复杂价格推送接口（全量更新）
*/
func (ability *Ability347) TaobaoXhotelMultiplerateUpdate(req *request.TaobaoXhotelMultiplerateUpdateRequest, session string) (*response.TaobaoXhotelMultiplerateUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.multiplerate.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelMultiplerateUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelMultiplerateUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
复杂房价推送接口（批量增量）
*/
func (ability *Ability347) TaobaoXhotelMultipleratesIncrement(req *request.TaobaoXhotelMultipleratesIncrementRequest, session string) (*response.TaobaoXhotelMultipleratesIncrementResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.multiplerates.increment", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelMultipleratesIncrementResponse{}
	if err != nil {
		log.Println("taobaoXhotelMultipleratesIncrement error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
复杂价格删除
*/
func (ability *Ability347) TaobaoXhotelMultiplerateDelete(req *request.TaobaoXhotelMultiplerateDeleteRequest, session string) (*response.TaobaoXhotelMultiplerateDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.multiplerate.delete", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelMultiplerateDeleteResponse{}
	if err != nil {
		log.Println("taobaoXhotelMultiplerateDelete error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
上传信息计算飞猪国际城市
*/
func (ability *Ability347) TaobaoXhotelCityCoordinatesBatchUpload(req *request.TaobaoXhotelCityCoordinatesBatchUploadRequest, session string) (*response.TaobaoXhotelCityCoordinatesBatchUploadResponse, error) {
	log.Println("in")
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.city.coordinates.batch.upload", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelCityCoordinatesBatchUploadResponse{}
	if err != nil {
		log.Println("taobaoXhotelCityCoordinatesBatchUpload error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
民宿营销活动更新
*/
func (ability *Ability347) TaobaoXhotelBnbpromoUpdate(req *request.TaobaoXhotelBnbpromoUpdateRequest, session string) (*response.TaobaoXhotelBnbpromoUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbpromo.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbpromoUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbpromoUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
不落库商家推送更新酒店rate
*/
func (ability *Ability347) TaobaoXhotelIntlRateUpdate(req *request.TaobaoXhotelIntlRateUpdateRequest, session string) (*response.TaobaoXhotelIntlRateUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.intl.rate.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelIntlRateUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelIntlRateUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
自促活动绑定接口
*/
func (ability *Ability347) TaobaoXhotelBnbpromoBind(req *request.TaobaoXhotelBnbpromoBindRequest, session string) (*response.TaobaoXhotelBnbpromoBindResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbpromo.bind", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbpromoBindResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbpromoBind error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
自促活动解绑接口
*/
func (ability *Ability347) TaobaoXhotelBnbpromoUnbind(req *request.TaobaoXhotelBnbpromoUnbindRequest, session string) (*response.TaobaoXhotelBnbpromoUnbindResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbpromo.unbind", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbpromoUnbindResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbpromoUnbind error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
酒店房型与房价查询
*/
func (ability *Ability347) TaobaoXhotelBaseinfoRoomGet(req *request.TaobaoXhotelBaseinfoRoomGetRequest, session string) (*response.TaobaoXhotelBaseinfoRoomGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.baseinfo.room.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBaseinfoRoomGetResponse{}
	if err != nil {
		log.Println("taobaoXhotelBaseinfoRoomGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
商家数据-选品整体概况
*/
func (ability *Ability347) TaobaoXhotelItemSelectionSellerStatSummary(req *request.TaobaoXhotelItemSelectionSellerStatSummaryRequest, session string) (*response.TaobaoXhotelItemSelectionSellerStatSummaryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.item.selection.seller.stat.summary", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelItemSelectionSellerStatSummaryResponse{}
	if err != nil {
		log.Println("taobaoXhotelItemSelectionSellerStatSummary error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
选品曝光趋势
*/
func (ability *Ability347) TaobaoXhotelItemSelectionSellerStatExposure(req *request.TaobaoXhotelItemSelectionSellerStatExposureRequest, session string) (*response.TaobaoXhotelItemSelectionSellerStatExposureResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.item.selection.seller.stat.exposure", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelItemSelectionSellerStatExposureResponse{}
	if err != nil {
		log.Println("taobaoXhotelItemSelectionSellerStatExposure error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
供应链选品热销标准酒店覆盖情况
*/
func (ability *Ability347) TaobaoXhotelItemSelectionSellerStatHotshid(req *request.TaobaoXhotelItemSelectionSellerStatHotshidRequest, session string) (*response.TaobaoXhotelItemSelectionSellerStatHotshidResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.item.selection.seller.stat.hotshid", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelItemSelectionSellerStatHotshidResponse{}
	if err != nil {
		log.Println("taobaoXhotelItemSelectionSellerStatHotshid error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
酒店价格库存轻量级增量接口
*/
func (ability *Ability347) TaobaoXhotelRatesLiteIncrUpdate(req *request.TaobaoXhotelRatesLiteIncrUpdateRequest, session string) (*response.TaobaoXhotelRatesLiteIncrUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rates.lite.incr.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRatesLiteIncrUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelRatesLiteIncrUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
民宿卖家活动删除
*/
func (ability *Ability347) TaobaoXhotelBnbpromoDelete(req *request.TaobaoXhotelBnbpromoDeleteRequest, session string) (*response.TaobaoXhotelBnbpromoDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbpromo.delete", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbpromoDeleteResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbpromoDelete error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
民宿查询营销活动
*/
func (ability *Ability347) TaobaoXhotelBnbpromoGet(req *request.TaobaoXhotelBnbpromoGetRequest, session string) (*response.TaobaoXhotelBnbpromoGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbpromo.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbpromoGetResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbpromoGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
直连营销活动解绑资源
*/
func (ability *Ability347) TaobaoXhotelPromotionResourceUnbind(req *request.TaobaoXhotelPromotionResourceUnbindRequest, session string) (*response.TaobaoXhotelPromotionResourceUnbindResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.promotion.resource.unbind", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelPromotionResourceUnbindResponse{}
	if err != nil {
		log.Println("taobaoXhotelPromotionResourceUnbind error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
新增/更新营销活动
*/
func (ability *Ability347) TaobaoXhotelPromotionCreate(req *request.TaobaoXhotelPromotionCreateRequest, session string) (*response.TaobaoXhotelPromotionCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.promotion.create", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelPromotionCreateResponse{}
	if err != nil {
		log.Println("taobaoXhotelPromotionCreate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
直连营销活动绑定资源
*/
func (ability *Ability347) TaobaoXhotelPromotionResourceBind(req *request.TaobaoXhotelPromotionResourceBindRequest, session string) (*response.TaobaoXhotelPromotionResourceBindResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.promotion.resource.bind", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelPromotionResourceBindResponse{}
	if err != nil {
		log.Println("taobaoXhotelPromotionResourceBind error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
直连营销活动和绑定资源查询
*/
func (ability *Ability347) TaobaoXhotelPromotionQuery(req *request.TaobaoXhotelPromotionQueryRequest, session string) (*response.TaobaoXhotelPromotionQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.promotion.query", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelPromotionQueryResponse{}
	if err != nil {
		log.Println("taobaoXhotelPromotionQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
rate删除接口
*/
func (ability *Ability347) TaobaoXhotelRateDelete(req *request.TaobaoXhotelRateDeleteRequest, session string) (*response.TaobaoXhotelRateDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rate.delete", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRateDeleteResponse{}
	if err != nil {
		log.Println("taobaoXhotelRateDelete error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
民宿房东删除接口
*/
func (ability *Ability347) TaobaoXhotelBnbownerDelete(req *request.TaobaoXhotelBnbownerDeleteRequest, session string) (*response.TaobaoXhotelBnbownerDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbowner.delete", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbownerDeleteResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbownerDelete error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
对外开放评论接口
*/
func (ability *Ability347) TaobaoXhotelBnbreviewReply(req *request.TaobaoXhotelBnbreviewReplyRequest, session string) (*response.TaobaoXhotelBnbreviewReplyResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbreview.reply", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbreviewReplyResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbreviewReply error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
酒店基础信息查询接口
*/
func (ability *Ability347) TaobaoXhotelBaseinfoGet(req *request.TaobaoXhotelBaseinfoGetRequest, session string) (*response.TaobaoXhotelBaseinfoGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.baseinfo.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBaseinfoGetResponse{}
	if err != nil {
		log.Println("taobaoXhotelBaseinfoGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
房型查询接口
*/
func (ability *Ability347) TaobaoXhotelRoomtypeGet(req *request.TaobaoXhotelRoomtypeGetRequest, session string) (*response.TaobaoXhotelRoomtypeGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.roomtype.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRoomtypeGetResponse{}
	if err != nil {
		log.Println("taobaoXhotelRoomtypeGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
房型更新接口（ID不存在自动新增）
*/
func (ability *Ability347) TaobaoXhotelRoomtypeUpdate(req *request.TaobaoXhotelRoomtypeUpdateRequest, session string) (*response.TaobaoXhotelRoomtypeUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.roomtype.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRoomtypeUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelRoomtypeUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
民宿新增房源
*/
func (ability *Ability347) TaobaoXhotelBnbroomtypeAdd(req *request.TaobaoXhotelBnbroomtypeAddRequest, session string) (*response.TaobaoXhotelBnbroomtypeAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbroomtype.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbroomtypeAddResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbroomtypeAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
房型新增接口（ID重复变更新）
*/
func (ability *Ability347) TaobaoXhotelRoomtypeAdd(req *request.TaobaoXhotelRoomtypeAddRequest, session string) (*response.TaobaoXhotelRoomtypeAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.roomtype.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRoomtypeAddResponse{}
	if err != nil {
		log.Println("taobaoXhotelRoomtypeAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
民宿门店信息添加
*/
func (ability *Ability347) TaobaoXhotelBnbhouseAdd(req *request.TaobaoXhotelBnbhouseAddRequest, session string) (*response.TaobaoXhotelBnbhouseAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbhouse.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbhouseAddResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbhouseAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
酒店查询接口
*/
func (ability *Ability347) TaobaoXhotelGet(req *request.TaobaoXhotelGetRequest, session string) (*response.TaobaoXhotelGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelGetResponse{}
	if err != nil {
		log.Println("taobaoXhotelGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
民宿房源删除接口
*/
func (ability *Ability347) TaobaoXhotelBnbroomtypeDelete(req *request.TaobaoXhotelBnbroomtypeDeleteRequest, session string) (*response.TaobaoXhotelBnbroomtypeDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbroomtype.delete", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbroomtypeDeleteResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbroomtypeDelete error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
酒店更新接口（ID不存在自动新增）
*/
func (ability *Ability347) TaobaoXhotelUpdate(req *request.TaobaoXhotelUpdateRequest, session string) (*response.TaobaoXhotelUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
民宿房东信息添加
*/
func (ability *Ability347) TaobaoXhotelBnbownerAdd(req *request.TaobaoXhotelBnbownerAddRequest, session string) (*response.TaobaoXhotelBnbownerAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbowner.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbownerAddResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbownerAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
酒店新增接口（ID重复自动更新）
*/
func (ability *Ability347) TaobaoXhotelAdd(req *request.TaobaoXhotelAddRequest, session string) (*response.TaobaoXhotelAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelAddResponse{}
	if err != nil {
		log.Println("taobaoXhotelAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
room实体查询
*/
func (ability *Ability347) TaobaoXhotelRoomGet(req *request.TaobaoXhotelRoomGetRequest, session string) (*response.TaobaoXhotelRoomGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.room.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRoomGetResponse{}
	if err != nil {
		log.Println("taobaoXhotelRoomGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
价格计划rateplan查询
*/
func (ability *Ability347) TaobaoXhotelRateplanGet(req *request.TaobaoXhotelRateplanGetRequest, session string) (*response.TaobaoXhotelRateplanGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rateplan.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRateplanGetResponse{}
	if err != nil {
		log.Println("taobaoXhotelRateplanGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
房型共享库存推送接口（批量全量）
*/
func (ability *Ability347) TaobaoXhotelRoomsUpdate(req *request.TaobaoXhotelRoomsUpdateRequest, session string) (*response.TaobaoXhotelRoomsUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rooms.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRoomsUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelRoomsUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
房型库存推送接口（全量推送）
*/
func (ability *Ability347) TaobaoXhotelRoomUpdate(req *request.TaobaoXhotelRoomUpdateRequest, session string) (*response.TaobaoXhotelRoomUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.room.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRoomUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelRoomUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
价格推送接口（批量全量）
*/
func (ability *Ability347) TaobaoXhotelRatesUpdate(req *request.TaobaoXhotelRatesUpdateRequest, session string) (*response.TaobaoXhotelRatesUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rates.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRatesUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelRatesUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
价格推送接口（全量更新）
*/
func (ability *Ability347) TaobaoXhotelRateUpdate(req *request.TaobaoXhotelRateUpdateRequest, session string) (*response.TaobaoXhotelRateUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rate.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRateUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelRateUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
酒店产品库rate查询
*/
func (ability *Ability347) TaobaoXhotelRateGet(req *request.TaobaoXhotelRateGetRequest, session string) (*response.TaobaoXhotelRateGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rate.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRateGetResponse{}
	if err != nil {
		log.Println("taobaoXhotelRateGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
新增专享房价
*/
func (ability *Ability347) TaobaoXhotelRateAdd(req *request.TaobaoXhotelRateAddRequest, session string) (*response.TaobaoXhotelRateAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rate.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRateAddResponse{}
	if err != nil {
		log.Println("taobaoXhotelRateAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
价格计划rateplan更新或添加
*/
func (ability *Ability347) TaobaoXhotelRateplanUpdate(req *request.TaobaoXhotelRateplanUpdateRequest, session string) (*response.TaobaoXhotelRateplanUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rateplan.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRateplanUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelRateplanUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
删除 x 元素
*/
func (ability *Ability347) TaobaoXhotelXitemDelete(req *request.TaobaoXhotelXitemDeleteRequest, session string) (*response.TaobaoXhotelXitemDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.xitem.delete", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelXitemDeleteResponse{}
	if err != nil {
		log.Println("taobaoXhotelXitemDelete error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
酒店产品库rateplan添加
*/
func (ability *Ability347) TaobaoXhotelRateplanAdd(req *request.TaobaoXhotelRateplanAddRequest, session string) (*response.TaobaoXhotelRateplanAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rateplan.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRateplanAddResponse{}
	if err != nil {
		log.Println("taobaoXhotelRateplanAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询 x 元素
*/
func (ability *Ability347) TaobaoXhotelXitemQuery(req *request.TaobaoXhotelXitemQueryRequest, session string) (*response.TaobaoXhotelXitemQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.xitem.query", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelXitemQueryResponse{}
	if err != nil {
		log.Println("taobaoXhotelXitemQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
飞猪商品各实体通用配置
*/
func (ability *Ability347) TaobaoXhotelEntityConfig(req *request.TaobaoXhotelEntityConfigRequest, session string) (*response.TaobaoXhotelEntityConfigResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.entity.config", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelEntityConfigResponse{}
	if err != nil {
		log.Println("taobaoXhotelEntityConfig error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据标签查询实体
*/
func (ability *Ability347) TaobaoXhotelGetEntityByTag(req *request.TaobaoXhotelGetEntityByTagRequest, session string) (*response.TaobaoXhotelGetEntityByTagResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.get.entity.by.tag", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelGetEntityByTagResponse{}
	if err != nil {
		log.Println("taobaoXhotelGetEntityByTag error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
库存更新接口
*/
func (ability *Ability347) TaobaoXhotelQuotaUpdate(req *request.TaobaoXhotelQuotaUpdateRequest, session string) (*response.TaobaoXhotelQuotaUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.quota.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelQuotaUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelQuotaUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
对外开放评论接口
*/
func (ability *Ability347) TaobaoXhotelBnbreviewAdd(req *request.TaobaoXhotelBnbreviewAddRequest, session string) (*response.TaobaoXhotelBnbreviewAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbreview.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbreviewAddResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbreviewAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
通用评价接口
*/
func (ability *Ability347) TaobaoXhotelCommentGetmixratelist(req *request.TaobaoXhotelCommentGetmixratelistRequest, session string) (*response.TaobaoXhotelCommentGetmixratelistResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.comment.getmixratelist", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelCommentGetmixratelistResponse{}
	if err != nil {
		log.Println("taobaoXhotelCommentGetmixratelist error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
删除酒店接口
*/
func (ability *Ability347) TaobaoXhotelDelete(req *request.TaobaoXhotelDeleteRequest, session string) (*response.TaobaoXhotelDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.delete", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelDeleteResponse{}
	if err != nil {
		log.Println("taobaoXhotelDelete error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
商家删除房型数据接口
*/
func (ability *Ability347) TaobaoXhotelRoomtypeDeletePublic(req *request.TaobaoXhotelRoomtypeDeletePublicRequest, session string) (*response.TaobaoXhotelRoomtypeDeletePublicResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.roomtype.delete.public", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRoomtypeDeletePublicResponse{}
	if err != nil {
		log.Println("taobaoXhotelRoomtypeDeletePublic error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
民宿门店删除接口
*/
func (ability *Ability347) TaobaoXhotelBnbhouseDelete(req *request.TaobaoXhotelBnbhouseDeleteRequest, session string) (*response.TaobaoXhotelBnbhouseDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbhouse.delete", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbhouseDeleteResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbhouseDelete error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
商家床型冲突数据接口
*/
func (ability *Ability347) TaobaoXhotelRoomtypeConflictData(req *request.TaobaoXhotelRoomtypeConflictDataRequest, session string) (*response.TaobaoXhotelRoomtypeConflictDataResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.roomtype.conflict.data", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRoomtypeConflictDataResponse{}
	if err != nil {
		log.Println("taobaoXhotelRoomtypeConflictData error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
通用调用top接口
*/
func (ability *Ability347) TaobaoXhotelBnbcommonAdd(req *request.TaobaoXhotelBnbcommonAddRequest, session string) (*response.TaobaoXhotelBnbcommonAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.bnbcommon.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelBnbcommonAddResponse{}
	if err != nil {
		log.Println("taobaoXhotelBnbcommonAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
价格推送接口（批量增量）
*/
func (ability *Ability347) TaobaoXhotelRatesIncrement(req *request.TaobaoXhotelRatesIncrementRequest, session string) (*response.TaobaoXhotelRatesIncrementResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rates.increment", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRatesIncrementResponse{}
	if err != nil {
		log.Println("taobaoXhotelRatesIncrement error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
房型库存推送接口（批量增量）
*/
func (ability *Ability347) TaobaoXhotelRoomsIncrement(req *request.TaobaoXhotelRoomsIncrementRequest, session string) (*response.TaobaoXhotelRoomsIncrementResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rooms.increment", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRoomsIncrementResponse{}
	if err != nil {
		log.Println("taobaoXhotelRoomsIncrement error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
酒店状态更新
*/
func (ability *Ability347) TaobaoXhotelStatusUpdate(req *request.TaobaoXhotelStatusUpdateRequest, session string) (*response.TaobaoXhotelStatusUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.status.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelStatusUpdateResponse{}
	if err != nil {
		log.Println("taobaoXhotelStatusUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
top房型状态修改
*/
func (ability *Ability347) TaobaoRoomtypeStatusUpdate(req *request.TaobaoRoomtypeStatusUpdateRequest, session string) (*response.TaobaoRoomtypeStatusUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.roomtype.status.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoRoomtypeStatusUpdateResponse{}
	if err != nil {
		log.Println("taobaoRoomtypeStatusUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
非标准民宿房源添加
*/
func (ability *Ability347) TaobaoXhotelHouseAdd(req *request.TaobaoXhotelHouseAddRequest, session string) (*response.TaobaoXhotelHouseAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.house.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelHouseAddResponse{}
	if err != nil {
		log.Println("taobaoXhotelHouseAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
民宿房型信息添加
*/
func (ability *Ability347) TaobaoXhotelHouseRoomtypeAdd(req *request.TaobaoXhotelHouseRoomtypeAddRequest, session string) (*response.TaobaoXhotelHouseRoomtypeAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.house.roomtype.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelHouseRoomtypeAddResponse{}
	if err != nil {
		log.Println("taobaoXhotelHouseRoomtypeAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据gid查询卖家下所有的rpId
*/
func (ability *Ability347) TaobaoXhotelRateRelationshipwithrpGet(req *request.TaobaoXhotelRateRelationshipwithrpGetRequest, session string) (*response.TaobaoXhotelRateRelationshipwithrpGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rate.relationshipwithrp.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRateRelationshipwithrpGetResponse{}
	if err != nil {
		log.Println("taobaoXhotelRateRelationshipwithrpGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
下载飞猪国际城市结果
*/
func (ability *Ability347) TaobaoXhotelCityCoordinatesBatchDownload(req *request.TaobaoXhotelCityCoordinatesBatchDownloadRequest, session string) (*response.TaobaoXhotelCityCoordinatesBatchDownloadResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.city.coordinates.batch.download", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelCityCoordinatesBatchDownloadResponse{}
	if err != nil {
		log.Println("taobaoXhotelCityCoordinatesBatchDownload error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询rpId
*/
func (ability *Ability347) TaobaoXhotelRateRelationshipwithroomGet(req *request.TaobaoXhotelRateRelationshipwithroomGetRequest, session string) (*response.TaobaoXhotelRateRelationshipwithroomGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability347 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.xhotel.rate.relationshipwithroom.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoXhotelRateRelationshipwithroomGetResponse{}
	if err != nil {
		log.Println("taobaoXhotelRateRelationshipwithroomGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
