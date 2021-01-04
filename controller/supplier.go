package controller

import (
	"context"

	supplierpb "github.com/kalpg69/itas_supplier_api/api"
	"github.com/kalpg69/itas_supplier_api/domain"
	"github.com/kalpg69/itas_supplier_api/service"
)

type supplierCtrl struct {
	iService service.IService
}

func NewController(iService service.IService) supplierpb.SupplierServer {
	return &supplierCtrl{iService: iService}
}

func (s *supplierCtrl) Create(ctx context.Context, req *supplierpb.CreateRequest) (*supplierpb.CreateResponse, error) {

	id, err := s.iService.Create(ctx, &domain.Supplier{
		SupplierCode: req.GetCode(),
		SupplierName: req.GetName(),
		Email:        req.GetEmail(),
	})
	if err != nil {
		return nil, err
	}
	return &supplierpb.CreateResponse{
		Id: *id,
	}, nil
}
