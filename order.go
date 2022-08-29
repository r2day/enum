package enum

// OrderStatus 订单状态
type OrderStatusEnum int

// 订单状态(1: 下单成功；
// 2: 已经支付；
// 3: 仓库无货;
// 4: 正在排队制作;
// 5: 已经制作完成;
// 6: 等待配送
// 7: 外卖已经接单
// 8: 已经送达
// 9: 已经完成
// 10: 发起退款
// 11: 进入退款流程

// 订单状态
const (
	Unknown OrderStatusEnum = iota // value --> 0
	// 1: 下单成功；
	Init
	// 2: 已经支付；
	Paid
	// 3: 仓库无货;
	Empty
	// 4: 正在排队制作;
	OnQueue
	// 5: 已经制作完成;
	CompleteProduction
	// 6: 等待配送
	WaitForTake
	// 7: 外卖已经接单
	TakeAway
	// 8: 已经送达
	Delivered
	// 9: 已经完成
	Done
	// 10: 发起退款
	Refund
	// 11: 进入退款流程
	DoneRefund

	// 12: 超时未支付
	OutOfPayTime
)
