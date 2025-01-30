package dto

import "github.com/denyherianto/go-fiber-clean-boilerplate/internal/entities"

type CompanyRequest struct {
	Name string `json:"name" validate:"required,min=3,max=50"`
}

type CompanyResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CompaniesResponse struct {
	Data []CompanyResponse `json:"data"`
	Meta Meta              `json:"meta"`
}

func ToCompany(in *CompanyRequest) *entities.Company {
	return &entities.Company{
		Name: in.Name,
	}
}

func ToCompaniesResponse(in *[]entities.Company) CompaniesResponse {
	out := CompaniesResponse{}
	out.Data = []CompanyResponse{}

	for _, company := range *in {
		out.Data = append(out.Data, CompanyResponse{
			ID:   company.ID,
			Name: company.Name,
		})
		out.Meta = Meta{
			Total: len(*in),
		}
	}
	return out
}
