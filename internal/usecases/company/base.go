package company

import (
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/dto"
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/repositories"
)

type UseCase interface {
	GetCompanies() (dto.CompaniesResponse, error)
	CreateCompany(in *dto.CompanyRequest) (*dto.CompanyResponse, error)
}

type useCase struct {
	companyRepository repositories.CompanyRepository
}


func NewUseCase(companyRepository repositories.CompanyRepository) UseCase {
	return &useCase{companyRepository: companyRepository}
}
