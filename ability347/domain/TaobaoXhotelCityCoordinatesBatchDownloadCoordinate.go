package domain


type TaobaoXhotelCityCoordinatesBatchDownloadCoordinate struct {
    /*
        批次号     */
    BatchId  *int64 `json:"batch_id,omitempty" `

    /*
        飞猪城市编码     */
    City  *int64 `json:"city,omitempty" `

    /*
        飞猪城市中文名称     */
    CityCnName  *string `json:"city_cn_name,omitempty" `

    /*
        飞猪城市英文名称     */
    CityEnName  *string `json:"city_en_name,omitempty" `

    /*
        飞猪国家编码     */
    Country  *int64 `json:"country,omitempty" `

    /*
        外部经纬度标识id，可以是酒店或城市的id     */
    OuterId  *string `json:"outer_id,omitempty" `

}

func (s *TaobaoXhotelCityCoordinatesBatchDownloadCoordinate) SetBatchId(v int64) *TaobaoXhotelCityCoordinatesBatchDownloadCoordinate {
    s.BatchId = &v
    return s
}
func (s *TaobaoXhotelCityCoordinatesBatchDownloadCoordinate) SetCity(v int64) *TaobaoXhotelCityCoordinatesBatchDownloadCoordinate {
    s.City = &v
    return s
}
func (s *TaobaoXhotelCityCoordinatesBatchDownloadCoordinate) SetCityCnName(v string) *TaobaoXhotelCityCoordinatesBatchDownloadCoordinate {
    s.CityCnName = &v
    return s
}
func (s *TaobaoXhotelCityCoordinatesBatchDownloadCoordinate) SetCityEnName(v string) *TaobaoXhotelCityCoordinatesBatchDownloadCoordinate {
    s.CityEnName = &v
    return s
}
func (s *TaobaoXhotelCityCoordinatesBatchDownloadCoordinate) SetCountry(v int64) *TaobaoXhotelCityCoordinatesBatchDownloadCoordinate {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelCityCoordinatesBatchDownloadCoordinate) SetOuterId(v string) *TaobaoXhotelCityCoordinatesBatchDownloadCoordinate {
    s.OuterId = &v
    return s
}
