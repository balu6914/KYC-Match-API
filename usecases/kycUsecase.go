package usecases

import (
	"context"
	"your_project/models"
)

type KYCUseCase interface {
	MatchCustomer(ctx context.Context, req models.KYCRequest) (*models.KYCResponse, error)
}
