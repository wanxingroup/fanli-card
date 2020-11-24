package card

type OrderStatus uint8

const (
	OrderStatusNotPay    OrderStatus = 1 //未支付
	OrderStatusClosed    OrderStatus = 5 //订单关闭
	OrderStatusCompleted OrderStatus = 6 //订单完成
	OrderStatusCancel    OrderStatus = 7 //订单取消
)
