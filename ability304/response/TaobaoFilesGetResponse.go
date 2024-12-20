package response

import (
	"github.com/benny502/topsdk/ability304/domain"
)

type TaobaoFilesGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   results
	*/
	Results []domain.TaobaoFilesGetTopDownloadRecordDo `json:"results,omitempty" `
}
