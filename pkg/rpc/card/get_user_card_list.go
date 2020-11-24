package card

import (
	"time"

	rpclog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func (_ Controller) GetUserCardList(ctx context.Context, req *protos.GetUserCardListRequest) (*protos.GetUserCardListReply, error) {
	logger := rpclog.WithRequestId(ctx, log.GetLogger()).WithField("req", req)
	logger.Debug("received a request")

	err := validateGetUserCardList(req)
	if err != nil {
		logger.WithError(err).Info("get user card list validate failed")
		return &protos.GetUserCardListReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	logger.Debug("validate success")
	var userCardList []*card.User

	db := database.GetDB(constant.DatabaseConfigKey).Where("`userId` = ? AND `expireTime` > ?", req.GetUserId(), time.Now().Unix())

	if !validation.IsEmpty(req.GetShopId()) {
		db = db.Where("`shopId` = ?", req.GetShopId())
	}

	err = db.Find(&userCardList).Error
	if err != nil {
		logger.WithError(err).Error("get user card list error")
		return &protos.GetUserCardListReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeGetUserCardListFailed,
				Message: constant.ErrorMessageGetUserCardListFailed,
			},
		}, nil
	}

	logger.Debug("get user card list from database succeed")

	userCardListProtos := make([]*protos.UserCard, 0, len(userCardList))
	for _, userCardModel := range userCardList {
		userCardListProtos = append(userCardListProtos, convertToUserCardProtos(userCardModel))
	}

	logger.WithField("userCardList", userCardListProtos).Debug("reply user card list succeed")

	return &protos.GetUserCardListReply{
		UserCardList: userCardListProtos,
	}, nil
}

func validateGetUserCardList(req *protos.GetUserCardListRequest) error {
	return validation.ValidateStruct(req,
		//validation.Field(&req.ShopId, ShopIdRule...),
		validation.Field(&req.UserId, UserIdRule...),
	)
}
