package request

import (
        "topsdk/ability347/domain"
        "topsdk/util"
    )

type TaobaoXhotelBnbreviewAddRequest struct {
    /*
        入住时间     */
    CheckInTime  *string `json:"check_in_time" required:"true" `
    /*
        评论来源ID，38-小猪 39-爱彼迎     */
    Source  *int64 `json:"source" required:"true" `
    /*
        图片地址     */
    PicInfoList  *[]domain.TaobaoXhotelBnbreviewAddReviewPicInfo `json:"pic_info_list,omitempty" required:"false" `
    /*
        飞猪侧房源ID     */
    Rid  *int64 `json:"rid" required:"true" `
    /*
        创建时间     */
    GmtCreate  *util.LocalTime `json:"gmt_create" required:"true" `
    /*
        评论内容     */
    Content  *string `json:"content" required:"true" `
    /*
        用户名称     */
    UserNick  *string `json:"user_nick" required:"true" `
    /*
        外部评论id     */
    OuterId  *int64 `json:"outer_id" required:"true" `
    /*
        评分细分     */
    ScoreDetail  *[]domain.TaobaoXhotelBnbreviewAddReviewDetailInfo `json:"score_detail" required:"true" `
    /*
        总评分,Double类型得     */
    TotalScore  *string `json:"total_score" required:"true" `
}

func (s *TaobaoXhotelBnbreviewAddRequest) SetCheckInTime(v string) *TaobaoXhotelBnbreviewAddRequest {
    s.CheckInTime = &v
    return s
}
func (s *TaobaoXhotelBnbreviewAddRequest) SetSource(v int64) *TaobaoXhotelBnbreviewAddRequest {
    s.Source = &v
    return s
}
func (s *TaobaoXhotelBnbreviewAddRequest) SetPicInfoList(v []domain.TaobaoXhotelBnbreviewAddReviewPicInfo) *TaobaoXhotelBnbreviewAddRequest {
    s.PicInfoList = &v
    return s
}
func (s *TaobaoXhotelBnbreviewAddRequest) SetRid(v int64) *TaobaoXhotelBnbreviewAddRequest {
    s.Rid = &v
    return s
}
func (s *TaobaoXhotelBnbreviewAddRequest) SetGmtCreate(v util.LocalTime) *TaobaoXhotelBnbreviewAddRequest {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoXhotelBnbreviewAddRequest) SetContent(v string) *TaobaoXhotelBnbreviewAddRequest {
    s.Content = &v
    return s
}
func (s *TaobaoXhotelBnbreviewAddRequest) SetUserNick(v string) *TaobaoXhotelBnbreviewAddRequest {
    s.UserNick = &v
    return s
}
func (s *TaobaoXhotelBnbreviewAddRequest) SetOuterId(v int64) *TaobaoXhotelBnbreviewAddRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoXhotelBnbreviewAddRequest) SetScoreDetail(v []domain.TaobaoXhotelBnbreviewAddReviewDetailInfo) *TaobaoXhotelBnbreviewAddRequest {
    s.ScoreDetail = &v
    return s
}
func (s *TaobaoXhotelBnbreviewAddRequest) SetTotalScore(v string) *TaobaoXhotelBnbreviewAddRequest {
    s.TotalScore = &v
    return s
}

func (req *TaobaoXhotelBnbreviewAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CheckInTime != nil) {
        paramMap["check_in_time"] = *req.CheckInTime
    }
    if(req.Source != nil) {
        paramMap["source"] = *req.Source
    }
    if(req.PicInfoList != nil) {
        paramMap["pic_info_list"] = util.ConvertStructList(*req.PicInfoList)
    }
    if(req.Rid != nil) {
        paramMap["rid"] = *req.Rid
    }
    if(req.GmtCreate != nil) {
        paramMap["gmt_create"] = *req.GmtCreate
    }
    if(req.Content != nil) {
        paramMap["content"] = *req.Content
    }
    if(req.UserNick != nil) {
        paramMap["user_nick"] = *req.UserNick
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.ScoreDetail != nil) {
        paramMap["score_detail"] = util.ConvertStructList(*req.ScoreDetail)
    }
    if(req.TotalScore != nil) {
        paramMap["total_score"] = *req.TotalScore
    }
    return paramMap
}

func (req *TaobaoXhotelBnbreviewAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}