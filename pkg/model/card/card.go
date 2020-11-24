package card

import database "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"

const TableNameCard = "card"

type Card struct {
	CardId          uint64  `gorm:"column:cardId;type:bigint unsigned;primary_key;comment:'卡ID'"`
	ShopId          uint64  `gorm:"column:shopId;type:bigint unsigned;not null;default: '0';comment:'店铺ID'"`
	Name            string  `gorm:"column:name;type:varchar(40);not null;default:'';comment:'卡名称'"`
	Description     string  `gorm:"column:description;type:varchar(1024);not null;default:'';comment:'卡描述'"`
	BackgroundImage string  `gorm:"column:backgroundImage;type:varchar(200);not null;default:'';comment:'卡背景图'"`
	Status          Status  `gorm:"column:status;type:tinyint unsigned;not null;default:'0';comment:'卡状态'"`
	Sort            uint32  `gorm:"column:sort;type:int unsigned;not null;default:'0';comment:'卡排序'"`
	Items           []*Item `gorm:"foreignkey:cardId"`
	database.Time
}

func (card *Card) TableName() string {
	return TableNameCard
}
