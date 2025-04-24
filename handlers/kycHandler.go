package handlers

import (
	"errors"
	"net/http"

	"github.com/balu6914/KYC-Match-API/models"
	"github.com/balu6914/KYC-Match-API/usecases"
	"github.com/labstack/echo/v4"
)

// KYCHandler defines the interface for handling KYC match requests
type KYCHandler interface {
	Match(c echo.Context) error
}

// kycHandler implements the KYCHandler interface
type kycHandler struct {
	useCase usecases.KYCUseCase
}

// NewKYCHandler creates a new instance of kycHandler
func NewKYCHandler(useCase usecases.KYCUseCase) KYCHandler {
	return &kycHandler{useCase: useCase}
}

// Match handles the POST /match endpoint to verify customer identity
func (h *kycHandler) Match(c echo.Context) error {
	var req models.KYCRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"}) // Improved error message
	}

	ctx := c.Request().Context()
	response, err := h.useCase.MatchCustomer(ctx, req)
	if err != nil {
		if errors.Is(err, usecases.ErrIdentifierNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"status":  "404",
				"code":    "IDENTIFIER_NOT_FOUND",
				"message": "No customer found for phoneNumber: " + req.PhoneNumber,
			})
		}
		switch err.Error() {
		case "at least one field besides phoneNumber must be provided":
			return c.JSON(http.StatusBadRequest, map[string]string{"status": "400", "code": "KNOW_YOUR_CUSTOMER.INVALID_PARAM_COMBINATION", "message": err.Error()})
		case "customer not found":
			return c.JSON(http.StatusNotFound, map[string]string{"status": "404", "code": "IDENTIFIER_NOT_FOUND", "message": "The phone number provided is not associated with a customer account"})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		}
	}

	return c.JSON(http.StatusOK, response)
}
