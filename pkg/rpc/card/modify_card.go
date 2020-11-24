package card

import (
	"fmt"

	rpclog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	context "golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func (_ Controller) ModifyCard(ctx context.Context, req *protos.ModifyCardRequest) (*protos.ModifyCardReply, error) {

	logger := rpclog.WithRequestId(ctx, log.GetLogger())

	if req == nil {
		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	err := validateModifyCard(req)
	if err != nil {
		logger.WithError(err).Error("validate modify card error")
		return &protos.ModifyCardReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	var cardId uint64
	cardId, err = ModifyCard(req)
	if err != nil {
		logger.WithError(err).Error("modify card error")
		return &protos.ModifyCardReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeModifyCardFailed,
				Message: constant.ErrorMessageModifyCardFailed,
				Stack:   nil,
			},
		}, nil
	}

	return &protos.ModifyCardReply{
		CardId: cardId,
	}, nil
}

func ModifyCard(req *protos.ModifyCardRequest) (uint64, error) {

	record := &card.Card{
		CardId:          req.Card.CardId,
		Name:            req.Card.Name,
		Description:     req.Card.Description,
		BackgroundImage: req.Card.BackgroundImage,
		Sort:            req.Card.Sort,
	}

	err := database.GetDB(constant.DatabaseConfigKey).Model(record).Update(record).Error
	if err != nil {
		log.GetLogger().WithField("card", record).WithError(err).Error("modify card record error")
		return 0, err
	}

	return record.CardId, nil
}

func validateModifyCard(req *protos.ModifyCardRequest) error {
	return validation.ValidateStruct(req.Card,
		validation.Field(&req.Card.CardId, CardIdRule...),
		validation.Field(&req.Card.Name, NameRule...),
		validation.Field(&req.Card.Description, DescriptionRule...),
		validation.Field(&req.Card.Sort, SortRule...),
		validation.Field(&req.Card.BackgroundImage, BackgroundImageRule...),
	)
}
