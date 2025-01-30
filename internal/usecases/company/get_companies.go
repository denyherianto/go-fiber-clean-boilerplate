package company

import (
	"github.com/inditekno/office-backend/internal/dto"
	"github.com/inditekno/office-backend/internal/entities"
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
