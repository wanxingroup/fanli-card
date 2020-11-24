package card

import database "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"

type User struct {
	UserId     uint64 `gorm:"column:userId;type:bigint unsigned;primary_key;comment:'用户ID'"`
	ShopId     uint64 `gorm:"column:shopId;type:bigint unsigned;primary_key;comment:'店铺ID'"`
	CardId     uint64 `gorm:"column:cardId;type:bigint unsigned;primary_key;comment:'卡ID'"`
	ExpireTime uint32 `gorm:"column:expireTime;type:int unsigned;not null;default:'0';comment:'截止时间戳'"`
	database.Time
}

func (user *User) TableName() string {
	return "user_card"
}
