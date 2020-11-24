package card

import (
	"fmt"

	rpclog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	idcreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func (_ Controller) CreateCard(ctx context.Context, req *protos.CreateCardRequest) (*protos.CreateCardReply, error) {

	logger := rpclog.WithRequestId(ctx, log.GetLogger())

	if req == nil {
		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	err := validateCreateCard(req)
	if err != nil {
		return &protos.CreateCardReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	var cardId uint64
	cardId, err = createCard(req)
	if err != nil {
		logger.WithError(err).Error("create card error")
		return &protos.CreateCardReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeCreateCardFailed,
				Message: constant.ErrorMessageCreateCardFailed,
				Stack:   nil,
			},
		}, nil
	}

	return &protos.CreateCardReply{
		CardId: cardId,
	}, nil
}

func createCard(req *protos.CreateCardRequest) (uint64, error) {

	record := &card.Card{
		CardId:          idcreator.NextID(),
		ShopId:          req.ShopId,
		Name:            req.Name,
		Description:     req.Description,
		BackgroundImage: req.BackgroundImage,
		Status:          card.StatusUnused,
		Sort:            req.Sort,
	}

	err := database.GetDB(constant.DatabaseConfigKey).Create(record).Error
	if err != nil {
		log.GetLogger().WithField("card", record).WithError(err).Error("create record error")
		return 0, err
	}

	return record.CardId, nil
}

func validateCreateCard(req *protos.CreateCardRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.ShopId, ShopIdRule...),
		validation.Field(&req.Name, NameRule...),
		validation.Field(&req.Description, DescriptionRule...),
		validation.Field(&req.Sort, SortRule...),
		validation.Field(&req.BackgroundImage, BackgroundImageRule...),
	)
}
