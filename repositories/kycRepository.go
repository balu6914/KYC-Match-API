package repositories

import (
	"context"
	"your_project/models"
)

type KYCRepository interface {
	FindCustomerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error)
}
