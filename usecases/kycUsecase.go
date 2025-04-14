package usecases

import (
	"context"
	"/models"
)

type KYCUseCase interface {
	MatchCustomer(ctx context.Context, req models.KYCRequest) (*models.KYCResponse, error)
}
