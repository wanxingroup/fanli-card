package client

import (
	"context"

	rebateProtos "dev-gitlab.wanxingrowth.com/fanli/rebate/pkg/rpc/protos"
	"google.golang.org/grpc"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/config"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

var rebateRPCService rebateProtos.RebateControllerClient

func InitRebateService() {
	rpcName := constant.RPCRebateServiceConfigKey
	log.GetLogger().Info("starting init rebate rpc service")

	var ctx = context.Background()
	var rpcConfig, exist = config.Config.RPCServices[rpcName]
	if !exist {
		log.GetLogger().Error(rpcName + " rpc service configuration not exist")
		return
	}

	if rpcConfig.GetConnectTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.TODO(), rpcConfig.GetConnectTimeout())
		defer cancel()
	}

	conn, err := grpc.DialContext(ctx, rpcConfig.GetAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.GetLogger().WithError(err).Error(rpcName + " rpc service connect failed")
		return
	}

	rebateRPCService = rebateProtos.NewRebateControllerClient(conn)

	log.GetLogger().Info(rpcName + " rpc service init succeed")
}

func GetRebateService() rebateProtos.RebateControllerClient {

	if rebateRPCService == nil {
		InitRebateService()
	}

	return rebateRPCService
}

func SetRebateMockClient(client rebateProtos.RebateControllerClient) {
	rebateRPCService = client
}
