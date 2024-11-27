package request


type TaobaoXhotelRoomtypeConflictDataRequest struct {
    /*
        查询第几页数据 defalutValue��1    */
    CurrentPage  *int64 `json:"current_page,omitempty" required:"false" `
}

func (s *TaobaoXhotelRoomtypeConflictDataRequest) SetCurrentPage(v int64) *TaobaoXhotelRoomtypeConflictDataRequest {
    s.CurrentPage = &v
    return s
}

func (req *TaobaoXhotelRoomtypeConflictDataRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CurrentPage != nil) {
        paramMap["current_page"] = *req.CurrentPage
    }
    return paramMap
}

func (req *TaobaoXhotelRoomtypeConflictDataRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}