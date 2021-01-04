package service

import (
	"context"

	"github.com/kalpg69/itas_supplier_api/domain"
)

type IService interface {
	Create(context.Context, *domain.Supplier) (*int64, error)
}

func NewService(isupplier domain.ISupplier) IService {
	return &service{isupplier: isupplier}
}

type service struct {
	isupplier domain.ISupplier
}

func (s *service) Create(ctx context.Context, sup *domain.Supplier) (*int64, error) {
	id, err := s.isupplier.Create(ctx, sup)
	if err != nil {
		return nil, err
	}
	return id, nil
}
