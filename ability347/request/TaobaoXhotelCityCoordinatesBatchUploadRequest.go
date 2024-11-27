package request

import (
	"github.com/benny502/topsdk/ability347/domain"
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelCityCoordinatesBatchUploadRequest struct {
	/*
	   经纬度列表     */
	CoordinateList *[]domain.TaobaoXhotelCityCoordinatesBatchUploadCoordinate `json:"coordinate_list" required:"true" `
}

func (s *TaobaoXhotelCityCoordinatesBatchUploadRequest) SetCoordinateList(v []domain.TaobaoXhotelCityCoordinatesBatchUploadCoordinate) *TaobaoXhotelCityCoordinatesBatchUploadRequest {
	s.CoordinateList = &v
	return s
}

func (req *TaobaoXhotelCityCoordinatesBatchUploadRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.CoordinateList != nil {
		paramMap["coordinate_list"] = util.ConvertStructList(*req.CoordinateList)
	}
	return paramMap
}

func (req *TaobaoXhotelCityCoordinatesBatchUploadRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
