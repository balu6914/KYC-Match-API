package handlers

import (
	"net/http"

	"your_project/models"
	"your_project/usecases"

	"github.com/labstack/echo/v4"
)

type KYCHandler interface {
	Match(c echo.Context) error
}

type kycHandler struct {
	useCase usecases.KYCUseCase
}

func NewKYCHandler(useCase usecases.KYCUseCase) KYCHandler {
	return &kycHandler{useCase: useCase}
}

func (h *kycHandler) Match(c echo.Context) error {
	var req models.KYCRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	ctx := c.Request().Context()
	response, err := h.useCase.MatchCustomer(ctx, req)
	if err != nil {
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
