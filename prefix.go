package enum

// redis 缓存key 前缀
const (
	// QueueSeq 排队号码前缀 例如: queue_seq_<20220825>_<store_id>_R
	QueueSeq = "queue_seq"

	// SmsCode 验证码 例如: sms_code_<kind>_282922
	SmsCode = "sms_code"

	// LoginToken 登录token
	LoginToken = "login_token"

	// CommodityStocks 商品库存 例如: commodity_stocks_<store_id>_<product_id>
	CommodityStocks = "commodity_stocks"

	// 订单下单计时器 (商品下单后需要在指定时间内完成支付，这里会倒计时，用于展示在待支付订单中，为时间戳)
	// 例如: order_place_time_ttl_<order_id>
	OrderPlaceTimeTTL = "order_place_time_ttl"
)
