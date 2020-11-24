package client

import (
	"context"

	couponProtos "dev-gitlab.wanxingrowth.com/fanli/coupon/pkg/rpc/protos"
	"google.golang.org/grpc"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/config"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

var couponRPCService couponProtos.ShopCouponControllerClient
var userCouponRPCService couponProtos.UserCouponControllerClient

func InitCouponService() {

	log.GetLogger().Info("starting init coupon rpc service")

	var ctx = context.Background()

	var rpcConfig, exist = config.Config.RPCServices[constant.RPCCouponServiceConfigKey]
	if !exist {
		log.GetLogger().Error("coupon rpc service configuration not exist")
		return
	}

	if rpcConfig.GetConnectTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.TODO(), rpcConfig.GetConnectTimeout())
		defer cancel()
	}

	conn, err := grpc.DialContext(ctx, rpcConfig.GetAddress(), grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.GetLogger().WithError(err).Error("coupon rpc service connect failed")
		return
	}

	couponRPCService = couponProtos.NewShopCouponControllerClient(conn)
	userCouponRPCService = couponProtos.NewUserCouponControllerClient(conn)

	log.GetLogger().Info("coupon rpc service init succeed")
}

func GetCouponService() couponProtos.ShopCouponControllerClient {
	if couponRPCService == nil {
		InitCouponService()
	}

	return couponRPCService
}

func GetUserCouponService() couponProtos.UserCouponControllerClient {
	if userCouponRPCService == nil {
		InitCouponService()
	}

	return userCouponRPCService
}
