package order

import (
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func GetOrderList(query *protos.GetOrderListRequest) (orderList []*card.Order, count uint64, err error) {

	db := database.GetDB(constant.DatabaseConfigKey).Model(&card.Order{})

	if query.Status >= 0 {
		db = db.Where("`status` = ?", query.Status)
	}

	if query.UserId > 0 {
		db = db.Where("`userId` = ?", query.UserId)
	}

	if query.ShopId > 0 {
		db = db.Where("`shopId` = ?", query.ShopId)
	}

	if query.OrderId > 0 {
		db = db.Where("`orderId` = ?", query.OrderId)
	}

	if query.CardId > 0 {
		db = db.Where("`cardId` = ?", query.CardId)
	}

	if query.ItemId > 0 {
		db = db.Where("`itemId` = ?", query.ItemId)
	}

	if query.CreateTime != nil {
		if query.CreateTime.StartTime > 0 {
			db = db.Where("createdAt >= ?", time.Unix(int64(query.CreateTime.StartTime), 0))
		}

		if query.CreateTime.EndTime > 0 {
			db = db.Where("createdAt <= ?", time.Unix(int64(query.CreateTime.EndTime), 0))
		}
	}

	if query.Page > 0 {
		db = db.Offset((query.Page - 1) * query.PageSize)
	}

	err = db.Order("orderId desc").Limit(query.PageSize).
		Find(&orderList).
		Error

	log.GetLogger().WithField("GetOrderList db", db)

	if err != nil {
		return
	}

	err = db.Count(&count).Error

	return orderList, count, err

}

func convertOrderList(orders []*card.Order) []*protos.Order {

	if len(orders) == 0 {
		return []*protos.Order{}
	}

	result := make([]*protos.Order, 0, len(orders))
	for _, orderData := range orders {

		result = append(result, convertOrder(orderData))
	}

	return result
}

func convertOrder(orderData *card.Order) *protos.Order {
	return &protos.Order{
		OrderId:    orderData.OrderId,
		ShopId:     orderData.ShopId,
		UserId:     orderData.UserId,
		Status:     uint32(orderData.Status),
		Summary:    orderData.Summary,
		CardId:     orderData.CardId,
		ItemId:     orderData.ItemId,
		Amount:     orderData.Amount,
		CreateTime: uint64(orderData.CreatedAt.Unix()),
	}

}

type TimeRange struct {
	StartTime time.Time
	EndTime   time.Time
}
