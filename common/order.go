package common

type OrderStatus string

const (
	PENDING OrderStatus = "pending"
	PARTIAL OrderStatus = "partial"
	OFFICE  OrderStatus = "office"
	PAID    OrderStatus = "paid"
)

type TestType string

const (
	ANALYZE  TestType = "analyze"
	HARDNESS TestType = "hardness"
)
