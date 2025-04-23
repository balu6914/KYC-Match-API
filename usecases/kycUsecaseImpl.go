package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/balu6914/KYC-Match-API/models"
	"github.com/balu6914/KYC-Match-API/repositories"
)

// kycUseCaseImpl implements the KYCUseCase interface
type kycUseCaseImpl struct {
	repo repositories.KYCRepository
}

// Custom errors
var ErrIdentifierNotFound = errors.New("IDENTIFIER_NOT_FOUND")

// NewKYCUseCaseImpl creates a new instance of kycUseCaseImpl
func NewKYCUseCaseImpl(repo repositories.KYCRepository) KYCUseCase {
	return &kycUseCaseImpl{repo: repo}
}

// MatchCustomer handles the business logic for matching customer data
func (u *kycUseCaseImpl) MatchCustomer(ctx context.Context, req models.KYCRequest) (*models.KYCResponse, error) {
	// Validate request: at least one field besides phoneNumber must be provided
	if req.PhoneNumber == "" && allFieldsEmpty(req) {
		fmt.Printf("Invalid request: phoneNumber empty and no other fields provided\n")
		return nil, fmt.Errorf("at least one field besides phoneNumber must be provided")
	}

	// Query repository for customer
	customer, err := u.repo.FindCustomerByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		fmt.Printf("Error finding customer for phoneNumber %s: %v\n", req.PhoneNumber, err)
		return nil, fmt.Errorf("failed to find customer: %v", err)
	}
	if customer == nil {
		fmt.Printf("No customer found for phoneNumber: %s\n", req.PhoneNumber)
		return nil, ErrIdentifierNotFound
	}

	fmt.Printf("Customer found for phoneNumber %s: %+v\n", req.PhoneNumber, customer)

	// Perform field matching
	response := &models.KYCResponse{
		IDDocumentMatch:           matchField(req.IDDocument, customer.IDDocument),
		NameMatch:                 matchField(req.Name, customer.Name),
		GivenNameMatch:            matchField(req.GivenName, customer.GivenName),
		FamilyNameMatch:           matchField(req.FamilyName, customer.FamilyName),
		NameKanaHankakuMatch:      matchField(req.NameKanaHankaku, customer.NameKanaHankaku),
		NameKanaZenkakuMatch:      matchField(req.NameKanaZenkaku, customer.NameKanaZenkaku),
		MiddleNamesMatch:          matchField(req.MiddleNames, customer.MiddleNames),
		FamilyNameAtBirthMatch:    matchField(req.FamilyNameAtBirth, customer.FamilyNameAtBirth),
		AddressMatch:              matchField(req.Address, customer.Address),
		StreetNameMatch:           matchField(req.StreetName, customer.StreetName),
		StreetNumberMatch:         matchField(req.StreetNumber, customer.StreetNumber),
		PostalCodeMatch:           matchField(req.PostalCode, customer.PostalCode),
		RegionMatch:               matchField(req.Region, customer.Region),
		LocalityMatch:             matchField(req.Locality, customer.Locality),
		CountryMatch:              matchField(req.Country, customer.Country),
		HouseNumberExtensionMatch: matchField(req.HouseNumberExtension, customer.HouseNumberExtension),
		BirthdateMatch:            matchField(req.Birthdate, customer.Birthdate),
		EmailMatch:                matchField(req.Email, customer.Email),
		GenderMatch:               matchField(req.Gender, customer.Gender),
	}

	return response, nil
}

// allFieldsEmpty checks if all request fields (except phoneNumber) are empty
func allFieldsEmpty(req models.KYCRequest) bool {
	return req.IDDocument == "" &&
		req.Name == "" &&
		req.GivenName == "" &&
		req.FamilyName == "" &&
		req.NameKanaHankaku == "" &&
		req.NameKanaZenkaku == "" &&
		req.MiddleNames == "" &&
		req.FamilyNameAtBirth == "" &&
		req.Address == "" &&
		req.StreetName == "" &&
		req.StreetNumber == "" &&
		req.PostalCode == "" &&
		req.Region == "" &&
		req.Locality == "" &&
		req.Country == "" &&
		req.HouseNumberExtension == "" &&
		req.Birthdate == "" &&
		req.Email == "" &&
		req.Gender == ""
}

// matchField compares two fields and returns a MatchResult
func matchField(input, stored string) models.MatchResult {
	if input == "" || stored == "" {
		return models.MatchResult{Value: "not_available"}
	}
	if input == stored {
		return models.MatchResult{Value: "true"}
	}
	return models.MatchResult{Value: "false", Score: 85, Reason: "partial match"}
}
