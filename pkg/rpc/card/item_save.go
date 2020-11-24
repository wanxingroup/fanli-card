package card

import (
	"fmt"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	idcreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

type ItemSaveOption func(save *ItemSave)

type ItemSave struct {
	logger           *logrus.Entry
	cardId           uint64
	database         *gorm.DB
	originalItemList []*card.Item
	expectItemList   []*protos.ItemInformation

	needCreateItemIds []uint64
	needUpdateItemIds []uint64
	needDeleteItemIds []uint64
	expectItemMap     map[uint64]*protos.ItemInformation
	originalItemMap   map[uint64]*card.Item
	finalItemList     []*card.Item
}

func NewItemSave(options ...ItemSaveOption) *ItemSave {

	save := &ItemSave{
		logger:           log.GetLogger(),
		cardId:           0,
		database:         database.GetDB(constant.DatabaseConfigKey),
		originalItemList: make([]*card.Item, 0),
		expectItemList:   make([]*protos.ItemInformation, 0),
	}

	for _, option := range options {

		option(save)
	}

	return save
}

func ItemSaveLogger(logger *logrus.Entry) ItemSaveOption {
	return func(save *ItemSave) {
		save.logger = logger
	}
}

func ItemSaveCardId(cardId uint64) ItemSaveOption {
	return func(save *ItemSave) {
		save.cardId = cardId
	}
}

func ItemSaveDatabase(db *gorm.DB) ItemSaveOption {
	return func(save *ItemSave) {
		save.database = db
	}
}

func ItemSaveOriginalItemList(originalItemList []*card.Item) ItemSaveOption {
	return func(save *ItemSave) {
		save.originalItemList = originalItemList
	}
}

func ItemSaveExpectItemList(expectItemList []*protos.ItemInformation) ItemSaveOption {
	return func(save *ItemSave) {
		save.expectItemList = expectItemList
	}
}

func (save *ItemSave) Save() ([]*card.Item, error) {

	if save.cardId <= 0 {
		return nil, fmt.Errorf("cardId is zero")
	}

	save.init()

	save.fillNewItemId()

	save.buildMap()

	save.computeOperationItemIds()

	if err := save.updateItem(); err != nil {

		return nil, err
	}

	if err := save.createItem(); err != nil {

		return nil, err
	}

	if err := save.deleteItem(); err != nil {

		return nil, err
	}

	return save.finalItemList, nil
}

func (save *ItemSave) init() {

	save.needCreateItemIds = make([]uint64, 0)
	save.needUpdateItemIds = make([]uint64, 0)
	save.needDeleteItemIds = make([]uint64, 0)
	save.expectItemMap = make(map[uint64]*protos.ItemInformation)
	save.originalItemMap = make(map[uint64]*card.Item)
	save.finalItemList = make([]*card.Item, 0)
}

func (save *ItemSave) buildMap() {

	for _, item := range save.expectItemList {

		save.expectItemMap[item.GetItemId()] = item
	}

	for _, item := range save.originalItemList {

		save.originalItemMap[item.ItemId] = item
	}
}

func (save *ItemSave) computeOperationItemIds() {

	for itemId := range save.expectItemMap {

		if _, exist := save.originalItemMap[itemId]; exist {

			save.needUpdateItemIds = append(save.needUpdateItemIds, itemId)
		} else {

			save.needCreateItemIds = append(save.needCreateItemIds, itemId)
		}
	}

	for itemId := range save.originalItemMap {

		if _, exist := save.expectItemMap[itemId]; exist {
			continue
		}

		save.needDeleteItemIds = append(save.needDeleteItemIds, itemId)
	}
}

func (save *ItemSave) updateItem() error {

	for _, itemId := range save.needUpdateItemIds {

		itemData := save.originalItemMap[itemId]
		item := save.expectItemMap[itemId]

		itemData.Name = item.GetName()
		itemData.Description = item.GetDescription()
		itemData.Price = item.GetPrice()
		itemData.RenewPrice = item.GetRenewPrice()
		itemData.ValidityPeriod = item.GetValidityPeriod()
		itemData.Sort = item.GetSort()

		if err := save.database.Save(itemData).Error; err != nil {
			save.logger.WithError(err).WithField("item", item).Error("update item error")
			return err
		}

		couponList, err := NewCouponSave(
			CouponSaveLogger(save.logger),
			CouponSaveItemId(itemData.ItemId),
			CouponSaveDatabase(save.database),
			CouponSaveExpectCouponList(item.GetCoupons()),
			CouponSaveOriginalCouponList(itemData.ItemCoupons),
		).Save()
		if err != nil {
			save.logger.WithError(err).WithField("item", item).Error("update coupon error")
			return err
		}

		goodsList, err := NewGoodsSave(
			GoodsSaveLogger(save.logger),
			GoodsSaveItemId(itemData.ItemId),
			GoodsSaveDatabase(save.database),
			GoodsSaveExpectGoodsList(item.GetGoods()),
			GoodsSaveOriginalGoodsList(itemData.ItemGoods),
		).Save()
		if err != nil {
			save.logger.WithError(err).WithField("item", item).Error("update goods error")
			return err
		}

		itemData.ItemCoupons = couponList
		itemData.ItemGoods = goodsList
		itemData.FirstRebateRatio = item.GetFirstRebateRatio()
		itemData.SecondRebateRatio = item.GetSecondRebateRatio()

		save.finalItemList = append(save.finalItemList, itemData)
	}

	return nil
}

func (save *ItemSave) createItem() error {

	for _, itemId := range save.needCreateItemIds {

		item := save.expectItemMap[itemId]
		itemData := &card.Item{
			ItemId:         itemId,
			CardId:         save.cardId,
			Name:           item.GetName(),
			Description:    item.GetDescription(),
			Price:          item.GetPrice(),
			RenewPrice:     item.GetRenewPrice(),
			ValidityPeriod: item.GetValidityPeriod(),
			Sort:           item.GetSort(),
			ItemCoupons:    make([]*card.ItemCoupon, 0),
			ItemGoods:      make([]*card.ItemGoods, 0),
		}

		err := save.database.Save(itemData).Error
		if err != nil {
			save.logger.WithError(err).WithField("item", itemData).Error("create item error")
			return err
		}

		couponList, err := NewCouponSave(
			CouponSaveLogger(save.logger),
			CouponSaveItemId(itemData.ItemId),
			CouponSaveDatabase(save.database),
			CouponSaveExpectCouponList(item.GetCoupons()),
			CouponSaveOriginalCouponList(itemData.ItemCoupons),
		).Save()
		if err != nil {
			save.logger.WithError(err).WithField("item", item).Error("create coupon error")
			return err
		}

		goodsList, err := NewGoodsSave(
			GoodsSaveLogger(save.logger),
			GoodsSaveItemId(itemData.ItemId),
			GoodsSaveDatabase(save.database),
			GoodsSaveExpectGoodsList(item.GetGoods()),
			GoodsSaveOriginalGoodsList(itemData.ItemGoods),
		).Save()
		if err != nil {
			save.logger.WithError(err).WithField("item", item).Error("create goods error")
			return err
		}

		itemData.ItemCoupons = couponList
		itemData.ItemGoods = goodsList
		itemData.FirstRebateRatio = item.GetFirstRebateRatio()
		itemData.SecondRebateRatio = item.GetSecondRebateRatio()

		save.finalItemList = append(save.finalItemList, itemData)
	}

	return nil
}

func (save *ItemSave) fillNewItemId() {

	for _, item := range save.expectItemList {

		if item.ItemId > 0 {
			continue
		}

		item.ItemId = idcreator.NextID()
	}
}

func (save *ItemSave) deleteItem() error {

	for _, itemId := range save.needDeleteItemIds {

		err := save.database.Delete(save.originalItemMap[itemId]).Error
		if err != nil {
			save.logger.WithError(err).WithField("item", save.originalItemMap[itemId]).Error("delete item error")
			return err
		}
	}

	return nil
}
