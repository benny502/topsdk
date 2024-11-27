package response

import (
	"github.com/benny502/topsdk/ability347/domain"
)

type TaobaoXhotelCityCoordinatesBatchDownloadResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   经纬度计算结果列表
	*/
	CoordinateList []domain.TaobaoXhotelCityCoordinatesBatchDownloadCoordinate `json:"coordinate_list,omitempty" `
}
