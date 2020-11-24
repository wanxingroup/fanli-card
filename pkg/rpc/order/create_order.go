package order

import (
	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	idcreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	cardRPC "dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func (_ Controller) CreateOrder(ctx context.Context, req *protos.CreateOrderRequest) (*protos.CreateOrderReply, error) {

	logger := rpcLog.WithRequestId(ctx, log.GetLogger())
	err := validateCreateOderRequest(req)
	if err != nil {
		return &protos.CreateOrderReply{
			Err: cardRPC.ConvertErrorToProtobuf(err),
		}, nil
	}

	var orderId uint64
	orderId, err = createOrder(req)
	if err != nil {
		logger.WithError(err).Error("create order error")
		return &protos.CreateOrderReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeCreateOrderFailed,
				Message: constant.ErrorMessageCreateOrderFailed,
				Stack:   nil,
			},
		}, nil
	}

	return &protos.CreateOrderReply{
		OrderId: orderId,
	}, nil

}

func createOrder(req *protos.CreateOrderRequest) (uint64, error) {

	record := &card.Order{
		OrderId: idcreator.NextID(),
		UserId:  req.UserId,
		ShopId:  req.ShopId,
		CardId:  req.CardId,
		ItemId:  req.ItemId,
		Amount:  req.Amount,
		Status:  card.OrderStatusNotPay,
		Summary: req.Summary,
	}

	err := database.GetDB(constant.DatabaseConfigKey).Create(record).Error
	if err != nil {
		log.GetLogger().WithField("order", record).WithError(err).Error("create record error")
		return 0, err
	}

	return record.OrderId, nil
}

func validateCreateOderRequest(req *protos.CreateOrderRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.ShopId, cardRPC.ShopIdRule...),
		validation.Field(&req.UserId, cardRPC.UserIdRule...),
	)
}
