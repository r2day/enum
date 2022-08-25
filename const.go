package enum

import "time"

const (
	// DefaultOrderExpireTime 订单默认超时时间
	DefaultOrderExpireTime = 30 * 60 * time.Second

	// SmsCodeExpireTime 验证码默认过期时间
	SmsCodeExpireTime = 5 * 60 * time.Second

	// OrderCounterMaxLimit 单个商品每次最大上限订单数量
	OrderCounterMaxLimit = 100

	// PlaceMaxLimitPerUser 单个用户每天最大下单次数
	PlaceMaxLimitPerUser = 20
)
