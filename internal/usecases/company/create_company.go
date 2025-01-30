package company

import (
	"github.com/inditekno/office-backend/internal/dto"
)

func (uc *useCase) CreateCompany(companyRequest *dto.CompanyRequest) (*dto.CompanyResponse, error) {
	newCompany := &dto.CompanyRequest{
		Name: companyRequest.Name,
	}
	companyId, err := uc.companyRepository.CreateCompany(newCompany)
	if err != nil {
		return nil, err
	}

	companyResponse := &dto.CompanyResponse{
		ID:   *companyId,
		Name: companyRequest.Name,
	}

	return companyResponse, nil
}
