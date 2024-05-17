package db

import "time"

type Variant struct {
	ID          int64
	VariantName string
	Quantity    int32
	ProductID   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Product struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductWithVariat struct {
	ProductID   int64
	ProductName string
	VariantID   int64
	VariantName string
	Quantity    int32
}
