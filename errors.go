package enum

import (
	"errors"
)

var (
	// ErrOrderHasExpire 订单未在规定时间内完成下单
	ErrOrderHasExpire = errors.New("You order has expire, please try to place again")

	// ErrOrderHasExpire 订单库存不足
	ErrOrderHasOutOfStacks = errors.New("You order has out of stocks, please try to place again")
)
