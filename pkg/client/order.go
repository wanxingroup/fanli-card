package client

import (
	"context"

	orderProtos "dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
	"google.golang.org/grpc"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/config"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

var orderRPCService orderProtos.OrderControllerClient
var merchantFreightRPCService orderProtos.MerchantFreightControllerClient

func InitOrderService() {

	log.GetLogger().Info("starting init order rpc service")

	var ctx = context.Background()
	var rpcConfig, exist = config.Config.RPCServices[constant.RPCOrderServiceConfigKey]
	if !exist {
		log.GetLogger().Error("order rpc service configuration not exist")
		return
	}

	if rpcConfig.GetConnectTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.TODO(), rpcConfig.GetConnectTimeout())
		defer cancel()
	}

	conn, err := grpc.DialContext(ctx, rpcConfig.GetAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.GetLogger().WithError(err).Error("order rpc service connect failed")
		return
	}

	orderRPCService = orderProtos.NewOrderControllerClient(conn)
	merchantFreightRPCService = orderProtos.NewMerchantFreightControllerClient(conn)

	log.GetLogger().Info("order rpc service init succeed")
}

func GetOrderService() orderProtos.OrderControllerClient {

	if orderRPCService == nil {
		InitOrderService()
	}

	return orderRPCService
}

func GetMerchantFreightService() orderProtos.MerchantFreightControllerClient {

	if merchantFreightRPCService == nil {
		InitOrderService()
	}

	return merchantFreightRPCService
}

func SetOrderMockClient(client orderProtos.OrderControllerClient) {
	orderRPCService = client
}

func SetMerchantFreightClient(client orderProtos.MerchantFreightControllerClient) {
	merchantFreightRPCService = client
}
