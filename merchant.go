package enum

// MerchantStatusEnum 商户状态
type MerchantStatusEnum int

// MerchantTypeEnum 商户类型
type MerchantTypeEnum int

// 商户状态
const (
	// MerchantUnknown 0: 商户未知类型；
	MerchantUnknown MerchantStatusEnum = iota
	// MerchantInit 1: 商户申请；
	MerchantInit
	// MerchantApproved 2: 商户审批通过；
	MerchantApproved
	// MerchantDisabled 3: 商户禁用；
	MerchantDisabled
)

// 商户类型
const (
	// MerchantTypeUnknown 0: 商户未知类型；
	MerchantTypeUnknown MerchantStatusEnum = iota
	// MerchantOrdinary 1: 普通类型；
	MerchantOrdinary

	// MerchantJoin 2: 加盟类型；
	MerchantJoin
	// MerchantOther 3: 其他类型；
	MerchantOther
)
