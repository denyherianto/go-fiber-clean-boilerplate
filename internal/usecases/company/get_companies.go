package company

import (
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/dto"
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/entities"
)

func (uc *useCase) GetCompanies() (dto.CompaniesResponse, error) {
	companies, err := uc.companyRepository.GetCompanies()

	if err != nil {
		emptyCompanies := []entities.Company{}
		return dto.ToCompaniesResponse(&emptyCompanies), err
	}

	companiesResponse := dto.ToCompaniesResponse(&companies)

	return companiesResponse, nil
}
