package application

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
)

func autoMigration() {
	db := database.GetDB(constant.DatabaseConfigKey)
	db.AutoMigrate(card.Card{})
	db.AutoMigrate(card.User{})
	db.AutoMigrate(card.Item{})
	db.AutoMigrate(card.ItemCoupon{})
	db.AutoMigrate(card.ItemGoods{})
	db.AutoMigrate(card.Order{})
}
