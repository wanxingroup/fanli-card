package card

import database "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"

const TableNameItem = "card_item"

type Item struct {
	ItemId            uint64        `gorm:"column:itemId;type:bigint unsigned;primary_key;comment:'权益卡选项ID'"`
	CardId            uint64        `gorm:"column:cardId;type:bigint unsigned;comment:'权益卡ID'"`
	Name              string        `gorm:"column:name;type:varchar(40);not null;default:'';comment:'名称'"`
	Description       string        `gorm:"column:description;type:varchar(300);not null;default:'';comment:'描述'"`
	Price             uint64        `gorm:"column:price;type:bigint unsigned;not null;default:'0';comment:'售价，单位分'"`
	RenewPrice        uint64        `gorm:"column:renewPrice;type:bigint unsigned;not null;default:'0';comment:'续费售价，单位：分'"`
	ValidityPeriod    uint32        `gorm:"column:validityPeriod;type:int unsigned;not null;default:'0';comment:'有效天数'"`
	Sort              uint32        `gorm:"column:sort;type:int unsigned;not null;default:'0';comment:'排序'"`
	ItemCoupons       []*ItemCoupon `gorm:"foreignkey:itemId"`
	ItemGoods         []*ItemGoods  `gorm:"foreignkey:itemId"`
	FirstRebateRatio  uint32        `gorm:"-"`
	SecondRebateRatio uint32        `gorm:"-"`
	database.Time
}

func (item *Item) TableName() string {
	return TableNameItem
}
