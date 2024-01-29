package services

import (
	"context"

	"github.com/eduardoths/finance-api/modules/ledger/repositories"
	"github.com/eduardoths/finance-api/modules/ledger/structs"
)

type AssetService struct {
	repo repositories.AssetRepository
}

func NewAssetService(repo repositories.AssetRepository) AssetService {
	return AssetService{
		repo: repo,
	}
}

func (as AssetService) Create(ctx context.Context, data structs.Asset) (structs.Asset, error) {
	return as.repo.Create(ctx, structs.Asset{Name: data.Name, Type: data.Type})
}

func (as AssetService) Get(ctx context.Context, id structs.ID) (structs.Asset, error) {
	return as.repo.Get(ctx, id)
}

func (as AssetService) GetAll(ctx context.Context) ([]structs.Asset, error) {
	return as.repo.GetAll(ctx)
}
