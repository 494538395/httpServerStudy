package controller

import "github.com/gogf/gf/v2/net/ghttp"

type healthCtrl struct{}

var HealthCtrl healthCtrl

// Check 服务健康检查
func (ctrl *healthCtrl) Check(r *ghttp.Request) {
	r.Response.Writeln("OK")
}
