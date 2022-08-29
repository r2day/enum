package enum

// 消息队列queue的名称
// QueueXXX 格式
const (
	// QueueOrder 订单队列
	QueueOrder = "queue_order"
)

// QueueOrder 的消息AppID
// AppID 对应于该队列中的消息类型
// QueueAppIdXXX 格式
const (

	// QueueAppIdPlace 下单成功
	QueueAppIdPlace = "place"

	// QueueAppIdOutOfPayTime 超过支付时间
	QueueAppIdOutOfPayTime = "out_of_pay_time"

	// QueueAppIdPayDone 完成支付
	QueueAppIdPayDone = "pay_done"
)
