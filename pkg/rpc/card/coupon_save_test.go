package card

import (
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

func TestCouponSaveItemId(t *testing.T) {

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

		couponSave := &CouponSave{}
		CouponSaveItemId(test.input)(couponSave)
		assert.Equal(t, test.want, couponSave.itemId, test)
	}
}

func TestCouponSaveDatabase(t *testing.T) {

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

		couponSave := &CouponSave{}
		CouponSaveDatabase(test.input)(couponSave)
		assert.Equal(t, test.want, couponSave.database, test)
	}
}

func TestCouponSaveLogger(t *testing.T) {

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

		couponSave := &CouponSave{}
		CouponSaveLogger(test.input)(couponSave)
		assert.Equal(t, test.want, couponSave.logger, test)
	}
}

func TestCouponSaveExpectCouponList(t *testing.T) {

	tests := []struct {
		input []*protos.ItemCouponInformation
		want  []*protos.ItemCouponInformation
	}{
		{
			input: []*protos.ItemCouponInformation{},
			want:  []*protos.ItemCouponInformation{},
		},
		{
			input: []*protos.ItemCouponInformation{
				{
					CouponId:    10,
					CouponCount: 100,
				},
			},
			want: []*protos.ItemCouponInformation{
				{
					CouponId:    10,
					CouponCount: 100,
				},
			},
		},
	}

	for _, test := range tests {

		couponSave := &CouponSave{}
		CouponSaveExpectCouponList(test.input)(couponSave)
		assert.Equal(t, test.want, couponSave.expectCouponList, test)
	}
}

func TestCouponSaveOriginalCouponList(t *testing.T) {

	now := time.Now()
	tests := []struct {
		input []*card.ItemCoupon
		want  []*card.ItemCoupon
	}{
		{
			input: []*card.ItemCoupon{},
			want:  []*card.ItemCoupon{},
		},
		{
			input: []*card.ItemCoupon{
				{
					ItemId:      1,
					CouponId:    10,
					CouponCount: 100,
					Time: databases.Time{
						BasicTimeFields: databases.BasicTimeFields{
							CreatedAt: now,
							UpdatedAt: now,
						},
						DeletedAt: nil,
					},
				},
			},
			want: []*card.ItemCoupon{
				{
					ItemId:      1,
					CouponId:    10,
					CouponCount: 100,
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

		couponSave := &CouponSave{}
		CouponSaveOriginalCouponList(test.input)(couponSave)
		assert.Equal(t, test.want, couponSave.originalCouponList, test)
	}
}

func TestNewCouponSave(t *testing.T) {

	tests := []struct {
		input []CouponSaveOption
		want  *CouponSave
	}{
		{
			input: []CouponSaveOption{
				CouponSaveItemId(10),
				CouponSaveLogger(log.GetLogger()),
				CouponSaveDatabase(database.GetDB(constant.DatabaseConfigKey)),
			},
			want: &CouponSave{
				logger:             log.GetLogger(),
				itemId:             10,
				database:           database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: make([]*card.ItemCoupon, 0),
				expectCouponList:   make([]*protos.ItemCouponInformation, 0),
			},
		},
	}

	for _, test := range tests {

		assert.Equal(t, test.want, NewCouponSave(test.input...), test)
	}
}

func TestCouponSave_init(t *testing.T) {

	tests := []struct {
		want *CouponSave
	}{
		{
			want: &CouponSave{
				logger:              log.GetLogger(),
				itemId:              0,
				database:            database.GetDB(constant.DatabaseConfigKey),
				originalCouponList:  make([]*card.ItemCoupon, 0),
				expectCouponList:    make([]*protos.ItemCouponInformation, 0),
				needCreateCouponIds: make([]uint64, 0),
				needUpdateCouponIds: make([]uint64, 0),
				needDeleteCouponIds: make([]uint64, 0),
				expectCouponMap:     make(map[uint64]*protos.ItemCouponInformation),
				originalCouponMap:   make(map[uint64]*card.ItemCoupon),
				finalCouponList:     make([]*card.ItemCoupon, 0),
			},
		},
	}

	for _, test := range tests {

		couponSave := NewCouponSave()
		couponSave.init()
		assert.Equal(t, test.want, couponSave, test)
	}
}

func TestCouponSave_computeOperationCouponIds(t *testing.T) {

	now := time.Now()
	tests := []struct {
		input *CouponSave
		want  *CouponSave
	}{
		{
			input: &CouponSave{
				logger:   log.GetLogger(),
				itemId:   0,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{
					{
						ItemId:      1001,
						CouponId:    2001,
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
				expectCouponList: []*protos.ItemCouponInformation{
					{
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				needCreateCouponIds: make([]uint64, 0),
				needUpdateCouponIds: []uint64{},
				needDeleteCouponIds: make([]uint64, 0),
				expectCouponMap: map[uint64]*protos.ItemCouponInformation{
					2001: {
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				originalCouponMap: map[uint64]*card.ItemCoupon{
					2001: {
						ItemId:      1001,
						CouponId:    2001,
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
				finalCouponList: make([]*card.ItemCoupon, 0),
			},
			want: &CouponSave{
				logger:   log.GetLogger(),
				itemId:   0,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{
					{
						ItemId:      1001,
						CouponId:    2001,
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
				expectCouponList: []*protos.ItemCouponInformation{
					{
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				needCreateCouponIds: make([]uint64, 0),
				needUpdateCouponIds: []uint64{2001},
				needDeleteCouponIds: make([]uint64, 0),
				expectCouponMap: map[uint64]*protos.ItemCouponInformation{
					2001: {
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				originalCouponMap: map[uint64]*card.ItemCoupon{
					2001: {
						ItemId:      1001,
						CouponId:    2001,
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
				finalCouponList: make([]*card.ItemCoupon, 0),
			},
		},
		{
			input: &CouponSave{
				logger:   log.GetLogger(),
				itemId:   0,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{
					{
						ItemId:      1001,
						CouponId:    2001,
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
				expectCouponList:    []*protos.ItemCouponInformation{},
				needCreateCouponIds: make([]uint64, 0),
				needUpdateCouponIds: []uint64{},
				needDeleteCouponIds: make([]uint64, 0),
				expectCouponMap:     map[uint64]*protos.ItemCouponInformation{},
				originalCouponMap: map[uint64]*card.ItemCoupon{
					2001: {
						ItemId:      1001,
						CouponId:    2001,
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
				finalCouponList: make([]*card.ItemCoupon, 0),
			},
			want: &CouponSave{
				logger:   log.GetLogger(),
				itemId:   0,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{
					{
						ItemId:      1001,
						CouponId:    2001,
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
				expectCouponList:    []*protos.ItemCouponInformation{},
				needCreateCouponIds: []uint64{},
				needUpdateCouponIds: []uint64{},
				needDeleteCouponIds: []uint64{2001},
				expectCouponMap:     map[uint64]*protos.ItemCouponInformation{},
				originalCouponMap: map[uint64]*card.ItemCoupon{
					2001: {
						ItemId:      1001,
						CouponId:    2001,
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
				finalCouponList: make([]*card.ItemCoupon, 0),
			},
		},
		{
			input: &CouponSave{
				logger:             log.GetLogger(),
				itemId:             0,
				database:           database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{},
				expectCouponList: []*protos.ItemCouponInformation{
					{
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				needCreateCouponIds: make([]uint64, 0),
				needUpdateCouponIds: []uint64{},
				needDeleteCouponIds: make([]uint64, 0),
				expectCouponMap: map[uint64]*protos.ItemCouponInformation{
					2001: {
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				originalCouponMap: map[uint64]*card.ItemCoupon{},
				finalCouponList:   make([]*card.ItemCoupon, 0),
			},
			want: &CouponSave{
				logger:             log.GetLogger(),
				itemId:             0,
				database:           database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{},
				expectCouponList: []*protos.ItemCouponInformation{
					{
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				needCreateCouponIds: []uint64{2001},
				needUpdateCouponIds: []uint64{},
				needDeleteCouponIds: make([]uint64, 0),
				expectCouponMap: map[uint64]*protos.ItemCouponInformation{
					2001: {
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				originalCouponMap: map[uint64]*card.ItemCoupon{},
				finalCouponList:   make([]*card.ItemCoupon, 0),
			},
		},
		{
			input: &CouponSave{
				logger:   log.GetLogger(),
				itemId:   0,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{
					{
						ItemId:      1001,
						CouponId:    2001,
						CouponCount: 1,
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
					{
						ItemId:      1001,
						CouponId:    2002,
						CouponCount: 2,
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectCouponList: []*protos.ItemCouponInformation{
					{
						CouponId:    2001,
						CouponCount: 1,
					},
					{
						CouponId:    2003,
						CouponCount: 3,
					},
				},
				needCreateCouponIds: []uint64{},
				needUpdateCouponIds: []uint64{},
				needDeleteCouponIds: []uint64{},
				originalCouponMap: map[uint64]*card.ItemCoupon{
					2001: {
						ItemId:      1001,
						CouponId:    2001,
						CouponCount: 1,
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
					2002: {
						ItemId:      1001,
						CouponId:    2002,
						CouponCount: 2,
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectCouponMap: map[uint64]*protos.ItemCouponInformation{
					2001: {
						CouponId:    2001,
						CouponCount: 1,
					},
					2003: {
						CouponId:    2003,
						CouponCount: 3,
					},
				},
				finalCouponList: make([]*card.ItemCoupon, 0),
			},
			want: &CouponSave{
				logger:   log.GetLogger(),
				itemId:   0,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{
					{
						ItemId:      1001,
						CouponId:    2001,
						CouponCount: 1,
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
					{
						ItemId:      1001,
						CouponId:    2002,
						CouponCount: 2,
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectCouponList: []*protos.ItemCouponInformation{
					{
						CouponId:    2001,
						CouponCount: 1,
					},
					{
						CouponId:    2003,
						CouponCount: 3,
					},
				},
				needCreateCouponIds: []uint64{2003},
				needUpdateCouponIds: []uint64{2001},
				needDeleteCouponIds: []uint64{2002},
				originalCouponMap: map[uint64]*card.ItemCoupon{
					2001: {
						ItemId:      1001,
						CouponId:    2001,
						CouponCount: 1,
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
					2002: {
						ItemId:      1001,
						CouponId:    2002,
						CouponCount: 2,
						Time: databases.Time{
							BasicTimeFields: databases.BasicTimeFields{
								CreatedAt: now,
								UpdatedAt: now,
							},
							DeletedAt: nil,
						},
					},
				},
				expectCouponMap: map[uint64]*protos.ItemCouponInformation{
					2001: {
						CouponId:    2001,
						CouponCount: 1,
					},
					2003: {
						CouponId:    2003,
						CouponCount: 3,
					},
				},
				finalCouponList: make([]*card.ItemCoupon, 0),
			},
		},
	}

	for _, test := range tests {

		test.input.computeOperationCouponIds()
		assert.Equal(t, test.want, test.input, test)
	}
}

func TestCouponSave_buildMap(t *testing.T) {

	now := time.Now()
	tests := []struct {
		input *CouponSave
		want  *CouponSave
	}{
		{
			input: &CouponSave{
				logger:   log.GetLogger(),
				itemId:   0,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{
					{
						ItemId:      1001,
						CouponId:    2001,
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
				expectCouponList: []*protos.ItemCouponInformation{
					{
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				needCreateCouponIds: []uint64{},
				needUpdateCouponIds: []uint64{},
				needDeleteCouponIds: []uint64{},
				expectCouponMap:     map[uint64]*protos.ItemCouponInformation{},
				originalCouponMap:   map[uint64]*card.ItemCoupon{},
				finalCouponList:     make([]*card.ItemCoupon, 0),
			},
			want: &CouponSave{
				logger:   log.GetLogger(),
				itemId:   0,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{
					{
						ItemId:      1001,
						CouponId:    2001,
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
				expectCouponList: []*protos.ItemCouponInformation{
					{
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				needCreateCouponIds: []uint64{},
				needUpdateCouponIds: []uint64{},
				needDeleteCouponIds: []uint64{},
				expectCouponMap: map[uint64]*protos.ItemCouponInformation{
					2001: {
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				originalCouponMap: map[uint64]*card.ItemCoupon{
					2001: {
						ItemId:      1001,
						CouponId:    2001,
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
				finalCouponList: make([]*card.ItemCoupon, 0),
			},
		},
		{
			input: &CouponSave{
				logger:   log.GetLogger(),
				itemId:   0,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{
					{
						ItemId:      1001,
						CouponId:    2001,
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
				expectCouponList:    []*protos.ItemCouponInformation{},
				needCreateCouponIds: []uint64{},
				needUpdateCouponIds: []uint64{},
				needDeleteCouponIds: []uint64{},
				expectCouponMap:     map[uint64]*protos.ItemCouponInformation{},
				originalCouponMap:   map[uint64]*card.ItemCoupon{},
				finalCouponList:     make([]*card.ItemCoupon, 0),
			},
			want: &CouponSave{
				logger:   log.GetLogger(),
				itemId:   0,
				database: database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{
					{
						ItemId:      1001,
						CouponId:    2001,
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
				expectCouponList:    []*protos.ItemCouponInformation{},
				needCreateCouponIds: []uint64{},
				needUpdateCouponIds: []uint64{},
				needDeleteCouponIds: []uint64{},
				originalCouponMap: map[uint64]*card.ItemCoupon{
					2001: {
						ItemId:      1001,
						CouponId:    2001,
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
				expectCouponMap: map[uint64]*protos.ItemCouponInformation{},
				finalCouponList: make([]*card.ItemCoupon, 0),
			},
		},
		{
			input: &CouponSave{
				logger:             log.GetLogger(),
				itemId:             0,
				database:           database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{},
				expectCouponList: []*protos.ItemCouponInformation{
					{
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				needCreateCouponIds: []uint64{},
				needUpdateCouponIds: []uint64{},
				needDeleteCouponIds: []uint64{},
				expectCouponMap:     map[uint64]*protos.ItemCouponInformation{},
				originalCouponMap:   map[uint64]*card.ItemCoupon{},
				finalCouponList:     make([]*card.ItemCoupon, 0),
			},
			want: &CouponSave{
				logger:             log.GetLogger(),
				itemId:             0,
				database:           database.GetDB(constant.DatabaseConfigKey),
				originalCouponList: []*card.ItemCoupon{},
				expectCouponList: []*protos.ItemCouponInformation{
					{
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				needCreateCouponIds: []uint64{},
				needUpdateCouponIds: []uint64{},
				needDeleteCouponIds: []uint64{},
				expectCouponMap: map[uint64]*protos.ItemCouponInformation{
					2001: {
						CouponId:    2001,
						CouponCount: 1,
					},
				},
				originalCouponMap: map[uint64]*card.ItemCoupon{},
				finalCouponList:   make([]*card.ItemCoupon, 0),
			},
		},
	}

	for _, test := range tests {

		test.input.buildMap()
		assert.Equal(t, test.want, test.input, test)
	}
}

func TestCouponSave_updateCoupon(t *testing.T) {

	tests := []struct {
		itemId   uint64
		initData []*card.ItemCoupon
		input    *CouponSave
		want     []*card.ItemCoupon
		err      error
	}{
		{
			itemId: 10,
			initData: []*card.ItemCoupon{
				{
					ItemId:      10,
					CouponId:    10001,
					CouponCount: 1,
				},
			},
			input: &CouponSave{
				logger:              log.GetLogger(),
				itemId:              10,
				database:            database.GetDB(constant.DatabaseConfigKey),
				needUpdateCouponIds: []uint64{10001},
				expectCouponMap: map[uint64]*protos.ItemCouponInformation{
					10001: {
						CouponId:    10001,
						CouponCount: 2,
					},
				},
				originalCouponMap: map[uint64]*card.ItemCoupon{
					10001: {
						ItemId:      10,
						CouponId:    10001,
						CouponCount: 1,
					},
				},
				finalCouponList: []*card.ItemCoupon{},
			},
			want: []*card.ItemCoupon{
				{
					ItemId:      10,
					CouponId:    10001,
					CouponCount: 2,
				},
			},
		},
	}

	for _, test := range tests {

		for _, coupon := range test.initData {

			err := database.GetDB(constant.DatabaseConfigKey).Create(coupon).Error
			assert.Nil(t, err, test)
			if err != nil {
				break
			}
		}

		err := test.input.updateCoupon()
		assert.Equal(t, test.err, err, test)
		var result []*card.ItemCoupon
		err = database.GetDB(constant.DatabaseConfigKey).Where(card.ItemCoupon{ItemId: test.itemId}).Find(&result).Error
		assert.Nil(t, err, test)
		if err != nil {
			continue
		}

		couponMap := map[uint64]*card.ItemCoupon{}
		for _, coupon := range result {

			couponMap[coupon.CouponId] = coupon
		}

		for _, coupon := range test.want {

			coupon.Time = couponMap[coupon.CouponId].Time
		}

		assert.Equal(t, test.want, result, test)
	}
}

func TestCouponSave_createCoupon(t *testing.T) {

	tests := []struct {
		itemId   uint64
		initData []*card.ItemCoupon
		input    *CouponSave
		want     []*card.ItemCoupon
		err      error
	}{
		{
			itemId:   11,
			initData: []*card.ItemCoupon{},
			input: &CouponSave{
				logger:              log.GetLogger(),
				itemId:              11,
				database:            database.GetDB(constant.DatabaseConfigKey),
				needCreateCouponIds: []uint64{10001},
				expectCouponMap: map[uint64]*protos.ItemCouponInformation{
					10001: {
						CouponId:    10001,
						CouponCount: 1,
					},
				},
				originalCouponMap: map[uint64]*card.ItemCoupon{},
				finalCouponList:   []*card.ItemCoupon{},
			},
			want: []*card.ItemCoupon{
				{
					ItemId:      11,
					CouponId:    10001,
					CouponCount: 1,
				},
			},
		},
	}

	for _, test := range tests {

		for _, coupon := range test.initData {

			err := database.GetDB(constant.DatabaseConfigKey).Create(coupon).Error
			assert.Nil(t, err, test)
			if err != nil {
				break
			}
		}

		err := test.input.createCoupon()
		assert.Equal(t, test.err, err, test)
		var result []*card.ItemCoupon
		err = database.GetDB(constant.DatabaseConfigKey).Where(card.ItemCoupon{ItemId: test.itemId}).Find(&result).Error
		assert.Nil(t, err, test)
		if err != nil {
			continue
		}

		couponMap := map[uint64]*card.ItemCoupon{}
		for _, coupon := range result {

			couponMap[coupon.CouponId] = coupon
		}

		for _, coupon := range test.want {

			coupon.Time = couponMap[coupon.CouponId].Time
		}

		assert.Equal(t, test.want, result, test)
	}
}

func TestCouponSave_deleteCoupon(t *testing.T) {

	tests := []struct {
		itemId   uint64
		initData []*card.ItemCoupon
		input    *CouponSave
		want     []*card.ItemCoupon
		err      error
	}{
		{
			itemId: 12,
			initData: []*card.ItemCoupon{
				{
					ItemId:      12,
					CouponId:    10001,
					CouponCount: 1,
				},
			},
			input: &CouponSave{
				logger:              log.GetLogger(),
				itemId:              12,
				database:            database.GetDB(constant.DatabaseConfigKey),
				needDeleteCouponIds: []uint64{10001},
				expectCouponMap:     map[uint64]*protos.ItemCouponInformation{},
				originalCouponMap: map[uint64]*card.ItemCoupon{
					10001: {
						ItemId:      12,
						CouponId:    10001,
						CouponCount: 1,
					},
				},
				finalCouponList: []*card.ItemCoupon{},
			},
			want: []*card.ItemCoupon{},
		},
	}

	for _, test := range tests {

		for _, coupon := range test.initData {

			err := database.GetDB(constant.DatabaseConfigKey).Create(coupon).Error
			assert.Nil(t, err, test)
			if err != nil {
				break
			}
		}

		err := test.input.deleteCoupon()
		assert.Equal(t, test.err, err, test)
		var result []*card.ItemCoupon
		err = database.GetDB(constant.DatabaseConfigKey).Where(card.ItemCoupon{ItemId: test.itemId}).Find(&result).Error
		assert.Nil(t, err, test)
		if err != nil {
			continue
		}

		couponMap := map[uint64]*card.ItemCoupon{}
		for _, coupon := range result {

			couponMap[coupon.CouponId] = coupon
		}

		for _, coupon := range test.want {

			coupon.Time = couponMap[coupon.CouponId].Time
		}

		assert.Equal(t, test.want, result, test)
	}
}

func TestCouponSave_Save(t *testing.T) {

	tests := []struct {
		itemId   uint64
		initData []*card.ItemCoupon
		input    *CouponSave
		want     []*card.ItemCoupon
		err      error
	}{
		{
			itemId: 13,
			initData: []*card.ItemCoupon{
				{
					ItemId:      13,
					CouponId:    10001,
					CouponCount: 1,
				},
				{
					ItemId:      13,
					CouponId:    10002,
					CouponCount: 1,
				},
			},
			input: NewCouponSave(
				CouponSaveItemId(13),
				CouponSaveExpectCouponList([]*protos.ItemCouponInformation{
					{
						CouponId:    10001,
						CouponCount: 2,
					},
					{
						CouponId:    10003,
						CouponCount: 1,
					},
				}),
				CouponSaveOriginalCouponList([]*card.ItemCoupon{
					{
						ItemId:      13,
						CouponId:    10001,
						CouponCount: 1,
					},
					{
						ItemId:      13,
						CouponId:    10002,
						CouponCount: 1,
					},
				}),
			),
			want: []*card.ItemCoupon{
				{
					ItemId:      13,
					CouponId:    10001,
					CouponCount: 2,
				},
				{
					ItemId:      13,
					CouponId:    10003,
					CouponCount: 1,
				},
			},
		},
	}

	for _, test := range tests {

		for _, coupon := range test.initData {

			err := database.GetDB(constant.DatabaseConfigKey).Create(coupon).Error
			assert.Nil(t, err, test)
			if err != nil {
				break
			}
		}

		var result []*card.ItemCoupon
		result, err := test.input.Save()
		assert.Equal(t, test.err, err, test)

		couponMap := map[uint64]*card.ItemCoupon{}
		for _, coupon := range result {

			couponMap[coupon.CouponId] = coupon
		}

		for _, coupon := range test.want {

			coupon.Time = couponMap[coupon.CouponId].Time
		}

		assert.Equal(t, test.want, result, test)

		err = database.GetDB(constant.DatabaseConfigKey).Where(card.ItemCoupon{ItemId: test.itemId}).Find(&result).Error
		assert.Nil(t, err, test)
		if err != nil {
			continue
		}

		for _, coupon := range result {

			couponMap[coupon.CouponId] = coupon
		}

		for _, coupon := range test.want {

			coupon.Time = couponMap[coupon.CouponId].Time
		}

		assert.Equal(t, test.want, result, test)
	}
}
