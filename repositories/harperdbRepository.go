package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/balu6914/KYC-Match-API/config"
	"github.com/balu6914/KYC-Match-API/models"
)

type HarperDBRepository struct {
	config *config.Config
}

func NewHarperDBRepository(cfg *config.Config) (*HarperDBRepository, error) {
	return &HarperDBRepository{config: cfg}, nil
}

func (r *HarperDBRepository) FindCustomerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Customer, error) {
	url := fmt.Sprintf("http://%s:%d", r.config.HarperDBHost, r.config.HarperDBPort)
	query := fmt.Sprintf(`{"operation":"sql","sql":"SELECT * FROM %s.customers WHERE phoneNumber = '%s'"}`, r.config.HarperDBSchema, phoneNumber)

	fmt.Printf("HarperDB Query: %s\n", query)

	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(query))
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("HarperDB Response: Status=%d, Body=%s\n", resp.StatusCode, string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("HarperDB request failed with status: " + resp.Status)
	}

	var customers []models.Customer
	if err := json.Unmarshal(body, &customers); err != nil {
		fmt.Printf("Unmarshal Error: %v\n", err)
		return nil, err
	}

	fmt.Printf("Unmarshalled Customers: %v\n", customers)

	if len(customers) == 0 {
		fmt.Println("No customers found")
		return nil, nil
	}

	return &customers[0], nil
}
