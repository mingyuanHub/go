package services

import (
	"golang.org/x/net/context"
)

type ProdService struct {

}

func (this *ProdService) GetProdStock(ctx context.Context, in *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 26}, nil
}
