package card

import (
	"fmt"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/errors"
	rpclog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func (_ Controller) SetCardItems(ctx context.Context, req *protos.SetCardItemsRequest) (reply *protos.SetCardItemsReply, _ error) {

	logger := rpclog.WithRequestId(ctx, log.GetLogger()).WithField("requestData", req)

	logger.Info("request set card items")

	if req == nil {
		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	var err error
	if err = validateSetItems(req); err != nil {
		logger.WithError(err).Info("validate failed")
		return &protos.SetCardItemsReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	cardData, err := getCard(req.GetCardId())
	if err != nil {
		logger.WithError(err).Error("get card data error")
		return &protos.SetCardItemsReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	if cardData == nil {
		logger.Warn("card data not found")
		return &protos.SetCardItemsReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeCardNotExist,
				Message: constant.ErrorMessageCardNotExist,
			},
		}, nil
	}

	transaction := database.GetDB(constant.DatabaseConfigKey).Begin()

	defer func() {
		var transactionError error
		if err != nil {
			transactionError = transaction.Rollback().Error
		} else {
			transactionError = transaction.Commit().Error
		}

		if transactionError != nil {

			logger.WithError(transactionError).Error("transaction error")
			reply = &protos.SetCardItemsReply{
				CardId: req.GetCardId(),
				Err: &protos.Error{
					Code:    constant.ErrorCodeTransactionError,
					Message: constant.ErrorMessageTransactionError,
				},
			}
		}
	}()

	itemList, err := NewItemSave(
		ItemSaveLogger(logger),
		ItemSaveCardId(cardData.CardId),
		ItemSaveDatabase(transaction),
		ItemSaveExpectItemList(req.Items),
		ItemSaveOriginalItemList(cardData.Items),
	).Save()
	if err != nil {
		logger.WithError(err).Error("get item data error")
		reply = &protos.SetCardItemsReply{
			CardId: cardData.CardId,
			Err: &protos.Error{
				Code:    errors.CodeServerInternalError,
				Message: err.Error(),
			},
		}
		return reply, nil
	}

	cardData.Items = itemList

	return &protos.SetCardItemsReply{
		CardId: cardData.CardId,
		Items:  convertModelCardItemsToProtobuf(cardData.Items),
	}, nil
}

func validateSetItems(req *protos.SetCardItemsRequest) error {

	return validation.ValidateStruct(req,
		validation.Field(&req.CardId, CardIdRule...),
		validation.Field(&req.Items, validation.By(func(value interface{}) error {
			itemList, ok := value.([]*protos.ItemInformation)
			if !ok {
				return validation.NewError(constant.ErrorCodeDataStructureInvalid, constant.ErrorMessageDataStructureInvalid)
			}
			for _, item := range itemList {

				err := validateItem(item)
				if err != nil {
					return err
				}
			}
			return nil
		})),
	)
}

func validateItem(req *protos.ItemInformation) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Name, NameRule...),
		validation.Field(&req.Description, DescriptionRule...),
		validation.Field(&req.Price, PriceRule...),
		validation.Field(&req.RenewPrice, RenewPriceRule...),
		validation.Field(&req.Sort, SortRule...),
		validation.Field(&req.ValidityPeriod, ValidityPeriodRule...),
		validation.Field(&req.Coupons, validation.By(func(value interface{}) error {
			couponList, ok := value.([]*protos.ItemCouponInformation)
			if !ok {
				return validation.NewError(constant.ErrorCodeDataStructureInvalid, constant.ErrorMessageDataStructureInvalid)
			}

			for _, coupon := range couponList {

				err := validateCouponRelated(coupon)
				if err != nil {
					return err
				}
			}

			return nil
		})),
	)
}

func validateCouponRelated(req *protos.ItemCouponInformation) error {

	return validation.ValidateStruct(req,
		validation.Field(&req.CouponId, CouponIdRule...),
		validation.Field(&req.CouponCount, CouponCountRule...),
	)
}
