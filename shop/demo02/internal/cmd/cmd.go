package cmd

import (
	"context"
	"goframe-shop-v2/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"goframe-shop-v2/internal/controller"
)

// 对外暴露接口

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				//下面是官方提供的
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				//自定义中间件
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				//绑定路由
				group.Bind(
					//示例
					controller.Rotation, // 轮播图
					controller.Position, // 手工位
					//下面的不需要url
					controller.Admin.Create,
					controller.Admin.Update,
					controller.Admin.Delete,
					controller.Admin.List,
					controller.Login,
				)
				// JWT自定义中间件
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.ALLMap(g.Map{
						"/backend/admin/info": controller.Admin.Info,
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
