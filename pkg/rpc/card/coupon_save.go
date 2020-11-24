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

type CouponSaveOption func(save *CouponSave)

type CouponSave struct {
	logger             *logrus.Entry
	itemId             uint64
	database           *gorm.DB
	originalCouponList []*card.ItemCoupon
	expectCouponList   []*protos.ItemCouponInformation

	needCreateCouponIds []uint64
	needUpdateCouponIds []uint64
	needDeleteCouponIds []uint64
	expectCouponMap     map[uint64]*protos.ItemCouponInformation
	originalCouponMap   map[uint64]*card.ItemCoupon
	finalCouponList     []*card.ItemCoupon
}

func NewCouponSave(options ...CouponSaveOption) *CouponSave {

	save := &CouponSave{
		logger:             log.GetLogger(),
		itemId:             0,
		database:           database.GetDB(constant.DatabaseConfigKey),
		originalCouponList: make([]*card.ItemCoupon, 0),
		expectCouponList:   make([]*protos.ItemCouponInformation, 0),
	}

	for _, option := range options {
		option(save)
	}

	return save
}

func CouponSaveLogger(logger *logrus.Entry) CouponSaveOption {
	return func(save *CouponSave) {
		save.logger = logger
	}
}

func CouponSaveItemId(itemId uint64) CouponSaveOption {
	return func(save *CouponSave) {
		save.itemId = itemId
	}
}

func CouponSaveDatabase(db *gorm.DB) CouponSaveOption {
	return func(save *CouponSave) {
		save.database = db
	}
}

func CouponSaveOriginalCouponList(originalCouponList []*card.ItemCoupon) CouponSaveOption {
	return func(save *CouponSave) {
		save.originalCouponList = originalCouponList
	}
}

func CouponSaveExpectCouponList(expectCouponList []*protos.ItemCouponInformation) CouponSaveOption {
	return func(save *CouponSave) {
		save.expectCouponList = expectCouponList
	}
}

func (save *CouponSave) Save() ([]*card.ItemCoupon, error) {

	if save.itemId <= 0 {
		return nil, fmt.Errorf("itemId is zero")
	}

	save.init()

	save.buildMap()

	save.computeOperationCouponIds()

	if err := save.updateCoupon(); err != nil {

		return nil, err
	}

	if err := save.createCoupon(); err != nil {

		return nil, err
	}

	if err := save.deleteCoupon(); err != nil {

		return nil, err
	}

	return save.finalCouponList, nil
}

func (save *CouponSave) buildMap() {

	for _, coupon := range save.expectCouponList {

		save.expectCouponMap[coupon.GetCouponId()] = coupon
	}

	for _, coupon := range save.originalCouponList {

		save.originalCouponMap[coupon.CouponId] = coupon
	}
}

func (save *CouponSave) init() {

	save.expectCouponMap = make(map[uint64]*protos.ItemCouponInformation)
	save.originalCouponMap = make(map[uint64]*card.ItemCoupon)
	save.finalCouponList = make([]*card.ItemCoupon, 0)
	save.needUpdateCouponIds = make([]uint64, 0)
	save.needCreateCouponIds = make([]uint64, 0)
	save.needDeleteCouponIds = make([]uint64, 0)
}

func (save *CouponSave) computeOperationCouponIds() {

	for couponId := range save.expectCouponMap {

		if _, exist := save.originalCouponMap[couponId]; exist {

			save.needUpdateCouponIds = append(save.needUpdateCouponIds, couponId)
		} else {

			save.needCreateCouponIds = append(save.needCreateCouponIds, couponId)
		}
	}

	for couponId := range save.originalCouponMap {

		if _, exist := save.expectCouponMap[couponId]; exist {
			continue
		}

		save.needDeleteCouponIds = append(save.needDeleteCouponIds, couponId)
	}
}

func (save *CouponSave) updateCoupon() error {

	for _, couponId := range save.needUpdateCouponIds {

		coupon := save.originalCouponMap[couponId]
		coupon.CouponCount = save.expectCouponMap[couponId].GetCouponCount()
		err := save.database.Save(coupon).Error
		if err != nil {
			save.logger.WithError(err).WithField("coupon", coupon).Error("update coupon error")
			return err
		}

		save.finalCouponList = append(save.finalCouponList, coupon)
	}

	return nil
}

func (save *CouponSave) createCoupon() error {

	for _, couponId := range save.needCreateCouponIds {

		coupon := &card.ItemCoupon{
			ItemId:      save.itemId,
			CouponId:    couponId,
			CouponCount: save.expectCouponMap[couponId].GetCouponCount(),
		}

		err := save.database.Save(coupon).Error
		if err != nil {
			save.logger.WithError(err).WithField("coupon", coupon).Error("create coupon error")
			return err
		}

		save.finalCouponList = append(save.finalCouponList, coupon)
	}

	return nil
}

func (save *CouponSave) deleteCoupon() error {

	for _, couponId := range save.needDeleteCouponIds {

		err := save.database.Delete(save.originalCouponMap[couponId]).Error
		if err != nil {
			save.logger.WithError(err).WithField("coupon", save.originalCouponMap[couponId]).Error("delete coupon error")
			return err
		}
	}

	return nil
}
