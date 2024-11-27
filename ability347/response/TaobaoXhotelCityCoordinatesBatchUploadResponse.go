package response

import (
)

type TaobaoXhotelCityCoordinatesBatchUploadResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        上传成功后的批次号
    */
    BatchId  int64 `json:"batch_id,omitempty" `
}
