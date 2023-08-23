package common

type OrderStatus string
type TestType string

const (
	PENDING OrderStatus = "pending"
	PARTIAL OrderStatus = "partial"
	OFFICE  OrderStatus = "office"
	PAID    OrderStatus = "paid"
)

const (
	ANALYZE  TestType = "analyze"
	HARDNESS TestType = "hardness"
	BOTH     TestType = "both"
)
