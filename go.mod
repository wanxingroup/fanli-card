module dev-gitlab.wanxingrowth.com/fanli/card

go 1.13

require (
	dev-gitlab.wanxingrowth.com/fanli/coupon v0.0.0-20200813032610-21c19c2b9aea
	dev-gitlab.wanxingrowth.com/fanli/goods/v2 v2.1.8
	dev-gitlab.wanxingrowth.com/fanli/order/v2 v2.0.13
	dev-gitlab.wanxingrowth.com/fanli/rebate v0.0.0-20200928130840-0530b217daf7
	dev-gitlab.wanxingrowth.com/fanli/user v0.0.2
	dev-gitlab.wanxingrowth.com/wanxin-go-micro/base v0.2.27
	github.com/go-ozzo/ozzo-validation/v4 v4.2.2
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.12
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace dev-gitlab.wanxingrowth.com/fanli/coupon => github.com/wanxingroup/fanli-coupon v0.0.1

replace dev-gitlab.wanxingrowth.com/wanxin-go-micro/base => github.com/wanxingroup/base v0.2.27

replace dev-gitlab.wanxingrowth.com/fanli/goods/v2 => github.com/wanxingroup/fanli-goods/v2 v2.0.0-20201124070303-ea0a037380c1

replace dev-gitlab.wanxingrowth.com/fanli/rebate => github.com/wanxingroup/fanli-rebate v0.0.2

replace dev-gitlab.wanxingrowth.com/fanli/order/v2 => github.com/wanxingroup/fanli-order/v2 v2.0.13

replace dev-gitlab.wanxingrowth.com/fanli/user => github.com/wanxingroup/fanli-user v0.0.2

replace dev-gitlab.wanxingrowth.com/fanli/fuyou-payment-gateway => github.com/wanxingroup/fanli-fuyou-payment-gateway v0.0.0-20200828

replace dev-gitlab.wanxingrowth.com/fanli/card => github.com/wanxingroup/fanli-card v0.0.0

replace dev-gitlab.wanxingrowth.com/fanli/payment => github.com/wanxingroup/fanli-payment v0.0.0

replace dev-gitlab.wanxingrowth.com/fanli/merchant => github.com/wanxingroup/fanli-merchant v0.0.0
