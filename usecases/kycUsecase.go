package usecases

import (
	"context"

	"github.com/balu6914/KYC-Match-API/models"
)

// KYCUseCase defines the interface for use case operations
type KYCUseCase interface {
	MatchCustomer(ctx context.Context, req models.KYCRequest) (*models.KYCResponse, error)
}
