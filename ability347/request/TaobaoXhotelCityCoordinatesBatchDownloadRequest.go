package request


type TaobaoXhotelCityCoordinatesBatchDownloadRequest struct {
    /*
        上传的经纬度批次号     */
    BatchId  *int64 `json:"batch_id,omitempty" required:"false" `
}

func (s *TaobaoXhotelCityCoordinatesBatchDownloadRequest) SetBatchId(v int64) *TaobaoXhotelCityCoordinatesBatchDownloadRequest {
    s.BatchId = &v
    return s
}

func (req *TaobaoXhotelCityCoordinatesBatchDownloadRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BatchId != nil) {
        paramMap["batch_id"] = *req.BatchId
    }
    return paramMap
}

func (req *TaobaoXhotelCityCoordinatesBatchDownloadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}