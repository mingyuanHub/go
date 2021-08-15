package services

import (
	"golang.org/x/net/context"
)

type MaService struct {

}

func (this *MaService) GetName(ctx context.Context, request *MaRequest) (*MaResponse, error) {
	return &MaResponse{Name:222}, nil
}
