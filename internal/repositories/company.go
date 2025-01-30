package repositories

import (
	"github.com/denyherianto/go-fiber-clean-boilerplate/database"
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/dto"
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/entities"
)

type CompanyRepository interface {
	GetCompanies() ([]entities.Company, error)
	GetCompany(id uint) (entities.Company, error)
	CreateCompany(in *dto.CompanyRequest) (*uint, error)
	UpdateCompany(id uint, in *dto.CompanyRequest) error
	DeleteCompany(id uint) error
}

type companyRepository struct {
	db database.Database
}

func (r *companyRepository) GetCompanies() ([]entities.Company, error) {
	var companies []entities.Company

	result := r.db.GetDb().Find(&companies)
	if result.Error != nil {
		return nil, result.Error
	}

	return companies, nil
}

func (r *companyRepository) GetCompany(id uint) (entities.Company, error) {
	var company entities.Company

	result := r.db.GetDb().First(&company, id)
	if result.Error != nil {
		return company, result.Error
	}

	return company, nil
}

func (r *companyRepository) CreateCompany(in *dto.CompanyRequest) (*uint, error) {
	company := dto.ToCompany(in)

	if err := r.db.GetDb().Create(&company).Error; err != nil {
		return nil, err
	}

	return &company.ID, nil
}

func (r *companyRepository) UpdateCompany(id uint, in *dto.CompanyRequest) error {
	company := dto.ToCompany(in)

	result := r.db.GetDb().Where("id = ?", id).Updates(&company)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *companyRepository) DeleteCompany(id uint) error {
	company := entities.Company{}

	result := r.db.GetDb().Delete(&company, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewCompanyRepository(db database.Database) CompanyRepository {
	return &companyRepository{db: db}
}
