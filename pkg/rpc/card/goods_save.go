package card

import (
	"fmt"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

type GoodsSaveOption func(save *GoodsSave)

type GoodsSave struct {
	logger            *logrus.Entry
	itemId            uint64
	database          *gorm.DB
	originalGoodsList []*card.ItemGoods
	expectGoodsList   []*protos.ItemGoodsInformation

	needCreateGoodsIds []uint64
	needUpdateGoodsIds []uint64
	needDeleteGoodsIds []uint64
	expectGoodsMap     map[uint64]*protos.ItemGoodsInformation
	originalGoodsMap   map[uint64]*card.ItemGoods
	finalGoodsList     []*card.ItemGoods
}

func NewGoodsSave(options ...GoodsSaveOption) *GoodsSave {

	save := &GoodsSave{
		logger:            log.GetLogger(),
		itemId:            0,
		database:          database.GetDB(constant.DatabaseConfigKey),
		originalGoodsList: make([]*card.ItemGoods, 0),
		expectGoodsList:   make([]*protos.ItemGoodsInformation, 0),
	}

	for _, option := range options {
		option(save)
	}

	return save
}

func GoodsSaveLogger(logger *logrus.Entry) GoodsSaveOption {
	return func(save *GoodsSave) {
		save.logger = logger
	}
}

func GoodsSaveItemId(itemId uint64) GoodsSaveOption {
	return func(save *GoodsSave) {
		save.itemId = itemId
	}
}

func GoodsSaveDatabase(db *gorm.DB) GoodsSaveOption {
	return func(save *GoodsSave) {
		save.database = db
	}
}

func GoodsSaveOriginalGoodsList(originalGoodsList []*card.ItemGoods) GoodsSaveOption {
	return func(save *GoodsSave) {
		save.originalGoodsList = originalGoodsList
	}
}

func GoodsSaveExpectGoodsList(expectGoodsList []*protos.ItemGoodsInformation) GoodsSaveOption {
	return func(save *GoodsSave) {
		save.expectGoodsList = expectGoodsList
	}
}

func (save *GoodsSave) Save() ([]*card.ItemGoods, error) {

	if save.itemId <= 0 {
		return nil, fmt.Errorf("itemId is zero")
	}

	save.init()

	save.buildMap()

	save.computeOperationGoodsIds()

	if err := save.updateGoods(); err != nil {

		return nil, err
	}

	if err := save.createGoods(); err != nil {

		return nil, err
	}

	if err := save.deleteGoods(); err != nil {

		return nil, err
	}

	return save.finalGoodsList, nil
}

func (save *GoodsSave) buildMap() {

	for _, Goods := range save.expectGoodsList {

		save.expectGoodsMap[Goods.GetGoodsId()] = Goods
	}

	for _, Goods := range save.originalGoodsList {

		save.originalGoodsMap[Goods.GoodsId] = Goods
	}
}

func (save *GoodsSave) init() {

	save.expectGoodsMap = make(map[uint64]*protos.ItemGoodsInformation)
	save.originalGoodsMap = make(map[uint64]*card.ItemGoods)
	save.finalGoodsList = make([]*card.ItemGoods, 0)
	save.needUpdateGoodsIds = make([]uint64, 0)
	save.needCreateGoodsIds = make([]uint64, 0)
	save.needDeleteGoodsIds = make([]uint64, 0)
}

func (save *GoodsSave) computeOperationGoodsIds() {

	for GoodsId := range save.expectGoodsMap {

		if _, exist := save.originalGoodsMap[GoodsId]; exist {

			save.needUpdateGoodsIds = append(save.needUpdateGoodsIds, GoodsId)
		} else {

			save.needCreateGoodsIds = append(save.needCreateGoodsIds, GoodsId)
		}
	}

	for GoodsId := range save.originalGoodsMap {

		if _, exist := save.expectGoodsMap[GoodsId]; exist {
			continue
		}

		save.needDeleteGoodsIds = append(save.needDeleteGoodsIds, GoodsId)
	}
}

func (save *GoodsSave) updateGoods() error {

	for _, GoodsId := range save.needUpdateGoodsIds {

		Goods := save.originalGoodsMap[GoodsId]
		Goods.GoodsCount = save.expectGoodsMap[GoodsId].GetGoodsCount()
		err := save.database.Save(Goods).Error
		if err != nil {
			save.logger.WithError(err).WithField("Goods", Goods).Error("update Goods error")
			return err
		}

		save.finalGoodsList = append(save.finalGoodsList, Goods)
	}

	return nil
}

func (save *GoodsSave) createGoods() error {

	for _, GoodsId := range save.needCreateGoodsIds {

		Goods := &card.ItemGoods{
			ItemId:     save.itemId,
			GoodsId:    GoodsId,
			GoodsCount: save.expectGoodsMap[GoodsId].GetGoodsCount(),
		}

		err := save.database.Save(Goods).Error
		if err != nil {
			save.logger.WithError(err).WithField("Goods", Goods).Error("create Goods error")
			return err
		}

		save.finalGoodsList = append(save.finalGoodsList, Goods)
	}

	return nil
}

func (save *GoodsSave) deleteGoods() error {

	for _, GoodsId := range save.needDeleteGoodsIds {

		err := save.database.Delete(save.originalGoodsMap[GoodsId]).Error
		if err != nil {
			save.logger.WithError(err).WithField("Goods", save.originalGoodsMap[GoodsId]).Error("delete Goods error")
			return err
		}
	}

	return nil
}
