package repositories

import (
	"context"

	"github.com/balu6914/KYC-Match-API/models"
)

type KYCRepository interface {
	FindCustomerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error)
}
