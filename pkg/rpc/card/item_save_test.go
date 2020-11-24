package card

import (
	"sort"
	"testing"
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func TestItemSaveCardId(t *testing.T) {

	tests := []struct {
		input uint64
		want  uint64
	}{
		{
			input: 10,
			want:  10,
		},
		{
			input: 0,
			want:  0,
		},
	}

	for _, test := range tests {

		ItemSave := &ItemSave{}
		ItemSaveCardId(test.input)(ItemSave)
		assert.Equal(t, test.want, ItemSave.cardId, test)
	}
}

func TestItemSaveDatabase(t *testing.T) {

	tests := []struct {
		input *gorm.DB
		want  *gorm.DB
	}{
		{
			input: database.GetDB(constant.DatabaseConfigKey),
			want:  database.GetDB(constant.DatabaseConfigKey),
		},
	}

	for _, test := range tests {

		ItemSave := &ItemSave{}
		ItemSaveDatabase(test.input)(ItemSave)
		assert.Equal(t, test.want, ItemSave.database, test)
	}
}

func TestItemSaveLogger(t *testing.T) {

	tests := []struct {
		input *logrus.Entry
		want  *logrus.Entry
	}{
		{
			input: log.GetLogger(),
			want:  log.GetLogger(),
		},
	}

	for _, test := range tests {

		ItemSave := &ItemSave{}
		ItemSaveLogger(test.input)(ItemSave)
		assert.Equal(t, test.want, ItemSave.logger, test)
	}
}

func TestItemSaveExpectItemList(t *testing.T) {

	tests := []struct {
		input []*protos.ItemInformation
		want  []*protos.ItemInformation
	}{
		{
			input: []*protos.ItemInformation{},
			want:  []*protos.ItemInformation{},
		},
		{
			input: []*protos.ItemInformation{
				{
					ItemId:         10,
					Name:           "item 10",
					Description:    "description 10",
					Price:          10,
					RenewPrice:     9,
					ValidityPeriod: 2,
					Sort:           10,
					Coupons: []*protos.ItemCouponInformation{
						{
							CouponId:    1002,
							CouponCount: 1,
						},
					},
					Goods: []*protos.ItemGoodsInformation{
						{
							GoodsId:    1002,
							GoodsCount: 1,
						},
					},
				},
			},
			want: []*protos.ItemInformation{
				{
					ItemId:         10,
					Name:           "item 10",
					Description:    "description 10",
					Price:          10,
					RenewPrice:     9,
					ValidityPeriod: 2,
					Sort:           10,
					Coupons: []*protos.ItemCouponInformation{
						{
							CouponId:    1002,
							CouponCount: 1,
						},
					},
					Goods: []*protos.ItemGoodsInformation{
						{
							GoodsId:    1002,
							GoodsCount: 1,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {

		ItemSave := &ItemSave{}
		ItemSaveExpectItemList(test.input)(ItemSave)
		assert.Equal(t, test.want, ItemSave.expectItemList, test)
	}
}

func TestItemSaveOriginalItemList(t *testing.T) {

	now := time.Now()
	tests := []struct {
		input []*card.Item
		want  []*card.Item
	}{
		{
			input: []*card.Item{},
			want:  []*card.Item{},
		},
		{
			input: []*card.Item{
				{
					ItemId:         1001,
					CardId:         10,
					Name:           "name 1001",
					Description:    "description 1001",
					Price:          100,
					RenewPrice:     100,
					ValidityPeriod: 100,
					Sort:           20,
					ItemCoupons: []*card.ItemCoupon{
						{
							ItemId:      1001,
							CouponId:    10001,
							CouponCount: 1,
							Time: databases.Time{
								BasicTimeFields: databases.BasicTimeFields{
									CreatedAt: now,
									UpdatedAt: now,
								},
								DeletedAt: nil,
							},
						},
					},
					ItemGoods: []*card.ItemGoods{
						{
							ItemId:     1001,
							GoodsId:    10001,
							GoodsCount: 1,
							Time: databases.Time{
								BasicTimeFields: databases.BasicTimeFields{
									CreatedAt: now,
									UpdatedAt: now,
								},
								DeletedAt: nil,
							},
						},
					},
					Time: databases.Time{
						BasicTimeFields: databases.BasicTimeFields{
							CreatedAt: now,
							UpdatedAt: now,
						},
						DeletedAt: nil,
					},
				},
			},
			want: []*card.Item{
				{
					ItemId:         1001,
					CardId:         10,
					Name:           "name 1001",
					Description:    "description 1001",
					Price:          100,
					RenewPrice:     100,
					ValidityPeriod: 100,
					Sort:           20,
					ItemCoupons: []*card.ItemCoupon{
						{
							ItemId:      1001,
							CouponId:    10001,
							CouponCount: 1,
							Time: databases.Time{
								BasicTimeFields: databases.BasicTimeFields{
									CreatedAt: now,
									UpdatedAt: now,
								},
								DeletedAt: nil,
							},
						},
					},
					ItemGoods: []*card.ItemGoods{
						{
							ItemId:     1001,
							GoodsId:    10001,
							GoodsCount: 1,
							Time: databases.Time{
								BasicTimeFields: databases.BasicTimeFields{
									CreatedAt: now,
									UpdatedAt: now,
								},
								DeletedAt: nil,
							},
						},
					},
					Time: databases.Time{
						BasicTimeFields: databases.BasicTimeFields{
							CreatedAt: now,
							UpdatedAt: now,
						},
						DeletedAt: nil,
					},
				},
			},
		},
	}

	for _, test := range tests {

		ItemSave := &ItemSave{}
		ItemSaveOriginalItemList(test.input)(ItemSave)
		assert.Equal(t, test.want, ItemSave.originalItemList, test)
	}
}

func TestNewItemSave(t *testing.T) {

	tests := []struct {
		input []ItemSaveOption
		want  *ItemSave
	}{
		{
			input: []ItemSaveOption{
				ItemSaveCardId(10),
				ItemSaveLogger(log.GetLogger()),
				ItemSaveDatabase(database.GetDB(constant.DatabaseConfigKey)),
			},
			want: &ItemSave{
				logger:           log.GetLogger(),
				cardId:           10,
				database:         database.GetDB(constant.DatabaseConfigKey),
				originalItemList: make([]*card.Item, 0),
				expectItemList:   make([]*protos.ItemInformation, 0),
			},
		},
	}

	for _, test := range tests {

		assert.Equal(t, test.want, NewItemSave(test.input...), test)
	}
}

func TestItemSave_init(t *testing.T) {

	tests := []struct {
		want *ItemSave
	}{
		{
			want: &ItemSave{
				logger:            log.GetLogger(),
				cardId:            0,
				database:          database.GetDB(constant.DatabaseConfigKey),
				originalItemList:  make([]*card.Item, 0),
				expectItemList:    make([]*protos.ItemInformation, 0),
				needCreateItemIds: make([]uint64, 0),
				needUpdateItemIds: make([]uint64, 0),
				needDeleteItemIds: make([]uint64, 0),
				expectItemMap:     make(map[uint64]*protos.ItemInformation),
				originalItemMap:   make(map[uint64]*card.Item),
				finalItemList:     make([]*card.Item, 0),
			},
		},
	}

	for _, test := range tests {

		ItemSave := NewItemSave()
		ItemSave.init()
		assert.Equal(t, test.want, ItemSave, test)
	}
}

func TestItemSave_computeOperationItemIds(t *testing.T) {

	now := time.Now()
	tests := []struct {
		input *ItemSave
		want  *ItemSave
	}{
		{
			input: &ItemSave{
				logger:   log.GetLogger(),
				cardId:   1001,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{
					{
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemList: []*protos.ItemInformation{
					{
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
				},
				needCreateItemIds: []uint64{},
				needUpdateItemIds: []uint64{},
				needDeleteItemIds: []uint64{},
				originalItemMap: map[uint64]*card.Item{
					2001: {
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemMap: map[uint64]*protos.ItemInformation{
					2001: {
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
				},
				finalItemList: make([]*card.Item, 0),
			},
			want: &ItemSave{
				logger:   log.GetLogger(),
				cardId:   1001,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{
					{
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemList: []*protos.ItemInformation{
					{
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
				},
				needCreateItemIds: []uint64{},
				needUpdateItemIds: []uint64{2001},
				needDeleteItemIds: []uint64{},
				originalItemMap: map[uint64]*card.Item{
					2001: {
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemMap: map[uint64]*protos.ItemInformation{
					2001: {
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
				},
				finalItemList: make([]*card.Item, 0),
			},
		},
		{
			input: &ItemSave{
				logger:   log.GetLogger(),
				cardId:   1001,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{
					{
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemList:    []*protos.ItemInformation{},
				needCreateItemIds: []uint64{},
				needUpdateItemIds: []uint64{},
				needDeleteItemIds: []uint64{},
				originalItemMap: map[uint64]*card.Item{
					2001: {
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemMap: map[uint64]*protos.ItemInformation{},
				finalItemList: make([]*card.Item, 0),
			},
			want: &ItemSave{
				logger:   log.GetLogger(),
				cardId:   1001,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{
					{
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemList:    []*protos.ItemInformation{},
				needCreateItemIds: []uint64{},
				needUpdateItemIds: []uint64{},
				needDeleteItemIds: []uint64{2001},
				originalItemMap: map[uint64]*card.Item{
					2001: {
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemMap: map[uint64]*protos.ItemInformation{},
				finalItemList: make([]*card.Item, 0),
			},
		},
		{
			input: &ItemSave{
				logger:           log.GetLogger(),
				cardId:           1001,
				database:         database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{},
				expectItemList: []*protos.ItemInformation{
					{
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
				},
				needCreateItemIds: []uint64{},
				needUpdateItemIds: []uint64{},
				needDeleteItemIds: []uint64{},
				originalItemMap:   map[uint64]*card.Item{},
				expectItemMap: map[uint64]*protos.ItemInformation{
					2001: {
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
				},
				finalItemList: make([]*card.Item, 0),
			},
			want: &ItemSave{
				logger:           log.GetLogger(),
				cardId:           1001,
				database:         database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{},
				expectItemList: []*protos.ItemInformation{
					{
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
				},
				needCreateItemIds: []uint64{2001},
				needUpdateItemIds: []uint64{},
				needDeleteItemIds: []uint64{},
				originalItemMap:   map[uint64]*card.Item{},
				expectItemMap: map[uint64]*protos.ItemInformation{
					2001: {
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
				},
				finalItemList: make([]*card.Item, 0),
			},
		},
		{
			input: &ItemSave{
				logger:   log.GetLogger(),
				cardId:   1001,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{
					{
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
					{
						CardId:         1001,
						ItemId:         2002,
						Name:           "name 2002",
						Description:    "description 2002",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2002,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemList: []*protos.ItemInformation{
					{
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
					{
						ItemId:         1001,
						Name:           "name 2003",
						Description:    "description 2003",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
				},
				needCreateItemIds: []uint64{},
				needUpdateItemIds: []uint64{},
				needDeleteItemIds: []uint64{},
				originalItemMap: map[uint64]*card.Item{
					2001: {
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
					2002: {
						CardId:         1001,
						ItemId:         2002,
						Name:           "name 2002",
						Description:    "description 2002",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2002,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemMap: map[uint64]*protos.ItemInformation{
					2001: {
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
					2003: {
						ItemId:         1001,
						Name:           "name 2003",
						Description:    "description 2003",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
				},
				finalItemList: make([]*card.Item, 0),
			},
			want: &ItemSave{
				logger:   log.GetLogger(),
				cardId:   1001,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{
					{
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
					{
						CardId:         1001,
						ItemId:         2002,
						Name:           "name 2002",
						Description:    "description 2002",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2002,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemList: []*protos.ItemInformation{
					{
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
					{
						ItemId:         1001,
						Name:           "name 2003",
						Description:    "description 2003",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
				},
				needCreateItemIds: []uint64{2003},
				needUpdateItemIds: []uint64{2001},
				needDeleteItemIds: []uint64{2002},
				originalItemMap: map[uint64]*card.Item{
					2001: {
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
					2002: {
						CardId:         1001,
						ItemId:         2002,
						Name:           "name 2002",
						Description:    "description 2002",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2002,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     2001,
								GoodsId:    10001,
								GoodsCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemMap: map[uint64]*protos.ItemInformation{
					2001: {
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
					2003: {
						ItemId:         1001,
						Name:           "name 2003",
						Description:    "description 2003",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
				},
				finalItemList: make([]*card.Item, 0),
			},
		},
	}

	for _, test := range tests {

		test.input.computeOperationItemIds()
		assert.Equal(t, test.want, test.input, test)
	}
}

func TestItemSave_buildMap(t *testing.T) {

	now := time.Now()
	tests := []struct {
		input *ItemSave
		want  *ItemSave
	}{
		{
			input: &ItemSave{
				logger:   log.GetLogger(),
				cardId:   1001,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{
					{
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemList: []*protos.ItemInformation{
					{
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
					},
				},
				needCreateItemIds: []uint64{},
				needUpdateItemIds: []uint64{},
				needDeleteItemIds: []uint64{},
				originalItemMap:   map[uint64]*card.Item{},
				expectItemMap:     map[uint64]*protos.ItemInformation{},
				finalItemList:     make([]*card.Item, 0),
			},
			want: &ItemSave{
				logger:   log.GetLogger(),
				cardId:   1001,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{
					{
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemList: []*protos.ItemInformation{
					{
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
					},
				},
				needCreateItemIds: []uint64{},
				needUpdateItemIds: []uint64{},
				needDeleteItemIds: []uint64{},
				originalItemMap: map[uint64]*card.Item{
					2001: {
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemMap: map[uint64]*protos.ItemInformation{
					2001: {
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
					},
				},
				finalItemList: make([]*card.Item, 0),
			},
		},
		{
			input: &ItemSave{
				logger:   log.GetLogger(),
				cardId:   1001,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{
					{
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemList:    []*protos.ItemInformation{},
				needCreateItemIds: []uint64{},
				needUpdateItemIds: []uint64{},
				needDeleteItemIds: []uint64{},
				originalItemMap:   map[uint64]*card.Item{},
				expectItemMap:     map[uint64]*protos.ItemInformation{},
				finalItemList:     make([]*card.Item, 0),
			},
			want: &ItemSave{
				logger:   log.GetLogger(),
				cardId:   1001,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{
					{
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemList:    []*protos.ItemInformation{},
				needCreateItemIds: []uint64{},
				needUpdateItemIds: []uint64{},
				needDeleteItemIds: []uint64{},
				originalItemMap: map[uint64]*card.Item{
					2001: {
						CardId:         1001,
						ItemId:         2001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      2001,
								CouponId:    10001,
								CouponCount: 1,
								Time: databases.Time{
									BasicTimeFields: databases.BasicTimeFields{
										CreatedAt: now,
										UpdatedAt: now,
									},
									DeletedAt: nil,
								},
							},
						},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemMap: map[uint64]*protos.ItemInformation{},
				finalItemList: make([]*card.Item, 0),
			},
		},
		{
			input: &ItemSave{
				logger:           log.GetLogger(),
				cardId:           1001,
				database:         database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{},
				expectItemList: []*protos.ItemInformation{
					{
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
					},
				},
				needCreateItemIds: []uint64{},
				needUpdateItemIds: []uint64{},
				needDeleteItemIds: []uint64{},
				originalItemMap:   map[uint64]*card.Item{},
				expectItemMap:     map[uint64]*protos.ItemInformation{},
				finalItemList:     make([]*card.Item, 0),
			},
			want: &ItemSave{
				logger:           log.GetLogger(),
				cardId:           1001,
				database:         database.GetDB(constant.DatabaseConfigKey),
				originalItemList: []*card.Item{},
				expectItemList: []*protos.ItemInformation{
					{
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
					},
				},
				needCreateItemIds: []uint64{},
				needUpdateItemIds: []uint64{},
				needDeleteItemIds: []uint64{},
				originalItemMap:   map[uint64]*card.Item{},
				expectItemMap: map[uint64]*protos.ItemInformation{
					1001: {
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
					},
				},
				finalItemList: make([]*card.Item, 0),
			},
		},
	}

	for _, test := range tests {

		test.input.buildMap()
		assert.Equal(t, test.want, test.input, test)
	}
}

func TestItemSave_updateItem(t *testing.T) {

	now := time.Now()
	tests := []struct {
		cardId   uint64
		initData []*card.Item
		input    *ItemSave
		want     []*card.Item
		err      error
	}{
		{
			cardId: 10,
			initData: []*card.Item{
				{
					CardId:         10,
					ItemId:         1001,
					Name:           "name 2001",
					Description:    "description 2001",
					Price:          100,
					RenewPrice:     100,
					ValidityPeriod: 100,
					Sort:           20,
					ItemCoupons:    []*card.ItemCoupon{},
					ItemGoods:      []*card.ItemGoods{},
					Time: databases.Time{
						BasicTimeFields: databases.BasicTimeFields{
							CreatedAt: now,
							UpdatedAt: now,
						},
						DeletedAt: nil,
					},
				},
			},
			input: &ItemSave{
				logger:            log.GetLogger(),
				cardId:            10,
				database:          database.GetDB(constant.DatabaseConfigKey),
				needUpdateItemIds: []uint64{1001},
				originalItemMap: map[uint64]*card.Item{
					1001: {
						CardId:         10,
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons:    []*card.ItemCoupon{},
						ItemGoods:      []*card.ItemGoods{},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemMap: map[uint64]*protos.ItemInformation{
					1001: {
						ItemId:         1001,
						Name:           "name 2001",
						Description:    "description 2001",
						Price:          90,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           25,
						Coupons:        []*protos.ItemCouponInformation{},
					},
				},
				finalItemList: []*card.Item{},
			},
			want: []*card.Item{
				{
					CardId:         10,
					ItemId:         1001,
					Name:           "name 2001",
					Description:    "description 2001",
					Price:          90,
					RenewPrice:     100,
					ValidityPeriod: 100,
					Sort:           25,
					ItemCoupons:    []*card.ItemCoupon{},
					ItemGoods:      []*card.ItemGoods{},
					Time: databases.Time{
						BasicTimeFields: databases.BasicTimeFields{
							CreatedAt: now,
							UpdatedAt: now,
						},
						DeletedAt: nil,
					},
				},
			},
		},
	}

	for _, test := range tests {

		for _, Item := range test.initData {

			err := database.GetDB(constant.DatabaseConfigKey).Create(Item).Error
			assert.Nil(t, err, test)
			if err != nil {
				break
			}
		}

		err := test.input.updateItem()
		assert.Equal(t, test.err, err, test)
		var result []*card.Item
		err = database.GetDB(constant.DatabaseConfigKey).Preload("ItemCoupons").Preload("ItemGoods").Where(card.Item{CardId: test.cardId}).Find(&result).Error
		assert.Nil(t, err, test)
		if err != nil {
			continue
		}

		itemMap := map[uint64]*card.Item{}
		for _, item := range result {

			itemMap[item.CardId] = item
		}

		for _, item := range test.want {

			item.Time = itemMap[item.CardId].Time
		}

		assert.Equal(t, test.want, result, test)
	}
}

func TestItemSave_createItem(t *testing.T) {

	now := time.Now()
	tests := []struct {
		cardId   uint64
		initData []*card.Item
		input    *ItemSave
		want     []*card.Item
		err      error
	}{
		{
			cardId:   11,
			initData: []*card.Item{},
			input: &ItemSave{
				logger:            log.GetLogger(),
				cardId:            11,
				database:          database.GetDB(constant.DatabaseConfigKey),
				needCreateItemIds: []uint64{1002},
				originalItemMap:   map[uint64]*card.Item{},
				expectItemMap: map[uint64]*protos.ItemInformation{
					1002: {
						ItemId:         1002,
						Name:           "name 1002",
						Description:    "description 1002",
						Price:          90,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           25,
						Coupons:        []*protos.ItemCouponInformation{},
						Goods:          []*protos.ItemGoodsInformation{},
					},
				},
				finalItemList: []*card.Item{},
			},
			want: []*card.Item{
				{
					CardId:         11,
					ItemId:         1002,
					Name:           "name 1002",
					Description:    "description 1002",
					Price:          90,
					RenewPrice:     100,
					ValidityPeriod: 100,
					Sort:           25,
					ItemCoupons:    []*card.ItemCoupon{},
					ItemGoods:      []*card.ItemGoods{},
					Time: databases.Time{
						BasicTimeFields: databases.BasicTimeFields{
							CreatedAt: now,
							UpdatedAt: now,
						},
						DeletedAt: nil,
					},
				},
			},
		},
	}

	for _, test := range tests {

		for _, Item := range test.initData {

			err := database.GetDB(constant.DatabaseConfigKey).Create(Item).Error
			assert.Nil(t, err, test)
			if err != nil {
				break
			}
		}

		err := test.input.createItem()
		assert.Equal(t, test.err, err, test)
		var result []*card.Item
		err = database.GetDB(constant.DatabaseConfigKey).Preload("ItemCoupons").Preload("ItemGoods").Where(card.Item{CardId: test.cardId}).Find(&result).Error
		assert.Nil(t, err, test)
		if err != nil {
			continue
		}

		itemMap := map[uint64]*card.Item{}
		for _, Item := range result {

			itemMap[Item.CardId] = Item
		}

		for _, item := range test.want {

			item.Time = itemMap[item.CardId].Time
		}

		assert.Equal(t, test.want, result, test)
	}
}

func TestItemSave_deleteItem(t *testing.T) {

	now := time.Now()
	tests := []struct {
		cardId   uint64
		initData []*card.Item
		input    *ItemSave
		want     []*card.Item
		err      error
	}{
		{
			cardId: 12,
			initData: []*card.Item{
				{
					CardId:         12,
					ItemId:         1003,
					Name:           "name 1003",
					Description:    "description 1003",
					Price:          100,
					RenewPrice:     100,
					ValidityPeriod: 100,
					Sort:           20,
					ItemCoupons:    []*card.ItemCoupon{},
					ItemGoods:      []*card.ItemGoods{},
					Time: databases.Time{
						BasicTimeFields: databases.BasicTimeFields{
							CreatedAt: now,
							UpdatedAt: now,
						},
						DeletedAt: nil,
					},
				},
			},
			input: &ItemSave{
				logger:            log.GetLogger(),
				cardId:            12,
				database:          database.GetDB(constant.DatabaseConfigKey),
				needDeleteItemIds: []uint64{1003},
				originalItemMap: map[uint64]*card.Item{
					1003: {
						CardId:         12,
						ItemId:         1003,
						Name:           "name 1003",
						Description:    "description 1003",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons:    []*card.ItemCoupon{},
						ItemGoods:      []*card.ItemGoods{},
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectItemMap: map[uint64]*protos.ItemInformation{},
				finalItemList: []*card.Item{},
			},
			want: []*card.Item{},
		},
	}

	for _, test := range tests {

		for _, Item := range test.initData {

			err := database.GetDB(constant.DatabaseConfigKey).Create(Item).Error
			assert.Nil(t, err, test)
			if err != nil {
				break
			}
		}

		err := test.input.deleteItem()
		assert.Equal(t, test.err, err, test)
		var result []*card.Item
		err = database.GetDB(constant.DatabaseConfigKey).Preload("ItemCoupons").Where(card.Item{CardId: test.cardId}).Find(&result).Error
		assert.Nil(t, err, test)
		if err != nil {
			continue
		}

		itemMap := map[uint64]*card.Item{}
		for _, item := range result {

			itemMap[item.CardId] = item
		}

		for _, item := range test.want {

			item.Time = itemMap[item.CardId].Time
		}

		assert.Equal(t, test.want, result, test)
	}
}

func TestItemSave_Save(t *testing.T) {

	tests := []struct {
		CardId   uint64
		initData []*card.Item
		input    *ItemSave
		want     []*card.Item
		err      error
	}{
		{
			CardId: 13,
			initData: []*card.Item{
				{
					CardId:         13,
					ItemId:         1004,
					Name:           "name 1004",
					Description:    "description 1004",
					Price:          100,
					RenewPrice:     100,
					ValidityPeriod: 100,
					Sort:           20,
					ItemCoupons: []*card.ItemCoupon{
						{
							ItemId:      1004,
							CouponId:    10001,
							CouponCount: 1,
						},
					},
					ItemGoods: []*card.ItemGoods{
						{
							ItemId:     1004,
							GoodsId:    10001,
							GoodsCount: 1,
						},
					},
				},
				{
					CardId:         13,
					ItemId:         1005,
					Name:           "name 1005",
					Description:    "description 1005",
					Price:          100,
					RenewPrice:     100,
					ValidityPeriod: 100,
					Sort:           20,
					ItemCoupons: []*card.ItemCoupon{
						{
							ItemId:      1005,
							CouponId:    10001,
							CouponCount: 1,
						},
						{
							ItemId:      1005,
							CouponId:    10002,
							CouponCount: 1,
						},
					},
					ItemGoods: []*card.ItemGoods{
						{
							ItemId:     1005,
							GoodsId:    10001,
							GoodsCount: 1,
						},
						{
							ItemId:     1005,
							GoodsId:    10002,
							GoodsCount: 1,
						},
					},
				},
			},
			input: NewItemSave(
				ItemSaveCardId(13),
				ItemSaveExpectItemList([]*protos.ItemInformation{
					{
						ItemId:         1004,
						Name:           "name 1004",
						Description:    "description 1004",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
					{
						Name:           "name 1006",
						Description:    "description 1006",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						Coupons: []*protos.ItemCouponInformation{
							{
								CouponId:    10001,
								CouponCount: 1,
							},
							{
								CouponId:    10003,
								CouponCount: 1,
							},
						},
						Goods: []*protos.ItemGoodsInformation{
							{
								GoodsId:    10001,
								GoodsCount: 1,
							},
							{
								GoodsId:    10003,
								GoodsCount: 1,
							},
						},
					},
				}),
				ItemSaveOriginalItemList([]*card.Item{
					{
						CardId:         13,
						ItemId:         1004,
						Name:           "name 1004",
						Description:    "description 1004",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      1004,
								CouponId:    10001,
								CouponCount: 1,
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     1004,
								GoodsId:    10001,
								GoodsCount: 1,
							},
						},
					},
					{
						CardId:         13,
						ItemId:         1005,
						Name:           "name 1005",
						Description:    "description 1005",
						Price:          100,
						RenewPrice:     100,
						ValidityPeriod: 100,
						Sort:           20,
						ItemCoupons: []*card.ItemCoupon{
							{
								ItemId:      1005,
								CouponId:    10001,
								CouponCount: 1,
							},
							{
								ItemId:      1005,
								CouponId:    10002,
								CouponCount: 1,
							},
						},
						ItemGoods: []*card.ItemGoods{
							{
								ItemId:     1005,
								GoodsId:    10001,
								GoodsCount: 1,
							},
							{
								ItemId:     1005,
								GoodsId:    10002,
								GoodsCount: 1,
							},
						},
					},
				}),
			),
			want: []*card.Item{
				{
					CardId:         13,
					ItemId:         1004,
					Name:           "name 1004",
					Description:    "description 1004",
					Price:          100,
					RenewPrice:     100,
					ValidityPeriod: 100,
					Sort:           20,
					ItemCoupons: []*card.ItemCoupon{
						{
							ItemId:      1004,
							CouponId:    10001,
							CouponCount: 1,
						},
					},
					ItemGoods: []*card.ItemGoods{
						{
							ItemId:     1004,
							GoodsId:    10001,
							GoodsCount: 1,
						},
					},
				},
				{
					CardId:         13,
					ItemId:         1006,
					Name:           "name 1006",
					Description:    "description 1006",
					Price:          100,
					RenewPrice:     100,
					ValidityPeriod: 100,
					Sort:           20,
					ItemCoupons: []*card.ItemCoupon{
						{
							ItemId:      1006,
							CouponId:    10001,
							CouponCount: 1,
						},
						{
							ItemId:      1006,
							CouponId:    10003,
							CouponCount: 1,
						},
					},
					ItemGoods: []*card.ItemGoods{
						{
							ItemId:     1006,
							GoodsId:    10001,
							GoodsCount: 1,
						},
						{
							ItemId:     1006,
							GoodsId:    10003,
							GoodsCount: 1,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {

		for _, Item := range test.initData {

			err := database.GetDB(constant.DatabaseConfigKey).Create(Item).Error
			assert.Nil(t, err, test)
			if err != nil {
				break
			}
		}

		var result []*card.Item
		result, err := test.input.Save()
		assert.Equal(t, test.err, err, test)

		itemMap := map[uint64]*card.Item{}
		for _, item := range result {

			itemMap[item.ItemId] = item

			for _, coupon := range item.ItemCoupons {
				coupon.Time = item.Time
			}

			for _, good := range item.ItemGoods {
				good.Time = item.Time
			}

			sort.SliceStable(item.ItemCoupons, func(i, j int) bool {
				return item.ItemCoupons[i].CouponId < item.ItemCoupons[j].CouponId
			})

			sort.SliceStable(item.ItemGoods, func(i, j int) bool {
				return item.ItemGoods[i].GoodsId < item.ItemGoods[j].GoodsId
			})
		}

		newIdIndex := 0

		for _, item := range test.want {

			var result *card.Item
			var exist bool
			if result, exist = itemMap[item.ItemId]; !exist {
				result = itemMap[test.input.needCreateItemIds[newIdIndex]]
				item.ItemId = test.input.needCreateItemIds[newIdIndex]
				newIdIndex++
			}
			item.Time = result.Time

			for _, coupon := range item.ItemCoupons {

				coupon.ItemId = result.ItemId
				coupon.Time = result.Time
			}

			for _, good := range item.ItemGoods {

				good.ItemId = result.ItemId
				good.Time = result.Time
			}
		}

		assert.Equal(t, test.want, result, test)

		err = database.GetDB(constant.DatabaseConfigKey).Preload("ItemCoupons").Preload("ItemGoods").Where(card.Item{CardId: test.CardId}).Find(&result).Error
		assert.Nil(t, err, test)
		if err != nil {
			continue
		}

		for _, item := range result {

			itemMap[item.ItemId] = item

			for _, coupon := range item.ItemCoupons {

				coupon.Time = item.Time
			}

			for _, good := range item.ItemGoods {

				good.Time = item.Time
			}

			sort.SliceStable(item.ItemCoupons, func(i, j int) bool {
				return item.ItemCoupons[i].CouponId < item.ItemCoupons[j].CouponId
			})

			sort.SliceStable(item.ItemGoods, func(i, j int) bool {
				return item.ItemGoods[i].GoodsId < item.ItemGoods[j].GoodsId
			})
		}

		newIdIndex = 0

		for _, item := range test.want {

			var result *card.Item
			var exist bool
			if result, exist = itemMap[item.ItemId]; !exist {
				result = itemMap[test.input.needCreateItemIds[newIdIndex]]
				item.ItemId = test.input.needCreateItemIds[newIdIndex]
				newIdIndex++
			}
			item.Time = result.Time

			for _, coupon := range item.ItemCoupons {

				coupon.ItemId = result.ItemId
				coupon.Time = result.Time
			}

			for _, good := range item.ItemGoods {

				good.ItemId = result.ItemId
				good.Time = result.Time
			}
		}

		assert.Equal(t, test.want, result, test)

	}
}
