package card

import database "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"

type ItemCoupon struct {
	ItemId      uint64 `gorm:"column:itemId;type:bigint unsigned;primary_key;comment:'权益卡选项 ID'"`
	CouponId    uint64 `gorm:"column:couponId;type:bigint unsigned;primary_key;comment:'优惠券ID'"`
	CouponCount uint64 `gorm:"column:couponCount;type:bigint unsigned;not null;default:'0';comment:'优惠券数量'"`
	database.Time
}

func (itemCoupon *ItemCoupon) TableName() string {
	return "card_item_coupon"
}
