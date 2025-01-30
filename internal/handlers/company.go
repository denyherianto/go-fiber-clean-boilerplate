package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inditekno/office-backend/internal/dto"
	"github.com/inditekno/office-backend/internal/usecases/company"
	validatorUtils "github.com/inditekno/office-backend/pkg/validator"
)

type CompanyHandler interface {
	GetCompanies(c *fiber.Ctx) error
	GetCompany(c *fiber.Ctx) error
	CreateCompany(c *fiber.Ctx) error
	UpdateCompany(c *fiber.Ctx) error
	DeleteCompany(c *fiber.Ctx) error
}

type companyHandler struct {
	companyUseCase company.UseCase
}

func (h *companyHandler) GetCompanies(c *fiber.Ctx) error {
	companiesResponse, err := h.companyUseCase.GetCompanies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(companiesResponse)
}

func (h *companyHandler) GetCompany(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *companyHandler) UpdateCompany(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *companyHandler) DeleteCompany(c *fiber.Ctx) error {
	panic("unimplemented")
}

// CreateCompany func for creates a new company.
// @Description Create a new company.
// @Summary create a new company
// @Tags Companies
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Success 200 {object} entities.Company
// @Security ApiKeyAuth
func (h *companyHandler) CreateCompany(c *fiber.Ctx) error {
	// Create new Company struct
	reqBody := new(dto.CompanyRequest)

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&reqBody); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Validate company fields.
	validate := validatorUtils.NewValidator()
	if err := validate.Struct(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": validatorUtils.ValidatorErrors(err),
		})
	}

	company, err := h.companyUseCase.CreateCompany(reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": company,
		"meta": fiber.Map{},
	})
}

func NewCompanyHandler(companyUseCase company.UseCase) CompanyHandler {
	return &companyHandler{
		companyUseCase,
	}
}

// GetCompanies func gets all exists companies.
// @Description Get all exists companies.
// @Summary get all exists companies
// @Tags Companies
// @Accept json
// @Produce json
// @Success 200 {array} entities.Company
// @Router /v1/companies [get]
// func (h *CompanyHandler) GetCompanies(c *fiber.Ctx) error {
// 	// Get all companies.
// 	companies, err := repositories.GetCompanies()
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error":   fiber.StatusInternalServerError,
// 			"message": err.Error(),
// 		})
// 	}

// 	// Return status 200 OK.
// 	return c.JSON(fiber.Map{
// 		"data": companies,
// 		"meta": fiber.Map{
// 			"total": len(companies),
// 		},
// 	})
// }

// func NewCompanyHandler(companyUseCase company.UseCase) CompanyHandler {
// 	return &companyHandler{
// 		companyUseCase: companyUseCase,
// 	}
// }
