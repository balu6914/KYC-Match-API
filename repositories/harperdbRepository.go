package repositories

import (
	"context"
	"fmt"
	"your_project/config"
	"your_project/models"

	"github.com/harperdb/harperdb-go"
)

type harperdbRepository struct {
	client *harperdb.Client
	config *config.Config
}

func NewHarperDBRepository(cfg *config.Config) (KYCRepository, error) {
	client := harperdb.NewClient(cfg.HarperDBHost, cfg.HarperDBPort, cfg.HarperDBUsername, cfg.HarperDBPassword)
	return &harperdbRepository{client: client, config: cfg}, nil
}

func (r *harperdbRepository) FindCustomerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error) {
	var customer models.Customer
	err := r.client.SQL(ctx, fmt.Sprintf("SELECT * FROM %s.customers WHERE phoneNumber = '%s'", r.config.HarperDBSchema, phoneNumber), &customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
//