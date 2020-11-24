package card

import (
	goodsProtos "dev-gitlab.wanxingrowth.com/fanli/goods/v2/pkg/rpc/protos"
	rpclog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/client"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/model/card"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/utils/log"
)

func (_ Controller) GetCard(ctx context.Context, req *protos.GetCardRequest) (*protos.GetCardReply, error) {

	logger := rpclog.WithRequestId(ctx, log.GetLogger())

	err := validateGetCard(req)
	if err != nil {
		return &protos.GetCardReply{
			Err: ConvertErrorToProtobuf(err),
		}, nil
	}

	var cardData *card.Card

	cardData, err = getCard(req.GetCardId())
	if err != nil {
		logger.WithError(err).Error("get card error")
		return &protos.GetCardReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeGetCardFailed,
				Message: constant.ErrorMessageGetCardFailed,
			},
		}, nil
	}

	spuApi := client.GetSPUService()
	if spuApi == nil {
		logger.Error("get spu rpc service client is nil")

		return &protos.GetCardReply{
			Err: &protos.Error{
				Code:    constant.ErrorCodeGetCardFailed,
				Message: constant.ErrorMessageGetCardFailed,
			},
		}, nil
	}

	// 获取所有商品数据
	returnCardData := convertToCardInformation(cardData)
	for _, items := range returnCardData.Items {
		for _, goods := range items.Goods {

			spuReply, err := spuApi.GetSku(context.Background(), &goodsProtos.GetSkuRequest{
				SkuId: goods.GoodsId,
			})

			if err != nil {

				logger.WithError(err).Error("call spu service error")
				return &protos.GetCardReply{
					Err: &protos.Error{
						Code:    constant.ErrorCodeGetCardFailed,
						Message: constant.ErrorMessageGetCardFailed,
					},
				}, nil
			}
			if spuReply.Err != nil {
				logger.WithError(err).Error("spu reply error")
				goods.GoodsName = ""
				goods.SPUID = 0
				goods.Images = []string{}
			} else {
				goods.GoodsName = spuReply.Sku.Name
				goods.SPUID = spuReply.Sku.SpuId
				goods.Images = spuReply.Sku.Spu.Images
			}

		}
	}

	return &protos.GetCardReply{
		CardInformation: returnCardData,
	}, nil
}

func validateGetCard(req *protos.GetCardRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.CardId, CardIdRule...),
		validation.Field(&req.ShopId, ShopIdRule...),
	)
}
