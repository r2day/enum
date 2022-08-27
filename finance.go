package enum

// Fkind Finance Kind  金融类型
type Fkind int

// 金融账号信息类型
const (
	FUnknown Fkind = iota // value --> 0
	// 1: 余额；
	Balance
	// 2: 积分
	Integral
	// 3: 优惠券;
	Coupon
	// 4: 纪念卡;
	Commemorative
)

// PayMethod 支付类型
type PayMethod int

// 金融账号信息类型
const (
	PUnknown PayMethod = iota // value --> 0
	// 1: 金融资产
	Fiance
	// 2: 微信
	Wechat
	// 3: 银行卡;
	Bank
	// 4: 现金;
	Cash
)
