package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/balu6914/KYC-Match-API/models"
	"github.com/balu6914/KYC-Match-API/usecases"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// MockKYCRepository mocks the KYCRepository interface
type MockKYCRepository struct{}

func (m *MockKYCRepository) FindByPhoneNumber(phoneNumber string) (*models.Customer, error) {
	if phoneNumber == "+34629255833" {
		return &models.Customer{
			ID:          "1",
			PhoneNumber: "+34629255833",
			IDDocument:  "66666666q",
			Name:        "Federica Sanchez Arjona",
			GivenName:   "Federica",
			FamilyName:  "Sanchez Arjona",
			Birthdate:   "1978-08-22",
			Email:       "abc@example.com",
		}, nil
	}
	return nil, nil // Simulate non-existent customer
}

// MockKYCUseCase mocks the KYCUseCase interface
type MockKYCUseCase struct {
	MatchCustomerFunc func(ctx context.Context, req models.KYCRequest) (*models.KYCResponse, error)
}

func (m *MockKYCUseCase) MatchCustomer(ctx context.Context, req models.KYCRequest) (*models.KYCResponse, error) {
	return m.MatchCustomerFunc(ctx, req)
}

func TestMatch(t *testing.T) {
	// Set up Echo
	e := echo.New()

	// Helper to create a MatchResult with a value
	matchTrue := models.MatchResult{Value: "true"}
	matchFalse := models.MatchResult{Value: "false"}

	// Test cases
	tests := []struct {
		name           string
		payload        interface{}
		mockResponse   *models.KYCResponse
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Matching request",
			payload: models.KYCRequest{
				PhoneNumber: "+34629255833",
				IDDocument:  "66666666q",
				GivenName:   "Federica",
				FamilyName:  "Sanchez Arjona",
				Name:        "Federica Sanchez Arjona",
				Birthdate:   "1978-08-22",
				Email:       "abc@example.com",
			},
			mockResponse: &models.KYCResponse{
				IDDocumentMatch:           matchTrue,
				NameMatch:                 matchTrue,
				GivenNameMatch:            matchTrue,
				FamilyNameMatch:           matchTrue,
				NameKanaHankakuMatch:      matchTrue,
				NameKanaZenkakuMatch:      matchTrue,
				MiddleNamesMatch:          matchTrue,
				FamilyNameAtBirthMatch:    matchTrue,
				AddressMatch:              matchTrue,
				StreetNameMatch:           matchTrue,
				StreetNumberMatch:         matchTrue,
				PostalCodeMatch:           matchTrue,
				RegionMatch:               matchTrue,
				LocalityMatch:             matchTrue,
				CountryMatch:              matchTrue,
				HouseNumberExtensionMatch: matchTrue,
				BirthdateMatch:            matchTrue,
				EmailMatch:                matchTrue,
				GenderMatch:               matchTrue,
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"idDocumentMatch": {"value": "true"},
				"nameMatch": {"value": "true"},
				"givenNameMatch": {"value": "true"},
				"familyNameMatch": {"value": "true"},
				"nameKanaHankakuMatch": {"value": "true"},
				"nameKanaZenkakuMatch": {"value": "true"},
				"middleNamesMatch": {"value": "true"},
				"familyNameAtBirthMatch": {"value": "true"},
				"addressMatch": {"value": "true"},
				"streetNameMatch": {"value": "true"},
				"streetNumberMatch": {"value": "true"},
				"postalCodeMatch": {"value": "true"},
				"regionMatch": {"value": "true"},
				"localityMatch": {"value": "true"},
				"countryMatch": {"value": "true"},
				"houseNumberExtensionMatch": {"value": "true"},
				"birthdateMatch": {"value": "true"},
				"emailMatch": {"value": "true"},
				"genderMatch": {"value": "true"}
			}`,
		},
		{
			name: "Non-matching request",
			payload: models.KYCRequest{
				PhoneNumber: "+34629255833",
				IDDocument:  "wrong_id",
				GivenName:   "Wrong",
				FamilyName:  "Wrong",
				Name:        "Wrong Name",
				Birthdate:   "2000-01-01",
				Email:       "wrong@example.com",
			},
			mockResponse: &models.KYCResponse{
				IDDocumentMatch:           matchFalse,
				NameMatch:                 matchFalse,
				GivenNameMatch:            matchFalse,
				FamilyNameMatch:           matchFalse,
				NameKanaHankakuMatch:      matchTrue,
				NameKanaZenkakuMatch:      matchTrue,
				MiddleNamesMatch:          matchTrue,
				FamilyNameAtBirthMatch:    matchTrue,
				AddressMatch:              matchTrue,
				StreetNameMatch:           matchTrue,
				StreetNumberMatch:         matchTrue,
				PostalCodeMatch:           matchTrue,
				RegionMatch:               matchTrue,
				LocalityMatch:             matchTrue,
				CountryMatch:              matchTrue,
				HouseNumberExtensionMatch: matchTrue,
				BirthdateMatch:            matchFalse,
				EmailMatch:                matchFalse,
				GenderMatch:               matchTrue,
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"idDocumentMatch": {"value": "false"},
				"nameMatch": {"value": "false"},
				"givenNameMatch": {"value": "false"},
				"familyNameMatch": {"value": "false"},
				"nameKanaHankakuMatch": {"value": "true"},
				"nameKanaZenkakuMatch": {"value": "true"},
				"middleNamesMatch": {"value": "true"},
				"familyNameAtBirthMatch": {"value": "true"},
				"addressMatch": {"value": "true"},
				"streetNameMatch": {"value": "true"},
				"streetNumberMatch": {"value": "true"},
				"postalCodeMatch": {"value": "true"},
				"regionMatch": {"value": "true"},
				"localityMatch": {"value": "true"},
				"countryMatch": {"value": "true"},
				"houseNumberExtensionMatch": {"value": "true"},
				"birthdateMatch": {"value": "false"},
				"emailMatch": {"value": "false"},
				"genderMatch": {"value": "true"}
			}`,
		},
		{
			name: "Invalid request - missing fields",
			payload: models.KYCRequest{
				PhoneNumber: "+34629255833",
			},
			mockResponse:   nil,
			mockError:      errors.New("at least one field besides phoneNumber must be provided"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"status":"400","code":"KNOW_YOUR_CUSTOMER.INVALID_PARAM_COMBINATION","message":"at least one field besides phoneNumber must be provided"}`,
		},
		{
			name: "Non-existent phone number",
			payload: models.KYCRequest{
				PhoneNumber: "+99999999999",
				IDDocument:  "12345678z",
			},
			mockResponse:   nil,
			mockError:      usecases.ErrIdentifierNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"status":"404","code":"IDENTIFIER_NOT_FOUND","message":"No customer found for phoneNumber: +99999999999"}`,
		},
		{
			name: "Empty phone number",
			payload: models.KYCRequest{
				PhoneNumber: " ",
			},
			mockResponse:   nil,
			mockError:      errors.New("at least one field besides phoneNumber must be provided"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"status":"400","code":"KNOW_YOUR_CUSTOMER.INVALID_PARAM_COMBINATION","message":"at least one field besides phoneNumber must be provided"}`,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock usecase
			mockUsecase := &MockKYCUseCase{
				MatchCustomerFunc: func(ctx context.Context, req models.KYCRequest) (*models.KYCResponse, error) {
					return tt.mockResponse, tt.mockError
				},
			}
			handler := NewKYCHandler(mockUsecase)

			// Convert payload to JSON
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(http.MethodPost, "/match", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Call the handler
			err := handler.Match(c)
			assert.NoError(t, err)

			// Check status code
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Check response body
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}
