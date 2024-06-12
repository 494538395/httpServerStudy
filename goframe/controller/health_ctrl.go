package controller

import "github.com/gogf/gf/v2/net/ghttp"

type healthCtrl struct{}

var HealthCtrl healthCtrl

type CommonRsp struct {
	Code int32
	Msg  string
	Data interface{}
}

type GetGroupInfoExResp struct {
	GroupInfos []GroupInfoEx `json:"groupIds"`
}

type GroupInfoEx struct {
	GroupId int32  `json:"groupId"`
	Ex      string `json:"ex"`
}

type GetGroupVerificationResultResp struct {
	Results []VerificationResult `json:"results"`
}

type VerificationResult struct {
	UserId string `json:"userId"`
	Pass   bool   `json:"pass"`
}

// Check 服务健康检查
func (ctrl *healthCtrl) Check(r *ghttp.Request) {
	r.Response.Writeln("OK")
}

// Check 服务健康检查
func (ctrl *healthCtrl) GetGroupExt(r *ghttp.Request) {
	commonResp := &CommonRsp{}

	commonResp.Code = 1
	commonResp.Msg = "success"
	commonResp.Data = &GetGroupInfoExResp{
		GroupInfos: []GroupInfoEx{
			{
				GroupId: 12731,
				Ex:      "{name:jerry}",
			},
			{
				GroupId: 12725,
				Ex:      "{name:jack}",
			},
		},
	}

	r.Response.WriteJson(commonResp)
}

// VerificationCheck 入群检测
func (ctrl *healthCtrl) VerificationCheck(r *ghttp.Request) {
	commonResp := &CommonRsp{}

	commonResp.Code = 1
	commonResp.Msg = "success"
	commonResp.Data = &GetGroupVerificationResultResp{
		Results: []VerificationResult{
			{
				UserId: "2011686",
				Pass:   false,
			},
			{
				UserId: "2011687",
				Pass:   false,
			},
		},
	}

	r.Response.WriteJson(commonResp)
}
