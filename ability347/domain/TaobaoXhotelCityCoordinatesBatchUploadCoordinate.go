package domain


type TaobaoXhotelCityCoordinatesBatchUploadCoordinate struct {
    /*
        飞猪国家编码     */
    Country  *int64 `json:"country,omitempty" `

    /*
        纬度     */
    Latitude  *string `json:"latitude,omitempty" `

    /*
        经度     */
    Longitude  *string `json:"longitude,omitempty" `

    /*
        外部经纬度标识id，可以是酒店或城市的id     */
    OuterId  *string `json:"outer_id,omitempty" `

}

func (s *TaobaoXhotelCityCoordinatesBatchUploadCoordinate) SetCountry(v int64) *TaobaoXhotelCityCoordinatesBatchUploadCoordinate {
    s.Country = &v
    return s
}
func (s *TaobaoXhotelCityCoordinatesBatchUploadCoordinate) SetLatitude(v string) *TaobaoXhotelCityCoordinatesBatchUploadCoordinate {
    s.Latitude = &v
    return s
}
func (s *TaobaoXhotelCityCoordinatesBatchUploadCoordinate) SetLongitude(v string) *TaobaoXhotelCityCoordinatesBatchUploadCoordinate {
    s.Longitude = &v
    return s
}
func (s *TaobaoXhotelCityCoordinatesBatchUploadCoordinate) SetOuterId(v string) *TaobaoXhotelCityCoordinatesBatchUploadCoordinate {
    s.OuterId = &v
    return s
}
