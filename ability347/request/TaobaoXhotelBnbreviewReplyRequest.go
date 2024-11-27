package request

import (
	"github.com/benny502/topsdk/ability347/domain"
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelBnbreviewReplyRequest struct {
	/*
	   外部评价ID     */
	OuterId *int64 `json:"outer_id" required:"true" `
	/*
	   评价内容     */
	Content *string `json:"content" required:"true" `
	/*
	   评价类型，1-主评，2-追评     */
	WrateType *int64 `json:"wrate_type" required:"true" `
	/*
	   评价ID     */
	WrateId *int64 `json:"wrate_id" required:"true" `
	/*
	   主评ID     */
	ItemWrateId *int64 `json:"item_wrate_id" required:"true" `
	/*
	   商品ID     */
	ItemId *int64 `json:"item_id" required:"true" `
	/*
	   媒体信息     */
	MediaInfo *domain.TaobaoXhotelBnbreviewReplyMediaInfo `json:"media_info" required:"true" `
}

func (s *TaobaoXhotelBnbreviewReplyRequest) SetOuterId(v int64) *TaobaoXhotelBnbreviewReplyRequest {
	s.OuterId = &v
	return s
}
func (s *TaobaoXhotelBnbreviewReplyRequest) SetContent(v string) *TaobaoXhotelBnbreviewReplyRequest {
	s.Content = &v
	return s
}
func (s *TaobaoXhotelBnbreviewReplyRequest) SetWrateType(v int64) *TaobaoXhotelBnbreviewReplyRequest {
	s.WrateType = &v
	return s
}
func (s *TaobaoXhotelBnbreviewReplyRequest) SetWrateId(v int64) *TaobaoXhotelBnbreviewReplyRequest {
	s.WrateId = &v
	return s
}
func (s *TaobaoXhotelBnbreviewReplyRequest) SetItemWrateId(v int64) *TaobaoXhotelBnbreviewReplyRequest {
	s.ItemWrateId = &v
	return s
}
func (s *TaobaoXhotelBnbreviewReplyRequest) SetItemId(v int64) *TaobaoXhotelBnbreviewReplyRequest {
	s.ItemId = &v
	return s
}
func (s *TaobaoXhotelBnbreviewReplyRequest) SetMediaInfo(v domain.TaobaoXhotelBnbreviewReplyMediaInfo) *TaobaoXhotelBnbreviewReplyRequest {
	s.MediaInfo = &v
	return s
}

func (req *TaobaoXhotelBnbreviewReplyRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.OuterId != nil {
		paramMap["outer_id"] = *req.OuterId
	}
	if req.Content != nil {
		paramMap["content"] = *req.Content
	}
	if req.WrateType != nil {
		paramMap["wrate_type"] = *req.WrateType
	}
	if req.WrateId != nil {
		paramMap["wrate_id"] = *req.WrateId
	}
	if req.ItemWrateId != nil {
		paramMap["item_wrate_id"] = *req.ItemWrateId
	}
	if req.ItemId != nil {
		paramMap["item_id"] = *req.ItemId
	}
	if req.MediaInfo != nil {
		paramMap["media_info"] = util.ConvertStruct(*req.MediaInfo)
	}
	return paramMap
}

func (req *TaobaoXhotelBnbreviewReplyRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
