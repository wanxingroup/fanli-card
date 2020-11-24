package order

import (
	"context"

	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	cardRPC "dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func (_ Controller) GetOrder(ctx context.Context, req *protos.GetOrderRequest) (*protos.GetOrderReply, error) {
	logger := rpcLog.WithRequestId(ctx, log.GetLogger())
	logger.Info(req)
	err := validateGetOderRequest(req)
	if err != nil {
		logger.WithError(err).Error("validateGetOderRequest error")
		return &protos.GetOrderReply{
			Err: cardRPC.ConvertErrorToProtobuf(err),
		}, nil
	}

	orderData, orderDataErr := GetOrder(req.OrderId)
	if orderDataErr != nil {
		logger.WithError(err).Error("get order error")
		return &protos.GetOrderReply{
			Err: cardRPC.ConvertErrorToProtobuf(orderDataErr),
		}, nil
	}

	return &protos.GetOrderReply{
		Order: cardRPC.ConvertModelOrderToProtobuf(orderData),
	}, nil
}

func validateGetOderRequest(req *protos.GetOrderRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.OrderId, cardRPC.OrderIdRule...),
	)
}
