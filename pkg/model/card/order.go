package card

import (
	database "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

type Order struct {
	OrderId uint64      `gorm:"column:orderId;type:bigint unsigned;primary_key;commit:'订单ID'"`
	UserId  uint64      `gorm:"column:userId;type:bigint unsigned;comment:'用户ID'"`
	ShopId  uint64      `gorm:"column:shopId;type:bigint unsigned;comment:'店铺ID'"`
	CardId  uint64      `gorm:"column:cardId;type:bigint unsigned;comment:'卡ID'"`
	ItemId  uint64      `gorm:"column:itemId;type:bigint unsigned;comment:'卡价格选项ID'"`
	Amount  uint64      `gorm:"column:amount;type:bigint unsigned;comment:'总价(分)'"`
	Summary string      `gorm:"column:summary;type:varchar(100);not null;default:'';comment:'简介'"`
	Status  OrderStatus `gorm:"column:status;type:tinyint unsigned;not null;default:'0';comment:'状态'"`

	database.Time
}

func (_ *Order) TableName() string {
	return "order_card"
}
