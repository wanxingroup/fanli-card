package card

import (
	"fmt"

	rpclog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	context "golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func (_ Controller) SetCardStatus(ctx context.Context, req *protos.SetCardStatusRequest) (*protos.SetCardStatusReply, error) {
	logger := rpclog.WithRequestId(ctx, log.GetLogger()).WithField("requestData", req)
	logger.Info("request set card status")
	if req == nil {
		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	cardId, err := SetCardStatus(req)
	if err != nil {
		logger.WithError(err).Error("set card status error")
		return &protos.SetCardStatusReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeSetCardStatusFailed,
				Message: constant.ErrorMessageSetCardStatusFailed,
				Stack:   nil,
			},
		}, nil
	}

	return &protos.SetCardStatusReply{
		CardId: cardId,
	}, nil
}

func SetCardStatus(req *protos.SetCardStatusRequest) (uint64, error) {
	status := card.StatusUnused
	if req.Status == protos.CardStatus_Inuse {
		status = card.StatusInuse
	}

	record := &card.Card{
		CardId: req.CardId,
	}

	log.GetLogger().WithField("status", status).Info("SetCardStatus")
	err := database.GetDB(constant.DatabaseConfigKey).Model(record).Update(card.Card{Status: status}).Error
	if err != nil {
		log.GetLogger().WithField("card", record).WithError(err).Error("set card status record error")
		return 0, err
	}

	return record.CardId, nil
}
