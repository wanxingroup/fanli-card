package constant

const (
	ErrorCodeCreateCardFailed    = 514001
	ErrorMessageCreateCardFailed = "创建权益卡失败，内部服务暂时不可用"
)

const (
	ErrorCodeModifyCardFailed    = 514002
	ErrorMessageModifyCardFailed = "更新权益卡价格项目失败，内部服务暂时不可用"
)

const (
	ErrorCodeGetCardFailed    = 514003
	ErrorMessageGetCardFailed = "获取权益卡失败，内部服务暂时不可用"
)

const (
	ErrorCodeGetCardListFailed    = 514004
	ErrorMessageGetCardListFailed = "获取权益卡列表失败，内部服务暂时不可用"
)

const (
	ErrorCodeGetUserCardListFailed    = 514005
	ErrorMessageGetUserCardListFailed = "获取用户权益卡失败，内部服务暂时不可用"
)

const (
	ErrorCodeCreateOrderFailed    = 514006
	ErrorMessageCreateOrderFailed = "创建订单失败，内部服务暂时不可用"
)

const (
	ErrorCodeOrderStatusError    = 514007
	ErrorMessageOrderStatusError = "订单状态不正确"
)

const (
	ErrorCodeGetCouponFailed    = 514008
	ErrorMessageGetCouponFailed = "获取优惠券服务错误"
)

const (
	ErrorCodeGetRebateFailed    = 514009
	ErrorMessageGetRebateFailed = "获取返利服务错误"
)

const (
	ErrorCodeShopIdEmpty    = "414001"
	ErrorMessageShopIdEmpty = "店铺 ID 为必填"
)

const (
	ErrorCodeNameEmpty    = "414002"
	ErrorMessageNameEmpty = "店铺名称为必填"
)

const (
	ErrorCodeNameLengthOutOfRange    = "414003"
	ErrorMessageNameLengthOutOfRange = "权益卡名称长度超出范围，只接受长度为2到40个字"
)

const (
	ErrorCodeDescriptionLengthOutOfRange    = "414004"
	ErrorMessageDescriptionLengthOutOfRange = "店铺描述长度超出范围，只接受最大1024个字"
)

const (
	ErrorCodeSortOutOfRange    = "414005"
	ErrorMessageSortOutOfRange = "店铺排序号超出范围，最大只能设置255"
)

const (
	ErrorCodeBackgroundImageLengthOutOfRange    = "414006"
	ErrorMessageBackgroundImageLengthOutOfRange = "背景图片长度超出范围，最大只能设置200个字节"
)

const (
	ErrorCodeValidityPeriodEmpty    = "414007"
	ErrorMessageValidityPeriodEmpty = "有效期为必填"
)

const (
	ErrorCodeCardIdEmpty    = "414008"
	ErrorMessageCardIdEmpty = "权益卡 ID 为必填"
)

const (
	ErrorCodeCouponIdEmpty    = "414009"
	ErrorMessageCouponIdEmpty = "优惠券 ID 为必填"
)

const (
	ErrorCodeCouponCountEmpty    = "414010"
	ErrorMessageCouponCountEmpty = "优惠券数量为必填"
)

const (
	ErrorCodeDataStructureInvalid    = "414011"
	ErrorMessageDataStructureInvalid = "输入的数据结构错误"
)

const (
	ErrorCodeCardNotExist    = 414012
	ErrorMessageCardNotExist = "权益卡不存在"
)

const (
	ErrorCodePriceRequired    = "414013"
	ErrorMessagePriceRequired = "权益卡选项购买金额为必填"
)

const (
	ErrorCodeRenewPriceRequired    = "414014"
	ErrorMessageRenewPriceRequired = "权益卡选项续卡金额为必填"
)

const (
	ErrorCodeSetCardStatusFailed    = 414015
	ErrorMessageSetCardStatusFailed = "设置权益卡状态失败"
)

const (
	ErrorCodeUserIdEmpty    = "414016"
	ErrorMessageUserIdEmpty = "用户 ID 为必填"
)

const (
	ErrorCodeOrderIdEmpty    = "414017"
	ErrorMessageOrderIdEmpty = "订单 ID 为必填"
)

const (
	ErrorCodeTransactionError    = 514005
	ErrorMessageTransactionError = "数据库事务出错"
)
