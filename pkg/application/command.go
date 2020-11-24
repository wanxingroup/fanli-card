package application

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/launcher"
	idcreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	"github.com/spf13/viper"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/client"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/config"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/order"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func Start() {

	app := launcher.NewApplication(
		launcher.SetApplicationDescription(
			&launcher.ApplicationDescription{
				ShortDescription: "card service",
				LongDescription:  "support card data management function.",
			},
		),
		launcher.SetApplicationLogger(log.GetLogger()),
		launcher.SetApplicationEvents(
			launcher.NewApplicationEvents(
				launcher.SetOnInitEvent(func(app *launcher.Application) {

					unmarshalConfiguration()

					registerCardRPCRouter(app)

					client.InitCouponService()
					client.InitRebateService()
					client.InitUserService()
					client.InitGoodsService()
					client.InitOrderService()

					idcreator.InitCreator(app.GetServiceId())
				}),
				launcher.SetOnStartEvent(func(app *launcher.Application) {

					autoMigration()
				}),
			),
		),
	)

	app.Launch()
}

func registerCardRPCRouter(app *launcher.Application) {

	rpcService := app.GetRPCService()
	if rpcService == nil {

		log.GetLogger().WithField("stage", "onInit").Error("get rpc service is nil")
		return
	}

	protos.RegisterCardControllerServer(rpcService.GetRPCConnection(), &card.Controller{})
	protos.RegisterOrderControllerServer(rpcService.GetRPCConnection(), &order.Controller{})
}

func unmarshalConfiguration() {
	err := viper.Unmarshal(config.Config)
	if err != nil {

		log.GetLogger().WithError(err).Error("unmarshal config error")
	}
}
