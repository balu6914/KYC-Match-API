package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/balu6914/KYC-Match-API/config"
	"github.com/balu6914/KYC-Match-API/models"
)

// harperdbRepository implements the KYCRepository interface using HarperDB REST API
type harperdbRepository struct {
	config *config.Config
}

// NewHarperDBRepository creates a new instance of harperdbRepository
func NewHarperDBRepository(cfg *config.Config) (KYCRepository, error) {
	return &harperdbRepository{config: cfg}, nil
}

// FindCustomerByPhoneNumber retrieves a customer by phone number from HarperDB
func (r *harperdbRepository) FindCustomerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error) {
	url := fmt.Sprintf("http://%s:%d/rest/%s/customers", r.config.HarperDBHost, r.config.HarperDBPort, r.config.HarperDBSchema)
	query := fmt.Sprintf(`{"operation":"sql","sql":"SELECT * FROM %s.customers WHERE phoneNumber = '%s'"}`, r.config.HarperDBSchema, phoneNumber)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(query))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(r.config.HarperDBUsername, r.config.HarperDBPassword)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HarperDB request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []models.Customer
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("customer not found")
	}

	return &result[0], nil
}
