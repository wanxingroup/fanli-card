package order

import (
	"context"
	"strconv"
	"time"

	orderProtos "dev-gitlab.wanxingrowth.com/fanli/order/v2/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/rebate/pkg/model/rebate"
	rebateProtos "dev-gitlab.wanxingrowth.com/fanli/rebate/pkg/rpc/protos"
	userProtos "dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/client"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	cardRPC "dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"

	couponProtos "dev-gitlab.wanxingrowth.com/fanli/coupon/pkg/rpc/protos"
)

func (_ Controller) FinishOrder(ctx context.Context, req *protos.FinishOrderRequest) (*protos.FinishOrderReply, error) {

	logger := rpcLog.WithRequestId(ctx, log.GetLogger())
	logger.Info(req)
	err := validateFinishOderRequest(req)
	if err != nil {
		logger.WithError(err).Error("validateFinishOderRequest error")
		return &protos.FinishOrderReply{
			Err: cardRPC.ConvertErrorToProtobuf(err),
		}, nil
	}

	orderData, orderDataErr := GetOrder(req.OrderId)
	if orderDataErr != nil {
		logger.WithError(err).Error("get order error")
		return &protos.FinishOrderReply{
			Err: cardRPC.ConvertErrorToProtobuf(orderDataErr),
		}, nil
	}

	if orderData.Status != card.OrderStatusNotPay {
		logger.WithError(err).Error("order status error")
		return &protos.FinishOrderReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeOrderStatusError,
				Message: constant.ErrorMessageOrderStatusError,
				Stack:   nil,
			},
		}, nil
	}

	getItemReply, getItemErr := cardRPC.GetItem(orderData.ItemId)
	if getItemErr != nil {
		logger.WithError(getItemErr).Error("getItemErr is not nil")
		return &protos.FinishOrderReply{
			Err: cardRPC.ConvertErrorToProtobuf(getItemErr),
		}, nil
	}

	if getItemReply.CardId <= 0 {

		logger.Error("getItemErr  getItemReply is nil")

		return &protos.FinishOrderReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeCardNotExist,
				Message: constant.ErrorMessageCardNotExist,
				Stack:   nil,
			},
		}, nil
	}

	_, finishOrderErr := finishOrder(req.OrderId)
	if finishOrderErr != nil {
		logger.WithError(err).Error("finish order error")
		return &protos.FinishOrderReply{
			Err: cardRPC.ConvertErrorToProtobuf(err),
		}, nil
	}

	// 设置邀请人
	userApi := client.GetUserService()
	if userApi == nil {
		logger.Error("Get user Service error")
	} else {

		// 直接怼数据进去
		_, setInviterReplyErr := userApi.SetInviter(ctx, &userProtos.SetInviterRequest{UserId: orderData.UserId})
		if setInviterReplyErr != nil {
			logger.WithError(setInviterReplyErr).Error("setInviterReplyErr")
		}

	}

	// 设置邀请人结束

	orderData.Status = card.OrderStatusCompleted

	//添加用户权益卡
	cardRPC.AddUserCard(orderData.UserId, orderData.ShopId, orderData.CardId, getItemReply.ValidityPeriod)

	//添加优惠券
	//item.ItemCoupons
	// 判断是否需要添加优惠券
	logger.Info("start stuffCoupon !")

	getItemCouponsReply, getItemCouponsReplyErr := cardRPC.GetItemCoupons(orderData.ItemId)
	if getItemCouponsReplyErr != nil {
		return nil, nil
	}

	if len(getItemCouponsReply) > 0 {
		userCouponApi := client.GetUserCouponService()
		if userCouponApi == nil {
			logger.Error("Get User Coupon Service error")
		} else {
			stuffCouponStuckList := make([]*couponProtos.StuffCouponStuck, 0, len(getItemReply.ItemCoupons))
			for _, itemCouponReply := range getItemCouponsReply {
				stuffCouponStuckList = append(stuffCouponStuckList, &couponProtos.StuffCouponStuck{
					CouponId: itemCouponReply.CouponId,
					Nums:     uint32(itemCouponReply.CouponCount),
				})
			}
			stuffUserCouponRequest := couponProtos.StuffUserCouponRequest{
				UserId:           orderData.UserId,
				StuffCouponStuck: stuffCouponStuckList,
			}
			logger.Info(stuffCouponStuckList)

			stuffUserCouponReply, stuffUserCouponReplyErr := userCouponApi.StuffUserCoupons(ctx, &stuffUserCouponRequest)
			if stuffUserCouponReplyErr != nil {
				logger.WithError(stuffUserCouponReplyErr).Error("stuffUserCouponReply error")
			}

			if stuffUserCouponReply == nil || stuffUserCouponReply.Err != nil {
				logger.WithError(stuffUserCouponReplyErr).Error("stuffUserCouponReply is nil or stuffUserCouponReply.Err is not nil")
			}
		}
	}

	// 完成赠送商品订单
	getItemGoodsReply, getItemGoodsReplyErr := cardRPC.GetItemGoods(orderData.ItemId)
	if getItemGoodsReplyErr != nil {
		return nil, nil
	}

	if len(getItemGoodsReply) > 0 {
		orderApi := client.GetOrderService()
		if orderApi == nil {
			logger.Error("Get Order Service error")
		} else {
			reply, err := orderApi.PaidOrder(ctx, &orderProtos.PaidOrderRequest{
				OrderId: orderData.OrderId,
				Payment: &orderProtos.Payment{
					TransactionId:  strconv.FormatUint(orderData.OrderId, 10),
					PaidPrice:      0,
					PaymentChannel: "fuyou",
					PaymentMode:    "miniProgram_card",
					PaymentProduct: "wechat",
					PaidTime:       uint64(time.Now().Unix()),
				},
			})
			if err != nil {
				logger.WithError(err).Error("get order detail error")
			} else if reply.Error != nil {
				logger.WithField("reply", reply).Info("get order detail reply error")
			}
		}
	}

	logger.Info("start rebate !")

	// 添加返利内容
	rebateApi := client.GetRebateService()

	if rebateApi == nil {
		logger.Error("Get rebate Service error")
	} else {
		rebateOrderData := &rebateProtos.RebateOrder{
			OrderId:   req.OrderId,
			UserId:    orderData.UserId,
			PaidPrice: orderData.Amount,
			PaidTime:  time.Now().Format("2006-01-02 15:04:05"),
			ItemType:  rebate.ItemTypeCard,
			ShopId:    orderData.ShopId,
			ItemId:    orderData.ItemId,
		}

		logger.WithField("rebateOrderData", rebateOrderData).Info("rebateOrder Data")

		// 直接怼数据进去
		_, createRebateOrderReplyErr := rebateApi.CreateRebateOrder(ctx, &rebateProtos.CreateRebateOrderRequest{
			RebateOrder: rebateOrderData,
		})

		if createRebateOrderReplyErr != nil {
			logger.WithError(createRebateOrderReplyErr).Error("createRebateOrderReplyErr")
		}
	}

	return &protos.FinishOrderReply{
		Order: cardRPC.ConvertModelOrderToProtobuf(orderData),
	}, nil
}

func finishOrder(orderId uint64) (bool, error) {
	record := &card.Order{
		OrderId: orderId,
		Status:  card.OrderStatusCompleted,
	}

	err := database.GetDB(constant.DatabaseConfigKey).Model(&card.Order{}).Update(record).Error
	if err != nil {
		log.GetLogger().WithField("order", record).WithError(err).Error("finish order record error")
		return false, err
	}

	return true, nil
}

func GetOrder(orderId uint64) (orderData *card.Order, err error) {
	orderData = new(card.Order)
	db := database.GetDB(constant.DatabaseConfigKey).
		Model(&card.Order{})

	err = db.Where(card.Order{OrderId: orderId}).First(&orderData).Error
	if err != nil {
		return nil, err
	}

	return
}

func validateFinishOderRequest(req *protos.FinishOrderRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.OrderId, cardRPC.ShopIdRule...),
	)
}
