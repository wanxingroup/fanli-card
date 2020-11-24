package card

import (
	"fmt"

	rpclog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func (_ Controller) RemoveCard(ctx context.Context, req *protos.RemoveCardRequest) (reply *protos.RemoveCardReply, _ error) {

	logger := rpclog.WithRequestId(ctx, log.GetLogger())

	if req == nil {
		logger.Warn("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	var err error
	err = validateRemoveCardData(req)
	if err != nil {
		return &protos.RemoveCardReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	var cardData *card.Card

	cardData, err = getCard(req.GetCardId())
	if err != nil {
		logger.WithError(err).Error("get card error")
		return &protos.RemoveCardReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeGetCardFailed,
				Message: constant.ErrorMessageGetCardFailed,
			},
		}, nil
	}

	transaction := database.GetDB(constant.DatabaseConfigKey).Begin()

	defer func() {
		if err != nil {
			transaction.Rollback()
			reply = &protos.RemoveCardReply{
				Err: ConvertErrorToProtobuf(err),
			}
		} else {
			transaction.Commit()
			reply = &protos.RemoveCardReply{}
		}
	}()

	err = transaction.Delete(cardData).Error
	if err != nil {
		return
	}

	for _, item := range cardData.Items {

		err = transaction.Delete(item).Error
		if err != nil {
			return
		}

		for _, coupon := range item.ItemCoupons {

			err = transaction.Delete(coupon).Error
			if err != nil {
				return
			}
		}
	}

	return reply, nil
}

func validateRemoveCardData(req *protos.RemoveCardRequest) error {

	return validation.ValidateStruct(req,
		validation.Field(&req.CardId, CardIdRule...),
		validation.Field(&req.ShopId, ShopIdRule...),
	)
}
