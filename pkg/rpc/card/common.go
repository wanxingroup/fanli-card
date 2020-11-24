package card

import (
	"strconv"
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/errors"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func convertToCardInformation(cardStruct *card.Card) *protos.CardInformation {
	return &protos.CardInformation{
		CardId:          cardStruct.CardId,
		Name:            cardStruct.Name,
		Description:     cardStruct.Description,
		BackgroundImage: cardStruct.BackgroundImage,
		Status:          convertModelCardStatusToProtobuf(cardStruct.Status),
		Sort:            cardStruct.Sort,
		Items:           convertModelCardItemsToProtobuf(cardStruct.Items),
		CreateTime:      cardStruct.CreatedAt.Unix(),
	}
}

func convertToUserCardProtos(userCardModel *card.User) *protos.UserCard {
	return &protos.UserCard{
		CardId:     userCardModel.CardId,
		ExpireTime: userCardModel.ExpireTime,
	}
}

func convertModelCardStatusToProtobuf(status card.Status) protos.CardStatus {

	switch status {
	case card.StatusUnused:
		return protos.CardStatus_Unused
	case card.StatusInuse:
		return protos.CardStatus_Inuse
	}
	return protos.CardStatus_Unused
}

func convertModelCardItemsToProtobuf(items []*card.Item) (list []*protos.ItemInformation) {

	list = make([]*protos.ItemInformation, 0, len(items))

	for _, item := range items {

		list = append(list, convertModelCardItemToProtobuf(item))
	}

	return
}

func convertModelCardItemToProtobuf(item *card.Item) *protos.ItemInformation {

	if item == nil {
		return nil
	}

	return &protos.ItemInformation{
		ItemId:            item.ItemId,
		Name:              item.Name,
		Description:       item.Description,
		Price:             item.Price,
		RenewPrice:        item.RenewPrice,
		ValidityPeriod:    item.ValidityPeriod,
		Sort:              item.Sort,
		Coupons:           convertModelCardItemCouponRelatedToProtobuf(item.ItemCoupons),
		Goods:             convertModelCardItemGoodsRelatedToProtobuf(item.ItemGoods),
		FirstRebateRatio:  item.FirstRebateRatio,
		SecondRebateRatio: item.SecondRebateRatio,
	}
}

func convertModelCardItemCouponRelatedToProtobuf(coupons []*card.ItemCoupon) (list []*protos.ItemCouponInformation) {

	list = make([]*protos.ItemCouponInformation, 0, len(coupons))

	for _, coupon := range coupons {

		list = append(list, &protos.ItemCouponInformation{
			CouponId:    coupon.CouponId,
			CouponCount: coupon.CouponCount,
		})
	}

	return
}

func convertModelCardItemGoodsRelatedToProtobuf(coupons []*card.ItemGoods) (list []*protos.ItemGoodsInformation) {

	list = make([]*protos.ItemGoodsInformation, 0, len(coupons))

	for _, coupon := range coupons {

		list = append(list, &protos.ItemGoodsInformation{
			GoodsId:    coupon.GoodsId,
			GoodsCount: coupon.GoodsCount,
		})
	}

	return
}

func getCard(cardId uint64) (cardData *card.Card, err error) {
	cardData = new(card.Card)
	db := database.GetDB(constant.DatabaseConfigKey).
		Model(&card.Card{}).
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order(card.TableNameItem + ".sort ASC")
		}).
		Preload("Items.ItemCoupons").
		Preload("Items.ItemGoods")

	err = db.Where(card.Card{CardId: cardId}).First(&cardData).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, err
	}

	return
}

func GetItem(itemId uint64) (item *card.Item, err error) {
	item = new(card.Item)
	db := database.GetDB(constant.DatabaseConfigKey)
	err = db.Where(card.Item{ItemId: itemId}).First(&item).Error
	if err != nil {
		return nil, err
	}

	return
}

func GetItemCoupons(itemId uint64) (itemCoupons []*card.ItemCoupon, err error) {
	db := database.GetDB(constant.DatabaseConfigKey)
	err = db.Where(card.ItemCoupon{ItemId: itemId}).Find(&itemCoupons).Error
	if err != nil {
		return nil, err
	}

	return
}

func GetItemGoods(itemId uint64) (itemGoods []*card.ItemGoods, err error) {
	db := database.GetDB(constant.DatabaseConfigKey)
	err = db.Where(card.ItemGoods{ItemId: itemId}).Find(&itemGoods).Error
	if err != nil {
		return nil, err
	}

	return
}

func AddUserCard(userId uint64, shopId uint64, cardId uint64, validityPeriod uint32) bool {
	var cardValue = &card.User{
		UserId: userId,
		ShopId: shopId,
		CardId: cardId,
	}

	var userCard = new(card.User)
	var db = database.GetDB(constant.DatabaseConfigKey)
	err := db.Where(cardValue).First(&userCard).Error
	if err != nil {
		log.GetLogger().WithError(err).Info("AddUserCard find UserCard error")
		if gorm.IsRecordNotFoundError(err) {
			cardValue.ExpireTime = uint32(time.Now().Unix()) + validityPeriod*86400
			err := db.Create(cardValue).Error
			if err != nil {
				log.GetLogger().WithError(err).Error("AddUserCard Create UserCard error")
				return false
			}

			log.GetLogger().Info("AddUserCard Create UserCard success")
			return true
		}

		log.GetLogger().WithError(err).Error("AddUserCard find UserCard error")
		return false
	}

	if int64(userCard.ExpireTime) > time.Now().Unix() {
		userCard.ExpireTime += validityPeriod * 86400
	} else {
		userCard.ExpireTime = uint32(time.Now().Unix()) + validityPeriod*86400
	}

	dbSaveErr := db.Save(userCard).Error
	if dbSaveErr != nil {
		log.GetLogger().WithError(dbSaveErr).Error("AddUserCard Update UserCard error")
		return false
	}

	log.GetLogger().Info("AddUserCard Update UserCard success")
	return true

}

func ConvertErrorToProtobuf(err error) *protos.Error {

	if validationError, ok := err.(validation.Error); ok {
		errorCode, convertError := strconv.Atoi(validationError.Code())
		if convertError != nil {
			errorCode = errors.CodeServerInternalError
		}
		return &protos.Error{
			Code:    int64(errorCode),
			Message: validationError.Error(),
		}
	}

	return &protos.Error{
		Code:    errors.CodeServerInternalError,
		Message: err.Error(),
	}
}

func ConvertModelOrderToProtobuf(order *card.Order) *protos.Order {

	return &protos.Order{
		OrderId:    order.OrderId,
		UserId:     order.UserId,
		ShopId:     order.ShopId,
		CardId:     order.CardId,
		ItemId:     order.ItemId,
		Amount:     order.Amount,
		Status:     uint32(order.Status),
		Summary:    order.Summary,
		CreateTime: uint64(order.CreatedAt.Unix()),
	}
}
