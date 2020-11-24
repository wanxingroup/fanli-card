package card

import database "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"

type ItemGoods struct {
	ItemId     uint64 `gorm:"column:itemId;type:bigint unsigned;primary_key;comment:'权益卡选项 ID'"`
	GoodsId    uint64 `gorm:"column:goodsId;type:bigint unsigned;primary_key;comment:'商品 Id'"`
	GoodsCount uint64 `gorm:"column:goodsCount;type:bigint unsigned;not null;default:'0';comment:'实物数量'"`
	database.Time
}

func (itemCoupon *ItemGoods) TableName() string {
	return "card_item_goods"
}
