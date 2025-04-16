package usecases

import (
	"context"
	"fmt"

	"github.com/balu6914/KYC-Match-API/models"
	"github.com/balu6914/KYC-Match-API/repositories"
)

// kycUseCaseImpl implements the KYCUseCase interface
type kycUseCaseImpl struct {
	repo repositories.KYCRepository
}

// NewKYCUseCaseImpl creates a new instance of kycUseCaseImpl
func NewKYCUseCaseImpl(repo repositories.KYCRepository) KYCUseCase {
	return &kycUseCaseImpl{repo: repo}
}

// MatchCustomer handles the business logic for matching customer data
func (u *kycUseCaseImpl) MatchCustomer(ctx context.Context, req models.KYCRequest) (*models.KYCResponse, error) {
	if req.PhoneNumber == "" && allFieldsEmpty(req) {
		return nil, fmt.Errorf("at least one field besides phoneNumber must be provided")
	}

	customer, err := u.repo.FindCustomerByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return nil, fmt.Errorf("customer not found")
	}

	response := &models.KYCResponse{}
	response.IDDocumentMatch = matchField(req.IDDocument, customer.IDDocument)
	response.NameMatch = matchField(req.Name, customer.Name)
	response.GivenNameMatch = matchField(req.GivenName, customer.GivenName)
	response.FamilyNameMatch = matchField(req.FamilyName, customer.FamilyName)
	response.NameKanaHankakuMatch = matchField(req.NameKanaHankaku, customer.NameKanaHankaku)
	response.NameKanaZenkakuMatch = matchField(req.NameKanaZenkaku, customer.NameKanaZenkaku)
	response.MiddleNamesMatch = matchField(req.MiddleNames, customer.MiddleNames)
	response.FamilyNameAtBirthMatch = matchField(req.FamilyNameAtBirth, customer.FamilyNameAtBirth)
	response.AddressMatch = matchField(req.Address, customer.Address)
	response.StreetNameMatch = matchField(req.StreetName, customer.StreetName)
	response.StreetNumberMatch = matchField(req.StreetNumber, customer.StreetNumber)
	response.PostalCodeMatch = matchField(req.PostalCode, customer.PostalCode)
	response.RegionMatch = matchField(req.Region, customer.Region)
	response.LocalityMatch = matchField(req.Locality, customer.Locality)
	response.CountryMatch = matchField(req.Country, customer.Country)
	response.HouseNumberExtensionMatch = matchField(req.HouseNumberExtension, customer.HouseNumberExtension)
	response.BirthdateMatch = matchField(req.Birthdate, customer.Birthdate.Format("2006-01-02"))
	response.EmailMatch = matchField(req.Email, customer.Email)
	response.GenderMatch = matchField(req.Gender, customer.Gender)

	return response, nil
}

// allFieldsEmpty checks if all request fields (except phoneNumber) are empty
func allFieldsEmpty(req models.KYCRequest) bool {
	return req.IDDocument == "" && req.Name == "" && req.GivenName == "" && req.FamilyName == "" &&
		req.NameKanaHankaku == "" && req.NameKanaZenkaku == "" && req.MiddleNames == "" &&
		req.FamilyNameAtBirth == "" && req.Address == "" && req.StreetName == "" &&
		req.StreetNumber == "" && req.PostalCode == "" && req.Region == "" && req.Locality == "" &&
		req.Country == "" && req.HouseNumberExtension == "" && req.Birthdate == "" &&
		req.Email == "" && req.Gender == ""
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
