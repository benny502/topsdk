package domain

import (
	"github.com/benny502/topsdk/util"
)

type TaobaoXhotelHouseRoomtypeAddXRoomType struct {
	/*
	   rid     */
	Rid *int64 `json:"rid,omitempty" `

	/*
	   hid     */
	Hid *int64 `json:"hid,omitempty" `

	/*
	   创建时间     */
	CreatedTime *util.LocalTime `json:"created_time,omitempty" `

	/*
	   修改时间     */
	ModifiedTime *util.LocalTime `json:"modified_time,omitempty" `

	/*
	   匹配状态: 0：待系统匹配 1：已系统匹配，匹配成功，待卖家确认 2：已系统匹配，匹配失败，待人工匹配 3：已人工匹配，匹配成功，待卖家确认 4：已人工匹配，匹配失败 5：卖家已确认，确认&ldquo;YES&rdquo; 6：卖家已确认，确认&ldquo;NO&rdquo; 7:已系统匹配，但是匹配重复，待人工确认     */
	MatchStatus *int64 `json:"match_status,omitempty" `

	/*
	   房型状态。0:正常，-1:删除，-2:停售     */
	Status *int64 `json:"status,omitempty" `

	/*
	   出错原因,没有匹配上标准房型时，小二拒绝的理由     */
	ErrorInfo *string `json:"error_info,omitempty" `

	/*
	   卖家系统id     */
	OuterId *string `json:"outer_id,omitempty" `

	/*
	   房型名称     */
	Name *string `json:"name,omitempty" `

	/*
	   最大入住人数     */
	MaxOccupancy *int64 `json:"max_occupancy,omitempty" `

	/*
	   可选值：A,B,C,D。分别代表： A：15平米以下，B：16－30平米，C：31－50平米，D：50平米以上 2）也可以自己定义，比如：40平方米     */
	Area *string `json:"area,omitempty" `

	/*
	   客房在建筑的第几层，隔层为1-2层，4-5层，7-8层     */
	Floor *string `json:"floor,omitempty" `

	/*
	   床型。按自己定义存储，比如：高低床、上下床     */
	BedType *string `json:"bed_type,omitempty" `

	/*
	   床宽。     */
	BedSize *string `json:"bed_size,omitempty" `

	/*
	   宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带     */
	Internet *string `json:"internet,omitempty" `

	/*
	   设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {&quot;bar&quot;:false,&quot;catv&quot;:false,&quot;ddd&quot;:false,&quot;idd&quot;:false,&quot;pubtoilet&quot;:false,&quot;toilet&quot;:false}     */
	Service *string `json:"service,omitempty" `

	/*
	   窗型,0：无窗/1：有窗     */
	WindowType *int64 `json:"window_type,omitempty" `

	/*
	   扩展信息的JSON。 注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析     */
	Extend *string `json:"extend,omitempty" `

	/*
	   标准房型信息     */
	SRoomtype *TaobaoXhotelHouseRoomtypeAddSRoomType `json:"s_roomtype,omitempty" `

	/*
	   卖家房型英文名称     */
	NameE *string `json:"name_e,omitempty" `
}

func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetRid(v int64) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.Rid = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetHid(v int64) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.Hid = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetCreatedTime(v util.LocalTime) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.CreatedTime = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetModifiedTime(v util.LocalTime) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.ModifiedTime = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetMatchStatus(v int64) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.MatchStatus = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetStatus(v int64) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.Status = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetErrorInfo(v string) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.ErrorInfo = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetOuterId(v string) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.OuterId = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetName(v string) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.Name = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetMaxOccupancy(v int64) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.MaxOccupancy = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetArea(v string) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.Area = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetFloor(v string) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.Floor = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetBedType(v string) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.BedType = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetBedSize(v string) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.BedSize = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetInternet(v string) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.Internet = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetService(v string) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.Service = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetWindowType(v int64) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.WindowType = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetExtend(v string) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.Extend = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetSRoomtype(v TaobaoXhotelHouseRoomtypeAddSRoomType) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.SRoomtype = &v
	return s
}
func (s *TaobaoXhotelHouseRoomtypeAddXRoomType) SetNameE(v string) *TaobaoXhotelHouseRoomtypeAddXRoomType {
	s.NameE = &v
	return s
}
