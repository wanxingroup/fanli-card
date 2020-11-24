package order

import (
	"context"

	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"

	cardRPC "dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func (_ Controller) GetOrderList(ctx context.Context, req *protos.GetOrderListRequest) (*protos.GetOrderListReply, error) {
	logger := rpcLog.WithRequestId(ctx, log.GetLogger())
	logger.Info(req)

	orderData, count, err := GetOrderList(req)
	if err != nil {
		logger.WithError(err).Error("get order list error")
		return &protos.GetOrderListReply{
			Err: cardRPC.ConvertErrorToProtobuf(err),
		}, nil
	}

	return &protos.GetOrderListReply{
		OrderList: convertOrderList(orderData),
		Count:     count,
	}, nil
}
