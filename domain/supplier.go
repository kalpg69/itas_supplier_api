package domain

import "context"

type ISupplier interface {
	Create(context.Context, *Supplier) (*int64, error)
}

type Supplier struct {
	ID           int64
	SupplierCode string
	SupplierName string
	Email        string
}
