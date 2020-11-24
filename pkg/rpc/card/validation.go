package card

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"dev-gitlab.wanxingrowth.com/fanli/card/pkg/constant"
)

var ShopIdRule = []validation.Rule{
	validation.Required.ErrorObject(
		validation.NewError(constant.ErrorCodeShopIdEmpty, constant.ErrorMessageShopIdEmpty),
	),
}

var UserIdRule = []validation.Rule{
	validation.Required.ErrorObject(
		validation.NewError(constant.ErrorCodeUserIdEmpty, constant.ErrorMessageUserIdEmpty),
	),
}

var CardIdRule = []validation.Rule{
	validation.Required.ErrorObject(
		validation.NewError(constant.ErrorCodeCardIdEmpty, constant.ErrorMessageCardIdEmpty),
	),
}

var OrderIdRule = []validation.Rule{
	validation.Required.ErrorObject(
		validation.NewError(constant.ErrorCodeOrderIdEmpty, constant.ErrorMessageOrderIdEmpty),
	),
}

var NameRule = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeNameEmpty, constant.ErrorMessageNameEmpty)),
	validation.RuneLength(2, 40).ErrorObject(validation.NewError(constant.ErrorCodeNameLengthOutOfRange, constant.ErrorMessageNameLengthOutOfRange)),
}

var DescriptionRule = []validation.Rule{
	validation.RuneLength(0, 1024).ErrorObject(validation.NewError(constant.ErrorCodeDescriptionLengthOutOfRange, constant.ErrorMessageDescriptionLengthOutOfRange)),
}

var PriceRule = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodePriceRequired, constant.ErrorMessagePriceRequired)),
}

var RenewPriceRule = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeRenewPriceRequired, constant.ErrorMessageRenewPriceRequired)),
}

var SortRule = []validation.Rule{
	validation.Max(uint32(255)).ErrorObject(validation.NewError(constant.ErrorCodeSortOutOfRange, constant.ErrorMessageSortOutOfRange)),
}

var BackgroundImageRule = []validation.Rule{
	validation.Length(0, 200).ErrorObject(validation.NewError(constant.ErrorCodeBackgroundImageLengthOutOfRange, constant.ErrorMessageBackgroundImageLengthOutOfRange)),
}

var ValidityPeriodRule = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeValidityPeriodEmpty, constant.ErrorMessageValidityPeriodEmpty)),
}

var CouponIdRule = []validation.Rule{
	validation.Required.ErrorObject(
		validation.NewError(constant.ErrorCodeCouponIdEmpty, constant.ErrorMessageCouponIdEmpty),
	),
}

var CouponCountRule = []validation.Rule{
	validation.Required.ErrorObject(
		validation.NewError(constant.ErrorCodeCouponCountEmpty, constant.ErrorMessageCouponCountEmpty),
	),
}
