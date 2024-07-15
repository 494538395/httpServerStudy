package cmd

import (
	"context"

	"goframe-study/controller"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetPort(9996)
			s.SetAccessLogEnabled(true)

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.ALLMap(
					g.Map{
						"/health": controller.HealthCtrl,
					},
				)
			})

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.ALLMap(
					g.Map{
						"/group/get/ex":          controller.HealthCtrl.GetGroupExt,
						"/group/verification":    controller.HealthCtrl.VerificationCheck,
						"/file/get/level/config": controller.FileCtrl.GetLevelConfig,
						"/file/get/image/*":      controller.FileCtrl.GetImage,
					},
				)
			})
			s.Run()
			return nil
		},
	}
)
