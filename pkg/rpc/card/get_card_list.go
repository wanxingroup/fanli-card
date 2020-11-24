package card

import (
	rpclog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func (_ Controller) GetCardList(ctx context.Context, req *protos.GetCardListRequest) (*protos.GetCardListReply, error) {
	logger := rpclog.WithRequestId(ctx, log.GetLogger())
	err := validateGetCardList(req)
	if err != nil {
		return &protos.GetCardListReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	var condition = card.Card{ShopId: req.GetShopId()}
	if req.GetStatus() > 0 {
		condition.Status = card.Status(req.GetStatus())
	}

	var cardList []*card.Card
	var count uint64

	db := database.GetDB(constant.DatabaseConfigKey).
		Model(&card.Card{}).Order("`sort` ASC")

	if req.WithItemDetail {
		db = db.Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order(card.TableNameItem + ".sort ASC")
		})

		db = db.Preload("Items.ItemCoupons")
		db = db.Preload("Items.ItemGoods")
	}

	err = db.Where(condition).Find(&cardList).Count(&count).Error
	if err != nil {
		logger.WithError(err).Error("get card list error")
		return &protos.GetCardListReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeGetCardListFailed,
				Message: constant.ErrorMessageGetCardListFailed,
			},
		}, nil
	}

	cardInformationList := make([]*protos.CardInformation, 0, len(cardList))
	for _, cardStruct := range cardList {
		cardInformationList = append(cardInformationList, convertToCardInformation(cardStruct))
	}

	return &protos.GetCardListReply{
		CardInformationList: cardInformationList,
		Count:               count,
	}, nil
}

func validateGetCardList(req *protos.GetCardListRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.ShopId, ShopIdRule...),
	)
}
